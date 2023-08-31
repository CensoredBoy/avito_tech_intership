package users_segments

import (
	"avito_task_segments/pkg/common/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
    "time"
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

    // Ищем истекшие сегменты пользователя

    var userSegmentsIds []uint

    for _, v := range userSegments {
        userSegmentsIds = append(userSegmentsIds, v.ID)
    }

    var associations []models.SegmentAssociation

    // Сначала ищем все сегменты пользователя и время их истечения в таблице связи между сегментом и пользователем

    if result := h.DB.Raw("SELECT segment_id, expires FROM user_segment WHERE user_id = ? AND segment_id IN (?)", user.ID, userSegmentsIds).Scan(&associations); result.Error != nil{
        ctx.JSON(http.StatusInternalServerError, gin.H{"type": "error", "message": "db error"})
        return
    }

    // Находим истекшие сегменты пользователя
    // Хеш таблица нужна для того, чтобы запомнить истекшие сегменты и иметь к ним быстрый доступ за О(1) потом

    var userSegmentsExpired []uint
    userSegmentsExpiredMap := make(map[uint]bool)

    currentTime := time.Now().Unix()

    for _, v := range associations{
        if v.Expires != nil {
            if currentTime > v.Expires.Unix() {
                userSegmentsExpired = append(userSegmentsExpired, v.SegmentId)
                userSegmentsExpiredMap[v.SegmentId] = true
            }
        }
    }

    // Удаляем истекшие связи

    if _, err := h.DB.Raw("DELETE FROM user_segment WHERE user_id = ? AND segment_id IN (?)", user.ID, userSegmentsExpired).Rows(); err != nil{
        ctx.JSON(http.StatusInternalServerError, gin.H{"type": "error", "message": "db error"})
        return
    }


    // Берем только названия (slugs) сегментов

	var userSlugs []string

	for _, v := range userSegments {

        // Здесь понадобился быстрый доступ к истекшим сегментам, чтобы лишний раз не делать запрос к БД, фильтруем имеющиеся сегменты

        if !userSegmentsExpiredMap[v.ID] {
		  userSlugs = append(userSlugs, v.Slug)
        }
	}

	responseBody := models.GetUserSegmentsResponseBody{"success", user.UserId, userSlugs}
	responseBodyJSON, err := json.Marshal(responseBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"type": "error", "message": "error while creating response"})
        return
	}

	ctx.Data(http.StatusOK, "application/json", responseBodyJSON)

}
