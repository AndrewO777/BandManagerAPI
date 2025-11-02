package models

import "time"

type Song struct {
	ID int `json:"id"`
	SongName string `json:"songName"`
	BandName string `json:"bandName"`
	PlayCount uint16 `json:"playCount"`
	CurrentConfidence byte `json:"currentConfidence"`
	LastPlayed time.Time `json:"lastPlayed"`
}
