package user

type User struct {
	Id              UserId           `json:"userId"`
	Name            UserName         `json:"userName"`
}

func NewUser(id UserId, name UserName) (*User) {
	return &User{Id: id, Name: name}
}

type Users []User