package models

type Device struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Status string `json:"status"`
}

func (d *Device) Register() {
	d.Status = "registered"
}

func (d *Device) Connect() {
	d.Status = "connected"
}

func (d *Device) Configure() {
	d.Status = "configured"
}