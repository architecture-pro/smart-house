package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TelemetryService struct {
	BaseURL    string
	HTTPClient *http.Client
}

type TelemetryResponse struct {
	ID        string    `json:"id" example:"6fc599a1-8763-411d-b203-0f947cefdc06"`
	DeviceID  string    `json:"deviceId" example:"device-1"`
	Type      string    `json:"type" example:"temperature"`
	Value     string    `json:"value" example:"23.5"`
	Timestamp time.Time `json:"timestamp" example:"2026-09-11T10:58:56.296Z"`
}

func NewTelemetryService(baseURL string) *TelemetryService {
	return &TelemetryService{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *TelemetryService) GetTelemetry() ([]TelemetryResponse, error) {
	url := fmt.Sprintf("%s/api/v1/telemetry", s.BaseURL)

	resp, err := s.HTTPClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching telemetry: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var telemetry []TelemetryResponse

	if err := json.NewDecoder(resp.Body).Decode(&telemetry); err != nil {
		return nil, fmt.Errorf("error decoding telemetry response: %w", err)
	}

	return telemetry, nil
}

func (s *TelemetryService) CreateTelemetry(
	deviceID string,
	telemetryType string,
	value string,
) (*TelemetryResponse, error) {
	requestBody := struct {
		DeviceID string `json:"deviceId"`
		Type     string `json:"type"`
		Value    string `json:"value"`
	}{
		DeviceID: deviceID,
		Type:     telemetryType,
		Value:    value,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error encoding telemetry request: %w", err)
	}

	url := fmt.Sprintf("%s/api/v1/telemetry", s.BaseURL)

	resp, err := s.HTTPClient.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, fmt.Errorf("error creating telemetry: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var telemetry TelemetryResponse

	if err := json.NewDecoder(resp.Body).Decode(&telemetry); err != nil {
		return nil, fmt.Errorf("error decoding telemetry response: %w", err)
	}

	return &telemetry, nil
}