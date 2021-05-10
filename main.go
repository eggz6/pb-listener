package main

import (
	"fmt"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", pingHandle)
	r.Run(":11001") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func pingHandle(c *gin.Context) {
	body := make(map[string]interface{})

	err := c.BindJSON(&body)
	if err != nil && err != io.EOF {
		log.Println("ping bind body failed. err=%v", err)
		c.AbortWithError(400, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "pong",
		"body":    fmt.Sprintf("%+v", body),
	})

	return
}
