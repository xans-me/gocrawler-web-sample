package kurs

import (
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
	data, err := delivery.service.IndexingKurs(r.Context())
	// http response when there is a business logic error
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// http response when the service executed successfully
	response.SendSuccessResponse(w, data)
}

// NewDelivery to init http delivery
func NewDelivery(service IKursService) *HTTPDelivery {
	return &HTTPDelivery{service: service}
}
