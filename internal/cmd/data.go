package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"path"
	"time"

	bolt "go.etcd.io/bbolt"

	"github.com/mchmarny/snip/pkg/snip"
)

const (
	dbFileName   = "snip.db"
	objectiveKey = "objective"
)

func getDB() *bolt.DB {
	dbPath := path.Join(getUserDirPath(), dbFileName)

	db, err := bolt.Open(dbPath, 0600, &bolt.Options{
		Timeout: 3 * time.Second,
	})
	if err != nil {
		log.Fatalf("error creating db: %v", err)
	}
	return db
}

func initDB() {

	db := getDB()
	defer db.Close()

	tx, err := db.Begin(true)
	if err != nil {
		log.Fatalf("error transaction: %v", err)
	}
	defer tx.Rollback()

	_, err = tx.CreateBucketIfNotExists([]byte("snippet"))
	if err != nil {
		log.Fatalf("error creating snippet bucket: %v", err)
	}

	_, err = tx.CreateBucketIfNotExists([]byte("objective"))
	if err != nil {
		log.Fatalf("error creating objective bucket: %v", err)
	}

	if err = tx.Commit(); err != nil {
		log.Fatalf("error commiting changes to db: %v", err)
	}

}

func saveSnippet(item *snip.Snippet) error {
	return getDB().Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("snippet"))
		buf, err := json.Marshal(item)
		if err != nil {
			return err
		}
		return b.Put(toByte(item.CreationTime), buf)
	})
}

func getPeriodSnippets(periodStart time.Time) (period *snip.Period, err error) {

	p := &snip.Period{
		PeriodStart:       periodStart,
		ObjectiveSnippets: make(map[string][]*snip.Snippet),
	}

	e := getDB().View(func(tx *bolt.Tx) error {

		c := tx.Bucket([]byte("snippet")).Cursor()

		min := toByte(periodStart)
		for k, v := c.Seek(min); k != nil; k, v = c.Next() {
			var s snip.Snippet
			if err := json.Unmarshal(v, &s); err != nil {
				return fmt.Errorf("error while decoding snippet")
			}
			if _, has := p.ObjectiveSnippets[s.Objective]; !has {
				p.ObjectiveSnippets[s.Objective] = []*snip.Snippet{}
			}
			p.ObjectiveSnippets[s.Objective] = append(p.ObjectiveSnippets[s.Objective], &s)
		}
		return nil
	})

	return p, e

}

func toByte(v time.Time) []byte {

	// b := make([]byte, 8)
	// binary.BigEndian.PutUint64(b, uint64(v.Unix()))
	// return b

	return []byte(fmt.Sprintf("%d", v.Unix()))

}
