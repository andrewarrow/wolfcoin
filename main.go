package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	"wolfcoin/args"
	"wolfcoin/network"
)

func PrintHelp() {
	fmt.Println("")
	fmt.Println("  wolfcoin help         # this menu")
	fmt.Println("  wolfcoin supply       # ")
	fmt.Println("")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	argMap := args.ToMap()

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]

	if command == "supply" {
		//2,147,483,647
		total, _ := strconv.ParseInt(argMap["total"], 10, 64)
		millionaires := float64(total) / 1000000.0
		m := map[string]int{}
		for i := 0; i < int(millionaires); i++ {
			addr := network.NewAddress()
			fmt.Println(addr)
			m[addr]++
		}
		for k, v := range m {
			if v > 1 {
				fmt.Println(k)
			}
		}
		fmt.Printf("%d %0.2f\n", total, millionaires)
	} else if command == "start" {
		network.Start()
	} else if command == "help" {
		PrintHelp()
	}
}
