package main

import "gopkg.in/gin-gonic/gin.v1"

func main() {
    r := gin.Default()

    r.GET("/debentures", GetDebentures)

    r.Run() // listen and serve on 0.0.0.0:8080
}
