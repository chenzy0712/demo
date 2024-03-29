package study

import (
	"sync/atomic"

	"git.kldmp.com/learning/demo/pkg/log"
)

type PayloadCollection struct {
	WindowsVersion string    `json:"version"`
	Token          string    `json:"token"`
	Payloads       []Payload `json:"data"`
}

type Payload struct {
	// [redacted]
	firstByte byte
}

func (p *Payload) recordSinglePieceOfWave() error {
	//DO WHAT YOU NEED
	log.Debug("%02x", p.firstByte)
	return nil
}

//payloadHandler handle payload
func payloadHandler(data []byte) {
	var payload Payload

	payload.firstByte = data[0]
	log.Debug("%02x", payload.firstByte)
	// let's create a job with the payload
	work := Job{Payload: payload}

	atomic.AddUint64(&pps, 1)

	// Push the work onto the queue.
	JobQueue <- work
	log.Debug("payloadHandler done")
}

// Job represents the job to be run
type Job struct {
	Payload Payload
}

// A buffered channel that we can send work requests on.
var JobQueue chan Job

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		log.Debug("Run worker")
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				// we have received a work request.
				if err := job.Payload.recordSinglePieceOfWave(); err != nil {
					log.Error("Error record single piece of wave: %s", err.Error())
				}

			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan Job

	maxWorkers int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, maxWorkers: maxWorkers}
}

func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 0; i < d.maxWorkers; i++ {
		log.Debug("Init %d worker", i)
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	log.Info("Start dispatch")
	for {
		select {
		case job := <-JobQueue:
			// a job request has been received
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}

//RoutineDemo routine demo for handling 1 million request per minute
func RoutineDemo(workers int) {
	JobQueue = make(chan Job)
	dispatcher := NewDispatcher(workers)
	dispatcher.Run()
}
