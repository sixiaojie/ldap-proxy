package accesslog

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func Logger(flow string,loglevel string,ip string,clientagent string,path string,platform string,msg string){
	if loglevel == "trace" {
		log.WithFields(log.Fields{
			"time":time.Now().Format("2006-01-02 15:04:05"),
			"flow":flow,
			"ip":ip,
			"agent":clientagent,
			"urlpath":path,
			"platform":platform,
		}).Trace(msg)
	}else if loglevel == "debug" {
		log.WithFields(log.Fields{
			"time":time.Now().Format("2006-01-02 15:04:05"),
			"flow":flow,
			"ip":ip,
			"agent":clientagent,
			"urlpath":path,
			"platform":platform,
		}).Debug(msg)
	}else if loglevel == "info" {
		log.WithFields(log.Fields{
			"time":time.Now().Format("2006-01-02 15:04:05"),
			"flow":flow,
			"ip":ip,
			"agent":clientagent,
			"urlpath":path,
			"platform":platform,
		}).Info(msg)
	}else if loglevel == "warn" {
		log.WithFields(log.Fields{
			"time":time.Now().Format("2006-01-02 15:04:05"),
			"flow":flow,
			"ip":ip,
			"agent":clientagent,
			"urlpath":path,
			"platform":platform,
		}).Warn(msg)
	}else if loglevel == "error" {
		log.WithFields(log.Fields{
			"time":time.Now().Format("2006-01-02 15:04:05"),
			"flow":flow,
			"ip":ip,
			"agent":clientagent,
			"urlpath":path,
			"platform":platform,
		}).Error(msg)
	}else if loglevel == "fatal" {
		log.WithFields(log.Fields{
			"time":time.Now().Format("2006-01-02 15:04:05"),
			"flow":flow,
			"ip":ip,
			"agent":clientagent,
			"urlpath":path,
			"platform":platform,
		}).Fatal(msg)
	}

}