package observer

import (
	"log"
)

// Observer interface
//
// implements ReceiveEvents which
// listens to events emitted by observables
type Observer interface {
	ReceiveEvents()
}

// Observable interface
//
// implements Attach(Observer) which
// attaches new observers to the observable
//
// implements Detach(Observer) which
// detaches observers from the observable
//
// implements Notify() which
// sends events to observers
type Observable interface {
	Attach(Observer)
	Detach(Observer)
	Notify()
}

// Default Observer struct
//
// contains go channel which receives events
type DefaultObserver struct {
	Events chan string
}

// New Default Observer
//
// returns a new default oberserver struct
func NewDefaultObserver() *DefaultObserver {
	return &DefaultObserver{
		Events: make(chan string),
	}
}

// Default Observer implements Receive Events
//
// receives events from observables it's
// subscribed to
func (this *DefaultObserver) ReceiveEvents() {
	s := <-this.Events
	log.Printf("Received Event: %s\n", s)
}

// Default Observable
//
// contains current state to send to observer
//
// contains list of observers to notify to
type DefaultObservable struct {
	State     string
	Observers map[chan string]bool
}

// New Default Oberservable
//
// returns a new Default Observable struct
func NewDefaultObservable() *DefaultObservable {
	return &DefaultObservable{
		State:     "Initial State",
		Observers: make(map[chan string]bool),
	}
}

// Default Observable implements Attach
//
// attaches a new observer to notify events
func (this *DefaultObservable) Attach(observer *DefaultObserver) {
	this.Observers[observer.Events] = true
	log.Println("Observable Attached.")
}

// Default Observable implements Dettach
//
// detaches observer from list of observers
func (this *DefaultObservable) Detach(observer *DefaultObserver) {
	delete(this.Observers, observer.Events)
	log.Println("Observable Detached.")
}

// Default Observable implements Notify
//
// notifies observers of current state
func (this *DefaultObservable) Notify() {
	for observer, _ := range this.Observers {
		observer <- this.State
	}
	log.Printf("Broadcast message %s to %d observers.\n", this.State, len(this.Observers))
}
