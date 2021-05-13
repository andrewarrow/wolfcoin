package main

import (
	"fmt"
	"time"
)

var nodes = []Node{}

type Node struct {
	Id        int
	Books     map[string]int
	Gossip    chan Tx
	Reject    chan Tx
	RejectMap map[string]int
}

type Tx struct {
	Id     string
	From   string
	To     string
	Amount int
	NodeId int
}

func Debug() {
	for _, n := range nodes {
		fmt.Printf("   %d %d %v\n", n.Id, len(n.Books), n.RejectMap)
	}
}
func (n Node) Others() []Node {
	list := []Node{}
	for _, other := range nodes {
		if other.Id == n.Id {
			continue
		}
		list = append(list, other)
	}
	return list
}

func NewNode(id int) Node {
	node := Node{}
	node.Id = id
	node.Books = map[string]int{}
	node.Books["ABC"] = 100
	node.Gossip = make(chan Tx, 1024)
	node.Reject = make(chan Tx, 1024)
	node.RejectMap = map[string]int{}
	go node.ListenToGossip()
	go node.ListenToRejects()
	return node
}

func (n Node) ListenToGossip() {
	for g := range n.Gossip {
		fmt.Println(n.Id, "ListenToGossip", g.Id)
		if n.Books[g.From]-g.Amount < 0 {
			for _, other := range n.Others() {
				g.NodeId = n.Id
				other.Reject <- g
			}
		} else {
			n.Books[g.From] -= g.Amount
			n.Books[g.To] += g.Amount
		}
		time.Sleep(time.Second)
	}
}
func (n Node) ListenToRejects() {
	for r := range n.Reject {
		//fmt.Println("---", n.Id, "ListenToRejects", r.Id, r.NodeId)
		n.RejectMap[r.Id]++
		time.Sleep(time.Second)
	}
}

func (n Node) AddTx(id, from, to string, amount int) {
	n.Books[from] -= amount
	n.Books[to] += amount
	tx := Tx{id, from, to, amount, n.Id}
	for _, other := range n.Others() {
		other.Gossip <- tx
	}
}

func main() {
	node1 := NewNode(1)
	node2 := NewNode(2)
	node3 := NewNode(3)

	nodes = []Node{node1, node2, node3}

	fmt.Println(node1, node2, node3)

	node2.AddTx("123", "ABC", "EFG", 100)
	node3.AddTx("124", "ABC", "XYZ", 100)

	for {
		time.Sleep(time.Second)
		Debug()
	}

}
