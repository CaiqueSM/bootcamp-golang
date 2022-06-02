package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/internal/usuarios"
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

type Usuario struct {
	service usuarios.Service
}

func NewUsuario(u usuarios.Service) *Usuario {
	return &Usuario{
		service: u,
	}
}

func (c *Usuario) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalido",
			})
			return
		}
		u, err := c.service.GetAll()

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, u)
	}
}

func (c *Usuario) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalido",
			})
			return
		}

		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		u, err := c.service.Store(req.Nome, req.Sobrenome, req.Email, req.Idade, req.Altura, req.Ativo, req.Data)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, u)
	}
}

func (c *Usuario) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalido",
			})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, fmt.Errorf("invalid ID: %d", id))
			return
		}
		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if req.Nome == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "O nome é obrigatório"})
			return
		}
		if req.Sobrenome == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "O sobrenome é obrigatório"})
			return
		}
		if req.Email == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "O email é obrigatório"})
			return
		}
		if req.Idade == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "A idade é obrigatória"})
			return
		}
		if req.Altura == 0.0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "A altura é obrigatória"})
			return
		}
		if !req.Ativo {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "É obrigatório estar ativo"})
			return
		}
		if req.Data == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "A data é obrigatória"})
			return
		}
		u, err := c.service.Update(id, req.Nome, req.Sobrenome, req.Email, req.Idade, req.Altura, req.Ativo, req.Data)
		if err!= nil{
			ctx.JSON(http.StatusNotFound, gin.H{"error":err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, u)
	}
}

func (c *Usuario) UpdateSobrenome() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalido",
			})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, fmt.Errorf("invalid ID: %d", id))
			return
		}
		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if req.Sobrenome == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "O sobrenome é obrigatório"})
			return
		}

		u, err := c.service.UpdateSobrenome(id, req.Sobrenome)
		if err!= nil{
			ctx.JSON(http.StatusNotFound, gin.H{"error":err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, u)
	}
}

func (c *Usuario) UpdateIdade() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalido",
			})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, fmt.Errorf("invalid ID: %d", id))
			return
		}
		var req request

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if req.Idade == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "A idade é obrigatória"})
			return
		}

		u, err := c.service.UpdateIdade(id, req.Idade)
		if err!= nil{
			ctx.JSON(http.StatusNotFound, gin.H{"error":err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, u)
	}
}

func (c *Usuario) Delete()gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalido",
			})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, fmt.Errorf("invalid ID: %d", id))
			return
		}

		err = c.service.Delete(id)
		if err!= nil{
			ctx.JSON(http.StatusNotFound, gin.H{"error":err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("O produto %d foi removido", id)})
	}
}