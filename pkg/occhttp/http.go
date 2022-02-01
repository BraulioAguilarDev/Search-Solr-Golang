package http

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/golang/glog"
)

var DebugMode = false

// Request struct
type request struct {
	Client *http.Client
}

// Options struct
type Options struct {
	Timeout  time.Duration
	Endpoint string
	Body     []byte
	Method   string
}

// SuccessResponse is the basic form of an informational reply upon success
type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// FailureResponse is the basic form of an informational reply upon failure
type FailureResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// PayloadResponseEnvelope standard format
type PayloadResponseEnvelope struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

// Failure is the basic failure HTTP response
func Failure(details string) FailureResponse {
	var resp = FailureResponse{
		Success: false,
		Message: "failed",
	}

	if DebugMode {
		resp.Details = details
	}

	return resp
}

// NewClient func
func NewClient() *request {
	return &request{
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

// MakeRequest func executes external request
// returns *http.Response and error
func (r *request) MakeRequest(o *Options) (*http.Response, error) {
	req, err := http.NewRequest(o.Method, o.Endpoint, bytes.NewBuffer(o.Body))
	if err != nil {
		glog.V(5).Info(fmt.Sprintf("Error request: %v", err))
		return nil, err
	}

	res, err := r.Client.Do(req)
	if err != nil {
		glog.V(5).Info(fmt.Sprintf("Error: %v", err))
		return nil, err
	}

	return res, nil
}

func SetDebugMode(debugMode bool) {
	DebugMode = debugMode
}
