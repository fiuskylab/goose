package components

import (
	"syscall/js"

	"github.com/fiuskylab/goose/builder/config"
)

// Div represents the <div> tag.
type Div struct {
	Base
	Component
}

// NewDiv returns an instance of Div
// with given attr
func NewDiv(attr Attributes, children ...Component) Component {
	return NewBase("div", attr, children...)
}

// GetIndex returns a Component instance
// of the index <div>.
func GetIndex() Component {
	doc := js.Global().Get("document")
	el := doc.Call("getElementById", config.IndexID)
	return &Base{
		doc:     doc,
		element: el,
		father:  el,
		id:      config.IndexID,
		Tag:     "div",
	}
}
