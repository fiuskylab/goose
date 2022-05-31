//go:build js && wasm
// +build js,wasm

package builder

import (
	"syscall/js"

	"github.com/fiuskylab/goose/builder/components"
)

type (
	// builder handlers the Goose bootstrap
	builder struct {
		global  js.Value
		console js.Value
		index   components.Component
	}
)

// Builder is an instance of builder, it will
// be responsible for handling all actions
// in DOM.
var Builder *builder

// bootstraps the builder package
func init() {
	g := js.Global()
	index := components.GetIndex()
	Builder = &builder{
		global:  g,
		console: g.Get("console"),
		index:   index,
	}
}

func (b *builder) GetComponent() components.Component {
	return b.index
}

func (b *builder) Build(cs ...components.Component) {
	for _, c := range cs {
		c.
			SetFather(b.index).
			Build()
	}
}

func (b *builder) G() js.Value {
	return b.global
}
