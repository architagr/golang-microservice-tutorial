package delete

import "github.com/gin-gonic/gin"

type Router struct {
	deleteHandler *Handler
}
func InitRouter(h *Handler) *Router{
	return &Router{
		deleteHandler: h,
	}
}
func (r *Router) RegisterRoutes(group *gin.RouterGroup) {
	group.DELETE("/:id", r.deleteHandler.Delete)
}
