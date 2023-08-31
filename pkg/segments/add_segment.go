package segments

import (
	"net/http"
    "gorm.io/gorm"
	"avito_task_segments/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) AddSegment(ctx *gin.Context) {

	body := models.SegmentRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"type": "error", "message": "bad input"})
		return
	}

	var segment models.Segment

	segment.Slug = body.Slug

    // Если сегмент есть в БД, записываем в контекст ответ о том, что он уже существует, иначе, создаем

	if result := h.DB.Where("slug = ?", segment.Slug).First(&segment); result.Error != nil {

        if result.Error == gorm.ErrRecordNotFound {
            segment.Slug = body.Slug
            if create := h.DB.Create(&segment); create.Error != nil {

                ctx.JSON(http.StatusInternalServerError, gin.H{"type": "error", "message": "db error"})
                return        

            }
        } else {

		    ctx.JSON(http.StatusInternalServerError, gin.H{"type": "error", "message": "db error"})
		    return

        }
	} else {

        ctx.JSON(http.StatusConflict, gin.H{"type": "error", "message": "this segment is exists"})
        return

    }

	ctx.JSON(http.StatusCreated, gin.H{"type": "success", "message": "success"})
}
