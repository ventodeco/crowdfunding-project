package handler

import (
	"crowfunding-api/helper"
	"crowfunding-api/travel"
	"net/http"

	"github.com/gin-gonic/gin"
)

type travelHandler struct {
	service travel.Service
}

func NewTravelHandler(service travel.Service) *travelHandler {
	return &travelHandler{service}
}

// api/v1/travels/locations
func (h *travelHandler) GetLocations(c *gin.Context) {
	travels, err := h.service.GetTravelLocations()
	if err != nil {
		response := helper.APIResponse(
			"Failed to get travel location",
			http.StatusBadRequest,
			"error",
			nil,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, travels)
}
