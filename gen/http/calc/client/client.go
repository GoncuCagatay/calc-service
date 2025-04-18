// Code generated by goa v3.20.1, DO NOT EDIT.
//
// calc client HTTP transport
//
// Command:
// $ goa gen github.com/GoncuCagatay/calc-service/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the calc service endpoint HTTP clients.
type Client struct {
	// Add Doer is the HTTP client used to make requests to the add endpoint.
	AddDoer goahttp.Doer

	// Subtract Doer is the HTTP client used to make requests to the subtract
	// endpoint.
	SubtractDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the calc service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		AddDoer:             doer,
		SubtractDoer:        doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Add returns an endpoint that makes HTTP requests to the calc service add
// server.
func (c *Client) Add() goa.Endpoint {
	var (
		decodeResponse = DecodeAddResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildAddRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AddDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("calc", "add", err)
		}
		return decodeResponse(resp)
	}
}

// Subtract returns an endpoint that makes HTTP requests to the calc service
// subtract server.
func (c *Client) Subtract() goa.Endpoint {
	var (
		decodeResponse = DecodeSubtractResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildSubtractRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SubtractDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("calc", "subtract", err)
		}
		return decodeResponse(resp)
	}
}
