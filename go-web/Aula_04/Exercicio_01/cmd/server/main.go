package main

import (
	"log"
	"os"

	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/cmd/server/handler/controller"
	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/docs"
	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/internal/usuarios/repository"
	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/internal/usuarios/service"
	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//@title MELI Bootcamp API
//@version 1.0
//@decripition This API Handle MELI Users.
//@termsOfService http://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

//@contact.name API Support
//@contact.url https://developers.mercadolivre.com.ar/suppoert

//@license.name Apache 2.0
//@license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("failed to load .env")
	}

	// db := store.New(store.FileType, "../../usuarios.json")
	// repo := repository.NewRepository(db)
	db.Init()
	repo := repository.NewRepositoryMariaDB(db.StorageDB)
	service := service.NewService(repo)
	u := controller.NewUsuario(service)
	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ur := r.Group("/usuarios")
	{
		ur.POST("/", u.Store())
		ur.GET("/", u.GetAll())
		ur.PUT("/:id", u.Update())
		ur.PATCH("/:id", u.UpdateSobrenomeIdade())
		ur.DELETE("/:id", u.Delete())
	}
	if err := r.Run(); err != nil{
		log.Println(err.Error())
	}
}
