// Package db ipmlements the caching boltdb
package db

import (
	"fmt"
	"log"
	"sync"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	"github.com/random-j-farmer/bobstore"
	"github.com/random-j-farmer/zkill-mirror/internal/zkb"
)

// DB is the database connection
var DB *bolt.DB

var kmByID = []byte("kmByID")
var kmByDate = []byte("kmByDate")
var kmBySystem = []byte("kmBySystem")
var kmCharID = []byte("kmCharID")
var kmCorpID = []byte("kmCharID")
var kmAlliID = []byte("kmAlliID")

// InitDB opens and initialize DB
func InitDB(dbname string) {
	var err error
	DB, err = bolt.Open(dbname, 0666, nil)
	if err != nil {
		panic(errors.Wrap(err, "bolt.Open"))
	}

	DB.Update(func(tx *bolt.Tx) error {
		for _, bucket := range [][]byte{kmByID, kmByDate, kmBySystem, kmCharID, kmCorpID, kmAlliID} {
			_, err2 := tx.CreateBucketIfNotExists(bucket)
			if err2 != nil {
				return err2
			}
		}
		return nil
	})
	if err != nil {
		panic(errors.Wrapf(err, "boltdb.Update %s", dbname))
	}
}

// CloseDB closes the database
func CloseDB() error {
	return DB.Close()
}

// IndexKillmail indexes the killmail
func IndexKillmail(ref bobstore.Ref, km *zkb.Killmail) error {
	err := DB.Batch(func(tx *bolt.Tx) error {
		refBytes := []byte(ref.String())

		type workitem struct {
			bucket []byte
			key    string
		}

		work := []workitem{
			{kmByID, fmt.Sprintf("%d", km.KillID)},
			{kmByDate, fmt.Sprintf("%s:%d", km.KillTime, km.KillID)},
			{kmBySystem, fmt.Sprintf("%d:%s:%d", km.SolarSystemID, km.KillTime, km.KillID)},
			{kmCharID, fmt.Sprintf("%d:%s:%d", km.Victim.CharID, km.KillTime, km.KillID)},
			{kmCorpID, fmt.Sprintf("%d:%s:%d", km.Victim.CorporationID, km.KillTime, km.KillID)},
			{kmAlliID, fmt.Sprintf("%d:%s:%d", km.Victim.AllianceID, km.KillTime, km.KillID)},
		}

		for _, attacker := range km.Attackers {
			work = append(work, workitem{kmCharID, fmt.Sprintf("%d:%s:%d", attacker.CharID, km.KillTime, km.KillID)})
			work = append(work, workitem{kmCorpID, fmt.Sprintf("%d:%s:%d", attacker.CorporationID, km.KillTime, km.KillID)})
			work = append(work, workitem{kmAlliID, fmt.Sprintf("%d:%s:%d", attacker.AllianceID, km.KillTime, km.KillID)})
		}
		for _, item := range work {
			b := tx.Bucket(item.bucket)
			err := b.Put([]byte(item.key), refBytes)
			if err != nil {
				return errors.Wrapf(err, "boldb.Put %s %s", b, item.key)
			}
		}

		log.Printf("indexed %s under pk:%d\n", ref, km.KillID)
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "boltdb.Batch")
	}
	return nil
}

// IndexWorker does the indexing
func IndexWorker(kmQueue <-chan *zkb.KillmailWithRef, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	for km := range kmQueue {
		err := IndexKillmail(km.Ref, km.Killmail)
		if err != nil {
			log.Printf("db.IndexWorker: error index killmail: %v", err)
		}
	}

	log.Printf("db.IndexWorker ... done")
}
