package components

import (
	"fmt"
	"syscall/js"

	"github.com/google/uuid"
)

// Button represents the <button> tag.
type Button struct {
	base
}

// NewButton returns an instance of Div
// with given attr
func NewButton(attr Attributes) Button {
	return Button{
		base: base{
			id:   uuid.NewString(),
			attr: attr,
			doc:  js.Global().Get("document"),
		},
	}
}

// Build creates the given Build element
func (b Button) Build() Component {
	el := b.doc.Call("createElement", "button")
	el.Set("innerText", "Lorem ipsum")
	b.father.Call("appendChild", el)

	b.element = el

	return b
}

// GetElement returns current Component's element.
func (b Button) GetElement() js.Value {
	return b.element
}

// SetFather receives a js.Value and assign
// it to `father` field.
func (b Button) SetFather(father Component) Component {
	b.father = father.GetElement()
	fmt.Println("Button.SetFather Received:", father.GetElement())
	fmt.Println("Button.SetFather: ", b.father)
	return b
}

// GetFather returns current component's father.
func (b Button) GetFather() js.Value {
	return b.father
}

// AddChild adds a new Component as a children
// to current Component.
func (b Button) AddChild(child Component) Component {
	b.children = append(b.children, child)
	return b
}
