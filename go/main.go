package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()

    r.GET("/go", func(c *gin.Context) {
        c.String(200, "Hello World! From Go")
    })

    r.Run(":4000")
}