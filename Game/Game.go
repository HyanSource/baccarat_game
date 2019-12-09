package Game

import (
	"20191118User_Manage/Player"
	"20191118User_Manage/Result"
	"20191118User_Manage/Rule"
	"fmt"
	"math/rand"
)

type IGamePlay interface {
	Shuffle()  //洗牌
	Init()     //初始化
	Draw() int //抽牌方法
}

type GamePlay struct {
	AllCards []int //公共牌
}

//抽牌的方法
func (t *GamePlay) Draw() int {
	if len(t.AllCards) <= 0 {
		panic("目前抽牌沒有牌")
	}

	ReturnCard := t.AllCards[0]

	t.AllCards = t.AllCards[1:]

	return ReturnCard
}

//洗牌方法
func (t *GamePlay) Shuffle() {

	for _, k := range t.AllCards {
		temp := k
		r := rand.Intn(len(t.AllCards))
		k = t.AllCards[r] //取隨機
		t.AllCards[r] = temp
	}

}

//初始化牌的方法
func (t *GamePlay) Init() {

	//make
	//append 可新增和
	//copy
	//移除元素 [1:]從1到上限

	//t.AllCards = make([]int, 0, 208)
	fmt.Println(t.AllCards == nil)
	fmt.Println(t.AllCards)
	for i := 0; i < 4; i++ {
		for j := 0; j < 52; j++ {
			t.AllCards = append(t.AllCards, j) //增加
		}
	}

}

//開獎的方法
func (t *GamePlay) Lottery() string {

	t.Init()
	t.Shuffle()

	var playerget = new(Player.CardStruct)
	var dealerget = new(Player.CardStruct)

	//初始化
	playerget.Init()
	dealerget.Init()

	//放入前兩張牌

	playerget.SetCards(t.Draw())
	dealerget.SetCards(t.Draw())
	playerget.SetCards(t.Draw())
	dealerget.SetCards(t.Draw())

	fmt.Println(playerget.Cards)
	fmt.Println(dealerget.Cards)

	//放入閒家第三張牌
	if Rule.Player_Draw(playerget.GetPoints(2)) {
		playerget.SetCards(t.Draw())
	}

	if Rule.Dealer_Draw(playerget.GetPoints(3), dealerget.GetPoints(2)) {
		dealerget.SetCards(t.Draw())
	}

	fmt.Println(playerget.Cards)
	fmt.Println(dealerget.Cards)

	//三種情況 4張 5張 6張

	R := &Result.Result{PlayerPoints: 1, DealerPoints: 1, PlayerCards: [3]int{1, 2, 3}, DealerCards: [3]int{4, 5, 6}, Win: "莊家勝"}

	R.GameMarshal()

	return ""
}
