package main

import (
	"fmt"
	"time"
)

var nodes = []Node{}

type Node struct {
	Id     int
	Books  map[string]int
	Gossip chan Tx
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
	node.Books["ABC"] = 100
	node.Gossip = make(chan Tx, 1024)
	go node.ListenToGossip()
	go node.ListenToRejects()
	return node
}

func (n Node) ListenToGossip() {
	for g := range n.Gossip {
		fmt.Println(n.Id, "ListenToGossip", g)
		time.Sleep(time.Second)
	}
}
func (n Node) ListenToRejects() {
	for {
		fmt.Println(n.Id, "ListenToRejects")
		time.Sleep(time.Second)
	}
}

func (n Node) AddTx(id int, from, to string, amount int) {
	n.Books[from] -= amount
	n.Books[to] += amount
	tx := Tx{id, from, to, amount}
	for _, other := range nodes {
		if other.Id == n.Id {
			continue
		}
		other.Gossip <- tx
	}
}

func main() {
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)

	nodes = []Node{node1, node2, node3}

	fmt.Println(node1, node2, node3)

	node2.AddTx(123, "ABC", "EFG", 100)
	node3.AddTx(124, "ABC", "XYZ", 100)

	for {
		time.Sleep(time.Second)
	}

}
