package handler

import (
	"encoding/json"
	"net/http"
	"ports"

	"github.com/gorilla/mux"
	"github.com/mig-elgt/sender"
	"github.com/mig-elgt/sender/codes"
	"github.com/sirupsen/logrus"
)

type handler struct {
	storage ports.StorageService
}

func New(storage ports.StorageService) http.Handler {
	r := mux.NewRouter()
	h := handler{storage}
	r.HandleFunc("/v1/ports", h.CreatePort).Methods("POST")
	return r
}

func (h handler) CreatePort(w http.ResponseWriter, r *http.Request) {
	var port ports.Port
	if err := json.NewDecoder(r.Body).Decode(&port); err != nil {
		logrus.Errorf("could not decode post body request: %v", err)
		sender.
			NewJSON(w, http.StatusBadRequest).
			WithError(codes.InvalidArgument, "The request body entity is bad format.").Send()
		return
	}
	if err := h.storage.CreateOrUpdate(&port); err != nil {
		logrus.Errorf("could not create port: %v", err)
		sender.
			NewJSON(w, http.StatusInternalServerError).
			WithError(codes.Internal, "Something went wrong...").Send()
		return
	}
	logrus.Infof("Port created: %v", port)
	sender.NewJSON(w, http.StatusOK).Send(&port)
}
