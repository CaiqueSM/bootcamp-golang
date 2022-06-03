package main

import (
	"log"

	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/cmd/server/handler"
	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/internal/usuarios"
	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("failed to load .env")
	}
	db := store.New(store.FileType, "usuarios.json")
	repo := usuarios.NewRepository(db)
	service := usuarios.NewService(repo)
	u := handler.NewUsuario(service)

	r := gin.Default()
	ur := r.Group("/usuarios")
	{
		ur.POST("/", u.Store())
		ur.GET("/", u.GetAll())
		ur.PUT("/:id", u.Update())
		ur.PATCH("/:id", u.UpdateSobrenomeIdade())
		ur.DELETE("/:id", u.Delete())
	}
	r.Run()
}
