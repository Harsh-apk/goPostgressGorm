package types

type User struct {
	ID        uint   `gorm:"primary key;autoIncrement" json:"id"`
	Name      string `gorm:"name" json:"name"`
	Email     string `gorm:"email;unique" json:"email"`
	EmailAuth bool   `gorm:"email_auth" json:"emailAuth"`
}
