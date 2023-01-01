package dto

type GroupEntity struct {
	ID    uint `json:"id" binding:"required"`
	Name  string
	Users []UserEntity `gorm:"many2many:group_users;"`
}

func (g *GroupEntity) TableName() string {
	return "groups"
}
