package usuarios_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/cmd/server/handler/controller"
	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/internal/usuarios"
	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, "../../usuarios.json")
	repo := usuarios.NewRepository(db)
	service := usuarios.NewService(repo)
	u := controller.NewUsuario(service)
	r := gin.Default()
	pr := r.Group("/usuarios")
	pr.DELETE("/:id", u.Delete())
	pr.GET("/", u.GetAll())
	pr.PATCH("/:id", u.UpdateSobrenomeIdade())
	pr.POST("/", u.Store())
	pr.PUT("/:id", u.Update())
	return r
}

func createRequestTest(method string, url string, body []byte) (*http.Request,
	*httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", "123456")
	return req, httptest.NewRecorder()
}

func TestDeleteUsuariosOK(t *testing.T) {

	r := createServer()
	req, rr := createRequestTest(http.MethodDelete, "/usuarios/1", nil)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetUsuariosOK(t *testing.T) {

	r := createServer()
	req, rr := createRequestTest(http.MethodGet, "/usuarios/", nil)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	type foo struct {
		Data []usuarios.Usuario
	}

	objRes := foo{}

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	assert.True(t, len(objRes.Data) > 0)
}

func TestPatchUsuariosOK(t *testing.T) {

	input := usuarios.Usuario{
		Id:        1,
		Sobrenome: "Patch Update",
		Idade:     45,
	}

	update, err := json.Marshal(input)

	if err != nil {
		log.Println(err.Error())
	}

	r := createServer()
	req, rr := createRequestTest(http.MethodPatch, "/usuarios/3", update)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestPostUsuariosOK(t *testing.T) {

	input := usuarios.Usuario{
		Nome:      "abcd",
		Sobrenome: "jdhfg",
		Email:     "generico@email.com",
		Idade:     25,
		Altura:    1.68,
		Ativo:     true,
		Data:      "10/06/2022",
	}

	create, err := json.Marshal(input)

	if err != nil {
		log.Println(err.Error())
	}

	r := createServer()
	req, rr := createRequestTest(http.MethodPost, "/usuarios/", create)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestPutUsuariosOK(t *testing.T) {

	input := usuarios.Usuario{
		Id:        4,
		Nome:      "abcd",
		Sobrenome: "Put Update",
		Email:     "generico@email.com",
		Idade:     25,
		Altura:    1.68,
		Ativo:     true,
		Data:      "09/06/2022",
	}

	update, err := json.Marshal(input)

	if err != nil {
		log.Println(err.Error())
	}

	r := createServer()
	req, rr := createRequestTest(http.MethodPut, "/usuarios/2", update)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
