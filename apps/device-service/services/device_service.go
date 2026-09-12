package services

import (
	"context"
	"device-service/db"
	"device-service/models"
	"github.com/google/uuid"
)

type DeviceService struct {
	DB *db.DB
}

func NewDeviceService(database *db.DB) *DeviceService {
	return &DeviceService{
		DB: database,
	}
}

func (s *DeviceService) Register(
	ctx context.Context,
	name string,
	deviceType string,
) (models.Device, error) {
	device := models.Device{
		ID:   uuid.New().String(),
		Name: name,
		Type: deviceType,
	}

	device.Register()

	return s.DB.CreateDevice(ctx, device)
}

func (s *DeviceService) Connect(
	ctx context.Context,
	id string,
) (models.Device, error) {
	device, err := s.DB.GetDeviceByID(ctx, id)
	if err != nil {
		return models.Device{}, err
	}

	device.Connect()

	return s.DB.UpdateDevice(ctx, device)
}

func (s *DeviceService) Configure(
	ctx context.Context,
	id string,
) (models.Device, error) {
	device, err := s.DB.GetDeviceByID(ctx, id)
	if err != nil {
		return models.Device{}, err
	}

	device.Configure()

	return s.DB.UpdateDevice(ctx, device)
}

func (s *DeviceService) GetDevices(
	ctx context.Context,
) ([]models.Device, error) {
	return s.DB.GetDevices(ctx)
}

func (s *DeviceService) GetDeviceByID(
	ctx context.Context,
	id string,
) (models.Device, error) {
	return s.DB.GetDeviceByID(ctx, id)
}

func (s *DeviceService) Delete(
	ctx context.Context,
	id string,
) error {
	return s.DB.DeleteDevice(ctx, id)
}