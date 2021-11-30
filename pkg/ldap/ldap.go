package ldap

import (
	"crypto/tls"
	"errors"
	"fmt"
	"gopkg.in/ldap.v2"
	"ldap-proxy/mvc/config"
)

//LdapServer
type LdapServer struct {
	Conn  *ldap.Conn
}

type LDAP_RESULT struct {
	DN         string              `json:"dn"`
	Attributes map[string][]string `json:"attributes"`
}

func (lc *LdapServer) Close() {
	if lc.Conn != nil {
		lc.Conn.Close()
		lc.Conn = nil
	}
}


func (lc *LdapServer) Connect() (err error) {
	if config.FileConfig.Ldap.Tls {
		lc.Conn,err = ldap.DialTLS(
			"tcp",
			config.FileConfig.Ldap.Addr,
			&tls.Config{
				InsecureSkipVerify: true,
			})
	}else{
		lc.Conn,err = ldap.Dial("tcp",config.FileConfig.Ldap.Addr)
	}
	if err != nil {
		return err
	}
	if ! config.FileConfig.Ldap.Tls && config.FileConfig.Ldap.StartTLS {
		err = lc.Conn.StartTLS(&tls.Config{InsecureSkipVerify: true})
		if err != nil {
			lc.Conn.Close()
			return err
		}
	}
	err = lc.Conn.Bind(config.FileConfig.Ldap.BindDn,config.FileConfig.Ldap.BindPass)
	if err != nil {
		lc.Conn.Close()
		return err
	}
	return err
}

func (lc *LdapServer) Auth(username,password string) (err error) {
	searchRequest := ldap.NewSearchRequest(
		config.FileConfig.Ldap.BaseDn,
		ldap.ScopeWholeSubtree,ldap.NeverDerefAliases,0,0,false,
		fmt.Sprintf(config.FileConfig.Ldap.AuthFilter,username),
		[]string{"uid", "cn", "mail"},nil,
		)
	sr,err := lc.Conn.Search(searchRequest)
	if err != nil {
		return err
	}
	if len(sr.Entries) == 0 {
		err = errors.New("cannot find such user")
		return
	}
	if len(sr.Entries) > 1 {
		err = errors.New("multi users in search")
		return
	}
	err = lc.Conn.Bind(sr.Entries[0].DN, password)
	if err != nil {
		return
	}
	//Rebind as the search user for any further queries
	err = lc.Conn.Bind(config.FileConfig.Ldap.BindDn, config.FileConfig.Ldap.BindPass)
	if err != nil {
		return
	}
	return
}

func NewLdapServer() *LdapServer {
	return &LdapServer{}
}


func LdapAuth( username, password string) (err error) {
	lc := NewLdapServer()
	err = lc.Connect()
	defer lc.Close()

	if err != nil {
		return
	}
	err = lc.Auth(username, password)
	return

}