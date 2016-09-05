// Package zkb pulls killmails from zkillboard
//
// total kills last 7 days: 73000
// averge killmail size (20 minute sample, 22:45): 12495.781553398057
// snappy compressed: 2707.3155339805826
//
// not practical to store the whole killmail
package zkb

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"sync"
	"time"

	"github.com/random-j-farmer/bobstore"
	"github.com/random-j-farmer/zkill-mirror/internal/config"
)

// PullKillmails until stop channel is closed
// this goroutine deletes non-critical work but does the db updates,
// so this goroutine is tracked via the waitgroup, while httpWorker is not
func PullKillmails(bdb *bobstore.DB, kmQueue chan<- *Killmail, stop <-chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	work := makeWorkQueue(stop)

	for {
		select {
		case <-stop:
			log.Print("zkb.PullKillmails: stop channel selects ... closing km indexing queue")
			close(kmQueue)
			return
		case work <- true:
			log.Print("zkb.PullKillmails: work channel selects ...")
			sleepy := pullKillmail(bdb, kmQueue)
			if !sleepy {
				pushBell(work, false)
			}
		}
	}
}

// make this general!
func makeWorkQueue(stop <-chan struct{}) chan bool {
	work := make(chan bool, 1)
	go wakeupWorkers(stop, work)
	return work
}

func pushBell(work <-chan bool, announce bool) {
	select {
	case <-work:
		if announce {
			log.Printf("zkb.pushBell: DING DONG DING DONG DING DONG")
		}
	default:
	}
}

func wakeupWorkers(stop <-chan struct{}, work <-chan bool) {
	for {
		select {
		case <-stop:
			log.Print("zkb.wakeupWorkers: stop channel selects ...")
			return
		default:
			time.Sleep(config.PullDelay())
			log.Printf("zkb.wakeupWorkers: should i wake them?")
			pushBell(work, true)
		}
	}
}

// client is a client with a cookiejar - this accepts the servers cookies
var client = &http.Client{
	Jar: func() *cookiejar.Jar {
		jar, _ := cookiejar.New(nil)
		return jar
	}(),
	Timeout: time.Second * 60,
}

var emptyRegex = regexp.MustCompile(`^\s*{\s*"?package"?:\s*null\s*}\s*`)

func pullKillmail(bdb *bobstore.DB, kmQueue chan<- *Killmail) (sleepy bool) {
	req, err := http.NewRequest("GET", "https://redisq.zkillboard.com/listen.php", nil)
	if err != nil {
		log.Printf("zkb.httpWorker: http.NewReqeuest: %v", err)
		return true
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("zkb.httpWorker: get error: %v", err)
		return true
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("zkb.httpWorker: error reading response body: %v", err)
		return true
	}

	status := resp.StatusCode
	if status/100 != 2 {
		log.Printf("zkb.httpWorker: status %d", status)
		return true
	}

	if emptyRegex.Match(body) {
		log.Printf("zkb.httpWorker: received empty package: %s", body)
		return true
	}

	// we parse before writing, so we don't store garbage
	km, err := Parse(body, bobstore.Ref{})
	if err != nil {
		log.Printf("zkb.httpWorker: error parsing killmail %v", err)
		return true
	}

	ref, err := bdb.WriteWithCodec(body, bobstore.CodecFor(config.BobsCodec()))
	if err != nil {
		log.Printf("zkb.httpWorker: error storing to blob store: %v", err)
		return true
	}

	km.Ref = ref

	kmQueue <- km
	log.Printf("received zkillmail and stored as %s", ref)

	return false
}
