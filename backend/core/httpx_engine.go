package core

import (
	cryptorand "crypto/rand"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func NewHTTPClient(timeout time.Duration) *http.Client {
	if timeout <= 0 {
		timeout = 1800 * time.Second
	}
	return &http.Client{Timeout: timeout}
}

func generateRequestID() string {
	b := make([]byte, 16)
	if _, err := cryptorand.Read(b); err != nil {
		return fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Int63())
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

func QwenHeaders(token string) http.Header {
	headers := http.Header{}
	headers.Set("Accept", "application/json, text/event-stream")
	headers.Set("Content-Type", "application/json")
	headers.Set("User-Agent", "Mozilla/5.0 qwen2api-go")
	if token != "" {
		headers.Set("Authorization", "Bearer "+token)
	}
	headers.Set("x-request-id", generateRequestID())
	return headers
}
