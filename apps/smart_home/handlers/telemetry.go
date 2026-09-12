package handlers

import (
	"net/http"
	"smarthome/services"
	"github.com/gin-gonic/gin"
)

type TelemetryHandler struct {
	TelemetryService *services.TelemetryService
}

type CreateTelemetryRequest struct {
	DeviceID string `json:"deviceId" binding:"required" example:"device-1"`
	Type     string `json:"type" binding:"required" example:"temperature"`
	Value    string `json:"value" binding:"required" example:"23.5"`
}

func NewTelemetryHandler(telemetryService *services.TelemetryService) *TelemetryHandler {
	return &TelemetryHandler{
		TelemetryService: telemetryService,
	}
}

func (h *TelemetryHandler) RegisterRoutes(router *gin.RouterGroup) {
	telemetry := router.Group("/telemetry")
	{
		telemetry.GET("", h.GetTelemetry)
		telemetry.POST("", h.CreateTelemetry)
	}
}

// GetTelemetry godoc
// @Summary Получить телеметрию
// @Description Возвращает данные телеметрии устройств из Telemetry Service
// @Tags Telemetry
// @Produce json
// @Success 200 {array} services.TelemetryResponse
// @Failure 500 {object} ErrorResponse
// @Router /telemetry [get]
func (h *TelemetryHandler) GetTelemetry(c *gin.Context) {
	telemetry, err := h.TelemetryService.GetTelemetry()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, telemetry)
}

// CreateTelemetry godoc
// @Summary Сохранить телеметрию
// @Description Передаёт новые данные телеметрии в Telemetry Service
// @Tags Telemetry
// @Accept json
// @Produce json
// @Param request body CreateTelemetryRequest true "Данные телеметрии"
// @Success 201 {object} services.TelemetryResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /telemetry [post]
func (h *TelemetryHandler) CreateTelemetry(c *gin.Context) {
	var request CreateTelemetryRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	telemetry, err := h.TelemetryService.CreateTelemetry(
		request.DeviceID,
		request.Type,
		request.Value,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, telemetry)
}