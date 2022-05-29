package components

import (
	"syscall/js"

	"github.com/fiuskylab/goose/builder/config"
	"github.com/google/uuid"
)

// Button represents the <button> tag.
type Button struct {
	base
}

// NewButton returns an instance of Div
// with given attr
func NewButton(attr Attributes) *Button {
	return &Button{
		base: base{
			id:   uuid.NewString(),
			attr: attr,
		},
	}
}

// Build creates the given Build element
func (b *Button) Build() error {
	doc := js.Global().Get("document")
	index := doc.Call("getElementById", config.IndexID)
	el := doc.Call("createElement", "button")
	el.Set("innerText", "Lorem ipsum")

	index.Call("appendChild", el)

	b.element = el

	return nil
}

// SetFather receives a js.Value and assign
// it to `father` field.
func (b *Button) SetFather(father js.Value) *Button {
	b.father = father
	return b
}

// GetFather returns current component's father.
func (b *Button) GetFather() js.Value {
	return b.father
}
