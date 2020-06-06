package crawl

const (
	WAITING uint8 = 0
	STOPPED uint8 = 1
	RUNNING uint8 = 2
)

// Worker is an interface that is used to manage
// agents that can work in different threads
type Worker interface {
	Run() error // starts the worker

	//Returns worker names
	//Example - Fetcher, Parser, Tracker
	Type() string

	// State returns the state that workers are in
	State() uint8
	SetState(state uint8) 
}

// Implementation of the worker interface
// AsWorker is the implementation, embedded in a new struct
type AsWorker struct {
	RunFunc func() error
	state uint8
	Quit chan uint8
	Name string
}

// NewAsWorker is a constructor for AsWorker
func NewAsWorker(name string) *AsWorker {
	quit := make(chan uint8)
	return &AsWorker{
		Name: name,
		Quit: quit,
	}
}

// Run the AsFetcher
func (w *AsWorker) Run() error {
	return w.RunFunc()
}

// Notify the channel to quit
func (w *AsWorker) Stop() {
	w.Quit <- 0
}


func (w *AsyncWorker) State() uint8 {
	return w.state
}

func (w *AsyncWorker) SetState(state uint8) {
	w.state = state
}

// Type returns the Name given to the Worker
// in initialisation
func (w *AsyncWorker) Type() string {
	return w.Name
}