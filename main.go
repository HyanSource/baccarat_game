package main

import (
	"fmt"
	"net/http"

	"github.com/garyburd/redigo/redis"

	"github.com/gorilla/websocket"

	"github.com/HyanSource/baccarat_server/Game"
	"github.com/HyanSource/baccarat_server/RedisPool"
)

//websocket的練習 百家樂製作

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/echo", echofunc)
	r.HandleFunc("/", routerfunc)
	r.HandleFunc("/baccarat", baccaratfunc)
	r.HandleFunc("/baccaratlog", baccaratlogfunc)

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
	//http.ServeFile(w, r, "websocket.html")//尚未完成20191226
}

//回傳baccarat資料
func baccaratfunc(w http.ResponseWriter, r *http.Request) {

	//r=客戶端的請求 w=服務器做出的回應
	/*
		w.Write([]byte(""))
		fmt.Fprintln(w, "")
	*/
	var t = Game.GetInstance().Lottery()
	w.Write(t)
	fmt.Println(string(t))

	//寫入紀錄
	conn := RedisPool.Pool.Get()
	defer conn.Close()
	_, err := redis.Int(conn.Do("LPUSHX", "baccaratlog", string(t)))

	if err != nil {
		fmt.Println(err)
		return
	}

}

func baccaratlogfunc(w http.ResponseWriter, r *http.Request) {

	conn := RedisPool.Pool.Get()
	defer conn.Close() //釋放資源回連接池

	m, err := redis.Strings(conn.Do("LRange", "baccaratlog", 0, 9)) //strings是切片

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range m {
		fmt.Fprintln(w, v)
	}

}
