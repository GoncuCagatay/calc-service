// Code generated by goa v3.20.1, DO NOT EDIT.
//
// calc endpoints
//
// Command:
// $ goa gen github.com/GoncuCagatay/calc-service/design

package calc

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "calc" service endpoints.
type Endpoints struct {
	Add      goa.Endpoint
	Subtract goa.Endpoint
}

// NewEndpoints wraps the methods of the "calc" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Add:      NewAddEndpoint(s),
		Subtract: NewSubtractEndpoint(s),
	}
}

// Use applies the given middleware to all the "calc" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Add = m(e.Add)
	e.Subtract = m(e.Subtract)
}

// NewAddEndpoint returns an endpoint function that calls the method "add" of
// service "calc".
func NewAddEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*AddPayload)
		return s.Add(ctx, p)
	}
}

// NewSubtractEndpoint returns an endpoint function that calls the method
// "subtract" of service "calc".
func NewSubtractEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*SubtractPayload)
		return s.Subtract(ctx, p)
	}
}
