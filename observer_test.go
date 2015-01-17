package observer

import (
	"testing"
	"time"
)

func setup() (*DefaultObserver, *DefaultObservable) {
	obs := NewDefaultObserver()
	observable := NewDefaultObservable()

	return obs, observable
}

func TestAttachDetach(t *testing.T) {
	obs, observable := setup()

	if len(observable.Observers) > 0 {
		t.Errorf("TestAttachDetach Failed: %s\n", "observers found before attach")
	}

	observable.Attach(obs)

	if len(observable.Observers) <= 0 {
		t.Errorf("TestAttachDetach Failed: %s\n", "no observers found after attach")
	}

	observable.Detach(obs)

	if len(observable.Observers) > 0 {
		t.Errorf("TestAttachDetach Failed: %s\n", "observers found after detach")
	}
}

func TestNotify(t *testing.T) {
	obs, observable := setup()

	if len(observable.Observers) > 0 {
		t.Errorf("TestNotify Failed: %s\n", "observers found before attach")
	}

	observable.Attach(obs)

	if len(observable.Observers) <= 0 {
		t.Errorf("TestNotify Failed: %s\n", "no observers found after attach")
	}

	go observable.Notify()
	s := <-obs.Events

	time.Sleep(1 * 1e9)

	if s == "" {
		t.Errorf("TestNotify Failed: %s\n", "no event received")
	}

	if s != observable.State {
		t.Errorf("TestNotify Failed: %s\n", "expected does not match actual")
	}
}
