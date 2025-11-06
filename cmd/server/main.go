package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	router.Any("/login", func(c *gin.Context) {
		c.String(200, "login")
	})
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
