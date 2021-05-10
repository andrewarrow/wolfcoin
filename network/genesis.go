package network

import (
	"fmt"
	"io/ioutil"
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
}
