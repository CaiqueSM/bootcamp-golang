package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type usuario struct {
	id        int     `json: "id"`
	nome      string  `json: "nome"`
	sobrenome string  `json: "sobrenome"`
	email     string  `json: "email"`
	idade     uint    `json: "idade"`
	altura    float64 `json: "altura"`
	ativo     bool    `json: "ativo"`
	data      string  `json: "data de criação"`
}

func GetAll(c *gin.Context) {
	//u := []usuario{{id: 1, nome: "A", sobrenome: "BCDE", email: "email@generico.com",
	//	idade: 25, altura: 1.68, ativo: true, data: "28-05-2022"},
	//	{id: 2, nome: "B", sobrenome: "CDEF", email: "email@generico.com", idade: 25, altura: 1.68,
	//		ativo: false, data: "28-05-2022"}}

	c.JSON(http.StatusOK, usuario{id: 2, nome: "B", sobrenome: "CDEF", email: "email@generico.com", idade: 25, altura: 1.68,
		ativo: false, data: "28-05-2022"})
}

func main() {

	router := gin.Default()

	router.GET("/usuarios", GetAll)
	router.Run()
}
