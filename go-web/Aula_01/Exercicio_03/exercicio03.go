package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Id        int       `json:"id"`
	Nome      string    `json:"nome"`
	Sobrenome string    `json:"sobrenome"`
	Email     string    `json:"email"`
	Idade     uint      `json:"idade"`
	Altura    float64   `json:"altura"`
	Ativo     bool      `json:"ativo"`
	Data      time.Time `json:"data de criação"`
}

func GetAll(c *gin.Context) {

	var usuarios = []Usuario{
		{Id: 1, Nome: "A", Sobrenome: "BCDE", Email: "email@generico.com", Idade: 25, Altura: 1.68, Ativo: true, Data: time.Now()},
		{Id: 2, Nome: "B", Sobrenome: "CDEF", Email: "email@generico.com", Idade: 25, Altura: 1.68, Ativo: false, Data: time.Now()}}

	c.IndentedJSON(http.StatusOK, usuarios)
}

func main() {
	router := gin.Default()
	router.GET("/usuarios", GetAll)
	router.Run()
}
