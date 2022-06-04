//go:build js && wasm
// +build js,wasm

package goose

import (
	"github.com/fiuskylab/goose/builder"
	"github.com/fiuskylab/goose/builder/components"
)

// Start receives components to be built
func Start(cs ...components.Component) {
	builder.Builder.Build(cs...)
}
