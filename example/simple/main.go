//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/fiuskylab/goose/builder"
	"github.com/fiuskylab/goose/builder/components"
)

func main() {
	b := builder.Builder

	index := b.GetComponent()
	n := components.Node{
		Element: index,
		Father:  index,
		Children: []components.Node{
			{
				//Father: components.Component{},
				Element: components.NewDiv(components.Attributes{}),
				Children: []components.Node{
					{
						Element: components.NewButton(components.Attributes{}),
					},
					{
						Element: components.NewButton(components.Attributes{}),
					},
				},
			},
			{
				Element: components.NewButton(components.Attributes{}),
			},
			{
				Element: components.NewDiv(components.Attributes{}),
				Children: []components.Node{
					{
						Element: components.NewButton(components.Attributes{}),
					},
					{
						Element: components.NewDiv(components.Attributes{}),
						Children: []components.Node{
							{
								Element: components.NewButton(components.Attributes{}),
							},
							{
								Element: components.NewButton(components.Attributes{}),
							},
						},
					},
				},
			},
		},
	}

	if err := n.Build(); err != nil {
		panic(err)
	}
}
