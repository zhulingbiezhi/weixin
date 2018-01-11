package tcpTransfer

import (
	"bytes"
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"
)

//func main() {
//	dstConn, err := sshConnect()
//	if err != nil {
//		panic(fmt.Sprintf("sshConnect error: %s", err.Error()))
//	}
//	serv, err := net.Listen("tcp", "127.0.0.1:3307")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer serv.Close()
//	for {
//		conn, err := serv.Accept()
//		if err != nil {
//			fmt.Println("建立连接错误:%v\n", err)
//			continue
//		}
//		defer conn.Close()
//		fmt.Println(conn.RemoteAddr(), conn.LocalAddr())
//		fmt.Println(dstConn.RemoteAddr(), dstConn.LocalAddr())
//		ExitChan := make(chan bool, 1)
//		go func(srcConn net.Conn, dstConn net.Conn, Exit chan bool) {
//			_, err := io.Copy(dstConn, srcConn)
//			fmt.Printf("往服务器发送数据失败:%v\n", err)
//			ExitChan <- true
//		}(conn, dstConn, ExitChan)

//		go func(srcConn net.Conn, dstConn net.Conn, Exit chan bool) {
//			_, err := io.Copy(srcConn, dstConn)
//			fmt.Printf("从服务器接收数据失败:%v\n", err)
//			ExitChan <- true
//		}(conn, dstConn, ExitChan)
//		<-ExitChan
//		dstConn.Close()
//	}
//}

func TcpTransferDial(addr string) (net.Conn, error) {
	// An SSH client is represented with a ClientConn. Currently only
	// the "password" authentication method is supported.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig.
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("h6100210050H"),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	client, err := ssh.Dial("tcp", "120.79.27.53:22", config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
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
