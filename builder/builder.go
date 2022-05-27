//go:build js && wasm
// +build js,wasm

package builder

import (
	"syscall/js"

	"github.com/fiuskylab/goose/builder/components"
)

type (
	// Builder handlers the Goose bootstrap
	Builder struct {
		global  js.Value
		console js.Value
	}
)

// New return an Empty page
func New() *Builder {
	g := js.Global()
	return &Builder{
		global:  g,
		console: g.Get("console"),
	}
}

func (b *Builder) Build(cs ...components.Components) {
	var err error
	for _, c := range cs {
		if err = c.Build(); err != nil {
			panic(err)
		}
	}
}

func (b *Builder) G() js.Value {
	return b.global
}
