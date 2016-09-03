// Package parallel implements parallel processing of bobstore refs
package parallel

import (
	"sync"

	"github.com/random-j-farmer/bobstore"
)

// Processor turns bobstore.Refs into some output
type Processor struct {
	numWorkers int
	fn         MappingFunction
	Input      chan bobstore.Ref
	output     chan *Output
	handler    OutputHandler
	wg         sync.WaitGroup
	hwg        sync.WaitGroup
	herr       []error
}

// MappingFunction maps a ref into something else (or an error)
type MappingFunction func(bobstore.Ref) (interface{}, error)

// OutputHandler handles the output
type OutputHandler func(chan *Output) []error

// Output of the mapping function
type Output struct {
	Value interface{}
	Err   error
}

// NewProcessor creates a new unordered processor
// the input queue will be processed in parallel and
// handler will be called on the result in one goroutine.
// ordering may not be preserved.
func NewProcessor(numWorkers int, fn MappingFunction, handler OutputHandler) *Processor {
	proc := &Processor{
		numWorkers: numWorkers,
		fn:         fn,
		handler:    handler,
		Input:      make(chan bobstore.Ref),
		output:     make(chan *Output, numWorkers),
	}

	proc.wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go runWorker(proc)
	}

	proc.hwg.Add(1)
	go runHandler(proc)

	return proc
}

func runWorker(proc *Processor) {
	for ref := range proc.Input {
		value, err := proc.fn(ref)
		proc.output <- &Output{Value: value, Err: err}
	}
	proc.wg.Done()
}

func runHandler(proc *Processor) {
	proc.herr = proc.handler(proc.output)
	proc.hwg.Done()
}

// Wait waits for completion of the workers and handler.
// It returns the handlers result.
func (p *Processor) Wait() []error {
	close(p.Input)
	p.wg.Wait()
	close(p.output)
	p.hwg.Wait()
	return p.herr
}
