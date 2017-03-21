package linkedList

type List struct {
	root   *Element
	end    *Element
	Length int
}

type Element struct {
	Value interface{}
	Next  *Element
}

func (l *List) Append(e *Element) {
	if l.Length == 0 {
		l.root = e
		l.end = e
	} else {
		l.end.Next = e
		l.end = e
	}
	l.Length++
}

func (l *List) Get(index int) *Element {
	e := l.root
	if index == 0 {
		return e
	}

	for i := 0; i < index; i++ {
		if i != index {
			e = e.Next
		}
	}
	return e
}
