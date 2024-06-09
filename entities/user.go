package entities

type User struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	FirstName string `gorm:"type:varchar(255)" json:"first_name"`
	LastName  string `gorm:"type:varchar(255)" json:"last_name"`
	UserName  string `gorm:"type:varchar(255)" json:"name"`
	Password  string `gorm:"not null" json:"password"`
	Email     string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Token     string `gorm:"-" json:"token"`
}
