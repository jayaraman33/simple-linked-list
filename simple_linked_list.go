package linkedlist

import "errors"

type Element struct {
	data int
	next *Element
}

type List struct {
	head *Element
	size int
}

func New(arr []int) *List {
	size := len(arr)
	if size == 0 {
		return &List{}
	}

	head := Element{data: arr[0]}
	curr := &head
	for i := 1; i < size; i++ {
		curr.next = &Element{data: arr[i]}
		curr = curr.next
	}
	return &List{&head, size}
}

func (l List) Size() int {
	return l.size
}

func (l *List) Push(data int) {
	l.size++
	new := Element{data: data}
	if l.head == nil {
		l.head = &new
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &new
}

func (l *List) Pop() (pop int, err error) {
	if l.size == 0 {
		err = errors.New("list underflow")
		return
	}

	l.size--
	if l.size == 0 {
		pop = l.head.data
		l.head = nil
		return
	}

	curr := l.head
	for i := 0; i < l.size-1; i++ {
		curr = curr.next
	}
	pop = curr.next.data
	curr.next = nil
	return
}

func (l List) Array() []int {
	var (
		arr  = make([]int, 0, l.size)
		curr = l.head
	)
	for curr != nil {
		arr = append(arr, curr.data)
		curr = curr.next
	}
	return arr
}
func (l List) Reverse() *List {
	arr := l.Array()
	for i, j := 0, l.size-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return New(arr)
}
