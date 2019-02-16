package ringo

import (
	"encoding/json"
	"io/ioutil"
)

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
