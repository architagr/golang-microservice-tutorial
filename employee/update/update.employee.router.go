package update

import "github.com/gin-gonic/gin"

type Router struct {
	updateHandler *Handler
}
func InitRouter(h *Handler) *Router{
	return &Router{
		updateHandler: h,
	}
}
func (r *Router) RegisterRoutes(group *gin.RouterGroup) {
	group.PUT("/:id", r.updateHandler.Update)
}
