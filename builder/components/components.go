package components

import "syscall/js"

// Components contains all base functions
// for a Component.
type Components interface {
	// Build builds the components
	Build() error
	// SetFather sets the father of
	// current Component
	SetFather(Components) Components
	// GetFather gets the father of
	// current Component
	GetFather() js.Value
}
