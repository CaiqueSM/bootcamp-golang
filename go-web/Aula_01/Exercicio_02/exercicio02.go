package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Olá Caíque"})
	})
	if err := router.Run(); err!= nil{
		log.Println(err.Error())
	}
}
