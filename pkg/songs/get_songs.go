package songs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/magnojunior07/golang-crud/pkg/common/models"
)

func (h handler) GetSongs(ctx *gin.Context) {
	var songs []models.CharlieBrownJrSongs
	if result := h.DB.Find(&songs); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	
	ctx.JSONP(http.StatusOK, &songs)
}