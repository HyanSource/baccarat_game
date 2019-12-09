package Player

import (
	"20191118User_Manage/Rule"
	"strconv"
)

type ICards interface {
	Init()                       //初始化
	GetPoints(count int) int     //取得目前點數
	GetCardPoints(count int) int //取得指定點數
	GetCardCount() int           //取得目前張數

	SetCards(cardnum int) //放入牌的點數
}

type CardStruct struct {
	Cards [3]int //取得牌
}

//初始化
func (t *CardStruct) Init() {
	for i := 0; i < len(t.Cards); i++ {
		t.Cards[i] = -1
	}
}

//取得目前點數
func (t *CardStruct) GetPoints(count int) int {
	var RetrurnPoints int

	if count < 0 && count > len(t.Cards) {
		panic("取得目前點數錯誤 count:" + strconv.Itoa(count))
	}
	for i := 0; i < count; i++ {
		if t.Cards[i] != -1 {
			RetrurnPoints += Rule.Point(t.Cards[i])
		}

	}
	return (RetrurnPoints % 10)
}

//取得指定點數
func (t *CardStruct) GetCardPoints(count int) int {

	if count < 0 {
		panic("取得指定牌點數錯誤1")
	} else if count >= len(t.Cards) {
		panic("取得指定牌點數錯誤1")
	}

	return Rule.Point(t.Cards[count])
}

//取得目前牌的張數
func (t *CardStruct) GetCardCount() int {
	var ReturnCount int
	for _, k := range t.Cards {
		if k != -1 {
			ReturnCount++
		}
	}

	return ReturnCount
}

//放入牌的方法
func (t *CardStruct) SetCards(cardsnum int) {
	for i := 0; i < len(t.Cards); i++ {
		if t.Cards[i] == -1 {
			t.Cards[i] = cardsnum
			return
		}
	}

	panic("放入牌不成功")
}

////////////////
