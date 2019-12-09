package Rule

import (
	"strconv"
)

//當 AB裡面有耦合時 編譯會沒辦法通過

//換算點數
func Point(cardnum int) int {

	var ReturnNum = cardnum%13 + 1
	if ReturnNum > 10 {
		return 10
	}
	return ReturnNum
}

//閒家抽第三張牌的方法
func Player_Draw(PlayerNowPoint int) bool {

	if PlayerNowPoint != 2 {
		panic("不是第二張牌的情況")
	}

	switch PlayerNowPoint {
	case 0, 1, 2, 3, 4, 5:
		return true
	case 6, 7, 8, 9:
		return false
	default:
		panic("抽第三張牌點數錯誤 points:" + strconv.Itoa(PlayerNowPoint))
	}
}

//莊家抽第三張牌的方法
func Dealer_Draw(PlayerNowPoint int, DealerNowPoint int) bool {

	if DealerNowPoint != 2 {
		panic("不是第二張牌的情況")
	}

	switch DealerNowPoint {
	case 0, 1, 2:
		return true
	case 3:

		switch PlayerNowPoint {
		case 8:
			return false
		case 0, 1, 2, 3, 4, 5, 6, 7, 9:
			return true
		}

	case 4:
		switch PlayerNowPoint {
		case 0, 1, 8, 9:
			return false
		case 2, 3, 4, 5, 6, 7:
			return true

		}
	case 5:
		switch PlayerNowPoint {
		case 0, 1, 2, 3, 8, 9:
			return false
		case 4, 5, 6, 7:
			return true
		}
	case 7, 8, 9:
		return false
	}
	panic("抽第三排點數錯誤 Playerpoints:" + strconv.Itoa(PlayerNowPoint) + " DealerPoints:" + strconv.Itoa(DealerNowPoint))
}

func WinType(PlayerPoint int, DealerPoint int) string {

	if PlayerPoint > DealerPoint {
		return "閒家勝"
	} else if PlayerPoint < DealerPoint {
		return "莊家勝"
	} else if PlayerPoint == DealerPoint {
		return "平手"
	} else {
		return ""
	}

}
