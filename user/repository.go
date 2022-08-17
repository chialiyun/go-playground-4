package user

import "fmt"

type Repository struct {
}

func (repo Repository) GetUser() {
	fmt.Println("Getting user from db")
}

func (repo Repository) SaveUser(user User) {
	fmt.Println("Inserting new user")
}
