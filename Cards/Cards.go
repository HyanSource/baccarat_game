package Cards

import (
	"math/rand"
)

type ICardManage interface {
	Shuffle()  //洗牌
	Init()     //初始化
	Draw() int //抽牌方法
}

type TableCard struct {
	AllCards []int //公共牌
}

//抽牌的方法
func (t *TableCard) Draw() int {
	if len(t.AllCards) <= 0 {
		panic("目前抽牌沒有牌")
	}

	ReturnCard := t.AllCards[0]

	t.AllCards = t.AllCards[1:]

	return ReturnCard
}

//洗牌方法
func (t *TableCard) Shuffle() {

	for _, k := range t.AllCards {
		temp := k
		r := rand.Intn(len(t.AllCards))
		k = t.AllCards[r] //取隨機
		t.AllCards[r] = temp
	}

}

//初始化牌的方法
func (t *TableCard) Init() {

	//make
	//append 可新增和
	//copy
	//移除元素 [1:]從1到上限

	//t.AllCards = make([]int, 0, 208)

	//if len(t.AllCards) < 104 {
	for i := 0; i < 4; i++ {
		for j := 0; j < 52; j++ {
			t.AllCards = append(t.AllCards, j) //增加
		}
	}

	//}

}
