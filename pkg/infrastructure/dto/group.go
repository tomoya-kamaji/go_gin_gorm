package dto

type GroupEntity struct {
	ID    uint `json:"id" binding:"required"`
	Name  string
	Users []UserEntity `gorm:"many2many:group_users;foreignKey:id;joinForeignKey:group_id;References:id;joinReferences:user_id"`
}

func (g *GroupEntity) TableName() string {
	return "groups"
}
