package component

type state struct {
	_        struct{}
	modified chan struct{}
}

func (s *state) run(f func()) {
	for range s.modified {
		f()
	}
}
