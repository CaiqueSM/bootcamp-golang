package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Id        int64   `json:"id"`
	Nome      string  `json:"nome"`
	Sobrenome string  `json:"sobrenome"`
	Email     string  `json:"email"`
	Idade     uint    `json:"idade"`
	Altura    float64 `json:"altura"`
	Ativo     bool    `json:"ativo"`
	Data      string  `json:"data"`
}

type Usuarios []Usuario

var ultimoID int64
var usuarios Usuarios

func ArmazenarUsuarios(c *gin.Context) {
	var usuario Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("campo(s) %s é(são) obrigatório(s)", err.Error()))
		return
	}
	ultimoID++
	usuario.Id = ultimoID

	usuarios = append(usuarios, usuario)

	c.JSON(http.StatusOK, gin.H{"Info:": usuario})
}

func main() {
	router := gin.Default()
	router.POST("/usuarios", ArmazenarUsuarios)
	if err := router.Run(); err!= nil{
		log.Println(err.Error())
	}
}
