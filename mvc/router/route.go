package router


import (
	"github.com/gin-gonic/gin"
	"ldap-proxy/mvc/controllers"
)


func Registry(r *gin.Engine) {
	r.LoadHTMLGlob("static/*")
	r.GET("/",controllers.Default)
	r.GET("/login",controllers.Login)
	r.POST("/login",controllers.Login)
	r.GET("/auth-proxy",controllers.Auth)
	//test
	r.GET("/ok",controllers.Ok)
	r.GET("/test",controllers.Test)
}