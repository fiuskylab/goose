package components

// Components contains all base functions
// for a Component.
type Components interface {
	Build() error
	SetFather(Components) Components
}
