package controllers

import (
	"api_com_gin/models"

	"github.com/gin-gonic/gin"
)

func ExibeTodosOsAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}
func Saldacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz: ": " E ai " + nome + ", tudo beleza?",
	})
}
