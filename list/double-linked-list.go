package list

import "fmt"

type Value interface {
	Great(v Value) bool
}

type ListNode struct {
	val Value
	lastNode *ListNode
	nextNode *ListNode
}

func NewListNode(v Value) *ListNode {
	return &ListNode{
		val:      v,
		lastNode: nil,
		nextNode: nil,
	}
}

type List struct {
	top *ListNode
	end *ListNode
}

func (l *List) FindByIndex(i int) *ListNode {
	curNode := l.top
	curIndex := 0
	for ; curIndex != i
	{
		curNode = curNode.nextNode
		curIndex ++
	}
	return curNode
}

func (l *List) Insert(v Value) {
	newNode := NewListNode(v)
	if l.top == nil {
		l.top = newNode
		l.end = newNode
	} else {
		l.end.nextNode = newNode
		newNode.lastNode = l.end
		l.end = newNode
	}
}

func (l *List) InsertByIndex(v Value, i int) {
	newNode := NewListNode(v)
	if l.top == nil {
		l.top = newNode
		l.end = newNode
	} else {
		curNode := l.FindByIndex(i)
		if curNode == l.end {
			l.Insert(v)
		} else {
			newNode.nextNode = curNode.nextNode
			newNode.lastNode = curNode
			curNode.nextNode.lastNode = newNode
			curNode.nextNode = newNode
		}
	}
}

func (l *List) String() string {
	curNode := l.top
	values := make([]Value, 0)
	for ;curNode!= nil
	{
		values = append(values, curNode.val)
		curNode = curNode.nextNode
	}
	return fmt.Sprint(values)
}

func NewList() *List {
	return &List{
		top: nil,
		end: nil,
	}
}



