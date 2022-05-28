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
	}
)

// Builder is an instance of builder, it will
// be responsible for handling all actions
// in DOM.
var Builder *builder

// New return an Empty page
func New() *builder {
	g := js.Global()
	return &builder{
		global:  g,
		console: g.Get("console"),
	}
}

func (b *builder) Build(cs ...components.Components) {
	var err error
	for _, c := range cs {
		if err = c.Build(); err != nil {
			panic(err)
		}
	}
}

func (b *builder) G() js.Value {
	return b.global
}
