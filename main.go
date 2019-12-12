package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/Hyan/baccarat_server/Game"
)

//websocket的練習 百家樂製作

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	fmt.Println("succeess")

	r := http.NewServeMux()

	r.HandleFunc("/echo", echofunc)
	r.HandleFunc("/", routerfunc)
	r.HandleFunc("/baccarat", baccaratfunc)

	http.ListenAndServe(":8080", r)

	/*
		http.HandleFunc("/echo", echofunc)

		http.HandleFunc("/", routerfunc)

		http.ListenAndServe(":8080", nil)
	*/
}

func echofunc(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)

	for {
		msgType, msg, err := conn.ReadMessage()

		if err != nil { //有錯誤
			return
		}

		fmt.Println("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}

}

func routerfunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
	http.ServeFile(w, r, "websocket.html")
}

//回傳baccarat資料
func baccaratfunc(w http.ResponseWriter, r *http.Request) {

	//r=客戶端的請求 w=服務器做出的回應
	/*
		w.Write([]byte(""))
		fmt.Fprintln(w, "")
	*/

	w.Write(Game.GetInstance().Lottery())

	//應該是沒有去做接口Init的關係
}
