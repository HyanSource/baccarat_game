package RedisPool

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	fmt.Println("開始")
	v, err := Get("t") // Get->抓1個值 其他的應該要不同method
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(v))
	fmt.Println("結束")
}
