package api

import (
	"AIAssistServer/controller"
	"AIAssistServer/db"
	"AIAssistServer/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// CreateProject 用户创建项目
// 1、拉取镜像创建服务、配置集群内访问路径、绑定访问IP
func CreateProject(context *gin.Context) {
	db1 := db.GetDb()
	var data models.ProjectRes
	context.ShouldBindJSON(&data)

	uuid := uuid.New().String()
	fmt.Println(data.User, uuid)
	url := controller.CreateProject(data.User, uuid)

	in := models.UserProjectRelation{
		UserName:  data.User,
		ProjectId: uuid,
	}

	err := db1.AutoMigrate(&in)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"message": "create db error",
		})
		return
	}

	if err = db1.Create(&in).Error; err != nil {
		context.JSON(http.StatusOK, gin.H{
			"message": "insert error",
		})
		return
	}

	insertData := models.Project{
		ProjectId:   uuid,
		ProjectName: data.ProjectName,
		Model:       data.ProjectModel,
		Name:        data.User,
		Url:         url,
	}

	err = db1.AutoMigrate(&insertData)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"message": "create db error",
		})
		return
	}

	if err = db1.Create(&insertData).Error; err != nil {
		context.JSON(http.StatusOK, gin.H{
			"message": "insert error",
		})
		return
	}

	result := models.ProjectReq{
		Id:  insertData.ProjectId,
		Url: url,
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    result,
	})
}

// ModifyDb 修改用户向量数据库内容
func ModifyDb(context *gin.Context) {
	//dbController(url, user, password, pdf)
}

// ModifyPrompt 修改用户提示词
func ModifyPrompt(context *gin.Context) {
	//promptController()
}
