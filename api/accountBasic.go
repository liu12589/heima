package api

import (
	"AIAssistServer/controller"
	"AIAssistServer/db"
	"AIAssistServer/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Login(context *gin.Context) {
	db1 := db.GetDb()
	var data models.LoginMessage
	context.ShouldBindJSON(&data)
	var result models.Login
	err := db1.AutoMigrate(&result)
	if err != nil {
		log.Fatal(err)
	}
	err = db1.Where("user = ?", data.User).First(&result).Error
	if err != nil {
		log.Println("query user account information error :", err)
		context.JSON(http.StatusOK, gin.H{
			"message": "query user account information error",
		})
		return
	}
	if result.Password != data.Password {
		context.JSON(http.StatusOK, gin.H{
			"message": "Incorrect password",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// Register 创建账户，同时在集群内创建命名空间，然后创建一个向量数据库
func Register(context *gin.Context) {
	db1 := db.GetDb()
	var data models.LoginMessage
	context.ShouldBindJSON(&data)

	var insertData models.Login
	err := db1.AutoMigrate(&insertData)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"message": "create db error",
		})
		return
	}

	insertData.User = data.User
	var result models.Login
	err = db1.Where("user = ?", data.User).First(&result).Error
	if err != nil {

	} else {
		context.JSON(http.StatusOK, gin.H{
			"message": "account is exit",
		})
		return
	}

	insertData.Password = data.Password
	if err = db1.Create(&insertData).Error; err != nil {
		context.JSON(http.StatusOK, gin.H{
			"message": "insert error",
		})
		return
	}

	err = controller.CreateNamespace(data.User)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"message": "create namespace error",
		})
		return
	}

	clusterIp, port, err := controller.CreateMilvus(data.User, data.User, data.Password)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"message": "create milvus error",
		})
	}

	// 这里插入数据库数据
	information := models.DbInformation{
		User:        data.User,
		MILVUS_HOSt: clusterIp,
		MILVUS_PORT: port,
	}

	err = db1.AutoMigrate(&models.DbInformation{})
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"message": "create db error",
		})
		return
	}
	if err = db1.Create(&information).Error; err != nil {
		context.JSON(http.StatusOK, gin.H{
			"message": "insert error",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func GetUserAllAccount(context *gin.Context) {
	db1 := db.GetDb()
	var data models.AllAccount
	context.ShouldBindJSON(&data)

	var result []models.Project
	err := db1.Where("user = ?", data.User).Find(&result).Error
	if err != nil {
		log.Println("query user account information error :", err)
		context.JSON(http.StatusOK, gin.H{
			"message": "query user account information error",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    result,
	})
}
