package service

import (
	"context"
)

type NetworkService interface {
	HTTPRequest(ctx context.Context, req *HTTPRequest) (*HTTPResponse, error)
}

// HTTPResponse is the structure received from HttpRequest
type HTTPResponse struct {
	ResponseCode int
	Data         []byte
	Headers      map[string][]string
}

// HTTPRequest is the structure received from HttpRequest
type HTTPRequest struct {
	URL     string
	Method  string
	Body    []byte
	Headers map[string][]string
}