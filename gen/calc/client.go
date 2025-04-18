// Code generated by goa v3.20.1, DO NOT EDIT.
//
// calc client
//
// Command:
// $ goa gen github.com/GoncuCagatay/calc-service/design

package calc

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "calc" service client.
type Client struct {
	AddEndpoint      goa.Endpoint
	SubtractEndpoint goa.Endpoint
}

// NewClient initializes a "calc" service client given the endpoints.
func NewClient(add, subtract goa.Endpoint) *Client {
	return &Client{
		AddEndpoint:      add,
		SubtractEndpoint: subtract,
	}
}

// Add calls the "add" endpoint of the "calc" service.
func (c *Client) Add(ctx context.Context, p *AddPayload) (res string, err error) {
	var ires any
	ires, err = c.AddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Subtract calls the "subtract" endpoint of the "calc" service.
func (c *Client) Subtract(ctx context.Context, p *SubtractPayload) (res string, err error) {
	var ires any
	ires, err = c.SubtractEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}
