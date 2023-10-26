package router

import (
	"AIAssistServer/api"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("login", api.Login)
	r.POST("register", api.Register)
	r.POST("createProject", api.CreateProject)
	r.POST("getUserAllAccount", api.GetUserAllAccount)
	r.Run(":8080")
	return r
}
