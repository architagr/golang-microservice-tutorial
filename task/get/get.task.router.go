package get

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	getHandler *Handler
}

func InitRouter(h *Handler) *Router{
	return &Router{
		getHandler: h,
	}
}

func (r *Router) RegisterRoutes(group *gin.RouterGroup) {
	group.GET("", r.getHandler.GetAll)
	group.GET("/:id", r.getHandler.GetById)
}
