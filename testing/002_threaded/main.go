package main

import (
	"fmt"
	"time"
)

var gossip []Tx = []Tx{}
var rejects []Tx = []Tx{}

type Node struct {
	Id    int
	Books map[string]int
}

type Tx struct {
	Id     int
	From   string
	To     string
	Amount int
}

func NewNode(id int) Node {
	node := Node{}
	node.Id = id
	node.Books = map[string]int{}
	return node
}

func (n Node) ListenToGossip() {
	for {
		fmt.Println(n.Id, "ListenToGossip")
		time.Sleep(time.Second)
	}
}

func main() {
	node1 := NewNode(1)
	node1.Books["ABC"] = 100
	node2 := NewNode(2)
	node2.Books["ABC"] = 100
	node3 := NewNode(3)
	node3.Books["ABC"] = 100

	fmt.Println(node1, node2, node3)

	go node1.ListenToGossip()

	for {
		time.Sleep(time.Second)
	}

}
