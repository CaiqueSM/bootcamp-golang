package main

import (
	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/cmd/server/handler"
	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/internal/usuarios"
	"github.com/gin-gonic/gin"
)

func main(){
	repo := usuarios.NewRepository()
	service := usuarios.NewService(repo)
	u := handler.NewUsuario(service)

	r := gin.Default()
	ur := r.Group("/usuarios")
	ur.POST("/", u.Store())
	ur.GET("/", u.GetAll())
	ur.PUT("/:id", u.Update())
	r.Run()
}