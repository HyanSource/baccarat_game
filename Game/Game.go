package Game

import (
	"fmt"
	"sync"

	"github.com/HyanSource/baccarat_server/Cards"
	"github.com/HyanSource/baccarat_server/Player"
	"github.com/HyanSource/baccarat_server/Result"
	"github.com/HyanSource/baccarat_server/Rule"
)

var (
	instance *GameTable
	once     sync.Once
)

type GameTable struct {
	CardManage Cards.ICardManage //管理牌的接口
}

func GetInstance() *GameTable {

	//只會執行一次的方法
	once.Do(func() {
		instance = new(GameTable)
		instance.CardManage = new(Cards.TableCard)
		fmt.Println("初始化成功")
	})

	return instance
}

type Play interface {
	Lottery() []byte //得到開獎資料的方法
}

//開獎的方法
func (t *GameTable) Lottery() []byte {

	t.CardManage.Init()
	t.CardManage.Shuffle()

	var playerget = new(Player.CardStruct)
	var dealerget = new(Player.CardStruct)

	//初始化
	playerget.Init()
	dealerget.Init()

	//放入前兩張牌

	playerget.SetCards(t.CardManage.Draw())
	dealerget.SetCards(t.CardManage.Draw())
	playerget.SetCards(t.CardManage.Draw())
	dealerget.SetCards(t.CardManage.Draw())

	//放入閒家第三張牌
	if Rule.Player_Draw(playerget.GetPoints(2)) {
		playerget.SetCards(t.CardManage.Draw())
	}

	if Rule.Dealer_Draw(playerget.GetPoints(3), dealerget.GetPoints(2)) {
		dealerget.SetCards(t.CardManage.Draw())
	}

	//三種情況 4張 5張 6張

	R := &Result.Result{PlayerPoints: playerget.GetPoints(3), DealerPoints: dealerget.GetPoints(3), PlayerCards: playerget.Cards, DealerCards: dealerget.Cards, Win: Rule.WinType(playerget.GetPoints(3), dealerget.GetPoints(3)), Desk: t.CardManage.CardsCount()}

	return R.GameMarshal()
}
