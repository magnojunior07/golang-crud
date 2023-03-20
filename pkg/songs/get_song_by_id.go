package songs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/magnojunior07/golang-crud/pkg/common/models"
)

func (h handler) GetSongById(ctx *gin.Context) {
	id := ctx.Param("id")

	var song models.CharlieBrownJrSongs

	if result := h.DB.First(&song, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
	}

	ctx.IndentedJSON(http.StatusOK, &song)
}
