package kurs

import "github.com/go-chi/chi"

// Routes struct contains http delivery
type Routes struct {
	delivery *HTTPDelivery
}

// RegisterRoutes to register all public routes
func (routes Routes) RegisterRoutes(r chi.Router) {

	r.Get("/api/indexing", routes.delivery.Indexing)

	r.Route("/api/kurs", func(r chi.Router) {
		r.Post("/", routes.delivery.InsertKurs)
		r.Put("/", routes.delivery.InsertKurs)
	})
}

// NewRoutes to init http delivery with registered routes
func NewRoutes(delivery *HTTPDelivery) *Routes {
	return &Routes{delivery: delivery}
}
