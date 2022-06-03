package components

// Button represents the <button> tag.
type Button struct {
	base
	Component
}

// NewButton returns an instance of Div
// with given attr
func NewButton(attr Attributes, children ...Component) Component {
	return NewBase("button", attr, children...)
}
