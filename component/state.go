package component

import "fmt"

type State struct {
	_        struct{}
	modified chan struct{}
}

func (s *State) run(f func() error) {
	i := 1
	for range s.modified {
		fmt.Printf("Ran %d times", i)
		i++
		if err := f(); err != nil {
			fmt.Println(err.Error())
		}
	}
}
