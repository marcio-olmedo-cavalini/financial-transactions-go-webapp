package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marcio-olmedo-cavalini/financial-transactions-go-webapp/controllers"
)

var r *gin.Engine

func HandleRequests() {
	r = gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.LoadHTMLGlob("html/*")
	handleHtml()
	handleServices()
	r.Run()
}

func handleHtml() {
	r.GET("/index", controllers.ShowIndexPage)
	r.GET("/user", controllers.ShowUserListPage)
	r.GET("/newuser", controllers.ShowNewUserPage)
	r.GET("/edituser", controllers.ShowEditUserPage)
	r.NoRoute(controllers.RouteNotFound)
}

func handleServices() {
	r.POST("/upload", controllers.UploadFile)
	r.POST("/insertuser", controllers.SaveNewUser)
	r.POST("/updateuser", controllers.UpdateUser)
}
