package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/mchmarny/snip/pkg/snip"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	bolt "go.etcd.io/bbolt"
)

const (
	dbFileName = "snip.db"
	dbFilePerm = 0600

	homeDirName = ".snip"
	homeDirPerm = 0700

	dbTimeoutSec = 3
)

var (
	dbFilePath = path.Join(getHomeDir(), dbFileName)
)

func Init() error {
	if _, err := os.Stat(dbFilePath); !errors.Is(err, os.ErrNotExist) {
		return nil
	}

	db := getDB()
	defer db.Close()

	tx, err := db.Begin(true)
	if err != nil {
		return errors.Wrap(err, "error creating transaction")
	}
	defer func() {
		if err = tx.Rollback(); err != nil && err != bolt.ErrTxClosed {
			log.Fatalf("error rolling back transaction: %v", err)
		}
	}()

	_, err = tx.CreateBucketIfNotExists([]byte("snippet"))
	if err != nil {
		return errors.Wrap(err, "error creating snippet bucket")
	}

	_, err = tx.CreateBucketIfNotExists([]byte("objective"))
	if err != nil {
		return errors.Wrap(err, "error creating objective bucket")
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, "error committing changes to db")
	}
	return nil
}

func getHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Printf("error getting home dir, using current dir instead: %v", err)
		return "."
	}
	log.Debugf("home dir: %s", home)

	dirPath := filepath.Join(home, homeDirName)
	if _, err := os.Stat(dirPath); errors.Is(err, os.ErrNotExist) {
		log.Debugf("creating dir: %s", dirPath)
		err := os.Mkdir(dirPath, homeDirPerm)
		if err != nil {
			log.Debugf("error creating dir: %s, using home: %s - %v", dirPath, home, err)
			return home
		}
	}
	return dirPath
}

func getDB() *bolt.DB {
	db, err := bolt.Open(dbFilePath, dbFilePerm, &bolt.Options{
		Timeout: dbTimeoutSec * time.Second,
	})
	if err != nil {
		log.Fatalf("error creating db: %v", err)
	}
	return db
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
	return []byte(fmt.Sprintf("%d", v.Unix()))
}
