package main

import "testing"

func TestCreateListHead(t *testing.T)  {
	var nodes []*Node
	nodes = []*Node{{data:0, next:nil}, {data:2, next:nil}, {data:3, next:nil}}
	l := CreateListHead(nodes)
	if l.next == nil {
		t.FailNow()
	}
}

func TestCreateListTail(t *testing.T) {
	var nodes []*Node
	nodes = []*Node{{data:0, next:nil}, {data:2, next:nil}, {data:3, next:nil}}
	l := CreateListTail(nodes)
	if l.next == nil {
		t.FailNow()
	}
}