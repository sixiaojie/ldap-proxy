package config

import "ldap-proxy/pkg/sessioncache"

var FileConfig = Config{}

var SessionCache sessioncache.Cache



type Config struct {
	Appname            string `json:"appname"`
	Port			   string `json:"port"`
	Expire             int    `json:"expire"`
	Loglevel           string `json:"loglevel"`
	LogFormat          string `json:"logformat"`
	LogPath            string `json:"logpath"`
	Store              *Store `json:"store"`
	Ldap               *LdapConfig `json:"ldap"`
}

type LdapConfig struct {
	Active             bool    `json:"active"`
	Addr               string  `json:"addr"`
	BaseDn             string  `json:"basedn"`
	BindDn             string  `json:"binddn"`
	BindPass		   string  `json:"bindpass"`
	AuthFilter         string  `json:"authfilter"`
	Tls      		   bool    `json:"tls"`
	StartTLS           bool    `json:"starttls"`
}

type Store struct {
	Name               string  `json:"name"`
	ConnInfo           string  `json:"conninfo"`
}