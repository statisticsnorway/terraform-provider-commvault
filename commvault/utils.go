package commvault

import (
	"crypto/tls"
	"net/http"
)

func newHTTPClient() *http.Client {
	return &http.Client{
		Transport: newTransport(),
	}
}

func newTransport() *http.Transport {
	return &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
}
