package ConnSQL

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DSN string
)

//DB的設計是用於嘗連線 不該頻繁的Open Close
//在主執行緒執行close 不然就是Mysql自己檢測

//初始化 會在main載入之前
func init() {
	DSN = "user:password@tcp(127.0.0.1:3306)/baccarat"

	db, err := sql.Open("mysql", DSN)//Open並未建立連線 提早準備實例

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	db.SetMaxIdleConns(10) //設定最大空閒連線數
	db.SetMaxOpenConns(10) //設定最大開啟連線數

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("成功")
}

