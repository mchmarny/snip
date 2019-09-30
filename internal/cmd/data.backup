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
			obj TEXT NOT NULL,
			ctx TEXT NOT NULL
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

	// transaction for when more then one table
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error creating transaction: %v", err)
	}

	_, err = tx.Exec(`INSERT INTO snippet
		(sid, stm, raw, txt, obj, ctx) VALUES (?,?,?,?,?,?)`,
		item.ID, item.CreationTime, item.Raw, item.Text,
		item.Objective,
		strings.Join(item.Contexts, ","))
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error inserting %+v to db: %v", item, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error commiting save to db: %v", err)
	}

	return nil
}

func getPeriodSnippets(periodStart time.Time) (period *snip.Period, err error) {

	db := getDB()
	defer db.Close()

	rows, err := db.Query(`SELECT
		sid, stm, raw, txt, obj, ctx
		FROM snippet
		WHERE stm >= ?
		ORDER BY obj, stm desc
	`, periodStart)

	if err != nil {
		return nil, fmt.Errorf("error selecting snippets: %v", err)
	}

	p := &snip.Period{
		PeriodStart:       periodStart,
		ObjectiveSnippets: make(map[string][]*snip.Snippet),
	}

	for rows.Next() {
		s := &snip.Snippet{}
		var ctxs string
		rows.Scan(&s.ID, &s.CreationTime, &s.Raw,
			&s.Text, &s.Objective, &ctxs)
		s.Contexts = strings.Split(ctxs, ",")

		if _, has := p.ObjectiveSnippets[s.Objective]; !has {
			p.ObjectiveSnippets[s.Objective] = []*snip.Snippet{}
		}

		p.ObjectiveSnippets[s.Objective] = append(p.ObjectiveSnippets[s.Objective], s)
	}

	return p, nil
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
