package main

import (
	"api_com_gin/database"
	"api_com_gin/models"
	"api_com_gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Gui Lima", CPF: "00000000000", RG: "47000000000"},
		{Nome: "ANA Lima", CPF: "00000000001", RG: "47000000001"},
	}
	routes.HandleRequest()
}
