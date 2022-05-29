package components

import "syscall/js"

type base struct {
	id       string
	doc      js.Value
	father   js.Value
	children []Component
	element  js.Value
	attr     Attributes
}
