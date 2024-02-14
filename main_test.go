package main

import (
	"api_com_gin/controllers"
	"api_com_gin/database"
	"api_com_gin/models"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaSatusCodeDaSaudacaoComParametro(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saldacao)
	req, _ := http.NewRequest("GET", "/Yuri", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")
	mockDaResposta := `{"API diz: ":" E ai Yuri, tudo beleza?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))
	fmt.Println(string(respostaBody))
	fmt.Println(mockDaResposta)
}
func TestListandoTodosOsAlunosHandle(t *testing.T) {
	//chamando o banco d edados de desenvolvimento
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	//busca as rotas de teste do gin
	r := SetupDasRotasDeTeste()
	//rota
	r.GET("/alunos", controllers.ExibeTodosOsAlunos)
	//registrar requisiccao
	req, _ := http.NewRequest("GET", "/alunos", nil)
	//resposta, armazena informacoes nessa variavel
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
	//imprime os alunos que estao no banco de dados
	fmt.Println(resposta.Body)
}
