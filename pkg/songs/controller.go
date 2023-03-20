package songs

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	router.GET("/song", h.GetSongs)
	router.GET("/song/:id", h.GetSongById)
	router.POST("/song", h.CreateSong)
	router.PUT("/song/:id", h.UpdateSong)
	router.DELETE("/song/:id", h.DeleteSong)

}