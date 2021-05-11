package network

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

var books map[string]int64 = map[string]int64{}

func ReadInGenesis() {
	b, _ := ioutil.ReadFile("genesis.v")
	for i, line := range strings.Split(string(b), "\n") {
		books[line] = 1000000 * 1000000
		fmt.Println(i, len(line))
	}
	fmt.Println(len(books))
	b, _ = ioutil.ReadFile("tx.log")
	for _, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 {
			break
		}
		var tx TxMessage
		json.Unmarshal([]byte(line), &tx)
		books[tx.From] -= tx.Amount
		books[tx.To] += tx.Amount
	}
	for k, v := range books {
		if v != 1000000*1000000 {
			fmt.Println(k, v)
		}
	}
}

func Random() (string, string) {
	b, _ := ioutil.ReadFile("genesis.v")
	lines := strings.Split(string(b), "\n")
	r := rand.Intn(len(lines))
	v := lines[r]
	b, _ = ioutil.ReadFile("/Users/aa/Documents/genesis.s")
	lines = strings.Split(string(b), "\n")
	s := lines[r]
	return v, s
}
