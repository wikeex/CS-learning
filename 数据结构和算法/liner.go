package main

import "fmt"

const MAXSIZE = 20

// 顺序表的插入和删除
type SqList struct {
	data [MAXSIZE]interface{}
	length int
}

func (s *SqList) ListInsert (i int, e interface{}) error {
	var k int
	if s.length >= MAXSIZE {
		return fmt.Errorf("顺序表长度过长！")
	}
	if i > s.length || i < 0 {
		return fmt.Errorf("插入位置错误")
	}
	if i < s.length {
		for k = s.length - 1; k > i; k-- {
			s.data[k+1] = s.data[k]
		}
	}
	s.data[i] = e
	s.length++
	return nil
}

func (s *SqList) ListDelete (i int) error {
	if s.length >= MAXSIZE {
		return fmt.Errorf("顺序表长度超出范围！")
	}
	if i < 0 || i > s.length {
		return fmt.Errorf("删除位置错误")
	}
	if i < s.length {
		for k := i; k < s.length; k++ {
			s.data[k] = s.data[k+1]
		}
		s.length--
	}
	return nil
}

// 单链表查询、插入和删除
type Node struct {
	data interface{}
	next *Node
}

type LinkList = Node

// 单链表查询，i从1开始数
func (l *LinkList) GetElem(i int) (interface{}, error) {
	p := l.next
	j := 1
	for j < i {
		p = p.next
		if p == nil {
			break
		}
		j++
	}
	if p == nil || j > i {
		return nil, fmt.Errorf("第i个节点不存在！")
	}
	return p.data, nil
}

// 单链表插入
func (l *LinkList) ListInsert(i int, e interface{}) error {
	p := l.next
	s := Node{}
	j := 1
	for j < i {
		p = p.next
		if p == nil {
			break
		}
		j++
	}
	if p == nil || j > i {
		return fmt.Errorf("第i个节点不存在")
	}
	s.data = e
	s.next = p.next
	p.next = &s
	return nil
}

// 单链表的删除
func (l *LinkList) ListDelete(i int) error {
	p := l.next
	j := 1
	for j < i {
		p = p.next
		if p == nil {
			break
		}
		j++
	}
	if p == nil && j > i {
		return fmt.Errorf("第i个节点不存在！")
	}
	q := p.next
	p.next = q.next

	return nil
}

// 单链表创建头插法
func CreateListHead(nodes []*Node) *LinkList {
	l := LinkList{next: nil}
	for _, node := range nodes {
		 node.next = l.next
		 l.next = node
	}
	return &l
}

// 单链表创建尾插法
func CreateListTail(nodes []*Node) *LinkList {
	l := LinkList{next: nil}
	tail := &l
	for _, node := range nodes {
		tail.next = node
		tail = node
	}
	tail.next = nil
	return &l
}

func main()  {
	var nodes []*Node
	nodes = []*Node{{data:0, next:nil}, {data:2, next:nil}, {data:3, next:nil}}
	l := CreateListHead(nodes)
	fmt.Print(l.data, l.next.data, l.next.next.data, l.next.next.next.data)
}