//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/fiuskylab/goose"
	"github.com/fiuskylab/goose/builder/components"
)

func main() {
	goose.Start(
		components.NewDiv(components.Attributes{
			"innerText": "Foo bar",
			"style":     "height: 100px; width: 200px; background-color: blue;",
		}, components.NewButton(components.Attributes{
			"innerText": "Button",
		}),
		),
	)
}
