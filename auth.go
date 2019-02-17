package ringo

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"client"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
}

type tokenRequest struct {
	Device *clientDevice `json:"device"`
}

type clientDevice struct {
	HardwareID string    `json:"hardware_id"`
	Metadata   *metadata `json:"metadata"`
	OS         string    `json:"os"`
}

type metadata struct {
	APIVersion int `json:"api_version"`
}

// Authenticate sets token based on login information.
func (r *Ringo) Authenticate() error {
	if r.Config.Token != "" {
		return nil
	}

	body := map[string]interface{}{
		"client_id":  "ring_official_android",
		"grant_type": "password",
		"username":   r.Config.Username,
		"password":   r.Config.Password,
		"scope":      "client",
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://oauth.ring.com/oauth/token", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/json")

	resp, err := r.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	resp.Body.Close()

	tr := new(tokenResponse)
	json.Unmarshal(content, tr)

	hardwareID, err := randomHex(16)
	if err != nil {
		return err
	}
	bodyBytes, err = json.Marshal(
		&tokenRequest{
			Device: &clientDevice{
				HardwareID: hardwareID,
				Metadata: &metadata{
					APIVersion: r.Config.APIVersion,
				},
				OS: "android",
			},
		})
	if err != nil {
		return err
	}

	rel, err := url.Parse(fmt.Sprintf("clients_api/session?api_version=%d", r.Config.APIVersion))
	if err != nil {
		return err
	}

	req, err = http.NewRequest("POST", r.BaseURL.ResolveReference(rel).String(), bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("%s %s", tr.TokenType, tr.AccessToken))
	req.Header.Add("content-type", "application/json")

	resp, err = r.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	content, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	reqMap := make(map[string]interface{})
	if err := json.Unmarshal(content, &reqMap); err != nil {
		return err
	}

	if profile, ok := reqMap["profile"]; ok {
		if token, ok := profile.(map[string]interface{})["authentication_token"]; ok {
			r.Config.Token = token.(string)
		}
	}

	return nil
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
