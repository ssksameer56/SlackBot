package utils

import (
	"net/http"
	"time"
)

type HttpRequest struct {
	Client  http.Client
	Timeout time.Duration
	BaseURL string
}

func (h *HttpRequest) Get() ([]byte, error) {
	return []byte{}, nil
}

func (h *HttpRequest) Post() ([]byte, error) {
	return []byte{}, nil
}
