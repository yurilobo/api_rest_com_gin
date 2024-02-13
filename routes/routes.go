package routes

import (
	"api_com_gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosOsAlunos)
	r.GET("/:nome", controllers.Saldacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.Run() //
}
