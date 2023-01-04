package dto

type UserEntity struct {
	ID   uint `json:"id" binding:"required"`
	Name string
}

func (u *UserEntity) TableName() string {
	return "users"
}
