// Package db ipmlements the caching boltdb
package db

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	"github.com/random-j-farmer/bobstore"
	"github.com/random-j-farmer/d64"
	"github.com/random-j-farmer/zkill-mirror/internal/config"
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
func InitDB(dbname string, noSync bool) {
	var err error
	DB, err = bolt.Open(dbname, 0666, nil)
	if err != nil {
		panic(errors.Wrap(err, "bolt.Open"))
	}
	DB.NoSync = noSync

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
		panic(errors.Wrapf(err, "bolt.Update %s", dbname))
	}
}

// CloseDB closes the database
func CloseDB() error {
	return DB.Close()
}

const d64Sep = "|"

// ids seem to have 8 characters.
// for 9 characters, they encode like this:
// 987654321|20160831214013|687654321 ===> 34
// uraXl|0Mln9S|czBLl ===> 18
// it seems some ids are longer ... but only solar system ids ??
func d64ID(id uint64) string {
	return d64.EncodeUInt64(id, 5)
}

func d64TimeID(dt string, id uint64) string {
	return strings.Join([]string{d64Time(dt), d64ID(id)}, d64Sep)
}

func d64IDTimeID(id uint64, dt string, id2 uint64) string {
	return strings.Join([]string{d64Time(dt), d64Time(dt), d64ID(id2)}, d64Sep)
}

func d64Ref(ref bobstore.Ref) string {
	return strings.Join([]string{
		d64.EncodeUInt64(uint64(ref.Fno), 3),
		d64.EncodeUInt64(uint64(ref.Pos), 5),
	}, d64Sep)
}

func dec64Ref(s string) (bobstore.Ref, error) {
	var ref bobstore.Ref

	parts := strings.Split(s, d64Sep)
	if len(parts) != 2 {
		return ref, fmt.Errorf("not a d64Ref: %s %d", s, len(parts))
	}

	fno, err := d64.DecodeUInt64(parts[0])
	if err != nil {
		return ref, errors.Wrapf(err, "dec64Ref %s", s)
	}
	pos, err := d64.DecodeUInt64(parts[1])
	if err != nil {
		return ref, errors.Wrapf(err, "dec64Ref %s", s)
	}

	ref.Fno = uint16(fno)
	ref.Pos = uint32(pos)

	return ref, nil
}

var d64Epoch uint64

func init() {
	dt, err := time.Parse(time.RFC3339, "2040-01-01T00:00:00Z")
	if err != nil {
		log.Panicf("can not initialize d64Epoch: %v", err)
	}

	d64Epoch = uint64(dt.UTC().Unix())
}

// d64Time is number of seconds UNTIL 2040-01-01
// this way, natural sort order for timestamps is most recent first
// current dates would encode with only 5 digits, but for 2006
// that does not work anymore ... maybe we'll need old killmails
// one day.
func d64Time(s string) string {
	// 01/02 03:04:05PM '06 -07:00  ===> JAN 01
	// "2016.08.28 18:10:28"
	dt, err := time.Parse("2006.01.02 15:04:05", s)
	if err != nil {
		log.Printf("warning: can not parse killtime %s", s)
		return d64.EncodeUInt64(0, 6)
	}

	seconds := d64Epoch - uint64(dt.UTC().Unix())
	return d64.EncodeUInt64(seconds, 6)
}

// IndexKillmail indexes the killmail
func IndexKillmail(ref bobstore.Ref, km *zkb.Killmail) error {
	err := DB.Batch(func(tx *bolt.Tx) error {
		refBytes := []byte(d64Ref(ref))

		type workitem struct {
			bucket []byte
			key    string
		}

		work := []workitem{
			{kmByID, d64ID(km.KillID)},
			{kmByDate, d64TimeID(km.KillTime, km.KillID)},
			{kmBySystem, d64IDTimeID(km.SolarSystemID, km.KillTime, km.KillID)},
			{kmCharID, d64IDTimeID(km.Victim.CharID, km.KillTime, km.KillID)},
			{kmCorpID, d64IDTimeID(km.Victim.CorporationID, km.KillTime, km.KillID)},
			{kmAlliID, d64IDTimeID(km.Victim.AllianceID, km.KillTime, km.KillID)},
		}

		for _, attacker := range km.Attackers {
			work = append(work, workitem{kmCharID, d64IDTimeID(attacker.CharID, km.KillTime, km.KillID)})
			work = append(work, workitem{kmCorpID, d64IDTimeID(attacker.CorporationID, km.KillTime, km.KillID)})
			work = append(work, workitem{kmAlliID, d64IDTimeID(attacker.AllianceID, km.KillTime, km.KillID)})
		}
		for _, item := range work {
			b := tx.Bucket(item.bucket)
			err := b.Put([]byte(item.key), refBytes)
			if config.Verbose() {
				log.Printf("indexing %s %s", item.key, refBytes)
			}
			if err != nil {
				return errors.Wrapf(err, "bolt.Put %s %s", b, item.key)
			}
		}
		log.Printf("indexed %s under pk:%d\n", ref, km.KillID)
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "bolt.Batch")
	}
	return nil
}

// ByKillID queries the DB by killID
func ByKillID(killID uint64) (bobstore.Ref, error) {
	var ref bobstore.Ref
	err := DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(kmByID)
		k := d64ID(killID)
		refStr := string(b.Get([]byte(k)))
		if refStr == "" {
			return fmt.Errorf("no such key: %s", k)
		}
		var err error
		ref, err = dec64Ref(refStr)
		return err
	})
	if err != nil {
		return ref, errors.Wrap(err, "bolt.View")
	}
	return ref, nil
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
