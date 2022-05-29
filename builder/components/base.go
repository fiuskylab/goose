package components

import "syscall/js"

type base struct {
	id       string
	father   js.Value
	children []js.Value
	element  js.Value
	attr     Attributes
}
