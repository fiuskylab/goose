package components

import (
	"syscall/js"

	"github.com/fiuskylab/goose/builder/config"
	"github.com/google/uuid"
)

type base struct {
	id       string
	doc      js.Value
	father   js.Value
	children []Component
	element  js.Value
	attr     Attributes
}

type Base struct {
	Tag      string
	id       string
	doc      js.Value
	father   js.Value
	children []Component
	element  js.Value
	Attr     Attributes
}

func NewBase(tag string, attr Attributes, children ...Component) *Base {
	return &Base{
		Tag:      tag,
		id:       uuid.NewString(),
		children: children,
		Attr:     attr,
		doc:      js.Global().Get("document"),
	}
}

func (b *Base) Build() Component {
	if b.id == config.IndexID {
		return b
	}

	el := b.doc.Call("createElement", b.Tag)

	b.element = el

	b.setAttributes()

	b.father.Call("appendChild", b.element)

	return b
}

func (b *Base) setAttributes() {
	for attr, value := range b.Attr {
		//b.element.Call("setAttributes", map[string]string{attr: value})
		b.element.Set(attr, value)
	}
}

func (b *Base) GetElement() js.Value {
	return b.element
}

func (b *Base) SetFather(father Component) Component {
	b.father = father.GetElement()
	return b
}

// GetFather returns current component's father.
func (b *Base) GetFather() js.Value {
	return b.father
}

// AddChild adds a new Component as a children
// to current Component.
func (b *Base) AddChild(child Component) Component {
	b.children = append(b.children, child)
	return b
}
