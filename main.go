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
	"strings"
	"time"
	"wolfcoin/args"
	"wolfcoin/files"
	"wolfcoin/network"
)

func PrintHelp() {
	fmt.Println("")
	fmt.Println("  wolfcoin help         # this menu")
	fmt.Println("  wolfcoin genesis      # ")
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
	port := argMap["port"]
	files.ReadyDir(port)

	if command == "genesis" {
		//30,000,000,000
		total := 30000000000
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
		ioutil.WriteFile(files.Path+"/genesis.v", []byte(strings.Join(vbuff, "\n")), 0755)
		ioutil.WriteFile(files.Path+"/genesis.s", []byte(strings.Join(sbuff, "\n")), 0755)
		fmt.Printf("%d %0.2f\n", total, millionaires)
	} else if command == "practice" {
		v, s, _ := ed25519.GenerateKey(nil)
		opts := SignerOptsThing{}
		message := []byte("hello")
		sig, _ := s.Sign(nil, message, opts)
		fmt.Println(sig)
		b := ed25519.Verify(v, message, sig)
		fmt.Println(b)
	} else if command == "loop" {
		network.ReadInGenesis()
		for {
			from, fromS := network.Random()
			to, _ := network.Random()
			amount := int64(rand.Intn(999999) * 1000000) // micro-wolf
			jsonString := network.CreateMessage(from, to, amount)
			fmt.Println(jsonString)
			opts := SignerOptsThing{}
			data, _ := hex.DecodeString(fromS)
			s := ed25519.PrivateKey(data)
			sig, _ := s.Sign(nil, []byte(jsonString), opts)
			sigString := strings.ToLower(fmt.Sprintf("%X", sig))
			thing := network.ValidatePayload{}
			thing.JsonString = jsonString
			thing.SigString = sigString
			asBytes, _ := json.Marshal(thing)
			network.DoPost("127.0.0.1:3001", "/validate", asBytes)
			time.Sleep(time.Second)
		}
	} else if command == "tx" {
		from := "cf90d4930c68918d6f73b6fa0d3780f6baed4cb0bbc4106ac93e54a963997707"
		to := "347b89a033993042a863886d39d97f1c9daa82d2d0a8e3ad49a37571451fc268"
		amount := int64(100 * 1000000) // micro-wolf
		jsonString := network.CreateMessage(from, to, amount)
		fmt.Println(jsonString)
		opts := SignerOptsThing{}
		shhh := ""
		data, _ := hex.DecodeString(shhh)
		s := ed25519.PrivateKey(data)
		sig, _ := s.Sign(nil, []byte(jsonString), opts)
		sigString := strings.ToLower(fmt.Sprintf("%X", sig))
		fmt.Println(sigString)
		network.Validate(jsonString, sigString)
	} else if command == "start" {
		port := argMap["port"]
		network.ReadInGenesis()
		network.Start(port)
	} else if command == "help" {
		PrintHelp()
	}
}

type SignerOptsThing struct {
}

func (s SignerOptsThing) HashFunc() crypto.Hash {
	return 0
}
