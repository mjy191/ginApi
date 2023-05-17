package models

type Phone struct {
	Id     int    `json:"id" gorm:"column:id;primaryKey"`
	Phone  string `json:"phone" gorm:"column:phone;default:null"`
	UserId int    `json:"userId" gorm:"column:userId;default:null"`
}

func (Phone) TableName() string {
	return "phone"
}
