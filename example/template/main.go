//go:build js && wasm
// +build js,wasm

package main

import (
	"bytes"
	"syscall/js"
	"text/template"

	"github.com/fiuskylab/goose/component"
)

func main() {
	doc := js.Global().Get("document")

	f := doc.Call("getElementById", "index")
	app := App{
		Counter: 100,
		father:  f,
	}

	component.Build(&app)

}

// App is the example component
type App struct {
	Counter  int
	Template string
	element  js.Value
	buf      *bytes.Buffer
	father   js.Value
}

// AddOne incremente the Counter
func (a *App) AddOne() {
	a.Counter++
}

// HTML return the HTML  template for the
func (a *App) setHTML() {
	a.Template = `<>
		<button onClick={}>Add Number</button>
		<div>Counter: {{.Counter}}</div>
	</>
	`
}

// Build d
func (a *App) Build() error {
	a.setHTML()
	t := template.Must(template.New("app").Parse(a.Template))
	a.buf = new(bytes.Buffer)
	if err := t.Execute(a.buf, a); err != nil {
		return err
	}
	return nil
}

// GetHTML returns the HTML string
// of the current Component
func (a *App) GetHTML() string {
	return a.buf.String()
}

func (a *App) Father() js.Value {
	return a.father
}

func (a *App) SetElement(el js.Value) {
	a.element = el
}
