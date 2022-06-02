package handler

import (
	"net/http"

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
