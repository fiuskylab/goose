package component

import "fmt"

// State controls the state and
// actions for the Components
type State[T funcs] struct {
	// Modified is a channel that
	// receives struct{}{} every time
	// an action occurs.
	Modified chan struct{}
	// Stop is a channel that
	// receives struct{}{} to Stop
	// the current process.
	Stop chan struct{}
	// Func is the func that
	// will be ran.
	Func T
}

// funcs is the Type Constraint to
// receive a limited type of functios.
type funcs interface {
	func() error |
		func(params ...any) error
}

// Run runs the process for the
// Component's state
func (s *State[T]) Run(params ...any) {
	for {
		select {
		case <-s.Modified:
			switch f := any(s.Func).(type) {
			case func() error:
				if err := f(); err != nil {
					fmt.Println(err.Error())
				}
			case func(params ...any) error:
				if err := f(params...); err != nil {
					fmt.Println(err.Error())
				}
			}
		case <-s.Stop:
			return
		}
	}
}
