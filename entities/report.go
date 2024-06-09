package entities

type Report struct {
	ID   int    `gorm:"primary_key:auto_increment" json:"id"`
	Name string `json:"name"`
	User User   `json:"user"`
}
