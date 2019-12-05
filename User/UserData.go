package User

import "fmt"

type Data struct {
	ID    int
	Name  string
	Money int
}

var Users map[int]Data

func init() {
	Users = make(map[int]Data, 100)
}

func AddUser(ClientID int, id int, Name string, Money int) bool {

	if _, ok := Users[ClientID]; ok {
		return false
	} else {
		Users[ClientID] = Data{id, Name, Money}
		return true
	}

}

func DelUser(ClientID int) {

	delete(Users, ClientID)
}

func ContainsUser(ClientID int) bool {

	_, ok := Users[ClientID]
	return ok
}

func AllUsersData() {
	for k, v := range Users {
		fmt.Println("clientID:", k, " Data:", v)
	}
}
