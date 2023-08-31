package users_segments

import (
	"avito_task_segments/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h handler) ChangeUserSegments(ctx *gin.Context) {

	body := models.ChangeUserSegmentsRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"type": "error", "message": "bad input"})
		return
	}

	var user models.User

	user.UserId = body.UserId

	// Если пользователя нет в БД, создаем его, иначе берем существующего

	if result := h.DB.Where("user_id = ?", user.UserId).FirstOrCreate(&user); result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"type": "error", "message": "db error"})
		return
	}

	var slugsToAdd []models.Segment
	var slugsToDelete []models.Segment
	var userSlugs []models.Segment

	// Выбираем существующие в БД сегменты, переданные нам в запросе для добавления и удаления

	if result := h.DB.Where("slug IN ?", body.SlugsToAdd).Find(&slugsToAdd); result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"type": "error", "message": "db error"})
		return
	}

	if result := h.DB.Where("slug IN ?", body.SlugsToDelete).Find(&slugsToDelete); result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"type": "error", "message": "db error"})
		return
	}

	// Выбираем сегменты, которые уже есть у пользователя

	h.DB.Model(&user).Association("Segments").Find(&userSlugs)

	// Вычисляем пересечение существующих в БД сегментов для добавления и удаления
	// Если мощность их пересечения > 0 то записываем в контекст соответствующий ответ

	interception := make([]models.Segment, 0)
	hash := make(map[uint]bool)

	for _, v := range slugsToAdd {
		hash[v.ID] = true
	}

	for _, v := range slugsToDelete {
		if hash[v.ID] {
			interception = append(interception, v)
		}
	}

	if len(interception) != 0 {
		ctx.JSON(http.StatusConflict, gin.H{"type": "error", "message": "segments to add are found in segments to delete"})
		return
	}

	// Выбираем сегменты для добавления пользователю такие, которых у него еще нет, чтобы не добавлять по второму кругу

	slugsToAddUnique := make([]models.Segment, 0)

	hash = make(map[uint]bool)

	for _, v := range userSlugs {
		hash[v.ID] = true
	}

	for _, v := range slugsToAdd {
		if !hash[v.ID] {
			slugsToAddUnique = append(slugsToAddUnique, v)
		}
	}

	h.DB.Model(&user).Association("Segments").Append(slugsToAddUnique)
	h.DB.Model(&user).Association("Segments").Delete(slugsToDelete)

	ctx.JSON(http.StatusCreated, gin.H{"type": "success", "message": "success"})

}
