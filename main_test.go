package main

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}
func TestFalhador(t *testing.T) {
	t.Fatalf("Testes falhou de proposito, não se preucupe")
}
