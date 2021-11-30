package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"ldap-proxy/mvc/config"
	"ldap-proxy/pkg/accesslog"
	"ldap-proxy/pkg/aescrypt"
	"ldap-proxy/pkg/authorization"
	"ldap-proxy/pkg/ldap"
	"ldap-proxy/pkg/random"
	"time"
)

func Login(c *gin.Context) {
	session := sessions.Default(c)
	auth := session.Get("sessionlogin")
	accesslog.Logger("login","trace",c.Request.RemoteAddr,c.Request.UserAgent(),c.Request.URL.Path,c.Request.Header.Get("sec-ch-ua-platform"),"get session from client")
	if c.Request.Method == "GET" {
		if auth == nil {
			sessionid := random.RandString(26)
			accesslog.Logger("login","trace",c.Request.RemoteAddr,c.Request.UserAgent(),c.Request.URL.Path,c.Request.Header.Get("sec-ch-ua-platform"),"set session for client")
			session.Set("sessionlogin","auth-"+sessionid)
			session.Save()
			c.HTML(200, "login.html", nil)
			return
		}
		//如果存在sessionid ，那么根据sessionid
		_,ok := config.SessionCache.Get(auth.(string))
		if !ok {
			accesslog.Logger("login","trace",c.Request.RemoteAddr,c.Request.UserAgent(),c.Request.URL.Path,c.Request.Header.Get("sec-ch-ua-platform"),"session id not auth")
			//如果发现sessionid 并没有验证存储到存储中
			c.HTML(200, "login.html", nil)
			return
		}
		//获取用户信息，并传递
		c.Redirect(302,"/")
		return
		/*
		if auth != "auth ok" {
			c.HTML(200, "login.html", nil)

		}else{
			username := session.Get("username")
			if username == nil {
				c.Redirect(302,"/login")
				return
			}
			c.Redirect(302,"/")

		}
		return
		 */
	}else if c.Request.Method == "POST" {
		//ip := c.Request.Host
		username := c.PostForm("username")
		password := aescrypt.AesDecrypt(c.PostForm("password"),"1010101010101010")
		accesslog.Logger("login","info",c.Request.RemoteAddr,c.Request.UserAgent(),"authing",c.Request.Header.Get("sec-ch-ua-platform"),fmt.Sprintf("%s authing",username))
		err := ldap.LdapAuth(username,password)
		if err == nil {
			auth := session.Get("sessionlogin")
			authstring := authorization.EntryBase64(username)
			accesslog.Logger("login","warn",c.Request.RemoteAddr,c.Request.UserAgent(),"auth success",c.Request.Header.Get("sec-ch-ua-platform"),fmt.Sprintf("%s auth success",username))
			config.SessionCache.SetValue(auth.(string),authstring,time.Duration(config.FileConfig.Expire) * time.Minute)
			c.Redirect(302,"/")
		}else {
			accesslog.Logger("login","warn",c.Request.RemoteAddr,c.Request.UserAgent(),"auth failed",c.Request.Header.Get("sec-ch-ua-platform"),fmt.Sprintf("%s auth failed",username))
			c.Redirect(302,"/login")
		}
	}
}




