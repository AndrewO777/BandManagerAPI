package models

type Setlist struct {
	SetlistID     int       `json:"setlist_id"`
	BandID        int       `json:"band_id"`
	Name          string    `json:"name" gorm:"not null"`
	Description   string    `json:"description"`
	IsTemplate    bool      `json:"is_template" gorm:"default:false"`
}

