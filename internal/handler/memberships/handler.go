package memberships

import (
	"github.com/gin-gonic/gin"
	"github.com/glng-swndru/catalog-music/internal/models/memberships"
)

type service interface {
	SignUp(request memberships.SignupRequest) error
}

type Handler struct {
	*gin.Engine
	service service
}

func NewHandler(api *gin.Engine, service service) *Handler {
	return &Handler{
		api,
		service,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("/memberships")
	route.POST("/sign_up", h.SignUp)
}
