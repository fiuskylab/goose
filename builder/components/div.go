package components

import (
	"syscall/js"

	"github.com/google/uuid"
)

// Div represents the <div> tag.
type Div struct {
	base
}

// NewDiv returns an instance of Div
// with given attr
func NewDiv(attr Attributes) *Div {
	return &Div{
		base: base{
			id:   uuid.NewString(),
			attr: attr,
		},
	}
}

// Build creates the given Div element
func (d *Div) Build() error {
	doc := js.Global().Get("document")
	index := doc.Call("getElementById", "index")
	el := doc.Call("createElement", "div")
	el.Set("innerText", "Lorem ipsum")

	index.Call("appendChild", el)

	d.element = el

	return nil
}

// SetFather receives a js.Value and assign
// it to `father` field.
func (d *Div) SetFather(father Components) *Div {
	d.father = father.GetFather()
	return d
}

// GetFather returns current component's father.
func (d *Div) GetFather() js.Value {
	return d.father
}
