// Package blobs has the handle for the bobstore blobstore
package blobs

import (
	"log"

	"github.com/random-j-farmer/bobstore"
	"github.com/random-j-farmer/zkill-mirror/internal/config"
)

// DB is the bobstore DB
var DB *bobstore.DB

// InitDB opens the bobstore blobstore
func InitDB(name string) {
	var err error
	DB, err = bobstore.OpenRW(name)
	if err != nil {
		log.Panicf("could not open blob store %s: %v", config.BobsName(), err)
	}
}

func Close() error {
	return DB.Close()
}
