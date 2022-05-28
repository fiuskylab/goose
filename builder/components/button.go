package components

import (
	"syscall/js"

	"github.com/google/uuid"
)

// Button represents the <button> tag.
type Button struct {
	id      string
	Text    string
	element js.Value
}

func NewButton() *Button {
	return &Button{
		id: uuid.NewString(),
	}
}

// Build creates the given Build element
func (b *Button) Build() error {
	doc := js.Global().Get("document")
	index := doc.Call("getElementById", "index")
	el := doc.Call("createElement", "button")
	el.Set("innerText", "salve cachorrada")

	index.Call("appendChild", el)

	b.element = el

	return nil
}
