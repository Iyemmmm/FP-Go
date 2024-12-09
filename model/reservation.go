package model

import "time"

type Reservation struct {
	ID      uint      `json:"id" gorm:"column:iduser;primaryKey;autoIncrement"`
	ID_User uint      `json:"-" gorm:"not null"`
	User    User      `json:"author" gorm:"foreignkey:id_user;references:iduser;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Nama    string    `json:"nama" gorm:"not null"`
	Email   string    `json:"email" gorm:"not null"`
	NoHP    string    `json:"nohp" gorm:"column:nohp;not null"`
	Guest	int		  `json:"guest" gorm:"not null"`
	Date    time.Time `json:"date" gorm:"type:date;not null"`
	Waktu   time.Time `json:"waktu" gorm:"type:time;not null"`	
	Note    string    `json:"note"`
}

func (Reservation) TableName() string {
	return "reservation"
}
