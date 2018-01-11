package models

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"golang.org/x/crypto/ssh"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/go-sql-driver/mysql"
)

var (
	serverIP    string
	sshUser     string
	sshPassword string
)

func init() {
	dburl := beego.AppConfig.String("dburl")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbName := beego.AppConfig.String("db")

	serverIP = beego.AppConfig.String("server_ip")
	sshUser = beego.AppConfig.String("ssh_user")
	sshPassword = beego.AppConfig.String("ssh_password")

	//注册mysql Driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	var conn string
	external_ip := get_external_ip()
	if external_ip != serverIP {
		mysql.RegisterDial("mysql+tcp", tcpTransferDial)
		conn = dbuser + ":" + dbpassword + "@mysql+tcp(" + dburl + ")/" + dbName + "?charset=utf8"
		fmt.Println("Foreigner ip: ", external_ip)
	} else {
		conn = dbuser + ":" + dbpassword + "@tcp(" + dburl + ")/" + dbName + "?charset=utf8"
		fmt.Println("Localhost ip")
	}

	err := orm.RegisterDataBase("default", "mysql", conn)

	if err != nil {
		panic(err)
	}
	fmt.Println("database connect success !")

}

func get_external_ip() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		panic("Failed get_external_ip: " + err.Error())
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	return string(data)
}

func tcpTransferDial(addr string) (net.Conn, error) {
	// An SSH client is represented with a ClientConn. Currently only
	// the "password" authentication method is supported.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig.
	config := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshPassword),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", serverIP), config)
	if err != nil {
		panic("Failed to dial: " + serverIP + err.Error())
	}
	mysqlConn, err := client.Dial("tcp", addr)
	if err != nil {
		panic(fmt.Sprintf("client.DialTCP error: %s", err.Error()))
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/whoami"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
	return mysqlConn, nil
}
