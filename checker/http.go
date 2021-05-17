package checker

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	// HTTPCheckerType name
	HTTPCheckerType = "http"
)

// HTTP checker implementation
type HTTP struct {
	URL            string `json:"url"`
	ExpectedStatus int    `json:"status_code"`
	Config
}

// Check the configuration via HTTP
func (h *HTTP) Check(ctx context.Context) (*Result, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, h.URL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create a request for URL %s: %w", h.URL, err)
	}

	result := &Result{}

	start := time.Now()
	response, err := http.DefaultClient.Do(request)
	result.RTT = time.Since(start)

	if err != nil {
		result.Message = err.Error()

		return result, nil
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		result.Message = fmt.Sprintf("failed to get a status 200, got %d", response.StatusCode)

		return result, nil
	}

	result.Success = true

	return result, nil
}
