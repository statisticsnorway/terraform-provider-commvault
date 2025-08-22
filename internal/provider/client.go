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

type loginResponse struct {
	Token string `xml:"token,attr"`
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
	defer resp.Body.Close()

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

func (c *APIClient) doXML(ctx context.Context, method, url string, xmlBody string) (*http.Response, error) {
	buf := bytes.NewBufferString(xmlBody)
	req, err := http.NewRequestWithContext(ctx, method, url, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Content-Type", "application/xml")
	if c.Token != "" {
		req.Header.Set("Authtoken", c.Token)
	}
	return c.HTTP.Do(req)
}

func (c *APIClient) findDefaultSubclientID(ctx context.Context, clientID string) (int64, error) {
	u := fmt.Sprintf("%s/Subclient?clientId=%s", c.BaseURL, clientID)
	resp, err := c.doJSON(ctx, http.MethodGet, u, nil)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return 0, nil
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("discover subclient: HTTP %d: %s", resp.StatusCode, string(b))
	}

	type ent struct {
		SubClientEntity struct {
			SubclientId   int64  `json:"subclientId"`
			SubclientName string `json:"subclientName"`
			AppName       string `json:"appName"`
		} `json:"subClientEntity"`
	}

	var shape1 struct {
		SubClients []ent `json:"subClientProperties"`
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&shape1); err == nil && len(shape1.SubClients) > 0 {
		for _, s := range shape1.SubClients {
			if s.SubClientEntity.SubclientName == "default" && s.SubClientEntity.AppName == "Cloud Apps" {
				return s.SubClientEntity.SubclientId, nil
			}
		}
		return 0, nil
	}

	var shape2 []ent
	resp.Body.Close()
	resp, err = c.doJSON(ctx, http.MethodGet, u, nil)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&shape2); err == nil {
		for _, s := range shape2 {
			if s.SubClientEntity.SubclientName == "default" && s.SubClientEntity.AppName == "Cloud Apps" {
				return s.SubClientEntity.SubclientId, nil
			}
		}
	}

	return 0, nil
}

func (c *APIClient) bindGCSBucket(ctx context.Context, subclientID int64, bucketName, projectName string) error {
	payload := map[string]any{
		"subClientProperties": map[string]any{
			"subClientEntity": map[string]any{"subclientId": subclientID},
			"useLocalContent": true,
			"cloudAppsSubClientProp": map[string]any{
				"instanceType": 20,
				"objectStorageSubclient": map[string]any{
					"contentOperationType": 2,
					"content": []any{
						map[string]any{
							"gcpContent": map[string]any{
								"bucketName":  bucketName,
								"projectName": projectName,
							},
						},
					},
				},
			},
		},
	}

	url := fmt.Sprintf("%s/Subclient/%d", c.BaseURL, subclientID)
	resp, err := c.doJSON(ctx, http.MethodPost, url, payload)
	if err != nil {
		return fmt.Errorf("bind bucket request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("bind bucket: HTTP %d: %s", resp.StatusCode, string(b))
	}
	var r struct {
		Response struct {
			WarningCode    int    `json:"warningCode"`
			ErrorCode      int    `json:"errorCode"`
			WarningMessage string `json:"warningMessage"`
		} `json:"response"`
	}
	_ = json.NewDecoder(resp.Body).Decode(&r)

	if r.Response.ErrorCode != 0 {
		return fmt.Errorf("bind bucket: Commvault errorCode=%d, warningCode=%d, msg=%q",
			r.Response.ErrorCode, r.Response.WarningCode, r.Response.WarningMessage)
	}

	return nil
}
