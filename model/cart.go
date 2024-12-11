package model

type Cart struct {
	ID       uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	ID_User  uint    `json:"-" gorm:"not null"`
	User     User    `json:"author" gorm:"foreignkey:id_user;references:iduser;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	ID_Menu  uint    `json:"-" gorm:"not null"`
	Menu     Menu    `json:"menu" gorm:"foreignkey:id_menu;references:id;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Quantity int     `json:"quantity" gorm:"not null"`
	Total    float64 `json:"total" gorm:"not null"`
	Note     string  `json:"note"`
	Status   string  `json:"status" gorm:"not null"`
}

func (Cart) TableName() string {
	return "cart"
}
