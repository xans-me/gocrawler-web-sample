package kurs

import (
	"encoding/json"
	"gocrawler-web-sample/shared/response"
	"net/http"
)

// HTTPDelivery struct contains service and newrelic
type HTTPDelivery struct {
	service IKursService
}

// Indexing controller
func (delivery HTTPDelivery) Indexing(w http.ResponseWriter, r *http.Request) {
	// business logic execution for updating product item
	data, err := delivery.service.IndexingKurs()
	// http response when there is a business logic error
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// http response when the service executed successfully
	response.SendSuccessResponse(w, data)
}

func (delivery HTTPDelivery) InsertKurs(w http.ResponseWriter, r *http.Request) {
	var dataKurs DataKurs
	if err := json.NewDecoder(r.Body).Decode(&dataKurs); err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
	}

	if err := delivery.service.InsertDataKurs(dataKurs); err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
	}

	response.SendSuccessResponse(w, http.StatusOK)

}

func (delivery HTTPDelivery) UpdateKurs(w http.ResponseWriter, r *http.Request) {
	var dataKurs DataKurs
	if err := json.NewDecoder(r.Body).Decode(&dataKurs); err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
	}

	// TODO 1 : Need to changes
	if err := delivery.service.InsertDataKurs(dataKurs); err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
	}

	response.SendSuccessResponse(w, http.StatusOK)

}

// NewDelivery to init http delivery
func NewDelivery(service IKursService) *HTTPDelivery {
	return &HTTPDelivery{service: service}
}
