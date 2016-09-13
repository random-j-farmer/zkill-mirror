package db

import (
	"bytes"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	"github.com/random-j-farmer/bobstore"
	"github.com/random-j-farmer/d64"
	"github.com/random-j-farmer/eveapi/mapdata"
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

// SystemStat is a statistics returned by Hot
type SystemStat struct {
	RegionID         uint64  `json:"regionID"`
	RegionName       string  `json:"regionName"`
	SolarSystemID    uint64  `json:"solarSystemID"`
	SolarSystemName  string  `json:"solarSystemName"`
	Security         float32 `json:"security"`
	Kills            int64   `json:"kills"`
	TotalValue       int64   `json:"totalValue"`
	RegionKills      int64   `json:"regionKills"`
	RegionTotalValue int64   `json:"regionTotalValue"`
}

// to implement sort.Interface on ...
type systemStats []*SystemStat

func (st *systemStats) Len() int {
	return len(*st)
}

func (st *systemStats) Less(i, j int) bool {
	if (*st)[i].Kills == (*st)[j].Kills {
		return (*st)[i].TotalValue > (*st)[j].TotalValue
	}
	return (*st)[i].Kills > (*st)[j].Kills
}

func (st *systemStats) Swap(i, j int) {
	(*st)[i], (*st)[j] = (*st)[j], (*st)[i]
}

// Activity returns the hot systems in the given period
func Activity(d time.Duration) ([]*SystemStat, error) {
	now := time.Now().UTC()
	end := now.Add(-d)

	endBytes := []byte(fmt.Sprintf("%s%s", d64TimeForTime(end), d64Sep))

	var killsByRegion = make(map[uint64]int64)
	var killsBySystem = make(map[uint64]int64)
	var iskByRegion = make(map[uint64]int64)
	var iskBySystem = make(map[uint64]int64)
	var cnt int

	err := DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(kmByDate)

		c := b.Cursor()
		k, v := c.First()
		for k != nil && bytes.Compare(k, endBytes) < 0 {
			parts := bytes.Split(v, []byte(d64Sep))
			reg, _ := d64.DecodeUInt64(string(parts[1]))
			sys, _ := d64.DecodeUInt64(string(parts[2]))
			isk, _ := d64.DecodeUInt64(string(parts[3]))

			killsByRegion[reg] = killsByRegion[reg] + 1
			killsBySystem[sys] = killsBySystem[sys] + 1
			iskByRegion[reg] = iskByRegion[reg] + int64(isk)
			iskBySystem[sys] = iskBySystem[sys] + int64(isk)

			cnt++

			k, v = c.Next()
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "bolt.View")
	}

	if config.Verbose() {
		log.Printf("activity %v: scanned %d records", d, cnt)
	}

	stats := systemStats(make([]*SystemStat, 0, len(killsBySystem)))
	for sysid, syskills := range killsBySystem {
		sd := mapdata.SolarSystemByID(sysid)
		regid, rname := mapdata.RegionBySolarSystem(sysid)

		stats = append(stats, &SystemStat{
			RegionID:         regid,
			RegionName:       rname,
			SolarSystemID:    sysid,
			SolarSystemName:  sd.SolarSystemName,
			Security:         sd.Security,
			Kills:            syskills,
			RegionKills:      killsByRegion[regid],
			TotalValue:       iskBySystem[sysid],
			RegionTotalValue: iskByRegion[regid],
		})
	}
	sort.Sort(&stats)

	return stats, nil
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

// CharCorpAllianceName does a combined lookup for the char, corp and alliance name
func CharCorpAllianceName(charID, corpID, allianceID uint64) (charName, corpName, allianceName string) {
	err := DB.View(func(tx *bolt.Tx) error {
		charName = string(tx.Bucket(charNameByID).Get([]byte(d64ID(charID))))
		corpName = string(tx.Bucket(corpNameByID).Get([]byte(d64ID(corpID))))
		allianceName = string(tx.Bucket(allNameByID).Get([]byte(d64ID(allianceID))))

		return nil
	})
	if err != nil {
		log.Printf("Error looking up char/corp/alliance: %v", err)
	}

	return // named return
}
