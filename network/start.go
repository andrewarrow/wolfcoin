package network

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"wolfcoin/files"

	"github.com/gin-gonic/gin"
)

var nodeName string

type ValidatePayload struct {
	JsonString string `json:"json"`
	SigString  string `json:"sig"`
}

func Start(port string) {
	fmt.Println("wolfcoin starting on port", port)
	nodeName = "127.0.0.1:" + port

	r := gin.Default()
	r.POST("/validate", func(c *gin.Context) {
		b, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body.Close()
		fmt.Println(string(b))
		var vp ValidatePayload
		json.Unmarshal(b, &vp)
		if Validate(vp.JsonString, vp.SigString) {
			c.JSON(200, gin.H{"ok": true})
			return
		}
		c.JSON(422, gin.H{"ok": false})
	})
	r.POST("/tx", func(c *gin.Context) {
		b, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body.Close()
		fmt.Println(string(b))
		f, _ := os.OpenFile(files.Path+"/tx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		f.WriteString(string(b) + "\n")
		f.Close()
		c.JSON(200, gin.H{"ok": true})
	})
	r.Run(":" + port)
}
