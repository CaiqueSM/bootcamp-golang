package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Id        int64   `json: "id"`
	Nome      string  `json: "nome"`
	Sobrenome string  `json: "sobrenome"`
	Email     string  `json: "email"`
	Idade     uint    `json: "idade"`
	Altura    float64 `json: "altura"`
	Ativo     bool    `json: "ativo"`
	Data      string  `json: "data de criação"`
}

var usuarios = map[string]Usuario{
	"1": {Id: 1, Nome: "A", Sobrenome: "BCDE", Email: "email@generico.com", Idade: 30, Altura: 1.80, Ativo: true, Data: "30-05-2022"},
	"2": {Id: 2, Nome: "B", Sobrenome: "CDEF", Email: "email@generico.com", Idade: 21, Altura: 1.70, Ativo: false, Data: "30-05-2022"},
	"3": {Id: 3, Nome: "c", Sobrenome: "DEFG", Email: "email@alternativo.com", Idade: 25, Altura: 1.68, Ativo: true, Data: "29-05-2022"},
	"4": {Id: 4, Nome: "d", Sobrenome: "DEFG", Email: "email@alternativo.com", Idade: 25, Altura: 1.68, Ativo: false, Data: "27-05-2022"},
	"5": {Id: 5, Nome: "e", Sobrenome: "CDEF", Email: "email@generico.com", Idade: 28, Altura: 1.91, Ativo: false, Data: "27-05-2022"},
}

func GetUser(c *gin.Context) {
	usuario, ok := usuarios[c.Param("id")]
	if ok {
		c.String(http.StatusOK, "Informações do usário %v:", usuario)
	} else {
		c.String(http.StatusNotFound, "Usário não encontrado.")
	}
}

func FilterUsers(c *gin.Context) {
	var filters []Usuario
	for _, u := range usuarios {
		if c.Query("id") == strconv.FormatInt(u.Id, 10) {
			filters = append(filters, u)
		}
		if c.Query("nome") == u.Nome {
			filters = append(filters, u)
		}
		if c.Query("sobrenome") == u.Sobrenome {
			filters = append(filters, u)
		}
		if c.Query("email") == u.Email {
			filters = append(filters, u)
		}
		if c.Query("idade") == strconv.FormatInt(int64(u.Idade), 10) {
			filters = append(filters, u)
		}
		if c.Query("altura") == fmt.Sprintf("%v", u.Altura) {
			filters = append(filters, u)
		}
		if c.Query("ativo") == strconv.FormatBool(u.Ativo) {
			filters = append(filters, u)
		}
		if c.Query("data") == u.Data {
			filters = append(filters, u)
		}

	}

	c.IndentedJSON(http.StatusOK, filters)
}

func main() {
	router := gin.Default()
	router.GET("/usuarios/:id", GetUser)
	router.GET("/usuarios", FilterUsers)
	router.Run()
}
