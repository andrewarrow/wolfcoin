package main

import (
	"crypto"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
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
		vbuff := []string{}
		sbuff := []string{}
		for i := 0; i < int(millionaires); i++ {
			v, s, _ := ed25519.GenerateKey(nil)
			vhex := strings.ToLower(fmt.Sprintf("%X", v))
			shex := strings.ToLower(fmt.Sprintf("%X", s))

			vbuff = append(vbuff, vhex)
			sbuff = append(sbuff, shex)
		}
		ioutil.WriteFile("genesis.v", []byte(strings.Join(vbuff, "\n")), 0755)
		ioutil.WriteFile("genesis.s", []byte(strings.Join(sbuff, "\n")), 0755)
		fmt.Printf("%d %0.2f\n", total, millionaires)
	} else if command == "practice" {
		v, s, _ := ed25519.GenerateKey(nil)
		opts := SignerOptsThing{}
		message := []byte("hello")
		sig, _ := s.Sign(nil, message, opts)
		fmt.Println(sig)
		b := ed25519.Verify(v, message, sig)
		fmt.Println(b)
	} else if command == "tx" {
		from := "wolf347b89a033993042a863886d39d97f1c9daa82d2d0a8e3ad49a37571451fc269"
		to := "wolf347b89a033993042a863886d39d97f1c9daa82d2d0a8e3ad49a37571451fc268"
		amount := int64(100000000) // micro-wolf
		jsonString := CreateMessage(from, to, amount)
		fmt.Println(jsonString)
		_, s, _ := ed25519.GenerateKey(nil)
		opts := SignerOptsThing{}
		sig, _ := s.Sign(nil, []byte(jsonString), opts)
		sigString := strings.ToLower(fmt.Sprintf("%X", sig))
		fmt.Println(sigString)
		v := from[4:]
		data, _ := hex.DecodeString(v)
		fmt.Println(data)
	} else if command == "start" {
		network.ReadInGenesis()
		network.Start()
	} else if command == "help" {
		PrintHelp()
	}
}

func CreateMessage(from, to string, amount int64) string {
	tx := TxMessage{}
	tx.From = from
	tx.To = to
	tx.Amount = amount
	s, _ := json.Marshal(tx)
	return string(s)
}

type TxMessage struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount int64  `json:"amount"`
}

type SignerOptsThing struct {
}

func (s SignerOptsThing) HashFunc() crypto.Hash {
	return 0
}
