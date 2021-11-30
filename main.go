package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
	"ldap-proxy/mvc/config"
	"ldap-proxy/mvc/router"
	"ldap-proxy/pkg/accesslog"
	"ldap-proxy/pkg/sessioncache"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	var err error = configor.Load(&config.FileConfig, "config/config.yaml")
	if err != nil {
		panic(err)
	}
	if config.FileConfig.Port == "" {
		config.FileConfig.Port = "9999"
	}
	config.SessionCache,err = sessioncache.NewCache(config.FileConfig.Store.Name,config.FileConfig.Store.ConnInfo)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	//config.FileConfig.Cache = sessioncache.NewCache()
}


func main() {
	var ginEngine *gin.Engine = gin.New()
	err := accesslog.LogInit(ginEngine,config.FileConfig.Loglevel,config.FileConfig.LogFormat,config.FileConfig.LogPath)
	if err != nil {
		log.Error(err.Error())
		os.Exit(2)
	}
	store := cookie.NewStore([]byte("secret"))
	ginEngine.Use(gin.Recovery())
	ginEngine.Use(sessions.Sessions("mysession-1", store))
	router.Registry(ginEngine)
	ginEngine.Run(":"+config.FileConfig.Port)
}