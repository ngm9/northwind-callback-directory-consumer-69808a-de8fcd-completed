package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"northwind/callbackdirectory/client"
	"northwind/callbackdirectory/service"
)

type stubClient struct {
	payload client.CountryPayload
}

func (s stubClient) FetchCountry(country string) (client.CountryPayload, error) {
	return s.payload, nil
}

func TestEndpointKnownCountry(t *testing.T) {
	stub := stubClient{payload: client.CountryPayload{Data: []client.CountryRecord{{Name: "Afghanistan", CallingCodes: []string{"93"}}}}}
	svc := service.NewDirectoryService(stub)
	h := NewHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/phone-numbers?country=Afghanistan&phone=656445445", nil)
	rec := httptest.NewRecorder()
	h.Routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}

func TestEndpointUnknowCountry(t *testing.T) {
	stub := stubClient{payload: client.CountryPayload{Data: []client.CountryRecord{{Name: "Afghan", CallingCodes: []string{}}}}}
	svc := service.NewDirectoryService(stub)
	h := NewHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/phone-numbers?country=Afghan&phone=656445445", nil)
	rec := httptest.NewRecorder()
	h.Routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}

func TestEndpointEmptyPhone(t *testing.T) {
	stub := stubClient{payload: client.CountryPayload{Data: []client.CountryRecord{{Name: "Afghan", CallingCodes: []string{}}}}}
	svc := service.NewDirectoryService(stub)
	h := NewHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/phone-numbers?country=&phone=656445445", nil)
	rec := httptest.NewRecorder()
	h.Routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rec.Code)
	}
}

func TestEndpointEmptyCountry(t *testing.T) {
	stub := stubClient{payload: client.CountryPayload{Data: []client.CountryRecord{{Name: "", CallingCodes: []string{}}}}}
	svc := service.NewDirectoryService(stub)
	h := NewHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/phone-numbers?country=Afghan&phone=", nil)
	rec := httptest.NewRecorder()
	h.Routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rec.Code)
	}
}
