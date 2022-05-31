package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Id        int64   `json:"id"`
	Nome      string  `json:"nome" binding:"required"`
	Sobrenome string  `json:"sobrenome" binding:"required"`
	Email     string  `json:"email" binding:"required"`
	Idade     uint    `json:"idade" binding:"required"`
	Altura    float64 `json:"altura" binding:"required"`
	Ativo     bool    `json:"ativo" binding:"required"`
	Data      string  `json:"data" binding:"required"`
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
	router.Run()
}
