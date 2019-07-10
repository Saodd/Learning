package learnWeb

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var quitSemaphore chan bool

func Main_ChatroomClient() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "192.168.1.213:9999")

	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	fmt.Println("connected!")

	go onMessageRecived(conn)

	// 控制台聊天功能加入
	for {
		var msg string
		myScanln(&msg)
		if msg == "quit" {
			break
		}
		b := []byte(msg + "\n")
		conn.Write(b)
	}
	<-quitSemaphore
}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println(msg)
		if err != nil {
			quitSemaphore <- true
			break
		}
	}
}

func myScanln(a *string) {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	*a = string(data)
}
