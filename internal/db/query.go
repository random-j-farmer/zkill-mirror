package db

import (
	"bytes"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	"github.com/random-j-farmer/bobstore"
	"github.com/random-j-farmer/zkill-mirror/internal/config"
)

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

// ByCharacterID gives the latest limit killmails of the character
func ByCharacterID(characterID uint64, limit int) ([]bobstore.Ref, error) {
	prefix := []byte(fmt.Sprintf("%s%s", d64ID(characterID), d64Sep))

	return byPrefix(prefix, kmCharID, limit)
}

// ByCorporationID gives the latest limit killmails of the Corporation
func ByCorporationID(corpID uint64, limit int) ([]bobstore.Ref, error) {
	prefix := []byte(fmt.Sprintf("%s%s", d64ID(corpID), d64Sep))

	return byPrefix(prefix, kmCorpID, limit)
}

// ByAllianceID gives the latest limit killmails of the alliance
func ByAllianceID(allianceID uint64, limit int) ([]bobstore.Ref, error) {
	prefix := []byte(fmt.Sprintf("%s%s", d64ID(allianceID), d64Sep))

	return byPrefix(prefix, kmAlliID, limit)
}

// BySystemID gives the latest limit killmails of the solarsystem
func BySystemID(systemID uint64, limit int) ([]bobstore.Ref, error) {
	prefix := []byte(fmt.Sprintf("%s%s", d64ID(systemID), d64Sep))

	return byPrefix(prefix, kmBySystem, limit)
}

// ByRegionID gives the latest limit killmails of the region
func ByRegionID(regionID uint64, limit int) ([]bobstore.Ref, error) {
	prefix := []byte(fmt.Sprintf("%s%s", d64ID(regionID), d64Sep))

	return byPrefix(prefix, kmByRegion, limit)
}

// Newest returns the newest limit killmail refs
func Newest(limit int) ([]bobstore.Ref, error) {
	return byPrefix([]byte(""), kmByDate, limit)
}

// byPrefix returns killmails by bucket and prefix
func byPrefix(prefix []byte, bucket []byte, limit int) ([]bobstore.Ref, error) {
	refs := make([]bobstore.Ref, 0, limit)

	if config.Verbose() {
		log.Printf("scanning for %s by prefix %s", bucket, prefix)
	}

	err := DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)

		c := b.Cursor()
		k, v := c.Seek(prefix)
		for i := 0; i < limit && k != nil; i++ {
			var ref bobstore.Ref

			if !bytes.HasPrefix(k, prefix) {
				break
			}

			ref, err := dec64Ref(string(v))
			if err != nil {
				return err
			}
			refs = append(refs, ref)

			k, v = c.Next()
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "bolt.View")
	}

	if config.Verbose() {
		log.Printf("scanning for %s by prefix %s: found %d", bucket, prefix, len(refs))
	}

	return refs, nil
}
