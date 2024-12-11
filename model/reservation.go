package model

type Reservation struct {
	ID      uint   `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	ID_User uint   `json:"-" gorm:"not null"`
	User    User   `json:"author" gorm:"foreignkey:id_user;references:iduser;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Nama    string `json:"nama" gorm:"not null"`
	Email   string `json:"email" gorm:"not null"`
	NoHP    string `json:"nohp" gorm:"column:nohp;not null"`
	Guest   int    `json:"guest" gorm:"not null"`
	Date    string `json:"date" gorm:"not null"`
	Waktu   string `json:"waktu" gorm:"not null"`
	Note    string `json:"note"`
	Status  string `json:"status" gorm:"not null"`
}

func (Reservation) TableName() string {
	return "reservation"
}
