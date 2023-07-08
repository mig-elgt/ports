package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"ports"
	"ports/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreatePort(t *testing.T) {
	type args struct {
		body                 *bytes.Buffer
		CreateOrUpdateMockFn func(port *ports.Port) error
	}
	testCases := map[string]struct {
		args           args
		wantStatusCode int
		wantResponse   string
	}{
		"body request bad format": {
			args: args{
				body: bytes.NewBufferString("sending plain text"),
			},
			wantStatusCode: http.StatusBadRequest,
			wantResponse: `
			{
				"error": {
					"status": 400,
					"error": "INVALID_ARGUMENT",
					"description": "The request body entity is bad format."
				}
			}`,
		},
		"database is not available": {
			args: args{
				body: bytes.NewBufferString(`
				{
					"name": "Ajman",
					"city": "Ajman",
					"country": "United Arab Emirates",
					"alias": [],
					"regions": [],
					"coordinates": [
					  55.5136433,
					  25.4052165
					],
					"province": "Ajman",
					"timezone": "Asia/Dubai",
					"unlocs": [
					  "AEAJM"
					],
					"code": "52000"
				}
				`),
				CreateOrUpdateMockFn: func(_ *ports.Port) error {
					return errors.New("failed to store port: could not exec query")
				},
			},
			wantStatusCode: http.StatusInternalServerError,
			wantResponse: `
			{
				"error": {
					"status":500,
					"error":"INTERNAL",
					"description":"Something went wrong..."
				}
			}`,
		},
		"base case": {
			args: args{
				body: bytes.NewBufferString(`
				{
					"name": "Ajman",
					"city": "Ajman",
					"country": "United Arab Emirates",
					"alias": [],
					"regions": [],
					"coordinates": [
					  55.5136433,
					  25.4052165
					],
					"province": "Ajman",
					"timezone": "Asia/Dubai",
					"unlocs": [
					  "AEAJM"
					],
					"code": "52000"
				}
				`),
				CreateOrUpdateMockFn: func(_ *ports.Port) error {
					return nil
				},
			},
			wantStatusCode: http.StatusOK,
			wantResponse: `
			{
				"name": "Ajman",
				"city": "Ajman",
				"country": "United Arab Emirates",
				"alias": [],
				"regions": [],
				"coordinates": [
				  55.5136433,
				  25.4052165
				],
				"province": "Ajman",
				"timezone": "Asia/Dubai",
				"unlocs": [
				  "AEAJM"
				],
				"code": "52000"
			}`,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodPost, "/v1/ports", tc.args.body)
			rec := httptest.NewRecorder()
			h := handler{
				storage: &mocks.StorageServiceMock{
					CreateOrUpdateFn: tc.args.CreateOrUpdateMockFn,
				},
			}
			h.CreatePort(rec, req)
			assert.Equal(t, tc.wantStatusCode, rec.Code)
			got := string(rec.Body.Bytes())
			assert.JSONEq(t, tc.wantResponse, got)
		})
	}
}
