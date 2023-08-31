package users_segments

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler {
		DB: db,
	}

	routes := router.Group("/users_segments")
	routes.POST("/change", h.ChangeUserSegments)
	routes.POST("/get", h.GetUserSegments)
}
