/*

zkill-mirror mirrors zkillboard data.

It uses the queue api to pull the killmails.
The json is stored in a bobstore and boltdb
is used to index the data.

*/
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/random-j-farmer/bobstore"
	"github.com/random-j-farmer/zkill-mirror/internal/blobs"
	"github.com/random-j-farmer/zkill-mirror/internal/config"
	"github.com/random-j-farmer/zkill-mirror/internal/db"
	"github.com/random-j-farmer/zkill-mirror/internal/server"
	"github.com/random-j-farmer/zkill-mirror/internal/zkb"
)

func main() {
	cmd := config.DefaultCommand()
	if len(os.Args) == 2 {
		cmd = os.Args[1]
	}

	switch cmd {
	case "help":
		fmt.Printf("Usage: %s COMMAND", os.Args[0])
		fmt.Println(`

Where COMMAND is one of:

* help: this help text
* serve: serve a zkillboard proxy with background fetching
* reindex: re-index the database
`)
		os.Exit(0)

	case "serve":
		myInit()
		serve()

	case "reindex":
		myInit()
		reindex()

	default:
		fmt.Printf("Unknown command: %s", cmd)
		os.Exit(1)
	}
}

var bobsDB *bobstore.DB

func myInit() {
	db.InitDB(config.DBName(), config.DBNoSync())
	blobs.InitDB(config.BobsName())
}

// cleanup all resources when shutting down
func cleanup(stop chan<- struct{}, wg *sync.WaitGroup) {
	log.Printf("cleanup: stopping go routines ...")
	close(stop)
	wg.Wait()
	log.Printf("cleanup: closing databases ...")
	closeDBs()
	log.Printf("cleanup: DONE")
}

// close db and bobstore
func closeDBs() {
	err := db.CloseDB()
	if err != nil {
		log.Printf("cleanup: error closing %s: %v", config.DBName(), err)
	}

	err = blobs.Close()
	if err != nil {
		log.Printf("cleanup: error closing bobs db: %v", err)
	}
}

func serve() {
	if config.PullEnabled() {
		var wg sync.WaitGroup
		stop := make(chan struct{})
		kmQueue := make(chan *zkb.KillmailWithRef, 1)

		go zkb.PullKillmails(bobsDB, kmQueue, stop, &wg)
		go db.IndexWorker(kmQueue, &wg)

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

		go func() {
			<-sigChan
			cleanup(stop, &wg)
			os.Exit(0)
		}()
	}

	// http server is not stopped - only worker goroutines and db handle
	server.Serve()
}

// reindex the database
// this reads the whole bobstore and puts it in the database
func reindex() {
	defer closeDBs()

	cursor := bobsDB.Cursor(bobstore.Ref{})
	for cursor.Next() {
		b, err := bobsDB.Read(cursor.Ref())
		if err != nil {
			log.Fatalf("reindex: bobsDB.Read %s: %v", cursor.Ref(), err)
		}

		km, err := zkb.Parse(b)
		if err != nil {
			log.Fatalf("error parsing killmail %s: %v\n\n%s", cursor.Ref(), err, b)
		}

		log.Printf("reindexing %s\n", cursor.Ref())
		err = db.IndexKillmail(cursor.Ref(), km)
		if err != nil {
			log.Fatalf("error storing to the database: %v", err)
		}

	}
}
