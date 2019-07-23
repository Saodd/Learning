package learnWeb

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var wssConnmap map[string]*websocket.Conn

func Main_WebSocketServer() {
	http.HandleFunc("/echo", wsAccept)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "C:/Users/lewin/github/Learning/Golang/src/learnWeb/20190710-WebSocketClient.html")
	})
	http.ListenAndServe("0.0.0.0:9999", nil)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsAccept(w http.ResponseWriter, r *http.Request) {
	fmt.Println("connect from: ", r.RemoteAddr)

	//conn, _ := upgrader.Upgrade(w,r,nil)

	fmt.Println(r.BasicAuth())



}

func wssBroadcast(s []byte) {
	//for _, conn := range wssConnmap {
	//	fmt.Println(string(s))
	//	conn.Write(s)
	//}
}


