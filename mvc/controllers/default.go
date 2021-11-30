package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"ldap-proxy/mvc/config"
	"ldap-proxy/pkg/accesslog"
)

func Default(c *gin.Context) {
	session := sessions.Default(c)
	auth := session.Get("sessionlogin")
	accesslog.Logger("default","trace",c.Request.RemoteAddr,c.Request.UserAgent(),c.Request.URL.Path,c.Request.Header.Get("sec-ch-ua-platform"),"get session from client")
	if auth == nil {
		c.Redirect(302,"/login")
		return
	}
	_,ok := config.SessionCache.Get(auth.(string))
	if !ok {
		accesslog.Logger("default","trace",c.Request.RemoteAddr,c.Request.UserAgent(),c.Request.URL.Path,c.Request.Header.Get("sec-ch-ua-platform"),"session auth failed")
		//如果发现sessionid 并没有验证存储到存储中
		c.Status(401)
		return
	}
	accesslog.Logger("default","trace",c.Request.RemoteAddr,c.Request.UserAgent(),c.Request.URL.Path,c.Request.Header.Get("sec-ch-ua-platform"),"already auth")
	c.Writer.WriteHeader(200)
}
