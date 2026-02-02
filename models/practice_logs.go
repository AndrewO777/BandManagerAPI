package models

import "time"

type UserPracticeLog struct {
	LogID          int       `json:"log_id"`
	UserID         int       `json:"user_id"`
	SongID         int       `json:"song_id"`
	DifficultyRating int8     `json:"difficulty_rating" gorm:"check:difficulty_rating BETWEEN 1 AND 10"`
	LastPracticedAt *time.Time `json:"last_practiced_at"`
}

type BandPracticeLog struct {
	LogID          int       `json:"log_id"`
	BandID         int       `json:"band_id"`
	SongID         int       `json:"song_id"`
	LastPracticedAt *time.Time `json:"last_practiced_at"`
}

