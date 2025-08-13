package commvault

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"time"
)

func login(cfg *CommvaultConfig) (string, error) {
	loginPayload := map[string]interface{}{
		"username":     cfg.Username,
		"password":     cfg.Password,
		"encodeBase64": true,
	}

	data, _ := json.Marshal(loginPayload)
	req, _ := http.NewRequest("POST", cfg.BaseURL+"/Login", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: newTransport(),
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", errors.New("login failed: non-2xx response")
	}

	var lr loginResponse
	if err := xml.NewDecoder(resp.Body).Decode(&lr); err != nil {
		return "", errors.New("login failed: could not parse token")
	}
	if lr.Token == "" {
		return "", errors.New("login failed: empty token")
	}

	return lr.Token, nil
}
