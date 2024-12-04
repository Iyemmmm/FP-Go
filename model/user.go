package model

type User struct {
    ID        uint   `json:"iduser" gorm:"column:iduser;primaryKey;autoIncrement"`
    Nama      string `json:"nama" gorm:"not null"`
    Email     string `json:"email" gorm:"not null;unique"`
    Password  string `json:"-" gorm:"not null"`
    NoHP      string `json:"nohp" gorm:"column:nohp;not null"`  
}

func (User) TableName() string {
    return "user"
}
