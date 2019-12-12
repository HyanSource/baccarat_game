package User

import (
	"fmt"
	"sync"
)

type Data struct {
	ID    int
	Name  string
	Money int
}

/////////////////////////////////

var (
	Instance *UsersInstance
	Once     sync.Once
)

type UsersInstance struct {
	Users map[int]Data
}

func GetInstance() *UsersInstance {

	Once.Do(func() {
		Instance.Users = make(map[int]Data, 100)
	})

	return Instance
}

func (t *UsersInstance) AddUser(ClientID int, id int, Name string, Money int) bool {

	if _, ok := t.Users[ClientID]; ok {
		return false
	} else {
		t.Users[ClientID] = Data{id, Name, Money}
		return true
	}

}

func (t *UsersInstance) DelUser(ClientID int) {

	delete(t.Users, ClientID)
}

func (t *UsersInstance) ContainsUser(ClientID int) bool {

	_, ok := t.Users[ClientID]
	return ok
}

func (t *UsersInstance) AllUsersData() {
	for k, v := range t.Users {
		fmt.Println("clientID:", k, " Data:", v)
	}
}
