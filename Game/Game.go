package Game

import (
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
		temp := k //紀錄本身
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

	var Player = new(CardStruct)
	var Dealer = new(CardStruct)

	//初始化
	Player.Init()
	Dealer.Init()

	//放入前兩張牌

	Player.SetCards(t.Draw())
	Dealer.SetCards(t.Draw())
	Player.SetCards(t.Draw())
	Dealer.SetCards(t.Draw())

	fmt.Println(Player.Cards)
	fmt.Println(Dealer.Cards)

	//放入閒家第三張牌
	if Player_Draw(Player) {
		fmt.Println("閒三")
		Player.SetCards(t.Draw())
	}

	if Dealer_Draw(Player, Dealer) {
		fmt.Println("莊三")
		Dealer.SetCards(t.Draw())
	}

	fmt.Println(Player.Cards)
	fmt.Println(Dealer.Cards)

	fmt.Println("閒家點數:", Player.GetPoints(3))
	fmt.Println("莊家點數:", Dealer.GetPoints(3))
	//三種情況 4張 5張 6張

	fmt.Println("開獎成功")

	return "" //回傳格式 閒家*3 莊家*3 閒家點數 莊家點數 中獎位置
}
