package ringo

import (
	"fmt"
	"net/http"
	"net/url"
)

func (r *Ringo) request(uri, method string, data interface{}) (*http.Request, error) {
	if err := r.Authenticate(); err != nil {
		return nil, err
	}

	rel, err := url.Parse(fmt.Sprintf("clients_api/%s?api_version=%d&auth_token=%s", uri, r.Config.APIVersion, r.Config.Token))
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, r.BaseURL.ResolveReference(rel).String(), nil)
	if err != nil {
		return nil, err
	}
	// req.Header.Add("content-type", "application/x-www-form-urlencoded")
	// req.Header.Add("user_agent", "android:com.ringapp:2.0.67(423)") // TODO: Is this needed?

	return req, nil
}
