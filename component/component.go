//go:build js && wasm
// +build js,wasm

package component

import (
	"syscall/js"
	"time"
)

// Component is the
type Component interface {
	Build() error
	GetHTML() string
	Father() js.Value
	SetElement(js.Value)
	GetState() State
}

// Base struct have the necessary
// fields to handle a Component
type Base struct {
	State State
}

// Ack will send a signal(chan struct{})
// to say that the component must be updated.
// e.g. A button was pressed
func (b *Base) Ack() {
	b.
		State.
		modified <- struct{}{}
}

// Build creates the given component
func Build(c Component) {
	state := c.GetState()
	state.modified = make(chan struct{}, 10)

	f := func() error {
		doc := js.Global().Get("document")
		el := doc.Call("createElement", "div")

		if err := c.Build(); err != nil {
			panic(err)
		}

		el.Set("innerHTML", c.GetHTML())

		c.Father().Call("appendChild", el)

		return nil
	}

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		if err := f(); err != nil {
			panic(err)
		}
	}

	state.run(f)
}
