package main

import "fmt"

// 链表节点结构体
type LinkNode struct {
	data interface{}
	next *LinkNode
}

// 链表结构体
type Link struct {
	head *LinkNode
	tail *LinkNode
}

// 从链表头部插入新的元素
func (p *Link) InsertHead(data interface{}) {
	// 创建一个新的节点
	node := &LinkNode{
		data: data,   // 设置新节点的数据
		next: p.head, // 设置新节点的 next 为当前节点的 head
	}
	if p.tail == nil && p.head == nil {
		p.tail = node
		p.head = node
		return
	}
	// node.next = p.head
	p.head = node // 设置新节点为整个链表的头节点
}

// 从尾部插入新节点
func (p *Link) InsertTail(data interface{}) {
	// 创建一个新的节点
	node := &LinkNode{
		data: data,
		next: nil,
	}

	if p.tail == nil && p.head == nil {
		p.tail = node
		p.head = node
		return
	}
	// 将现在尾部的 next 设置为新的节点
	p.tail.next = node
	// 设置新的节点为整个链表的尾节点
	p.tail = node
}

func (p *Link) Trans() {
	node := p.head
	for node != nil {
		fmt.Println(node.data)
		node = node.next
	}
}

// 链表
func main() {
	var link Link
	link.InsertHead(10)
	link.InsertHead(11)
	link.InsertHead(12)
	link.InsertHead(13)
	link.InsertHead(14)
	link.InsertHead(15)

	link.InsertTail(9)
	link.InsertTail(8)
	link.InsertTail(7)

	link.Trans()

}
