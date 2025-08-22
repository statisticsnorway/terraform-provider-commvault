package provider

import (
	"crypto/tls"
	"net/http"
	"time"
)

func NewHTTPClient(timeout time.Duration, skipVerifyTLS bool) *http.Client {
	return &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: skipVerifyTLS,
			},
		},
	}
}
