//go:build js && wasm
// +build js,wasm

package component

import "syscall/js"

// Component is the
type Component interface {
	Build() error
	GetHTML() string
	Father() js.Value
	SetElement(js.Value)
}

// Build creates the given component
func Build(c Component) {
	doc := js.Global().Get("document")

	el := doc.Call("createElement", "div")

	if err := c.Build(); err != nil {
		panic(err)
	}

	el.Set("innerHTML", c.GetHTML())

	c.Father().Call("appendChild", el)
}
