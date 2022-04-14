package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/hi", func(c *gin.Context) {
		c.String(200, "bye")
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
