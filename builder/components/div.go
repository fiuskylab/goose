package components

import (
	"syscall/js"

	"github.com/google/uuid"
)

// Div represents the <div> tag.
type Div struct {
	id      string
	element js.Value
}

func NewDiv() *Div {
	return &Div{
		id: uuid.NewString(),
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
