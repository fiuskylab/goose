package component

import "fmt"

type State struct {
	_        struct{}
	Modified chan struct{}
	Stop     chan struct{}
}

func (s *State) Run(f func() error) {
	for {
		select {
		case <-s.Modified:
			if err := f(); err != nil {
				fmt.Println(err.Error())
			}
		case <-s.Stop:
			return
		}
	}
}
