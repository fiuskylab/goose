package components

import "syscall/js"

// Component contains all base functions
// for a Component.
type Component interface {
	// Build builds the components
	Build() Component
	// GetElement returns current Component's element.
	GetElement() js.Value
	// SetFather sets the father of
	// current Component
	SetFather(Component) Component
	// GetFather gets the father of
	// current Component
	GetFather() js.Value
	// AddChildren adds a new Component as a chieldren
	// to current Component.
	AddChild(Component) Component
}
