package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"path"
	"strings"
	"time"

	// sqlite provides DB driver implementation
	_ "github.com/mattn/go-sqlite3"

	"github.com/mchmarny/snip/pkg/snip"
)

const (
	dbFileName   = "snip.db"
	objectiveKey = "objective"
)

func getDB() *sql.DB {
	dbPath := path.Join(getUserDirPath(), dbFileName)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("error connecting to db %s: %v", dbPath, err)
	}
	return db
}

func initDB() {

	db := getDB()
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("error pinging db: %v", err)
	}

	// snippet
	if err := makeTable(db, `CREATE TABLE IF NOT EXISTS
		snippet (
			sid TEXT PRIMARY KEY,
			stm DATETIME NOT NULL,
			raw TEXT NOT NULL,
			txt TEXT NOT NULL,
			ctx TEXT NOT NULL,
			obj TEXT NOT NULL
		)`); err != nil {
		log.Fatalf("error creating snippet table: %v", err)
	}

	// metric
	if err := makeTable(db, `CREATE TABLE IF NOT EXISTS
		metric (
			sid TEXT NOT NULL,
			key TEXT NOT NULL,
			val TEXT NOT NULL,
			PRIMARY KEY (sid, key)
		)`); err != nil {
		log.Fatalf("error creating metric table: %v", err)
	}

}

func saveSnippet(item *snip.Snippet) error {

	db := getDB()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error creating transaction: %v", err)
	}

	_, err = tx.Exec(`INSERT INTO snippet
		(sid, stm, raw, txt, ctx, obj) VALUES (?,?,?,?,?,?)`,
		item.ID, item.CreationTime, item.Raw, item.Text,
		strings.Join(item.Contexts, ","),
		strings.Join(item.Objectives, ","))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error inserting %+v to db: %v", item, err)
	}

	for _, obj := range item.Objectives {
		_, err = tx.Exec("INSERT INTO metric (sid, key, val) VALUES (?,?,?)",
			item.ID, objectiveKey, obj)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting objective %s to db: %v", obj, err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error commiting save to db: %v", err)
	}

	return nil
}

func getWeekSnippets(weekStart time.Time) (items []*snip.Snippet, err error) {

	weekEnd := weekStart.AddDate(0, 0, 7)

	db := getDB()
	defer db.Close()

	rows, err := db.Query(`SELECT
		sid, stm, raw, txt, ctx, obj
		FROM snippet
		WHERE stm >= ? AND stm <= ?
		ORDER BY stm
	`, weekStart, weekEnd)

	if err != nil {
		return nil, fmt.Errorf("error selecting snippets: %v", err)
	}

	snips := []*snip.Snippet{}
	for rows.Next() {
		snip := &snip.Snippet{}
		var ctxs, objs string
		rows.Scan(&snip.ID, &snip.CreationTime, &snip.Raw, &snip.Text, &ctxs, &objs)
		snip.Objectives = strings.Split(objs, ",")
		snip.Contexts = strings.Split(ctxs, ",")
		snips = append(snips, snip)
	}

	return snips, nil
}

func makeTable(db *sql.DB, sql string) error {

	stmt, err := db.Prepare(sql)
	if err != nil {
		return fmt.Errorf("error prepering table statement: %v", err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("error creating table: %v", err)
	}

	return nil

}
