package Game

import (
	"testing"
)

/*
func TestGame(t *testing.T) {

	a := new(GamePlay)
	a.Init()

	fmt.Println("測試成功")
}
*/
/*
func TestPoint(t *testing.T) {

	a := new(GamePlay)

	fmt.Println("莊家:", a.GetDealerCards())
	fmt.Println("閒家:", a.GetPlayerPoints())

}
*/
/*
func TestCardCount(t *testing.T) {
	a := new(GamePlay)
	a.Init()
	fmt.Println(a.GetCardCount())
}
*/

func TestPlay(t *testing.T) {
	a := new(GamePlay)
	a.Init()

	a.Lottery()
	/*
		fmt.Println(a.GetPlayerPoints())
		fmt.Println(a.GetDealerPoints())
		fmt.Println(a.GetCardCount())
		fmt.Println("測試成功2")
	*/
}

func TestShuffle(t *testing.T) {
	a := new(GamePlay)
	a.Init()
	a.Shuffle()

}

func TestDraw(t *testing.T) {
	a := new(GamePlay)
	a.Init()
	a.Shuffle()
}
