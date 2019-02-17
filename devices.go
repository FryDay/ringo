package ringo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Devices returns all devices connected to logged in account.
func (r *Ringo) Devices() (*Devices, error) {
	req, err := r.request("ring_devices", "GET", nil)
	if err != nil {
		return nil, err
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	devices := new(Devices)
	if err := json.Unmarshal(content, devices); err != nil {
		return nil, err
	}

	return devices, nil
}

// Health returns the health of a device from it's ID.
func (r *Ringo) Health(id int64) (*DeviceHealth, error) {
	req, err := r.request(fmt.Sprintf("doorbots/%d/health", id), "GET", nil)
	if err != nil {
		return nil, err
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	health := struct {
		*DeviceHealth `json:"device_health"`
	}{}
	if err := json.Unmarshal(content, &health); err != nil {
		return nil, err
	}

	return health.DeviceHealth, nil
}
