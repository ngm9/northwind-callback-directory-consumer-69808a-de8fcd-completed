package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"northwind/callbackdirectory/response"
	"northwind/callbackdirectory/service"
)

type Handler struct {
	svc *service.DirectoryService
}

func NewHandler(svc *service.DirectoryService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) PhoneNumbers(w http.ResponseWriter, r *http.Request) {
	country := r.URL.Query().Get("country")
	if country == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.BuildResponseBody(http.StatusBadRequest, nil, errors.New("country is blank, please provide a valid input")))
		return
	}

	phone := r.URL.Query().Get("phone")
	if phone == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.BuildResponseBody(http.StatusBadRequest, nil, errors.New("phone is blank, please provide a valid input")))
		return
	}

	resolved, err := h.svc.GetPhoneNumbers(country, phone)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.BuildResponseBody(http.StatusBadRequest, nil, errors.New("phone is blank, please provide a valid input")))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.BuildResponseBody(http.StatusOK, &resolved, nil))
}

func (h *Handler) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/phone-numbers", h.PhoneNumbers)
	return mux
}
