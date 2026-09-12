package handlers

import (
	"net/http"

	"smarthome/services"

	"github.com/gin-gonic/gin"
)

type DeviceHandler struct {
	DeviceService *services.DeviceService
}


type CreateDeviceRequest struct {
	Name string `json:"name" binding:"required" example:"Temperature Sensor"`
	Type string `json:"type" binding:"required" example:"temperature"`
}

func NewDeviceHandler(deviceService *services.DeviceService) *DeviceHandler {
	return &DeviceHandler{
		DeviceService: deviceService,
	}
}

func (h *DeviceHandler) RegisterRoutes(router *gin.RouterGroup) {
	devices := router.Group("/devices")
	{
		devices.GET("", h.GetDevices)
		devices.POST("", h.CreateDevice)
	}
}

// GetDevices godoc
// @Summary Получить список устройств
// @Description Возвращает список устройств из Device Service
// @Tags Devices
// @Produce json
// @Success 200 {array} services.DeviceResponse
// @Failure 500 {object} ErrorResponse
// @Router /devices [get]
func (h *DeviceHandler) GetDevices(c *gin.Context) {
	devices, err := h.DeviceService.GetDevices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, devices)
}


// CreateDevice godoc
// @Summary Зарегистрировать устройство
// @Description Регистрирует новое устройство через Device Service
// @Tags Devices
// @Accept json
// @Produce json
// @Param request body CreateDeviceRequest true "Данные устройства"
// @Success 201 {object} services.DeviceResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /devices [post]
func (h *DeviceHandler) CreateDevice(c *gin.Context) {
	var request CreateDeviceRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	device, err := h.DeviceService.Register(
		request.Name,
		request.Type,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, device)
}
