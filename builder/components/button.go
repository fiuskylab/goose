package components

import (
	"syscall/js"

	"github.com/google/uuid"
)

// Button represents the <button> tag.
type Button struct {
	base
	Component
}

// NewButton returns an instance of Div
// with given attr
func NewButton(attr Attributes, children ...Component) Button {
	return Button{
		base: base{
			id:       uuid.NewString(),
			attr:     attr,
			doc:      js.Global().Get("document"),
			children: children,
		},
	}
}

// Build creates the given Build element
func (b Button) Build() Component {
	return NewBase("button", b.attr, b.children...).Build()
}
