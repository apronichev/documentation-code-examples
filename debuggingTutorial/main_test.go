package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	tests := []struct {
		name           string
		r              *http.Request
		w              *httptest.ResponseRecorder
		expectedStatus int
	}{
		{
			name:           "good",
			r:              httptest.NewRequest("GET", "/", nil),
			w:              httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			indexHandler(test.w, test.r)
			if test.w.Code != test.expectedStatus {
				t.Errorf("Failed to produce expected status code %d, got %d", test.expectedStatus, test.w.Code)
			}
		})
	}
}