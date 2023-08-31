package segments

import (
	"avito_task_segments/pkg/common/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
    "gorm.io/gorm"
	"net/http"
)

func (h handler) DeleteSegment(ctx *gin.Context) {

	body := models.SegmentRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"type": "error", "message": "bad input"})
		return
	}

	var segment models.Segment

    // Если сегмента нет в БД, записываем в контекст ответ о том, что такого сегмента нет

	if result := h.DB.Where("slug = ?", body.Slug).First(&segment); result.Error != nil {

		if result.Error == gorm.ErrRecordNotFound {

			ctx.JSON(http.StatusNotFound, gin.H{"type": "error", "message": "there is no segment with this slug"})
            return

		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"type": "error", "message": "db error"})
		return

	}

    // Если сегмент есть в БД и все хорошо, удаляем его

	h.DB.Select(clause.Associations).Delete(&segment)

	ctx.JSON(http.StatusOK, gin.H{"type": "success", "message": "success"})
}
