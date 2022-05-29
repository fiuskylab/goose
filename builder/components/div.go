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
			doc:  js.Global().Get("document"),
		},
	}
}

// GetIndex returns a Component instance
// of the index <div>.
func GetIndex() Div {
	doc := js.Global().Get("document")
	el := doc.Call("getElementById", config.IndexID)
	return Div{
		base: base{
			doc:     doc,
			element: el,
		},
	}
}

// Build creates the given Div element
func (d Div) Build() error {
	el := d.doc.Call("createElement", "div")
	el.Set("innerText", "Lorem ipsum")

	d.GetFather().Call("appendChild", el)

	d.element = el

	return nil
}

// GetElement returns current Component's element.
func (d Div) GetElement() js.Value {
	return d.element
}

// SetFather receives a js.Value and assign
// it to `father` field.
func (d Div) SetFather(father Component) Component {
	d.father = father.GetElement()
	return d
}

// GetFather returns current component's father.
func (d Div) GetFather() js.Value {
	return d.father
}

// AddChild adds a new Component as a children
// to current Component.
func (d Div) AddChild(child Component) Component {
	d.children = append(d.children, child)
	return d
}
