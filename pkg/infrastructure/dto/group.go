package dto

type GroupEntity struct {
	ID   uint `json:"id" binding:"required"`
	Name string
}

func (g *GroupEntity) TableName() string {
	return "groups"
}
