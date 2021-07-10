package add

import "github.com/gin-gonic/gin"

type Router struct {
	addHandler *Handler
}
func InitRouter(h *Handler) *Router{
	return &Router{
		addHandler: h,
	}
}
func (r *Router) RegisterRoutes(group *gin.RouterGroup) {
	group.POST("", r.addHandler.Add)
}
