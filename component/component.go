package component

import (
	"bytes"
	"html/template"
	"syscall/js"

	"github.com/fiuskylab/goose/builder/config"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type (
	// Component ...
	Component interface {
		HTML() string
	}

	// Base is the base struct for building
	// a component.
	Base struct {
		el   js.Value
		id   string
		html string
	}
)

var (
	g     js.Value
	doc   js.Value
	index js.Value
	log   *zap.Logger
)

func init() {
	g = js.Global()
	doc = g.Get("document")
	index = doc.Call("getElementById", config.IndexID)
	log, _ = zap.NewProduction()
}

type components []Component

// Build ...
func Build(cs ...Component) {
	for _, c := range cs {
		id := uuid.NewString()
		t, err := template.New(id).Parse(c.HTML())
		if err != nil {
			log.Error(err.Error())
			continue
		}

		writer := &bytes.Buffer{}
		if err := t.Execute(writer, c); err != nil {
			log.Error(err.Error())
			continue
		}

		el := doc.Call("createElement", "div")

		el.Set("innerHTML", writer.String())

		index.Call("appendChild", el)
	}
}
