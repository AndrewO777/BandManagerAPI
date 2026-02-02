package models

type Band struct {
	BandID       int       `json:"band_id"`
	Name         string    `json:"name" gorm:"not null"`
	Description  string    `json:"description"`
}
