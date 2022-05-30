package components

import "fmt"

type Node struct {
	Element  Component
	Father   Component
	Children []Node
}

func (n Node) Build() error {
	n.Element = n.Element.Build()

	for _, child := range n.Children {
		child.Father = n.Element
		child.Element = child.Element.SetFather(n.Element)
		fmt.Println("Element", n.Element)
		fmt.Println("Child.GetFather", child.Element.GetFather())
		if err := child.Build(); err != nil {
			panic(err)
		}
	}

	return nil
}
