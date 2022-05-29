package components

import "syscall/js"

// Components
type Components interface {
	Build() error
	SetFather(js.Value) *Components
}
