package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // ブラウザで http://localhost/ にアクセスしたとき "Hello, World!" を返す
    r.GET("/", func(c *gin.Context) {
        c.String(200, "Hello, World!")
    })

		// ここでWebサーバーを立ち上げる
    r.Run(":8080")
}
