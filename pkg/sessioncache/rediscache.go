package sessioncache

import (
	"context"
	"errors"
	"strconv"
	"time"
	"github.com/go-redis/redis/v8"
	"encoding/json"
)



type RedisCache struct {
	p        *redis.Client // redis connection pool
	conninfo string
	dbNum    int
	username string
	password string
	Ctx context.Context
}

func NewRedisCache() (Cache){
	ctx := context.Background()
	return &RedisCache{Ctx: ctx}
}

func (r *RedisCache) SetValue(k string,v interface{},d time.Duration)(error){
	err := r.p.SetNX(r.Ctx,k,v,d).Err()
	return err
}

func (r *RedisCache) Get(k string)(interface{},bool){
	res,err := r.p.Get(r.Ctx,k).Result()
	if err != nil {
		return res,false
	}
	return res,true
}

func (r *RedisCache) StartService(config string) (error){
	//格式： "{'addr':'127.0.0.1:6379','dbnum':1,'username':'test','password':'password'}"
	var cf map[string]interface{}
	err := json.Unmarshal([]byte(config), &cf)
	if err != nil {
		return errors.New("格式如下： {'addr':'127.0.0.1:6379','dbnum':1,'username':'test','password':'password'}")
	}
	var conninfo string
	var dbnum int
	//连接信息
	if conninfo ,ok := cf["addr"]; ok {
		conninfo = conninfo.(string)
	}else{
		conninfo = "127.0.0.1:6379"
	}
	r.conninfo = conninfo
	if dbnum,ok := cf["dbnum"];ok {
		dbnumstring,err:=strconv.Atoi(dbnum.(string))
		if err != nil {
			dbnum = 0
		}else{
			dbnum = dbnumstring
		}
	}else{
		dbnum = 0
	}
	r.dbNum = dbnum
	if username,ok := cf["username"]; ok {
		r.username = username.(string)
	}
	if password,ok := cf["password"]; ok {
		r.password = password.(string)
	}else{
		r.password = ""
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     r.conninfo,
		Password: r.password, // no password set
		DB:       r.dbNum,  // use default DB
	})
	r.p = rdb
	return nil
}

func init() {
	Register("redis",NewRedisCache)
}

