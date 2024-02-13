package controllers

import (
	"api_com_gin/database"
	"api_com_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibeTodosOsAlunos(c *gin.Context) {

	//exibe uma lista de alunos
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}
func Saldacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz: ": " E ai " + nome + ", tudo beleza?",
	})
}
func BuscaAlunoPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado!"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)

}
