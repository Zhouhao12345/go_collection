package tests

import (
	"fmt"
	"testing"
	"zhouhao.com/collections/list"
)

type valueItem int

func (v valueItem) Great(value list.Value) bool {
	return !value.Great(v)
}

func TestDoubleLinkedList(t *testing.T) {
		l := list.NewList()
		l.Insert(valueItem(1))
		l.Insert(valueItem(2))
		l.Insert(valueItem(3))
		l.Insert(valueItem(4))
		l.Insert(valueItem(5))
		l.InsertByIndex(valueItem(5), 4)
		fmt.Print(l)
	}

