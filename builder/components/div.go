package components

import (
	"syscall/js"

	"github.com/fiuskylab/goose/builder/config"
	"github.com/google/uuid"
)

// Div represents the <div> tag.
type Div struct {
	base
}

// NewDiv returns an instance of Div
// with given attr
func NewDiv(attr Attributes, el ...js.Value) Div {
	return Div{
		base: base{
			id:   uuid.NewString(),
			attr: attr,
		},
	}
}

func GetIndex() Div {
	index := js.Global().Call("getElementById", config.IndexID)
	return Div{
		base: base{
			element: index,
		},
	}
}

// Build creates the given Div element
func (d Div) Build() error {
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
func (d Div) SetFather(father Components) Components {
	d.father = father.GetFather()
	return d
}

// GetFather returns current component's father.
func (d Div) GetFather() js.Value {
	return d.father
}
