package segments

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

	routes := router.Group("/segment")
	routes.POST("/add", h.AddSegment)
	routes.POST("/delete", h.DeleteSegment)
}
