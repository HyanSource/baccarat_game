package User

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	fmt.Println("測試")

	AddUser(1, 1, "H", 10000)
}

func TestAll(t *testing.T) {
	AddUser(1, 1, "H", 10000)
	AddUser(2, 2, "A", 999)

	DelUser(1)

	AllUsersData()
}
