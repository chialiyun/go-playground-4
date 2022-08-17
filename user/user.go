package user

type User struct {
	name  string
	email string
}

func (*User) DoSomething() int {
	return 1 + 2
}
