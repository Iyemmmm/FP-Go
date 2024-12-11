package model

type Menu struct {
	ID    uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Nama  string  `json:"nama" gorm:"not null"`
	Img   string  `json:"img" gorm:"not null"`
	Harga float64 `json:"harga" gorm:"not null"`
}

func (Menu) TableName() string {
	return "menu"
}
