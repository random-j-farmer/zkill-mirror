package db

import (
	"bytes"
	"log"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	"github.com/random-j-farmer/eveapi/mapdata"
	"github.com/random-j-farmer/zkill-mirror/internal/config"
)

// SearchResultItem is a name/id pair
type SearchResultItem struct {
	Name string
	ID   uint64
}

// SearchResult contains character, corporation and alliance hits
type SearchResult struct {
	Characters   []SearchResultItem
	Corporations []SearchResultItem
	Alliances    []SearchResultItem
	SolarSystems []SearchResultItem
	Regions      []SearchResultItem
}

// Search for character, corporation or alliance
func Search(q string) (result SearchResult, err error) {
	result.Characters, err = searchByPrefix([]byte(q), charIDByName)
	if err != nil {
		return
	}

	result.Corporations, err = searchByPrefix([]byte(q), corpIDByName)
	if err != nil {
		return
	}

	result.Alliances, err = searchByPrefix([]byte(q), allIDByName)
	if err != nil {
		return
	}

	for _, system := range mapdata.FindSolarSystems(q) {
		result.SolarSystems = append(result.SolarSystems, SearchResultItem{Name: system.Name, ID: system.ID})
	}
	for _, region := range mapdata.FindRegions(q) {
		result.Regions = append(result.Regions, SearchResultItem{Name: region.Name, ID: region.ID})
	}

	return
}

// searchByPrefix returns ids by name
func searchByPrefix(prefix []byte, bucket []byte) ([]SearchResultItem, error) {
	results := make([]SearchResultItem, 0, 16)

	if config.Debug() {
		log.Printf("scanning for %s by prefix %s", bucket, prefix)
	}

	err := DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)

		c := b.Cursor()
		k, v := c.Seek(prefix)
		for i := 0; k != nil; i++ {
			if !bytes.HasPrefix(k, prefix) {
				break
			}

			results = append(results, SearchResultItem{Name: string(k), ID: dec64ID(string(v))})

			k, v = c.Next()
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "bolt.View")
	}

	if config.Debug() {
		log.Printf("scanning for %s by prefix %s: found %d", bucket, prefix, len(results))
	}

	return results, nil
}
