package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/AndrewO777/BandPracticeAPI/models"
)

func RegisterSongRoutes(r *gin.Engine) {
	r.GET("/songs", GetSongs)
	r.PUT("/songs/:id", UpdateSong)
}

func GetSongs(c *gin.Context) {
	songs, err := models.GetSongs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
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
	
	err := models.UpdateSong(song, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song updated successfully"})
}
