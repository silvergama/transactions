package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestIDMiddleware(t *testing.T) {
	testCases := []struct {
		name            string
		requestIDHeader string
		expectedHeader  string
	}{
		{
			name:            "NoHeader",
			requestIDHeader: "",
			expectedHeader:  "uuid string",
		},
		{
			name:            "WithHeader",
			requestIDHeader: "custom-request-id",
			expectedHeader:  "custom-request-id",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			if tc.requestIDHeader != "" {
				req.Header.Set("X-Request-ID", tc.requestIDHeader)
			}

			rr := httptest.NewRecorder()
			handler := RequestIDMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

			handler.ServeHTTP(rr, req)

			if tc.requestIDHeader == "" {
				assert.NotEmpty(t, rr.Header().Get("X-Request-ID"))
				return
			}
			assert.Equal(t, tc.expectedHeader, rr.Header().Get("X-Request-ID"))

		})
	}
}
