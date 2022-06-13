package handler

import (
	"crowfunding-api/helper"
	"crowfunding-api/travel"
	"net/http"
	"strconv"

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

// api/v1/travels/locations
func (h *travelHandler) GetFavoriteTravel(c *gin.Context) {
	travels, err := h.service.GetFavoriteTravel()
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

// api/v1/travels/locations/:id/update
func (h *travelHandler) UpdateLike(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
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
	like, err := strconv.ParseBool(c.Request.URL.Query()["like"][0])

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

	travel, err := h.service.UpdateTravel(id, like)
	if err != nil {
		response := helper.APIResponse(
			"Failed to update travel",
			http.StatusBadRequest,
			"error",
			nil,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, travel)
}
