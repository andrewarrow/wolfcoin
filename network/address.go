package network

import (
	"crypto/ed25519"
	"fmt"
	"strings"
)

func NewAddress() string {
	v, s, _ := ed25519.GenerateKey(nil)
	name := fmt.Sprintf("%X", v)
	fmt.Printf("%d\n", len(s))
	return "wolf" + strings.ToLower(name)
}
