package songs

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/magnojunior07/golang-crud/pkg/common/models"
)

type UpdateSongRequestBody struct {
	SongName string `json:"song_name"`
	Album string `json:"album"`
	Duration string `json:"duration"`
	LinkSpotify string `json:"link_spotify"`
	LinkYouTube string `json:"link_you_tube"`
}

func (h handler) UpdateSong(ctx *gin.Context) {
	id := ctx.Param("id")
	body := UpdateSongRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var song models.CharlieBrownJrSongs

	if result := h.DB.First(&song, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
	}

	song.SongName = body.SongName
	song.Album = body.Album
	song.Duration = body.Duration
	song.LinkSpotify = body.LinkSpotify
	song.LinkYouTube = body.LinkYouTube
	song.UpdatedAt = time.Now()

	h.DB.Save(&song)

	ctx.IndentedJSON(http.StatusOK, &song)

}