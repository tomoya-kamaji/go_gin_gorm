package group

type Group struct {
	Id   GroupId `json:"groupId"`
	Name string  `json:"groupName"`
}

func NewUser(id GroupId, name string) *Group {
	return &Group{Id: id, Name: name}
}
