package Game

import (
	"fmt"
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

	fmt.Println("初始化成功")
}

//取得目前點數
func (t *CardStruct) GetPoints(count int) int {
	var RetrurnPoints int

	if count < 0 && count > len(t.Cards) {
		panic("取得目前點數錯誤 count:" + strconv.Itoa(count))
	}
	for i := 0; i < count; i++ {
		RetrurnPoints += Point(t.Cards[i])
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

	return Point(t.Cards[count])
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

//換算點數
func Point(cardnum int) int {

	var ReturnNum = cardnum%13 + 1
	if ReturnNum > 10 {
		return 10
	}
	return ReturnNum
}

//閒家抽第三張牌的方法
func Player_Draw(p ICards) bool {

	if p.GetCardCount() != 2 {
		panic("不是第二張牌的情況 cardcount:" + strconv.Itoa(p.GetCardCount()))
	}

	switch p.GetPoints(2) {
	case 0, 1, 2, 3, 4, 5:
		return true
	case 6, 7, 8, 9:
		return false
	default:
		panic("抽第三張牌點數錯誤 points:" + strconv.Itoa(p.GetPoints(2)))
	}
}

//莊家抽第三張牌的方法
func Dealer_Draw(p ICards, d ICards) bool {

	if d.GetCardCount() != 2 {
		panic("不是第二張牌的情況 cardcount:" + strconv.Itoa(d.GetCardCount()))
	}

	switch d.GetPoints(2) {
	case 0, 1, 2:
		return true
	case 3:

		switch p.GetPoints(2) {
		case 8:
			return false
		default:
			return true
		}

	case 4:
		switch p.GetPoints(2) {
		case 0, 1, 8, 9:
			return false
		default:
			return true

		}
	case 5:
		switch p.GetPoints(2) {
		case 0, 1, 2, 3, 8, 9:
			return false
		default:
			return true

		}
	case 7, 8, 9:
		return false
	default:
		panic("抽第三排點數錯誤 points:" + strconv.Itoa(d.GetPoints(2)))
	}
}
