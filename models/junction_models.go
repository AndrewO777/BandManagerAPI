package models

import "time"

type UserBand struct {
	UserID  int `json:"user_id"`
	BandID  int `json:"band_id"`
	JoinedAt time.Time `json:"joined_at"`
}

type SetlistSong struct {
	SetlistID int `json:"setlist_id"`
	SongID    int `json:"song_id"`
	Position  int `json:"position"`
}
