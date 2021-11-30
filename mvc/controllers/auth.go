package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"ldap-proxy/mvc/config"
	"ldap-proxy/pkg/random"
	"ldap-proxy/pkg/accesslog"
)


func Auth(c *gin.Context) {
	session := sessions.Default(c)
	//获取sessionID
	auth := session.Get("sessionlogin")
	accesslog.Logger("auth","trace",c.Request.RemoteAddr,c.Request.UserAgent(),c.Request.URL.Path,c.Request.Header.Get("sec-ch-ua-platform"),"get session from client")
	//如果获取不到，那么代表不存在，需要登陆，并设置客户端sessionid
	if auth == nil {
		sessionid := random.RandString(26)
		accesslog.Logger("auth","trace",c.Request.RemoteAddr,c.Request.UserAgent(),c.Request.URL.Path,c.Request.Header.Get("sec-ch-ua-platform"),"set session for client")
		session.Set("sessionlogin","auth-"+sessionid)
		session.Save()
		c.Status(401)
		return
	}
	//如果存在sessionid ，那么根据sessionid 查看是否登陆
	accesslog.Logger("auth","trace",c.Request.RemoteAddr,c.Request.UserAgent(),c.Request.URL.Path,c.Request.Header.Get("sec-ch-ua-platform"),"get session from client successfully")
	v,ok := config.SessionCache.Get(auth.(string))
	if !ok {
		//如果发现sessionid 并没有验证存储到存储中
		accesslog.Logger("auth","info",c.Request.RemoteAddr,c.Request.UserAgent(),c.Request.URL.Path,c.Request.Header.Get("sec-ch-ua-platform"),"sessionid not auth")
		c.Status(401)
		return
	}
	//获取用户信息，并传递
	username := v.(string)
	accesslog.Logger("auth","info",c.Request.RemoteAddr,c.Request.UserAgent(),c.Request.URL.Path,c.Request.Header.Get("sec-ch-ua-platform"),"set auth to client")
	c.Header("Auth_User",username)
	c.Status(200)
	return
	/*
	if auth == "auth ok" {
		username := session.Get("username")
		if username == nil {
			c.Status(401)
			return
		}
		user := username.(string)
		if  ! strings.Contains(user,"Basic") {
			//c.Redirect(302,"/")
			c.Status(401)
			return
		}
		//这里的参数主要是为了后端能够接受用户信息
		c.Header("Auth_User",user)
		c.Status(200)
		return
	}
	c.Status(401)
	*/
}