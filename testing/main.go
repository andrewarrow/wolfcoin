package main

import "fmt"

var gossip []Tx = []Tx{}
var rejects []Tx = []Tx{}

type Node struct {
	Books map[string]int
}

type Tx struct {
	Id     int
	From   string
	To     string
	Amount int
}

func NewNode() Node {
	node := Node{}
	node.Books = map[string]int{}
	return node
}

func (n Node) AddTx(id int, from, to string, amount int) {
	n.Books[from] -= amount
	n.Books[to] += amount
	tx := Tx{id, from, to, amount}
	gossip = append(gossip, tx)
}

func main() {
	node1 := NewNode()
	node1.Books["ABC"] = 100
	node2 := NewNode()
	node2.Books["ABC"] = 100
	node3 := NewNode()
	node3.Books["ABC"] = 100

	fmt.Println(node1, node2, node3)

	node2.AddTx(123, "ABC", "EFG", 100)
	fmt.Println(node1, node2, node3)
	node3.AddTx(124, "ABC", "XYZ", 100)
	fmt.Println(node1, node2, node3)

	for _, g := range gossip {

		if node1.Books[g.From]-g.Amount < 0 {
			fmt.Println("reject tx", g.Id)
			rejects = append(rejects, g)
			continue
		}

		node1.Books[g.From] -= g.Amount
		node1.Books[g.To] += g.Amount
	}

	fmt.Println(node1, node2, node3)

	for _, r := range rejects {
		node3.Books[r.From] += r.Amount
		node3.Books[r.To] -= r.Amount
	}
	fmt.Println(node1, node2, node3)
	node3.AddTx(123, "ABC", "EFG", 100)
	fmt.Println(node1, node2, node3)

}
