package controller

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/internal/usuarios/domain"
	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nome      string  `json:"nome" binding:"required"`
	Sobrenome string  `json:"sobrenome" binding:"required"`
	Email     string  `json:"email" binding:"required"`
	Idade     uint    `json:"idade" binding:"required"`
	Altura    float64 `json:"altura" binding:"required"`
	Ativo     bool    `json:"ativo" binding:"required"`
	Data      string  `json:"data" binding:"required"`
}

type requestUpdate struct {
	Nome      string  `json:"nome"`
	Sobrenome string  `json:"sobrenome"`
	Email     string  `json:"email"`
	Idade     uint    `json:"idade"`
	Altura    float64 `json:"altura"`
	Ativo     bool    `json:"ativo"`
	Data      string  `json:"data"`
}

type Usuario struct {
	service domain.Service
}

// ListUsers godoc
// @Summary List Users
// @Tags Users
// @Description get users
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /usuarios [get]
func (c *Usuario) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewRespponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		}
		u, err := c.service.GetAll()

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewRespponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewRespponse(http.StatusOK, u, ""))
	}
}

//StoreUsers godoc
//@Summary Store Users
//@Tags Users
//@Description get users
//@Accept json
//@Produce json
//@Param token header string true "token"
//@Param usuarios body request true "usuario to store"
//@Success 200 {object} web.Response
//@Router /usuarios [post]
func (c *Usuario) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewRespponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusNotFound, web.NewRespponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		u, err := c.service.Store(req.Nome, req.Sobrenome, req.Email, req.Idade, req.Altura, req.Ativo, req.Data)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewRespponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewRespponse(http.StatusOK, u, ""))
	}
}

//StoreUsers godoc
//@Summary Update Users
//@Tags Users
//@Description get users
//@Accept json
//@Produce json
//@Param token header string true "token"
//@Param usuarios body request true "usuario to update"
//@Success 200 {object} web.Response
//@Router /usuarios/:id [put]
func (c *Usuario) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewRespponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewRespponse(http.StatusBadRequest, nil, "Erro: Id inválido"))
			return
		}
		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewRespponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		if req.Nome == "" {
			ctx.JSON(http.StatusBadRequest, web.NewRespponse(http.StatusBadRequest, nil, "Error: O nome é obrigatório"))
			return
		}
		if req.Sobrenome == "" {
			ctx.JSON(http.StatusBadRequest, web.NewRespponse(http.StatusBadRequest, nil, "error: O sobrenome é obrigatório"))
			return
		}
		if req.Email == "" {
			ctx.JSON(http.StatusBadRequest, web.NewRespponse(http.StatusBadRequest, nil, "error: O email é obrigatório"))
			return
		}
		if req.Idade == 0 {
			ctx.JSON(http.StatusBadRequest, web.NewRespponse(http.StatusBadRequest, nil, "error: A idade é obrigatória"))
			return
		}
		if req.Altura == 0.0 {
			ctx.JSON(http.StatusBadRequest, web.NewRespponse(http.StatusBadRequest, nil, "error: A altura é obrigatória"))
			return
		}
		if !req.Ativo {
			ctx.JSON(http.StatusBadRequest, web.NewRespponse(http.StatusBadRequest, nil, "error: É obrigatório estar ativo"))
			return
		}
		if req.Data == "" {
			ctx.JSON(http.StatusBadRequest, web.NewRespponse(http.StatusBadRequest, nil, "error: A data é obrigatória"))
			return
		}
		u, err := c.service.Update(id, req.Nome, req.Sobrenome, req.Email, req.Idade, req.Altura, req.Ativo, req.Data)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewRespponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewRespponse(http.StatusOK, u, ""))
	}
}

//StoreUsers godoc
//@Summary UpdateSobrenomeIdade Users
//@Tags Users
//@Description Update users
//@Accept json
//@Produce json
//@Param token header string true "token"
//@Param usuarios body request true "usuario to update"
//@Success 200 {object} web.Response
//@Router /usuarios/:id [patch]
func (c *Usuario) UpdateSobrenomeIdade() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewRespponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewRespponse(http.StatusBadRequest, nil, "Erro: Id inválido"))
			return
		}
		var req requestUpdate

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewRespponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		if req.Nome == "" && req.Sobrenome == "" && req.Email == "" && req.Idade == 0 &&
			req.Altura == 0.0 && !req.Ativo && req.Data == "" {
			ctx.JSON(http.StatusBadRequest, web.NewRespponse(http.StatusBadRequest, nil, "error: solicitação vazia."))
			return
		}

		u, err := c.service.UpdateSobrenomeIdade(id, req.Sobrenome, req.Idade)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewRespponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewRespponse(http.StatusOK, u, ""))
	}
}

//StoreUsers godoc
//@Summary Delete Users
//@Tags Users
//@Description delete users
//@Accept json
//@Produce json
//@Param token header string true "token"
//@Param usuarios body request true "usuario to delete"
//@Success 200 {object} web.Response
//@Router /usuarios [delete]
func (c *Usuario) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewRespponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewRespponse(http.StatusBadRequest, nil, "Erro: Id inválido"))
			return
		}

		err = c.service.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewRespponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewRespponse(http.StatusOK, fmt.Sprintf("O Usuário %d foi removido", id), ""))
	}
}

func NewUsuario(u domain.Service) *Usuario {
	return &Usuario{
		service: u,
	}
}
