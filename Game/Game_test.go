package Game

import (
	"fmt"
	"testing"
)

func TestTableGame(t *testing.T) {
	fmt.Println(string(GetInstance().Lottery()))
}
