package users_segments

import (
	"avito_task_segments/pkg/common/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func (h handler) GetUserSegments(ctx *gin.Context) {
	body := models.GetUserSegmentsRequestBody{}

	if error := ctx.BindJSON(&body); error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"type": "error", "message": "bad input"})
		return
	}

	var user models.User
	var userSegments []models.Segment

	user.UserId = body.UserId

    // Если пользователя нет в БД, записываем в контекст соответствующий ответ

	if result := h.DB.Where("user_id = ?", user.UserId).First(&user); result.Error != nil {

		if result.Error == gorm.ErrRecordNotFound {

			ctx.JSON(http.StatusNotFound, gin.H{"type": "error", "message": "this user is not exists"})
            return

		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"type": "error", "message": result.Error.Error()})
		return
	}

	h.DB.Model(&user).Association("Segments").Find(&userSegments)

    // Если у пользователя нет сегментов, записываем в контекст соответствующий ответ

	if len(userSegments) == 0 {

		ctx.JSON(http.StatusNoContent, gin.H{"type": "info", "message": "this user has no segments"})
		return

	}

    // Берем только названия (slugs) сегментов

	var userSlugs []string

	for _, v := range userSegments {
		userSlugs = append(userSlugs, v.Slug)
	}

	responseBody := models.GetUserSegmentsResponseBody{"success", user.UserId, userSlugs}
	responseBodyJSON, err := json.Marshal(responseBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"type": "error", "message": "error while creating response"})
        return
	}

	ctx.Data(http.StatusOK, "application/json", responseBodyJSON)

}
