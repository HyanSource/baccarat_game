package Result

import (
	"encoding/json"
)

//回傳給client的協定
type Result struct {
	//string tag
	PlayerPoints int    `json:"playerpoints"`
	DealerPoints int    `json:"dealerpoints"`
	PlayerCards  [3]int `json:"playercards"`
	DealerCards  [3]int `json:"dealercards"`
	Win          string `json:"win"`
	Desk         int    `json:"desk"`
}

//序列化
func (t *Result) GameMarshal() []byte {
	data, err := json.Marshal(&t)

	if err != nil {
		panic(err)
	}

	return data
}
