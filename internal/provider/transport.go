package provider

import (
	"crypto/tls"
	"net/http"
)

func NewHTTPClientInsecure() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}
