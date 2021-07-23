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
	response.SendSuccessResponse(w, "OK")
}

// NewDelivery to init http delivery
func NewDelivery(service IKursService) *HTTPDelivery {
	return &HTTPDelivery{service: service}
}
