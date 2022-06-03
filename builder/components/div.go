package components

import (
	"syscall/js"

	"github.com/fiuskylab/goose/builder/config"
)

// Div represents the <div> tag.
type Div struct {
	base
	Component
}

// NewDiv returns an instance of Div
// with given attr
func NewDiv(attr Attributes, children ...Component) Component {
	return NewBase("div", attr, children...)
}

// Build creates the given Div element
func (d Div) Build() Component {
	b := Base{
		id:       d.id,
		doc:      d.doc,
		father:   d.father,
		children: d.children,
		element:  d.element,
		Tag:      "div",
		Attr:     d.attr,
	}
	return b.Build()
}

// GetIndex returns a Component instance
// of the index <div>.
func GetIndex() Div {
	doc := js.Global().Get("document")
	el := doc.Call("getElementById", config.IndexID)
	return Div{
		base: base{
			doc:     doc,
			element: el,
			father:  el,
			id:      config.IndexID,
		},
	}
}
