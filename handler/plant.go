package handler

import (
	"garden-quest/helper"
	"garden-quest/plant"
	"garden-quest/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type plantHandler struct {
	plantService plant.Service
}

func NewPlantHandler(plantService plant.Service) *plantHandler {
	return &plantHandler{plantService}
}

func (h *plantHandler) GetUserPlant(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	plant, err := h.plantService.GetUserPlant(userID)
	if err != nil {
		response := helper.APIResponse("Failed get user plant", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Successfuly get user plant's", http.StatusOK, "success", plant)
	c.JSON(http.StatusOK, response)
}

