package models

import (
	"time"
	"github.com/AndrewO777/BandPracticeAPI/db"
)

type Song struct {
	SongID int `json:"song_id"`
	Title string `json:"title" gorm:"not null"`
	Artist string `json:"artist"`
	Key string `json:"key"`
	TimeSignature string `json:"time_signature"`
	BPM float64 `json:"bpm"`
	Genre string `json:"genre"`
	CreatedAt time.Time `json:"created_at"`
}

func GetSongs() ([]Song, error) {
	var songs []Song
	rows, err := db.DB.Query("SELECT * FROM Songs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var song Song
		if err := rows.Scan(&song.SongID, &song.Title, &song.Artist, &song.Key, &song.TimeSignature, &song.BPM, &song.Genre, &song.CreatedAt); err != nil {
			return nil, err
		}
		songs = append(songs,song)
	}
	return songs, nil
}

func UpdateSong(song Song, id string) (error) { 
	query := `UPDATE LearnedSongs 
		SET Title = ?, Artist = ?, Key = ?, TimeSignature = ?, BPM = ?, Genre = ?
		WHERE ID = ?`
	
	_, err := db.DB.Exec(query, song.Title, song.Artist, song.Key, song.TimeSignature, song.BPM, song.Genre, id)
	if err != nil {
		return err
	}
	return nil
}
