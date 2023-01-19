package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	front *ListItem
	back  *ListItem
	len   int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	if l.len == 0 {
		l.front = &ListItem{Value: v}
		l.back = l.front
	} else {
		l.front.Prev = &ListItem{Value: v, Next: l.front}
		l.front = l.front.Prev
	}
	l.len++
	return l.Front()
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.len == 0 {
		l.PushFront(v)
	} else {
		l.back.Next = &ListItem{Value: v, Prev: l.back}
		l.back = l.back.Next
		l.len++
	}
	return l.Back()
}

func (l *list) Remove(i *ListItem) {
	switch {
	case l.len == 1:
		l.front = nil
		l.back = nil
	case i == l.front:
		i.Next.Prev = i.Prev
		l.front = i.Next
	case i == l.back:
		i.Prev.Next = i.Next
		l.back = i.Prev
	default:
		i.Next.Prev = i.Prev
		i.Prev.Next = i.Next
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return new(list)
}
