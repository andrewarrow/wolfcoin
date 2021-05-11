package network

import (
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func Transfer() {
}
func CreateMessage(from, to string, amount int64) string {
	tx := TxMessage{}
	tx.From = from
	tx.To = to
	tx.Amount = amount
	s, _ := json.Marshal(tx)
	return string(s)
}
func Validate(jsonString, sig string) {
	var tx TxMessage
	json.Unmarshal([]byte(jsonString), &tx)
	data, _ := hex.DecodeString(tx.From)
	v := ed25519.PublicKey(data)
	sigData, _ := hex.DecodeString(sig)
	b := ed25519.Verify(v, []byte(jsonString), sigData)
	if b == false {
		fmt.Println("sig is not right")
		return
	}
	if books[tx.From] < tx.Amount {
		//TODO
	}
	books[tx.From] -= tx.Amount
	books[tx.To] += tx.Amount

	tx.Ts = time.Now().UnixNano()
	asBytes, _ := json.Marshal(tx)
	go SendToPeers(asBytes)
	f, _ := os.OpenFile("tx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.WriteString(string(asBytes) + "\n")
	f.Close()
}

func SendToPeers(msg []byte) {
	DoPost("/tx", msg)
}

type TxMessage struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount int64  `json:"amount"`
	Ts     int64  `json:"ts"`
}
