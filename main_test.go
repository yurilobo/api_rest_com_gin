package main

import (
	"api_com_gin/controllers"
	"api_com_gin/database"
	"api_com_gin/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "12345678902", RG: "223456789"}
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
	//fmt.Println(resposta.Body)
}
func TestBucaAlunoPorCPFHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCpf)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678902", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
func TestBucaAlunoPorIdHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	//fmt.Println(alunoMock.Nome)
	assert.Equal(t, "Nome do Aluno Teste", alunoMock.Nome)
	assert.Equal(t, "12345678902", alunoMock.CPF)

}
func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	pathDeBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathDeBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
func TestEditaUmAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	aluno := models.Aluno{Nome: "Nome do Aluno Teste3", CPF: "47123456789", RG: "123456700"}
	valorJson, _ := json.Marshal(aluno)
	pathParaEditar := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathParaEditar, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMockAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)
	assert.Equal(t, "47123456789", alunoMockAtualizado.CPF)
	assert.Equal(t, "123456700", alunoMockAtualizado.RG)
	assert.Equal(t, "Nome do Aluno Teste3", alunoMockAtualizado.Nome)
	fmt.Println(alunoMockAtualizado.Nome)
}
