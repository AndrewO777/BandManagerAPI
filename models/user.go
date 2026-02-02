package models

type User struct {
	UserID     int       `json:"user_id"`
	Username   string    `json:"username" gorm:"unique;not null"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Bio        string    `json:"bio"`
}
