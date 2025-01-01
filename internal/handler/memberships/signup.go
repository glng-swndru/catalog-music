package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glng-swndru/catalog-music/internal/models/memberships"
)

func (h *Handler) SignUp(c *gin.Context) {
	var req memberships.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.SignUp(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}
