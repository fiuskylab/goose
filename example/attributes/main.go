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
				Element: components.NewBase("a", components.Attributes{
					"href": "https://twitch.tv/rafiusky",
				}),
				Children: []components.Node{
					{
						Element: components.NewDiv(components.Attributes{
							"style": "height: 100px; width: 100px; background-color: blue;",
						}),
					},
					{
						Element: components.NewButton(components.Attributes{
							"style": "height: 100px; width: 100px; background-color: blue;",
						}),
					},
				},
			},
			{
				Element: components.NewBase("a", components.Attributes{
					"href": "https://twitch.tv/rafiusky",
				}),
			},
		},
	}

	if err := n.Build(); err != nil {
		panic(err)
	}
}
