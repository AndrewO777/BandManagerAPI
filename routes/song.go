package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/AndrewO777/BandPracticeAPI/models"
	"github.com/AndrewO777/BandPracticeAPI/db"
)

func RegisterSongRoutes(r *gin.Engine) {
	r.GET("/songs", GetSongs)
	r.PUT("/songs/:id", UpdateSong)
}

func GetSongs(c *gin.Context) {
	rows, err := db.DB.Query("SELECT * FROM LearnedSongs")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var songs []models.Song
	for rows.Next() {
		var song models.Song
		if err := rows.Scan(&song.ID, &song.SongName, &song.BandName, &song.PlayCount, &song.CurrentConfidence, &song.LastPlayed); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		songs = append(songs, song)
	}
	c.JSON(http.StatusOK, songs)
}

func UpdateSong(c *gin.Context) {
	id := c.Param("id")
	
	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	query := `UPDATE LearnedSongs 
		SET SongName = ?, BandName = ?, PlayCount = ?, CurrentConfidence = ?, LastPlayed = ? 
		WHERE ID = ?`
	
	_, err := db.DB.Exec(query, song.SongName, song.BandName, song.PlayCount, song.CurrentConfidence, song.LastPlayed, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song updated successfully"})
}
