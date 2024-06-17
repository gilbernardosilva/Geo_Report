package entities

type User struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	RoleID    uint64 `gorm:"int" json:"roleID"`
	FirstName string `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName  string `gorm:"type:varchar(255);not null" json:"last_name"`
	UserName  string `gorm:"uniqueIndex;type:varchar(255);not null" json:"username"`
	Password  string `gorm:"not null" json:"password"`
	Email     string `gorm:"uniqueIndex;type:varchar(255);not null" json:"email"`
	Token     string `gorm:"-" json:"token"`
}
