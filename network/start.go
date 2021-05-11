package network

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func Start(port string) {
	fmt.Println("wolfcoin starting on port", port)

	r := gin.Default()
	r.POST("/tx", func(c *gin.Context) {
		b, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body.Close()
		fmt.Println(string(b))
		var tx TxMessage
		json.Unmarshal(b, &tx)
		c.JSON(200, gin.H{"ok": true})
	})
	r.Run(":" + port)
}
