package main

import (
	"fmt"
	"github.com/BrentBrown/GoSandbox/challenges/linkedList"
)

func main() {
	l := new(linkedList.List)
	e := new(linkedList.Element)
	e.Value = 42

	l.Append(e)

	f := new(linkedList.Element)
	f.Value = 27
	l.Append(f)

	g := new(linkedList.Element)
	g.Value = 1
	l.Append(g)

	m := l.Get(2)
	fmt.Printf("%v", m.Value)
}
