# observer

Observer Pattern written in Go using channels.


## API

### Observer

```
type Observer interface {
	Listen()
}
```

func NewDefaultObserver() *DefaultObserver

	Returns a DefaultObserver struct


func (this *DefaultObserver) Listen()

	Listens for events (run in goroutine)


### Observable

```
type Observable interface {
	Attach(Observer)
	Detach(Observer)
	Notify()
}
```


func NewDefaultObservable() *DefaultObservable

	Returns a DefaultObservable


func (this *DefaultObservable) Attach(observer *DefaultObserver)

	Attaches an observer to observable


func (this *DefaultObservable) Detach(observer *DefaultObserver)

	Detaches an observer to observable


func (this *DefaultObservable) Notify()

	Notifies all subscribed observers



## Working Example

```
package main

import (
	"fmt"
	"github.com/collinglass/patterns/observer"
	"time"
)

func main() {
	obs := observer.NewDefaultObserver()
	observable := observer.NewDefaultObservable()

	observable.Attach(obs)
	go observable.Notify()
	go obs.Listen()
	time.Sleep(1 * 1e9)
}
```