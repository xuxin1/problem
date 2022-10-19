package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	number := 0
	r.GET("/ping", func(context *gin.Context) {
		number += 1
		fmt.Printf("%d\n", number)
		context.String(http.StatusOK, "hello!")
	})
	r.Run("localhost:8081")
}
