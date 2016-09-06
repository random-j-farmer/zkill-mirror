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

	"github.com/pkg/errors"
	"github.com/random-j-farmer/bobstore"
	"github.com/random-j-farmer/zkill-mirror/internal/blobs"
	"github.com/random-j-farmer/zkill-mirror/internal/config"
	"github.com/random-j-farmer/zkill-mirror/internal/db"
	"github.com/random-j-farmer/zkill-mirror/internal/parallel"
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

// handleSignals is always called, wg.Wait just does not do anything
// if pulling is disabled.  this lets us test the shutdown procedure without
// pulling data
func handleSignals(sigChan <-chan os.Signal, stop chan<- struct{}, wg *sync.WaitGroup) {
	sig := <-sigChan
	log.Printf("signal caught: %v", sig)
	cleanup(stop, wg)
	os.Exit(0)
}

func serve() {
	var wg sync.WaitGroup
	stop := make(chan struct{})

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go handleSignals(sigChan, stop, &wg)

	if config.PullEnabled() {
		kmQueue := make(chan *zkb.Killmail, 1)
		go zkb.PullKillmails(blobs.DB, kmQueue, stop, &wg)
		go db.IndexWorker(kmQueue, &wg)
	}

	// http server is not stopped - only worker goroutines and db handle
	err := server.Serve()
	log.Printf("server.Serve(): %v", err)
}

// reindex the database
// this reads the whole bobstore and puts it in the database
func reindex() {
	defer closeDBs()

	cursor := blobs.DB.Cursor(bobstore.Ref{})
	proc := parallel.NewProcessor(config.ReindexWorkers(), readAndParse, indexKillmails)
	for cursor.Next() {
		proc.Input <- cursor.Ref()
	}
	errs := proc.Wait()
	if len(errs) > 0 {
		log.Fatalf("reindex: errors %v", errs)
	}
	log.Printf(`done.  parameters were:

		db_nosync: %v
		reindex_workers: %v
		reindex_batch_size: %v`, config.DBNoSync(), config.ReindexWorkers(), config.ReindexBatchSize())
}

func readAndParse(ref bobstore.Ref) (interface{}, error) {
	b, err := blobs.DB.Read(ref)
	if err != nil {
		return nil, err
	}

	km, err := zkb.Parse(b, ref)
	if err != nil {
		return nil, errors.Wrapf(err, "zkb.Parse %s", ref)
	}
	// if config.Verbose() {
	// 	log.Printf("parsed killmail: %#v", km)
	// }

	return km, nil
}

func indexKillmails(kmchan chan *parallel.Output) []error {
	groupSize := config.ReindexBatchSize()
	errs := make([]error, 0, groupSize)
	killmails := make([]*zkb.Killmail, 0, groupSize)
	for out := range kmchan {
		if out.Err != nil {
			errs = append(errs, out.Err)
		} else {
			km := out.Value.(*zkb.Killmail)
			// log.Printf("indexKillmails kmr: %v %#v", kmr.Ref, kmr.Killmail)
			killmails = append(killmails, km)
		}

		if len(killmails) >= groupSize {
			err := db.IndexKillmails(killmails)
			if err != nil {
				errs = append(errs, err)
			}
			killmails = killmails[0:0]
		}
	}

	if len(killmails) > 0 {
		err := db.IndexKillmails(killmails)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}
