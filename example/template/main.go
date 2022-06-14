//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/fiuskylab/goose/component"
)

// App is our app.
type App struct {
	component.Base
	Text    string
	Counter *int
}

// HTML will set current Component HTML
func (a *App) HTML() string {
	return `
		{{.Text}} something {{.Counter}}
	`
}

func main() {
	i := 20
	app := &App{
		Text:    "Sumemo Doido",
		Counter: &i,
	}

	component.Build(app)
}
