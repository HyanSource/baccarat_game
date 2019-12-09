package Result

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestResultMarshal(t *testing.T) {

	//第一種初始化的方法
	R := &Result{PlayerPoints: 1, DealerPoints: 1, PlayerCards: [3]int{1, 2, 3}, DealerCards: [3]int{4, 5, 6}, Win: "莊家勝"}

	//第二種
	//R := new(Result)

	fmt.Println(string(R.GameMarshal()))

	//反序列化
	var data Result
	err := json.Unmarshal(R.GameMarshal(), &data)

	if err != nil {
		panic(err)
	}

	fmt.Println("結果:", data)

}

func TestPlay(t *testing.T) {

}
