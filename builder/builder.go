//go:build js && wasm
// +build js,wasm

package builder

import (
	"syscall/js"

	"github.com/fiuskylab/goose/builder/components"
	"github.com/fiuskylab/goose/builder/config"
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

// bootstraps the builder package
func init() {
	g := js.Global()
	Builder = &builder{
		global:  g,
		console: g.Get("console"),
	}
}

func (b *builder) Build(cs ...components.Components) {
	var err error
	body := b.G().Call("getElementById", config.IndexID)
	for _, c := range cs {
		if err = c.
			SetFather(body).
			Build(); err != nil {
			panic(err)
		}
	}
}

func (b *builder) G() js.Value {
	return b.global
}
