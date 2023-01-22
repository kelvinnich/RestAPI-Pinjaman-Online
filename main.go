package main

import (
	"pinjaman-online/config"
	"github.com/gin-gonic/gin"
)
var db = config.ConnectDB()

func main(){
	defer config.CloseDB(db)

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(201, gin.H{
			"message" : "test",
		})
	})

	r.Run(":3000")
}