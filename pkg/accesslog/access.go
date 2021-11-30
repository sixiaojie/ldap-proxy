package accesslog

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/lestrrat-go/file-rotatelogs"
	"os"
	"path/filepath"
	"time"
)


func LogInit(ginEngine *gin.Engine,loglevel string,format string,path string) (error){
	//这是gin的日志
	ginEngine.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("{\"clientip\":\"%s\",\"time\":\"%s\",\"method\":\"%s\",\"path\":\"%s\",\"proto\":\"%s\",\"statuscode\":\"%d\",\"Latency\":\"%s\",\"agent\":\"%s\",\"Error\":\"%s\"}\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	//这里业务的日志
	if path == "" {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		path = dir
	}
	writer, _ := rotatelogs.New(
		path+".%Y%m%d%H",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(120)*time.Minute),
		rotatelogs.WithRotationTime(time.Duration(60)*time.Minute),
	)
	log.SetOutput(writer)

	if format == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}
	level,err := log.ParseLevel(loglevel)
	if err != nil{
		return err
	}
	log.SetLevel(level)
	return nil
}
