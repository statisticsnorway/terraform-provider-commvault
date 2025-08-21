package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type APIClient struct {
	BaseURL   string
	Username  string
	Password  string
	Token     string
	HTTP      *http.Client
	Transport http.RoundTripper
}

func (c *APIClient) Login(ctx context.Context) (string, error) {
	payload := map[string]any{
		"username":     c.Username,
		"password":     c.Password,
		"encodeBase64": true,
	}

	data, _ := json.Marshal(payload)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.BaseURL+"/Login", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	c.HTTP.Timeout = 30 * time.Second
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("login: non-2xx status %d", resp.StatusCode)
	}

	var lr loginResponse
	if err := xml.NewDecoder(resp.Body).Decode(&lr); err != nil {
		return "", errors.New("login: could not parse token")
	}
	if lr.Token == "" {
		return "", errors.New("login: empty token")
	}
	return lr.Token, nil
}

func (c *APIClient) doJSON(ctx context.Context, method, url string, body any) (*http.Response, error) {
	var buf *bytes.Buffer
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	} else {
		buf = bytes.NewBuffer(nil)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	if c.Token != "" {
		req.Header.Set("Authtoken", c.Token)
	}
	return c.HTTP.Do(req)
}
