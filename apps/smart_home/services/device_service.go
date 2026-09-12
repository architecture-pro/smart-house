package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type DeviceService struct {
	BaseURL    string
	HTTPClient *http.Client
}

type DeviceResponse struct {
	ID     string `json:"id" example:"81bb755d-5eb1-4a89-90d2-c2ffce58fa5a"`
	Name   string `json:"name" example:"Temperature Sensor"`
	Type   string `json:"type" example:"temperature"`
	Status string `json:"status" example:"registered"`
}

func NewDeviceService(baseURL string) *DeviceService {
	return &DeviceService{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *DeviceService) GetDevices() ([]DeviceResponse, error) {
	url := fmt.Sprintf("%s/api/v1/devices", s.BaseURL)

	resp, err := s.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching devices: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"unexpected status code: %d",
			resp.StatusCode,
		)
	}

	var devices []DeviceResponse

	if err := json.NewDecoder(resp.Body).Decode(&devices); err != nil {
		return nil, fmt.Errorf(
			"error decoding devices response: %w",
			err,
		)
	}

	return devices, nil
}

func (s *DeviceService) Register(name, deviceType string) (*DeviceResponse, error) {
	requestBody := struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}{
		Name: name,
		Type: deviceType,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error encoding device request: %w", err)
	}

	url := fmt.Sprintf("%s/api/v1/devices/register", s.BaseURL)

	resp, err := s.HTTPClient.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, fmt.Errorf("error registering device: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var device DeviceResponse
	if err := json.NewDecoder(resp.Body).Decode(&device); err != nil {
		return nil, fmt.Errorf("error decoding device response: %w", err)
	}

	return &device, nil
}