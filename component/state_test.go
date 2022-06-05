package component

import (
	"fmt"
	"testing"
	"time"
)

func TestState(t *testing.T) {
	s := State{
		Modified: make(chan struct{}, 10),
		Stop:     make(chan struct{}, 10),
	}

	f := func() error {
		fmt.Println("Test")
		return nil
	}

	go func() {
		for i := 0; i < 10; i++ {
			s.Modified <- struct{}{}
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		time.Sleep(5 * time.Second)
		s.Stop <- struct{}{}
	}()

	s.Run(f)
}
