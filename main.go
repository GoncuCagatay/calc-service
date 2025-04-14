package main

import (
	"context"
	"log"
	"net/http"
	"os"

	calc "github.com/GoncuCagatay/calc-service/gen/calc"
	calcsvr "github.com/GoncuCagatay/calc-service/gen/http/calc/server"
	goahttp "goa.design/goa/v3/http"
)

type statusCode int

func (s statusCode) StatusCode() int {
	return int(s)
}

// Service logic
type calcService struct{}

func (s *calcService) Add(ctx context.Context, p *calc.AddPayload) (int64, error) {
	return p.A + p.B, nil
}

func (s *calcService) Subtract(ctx context.Context, p *calc.SubtractPayload) (int64, error) {
	return p.A - p.B, nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create service implementation
	svc := &calcService{}

	// Create endpoints
	endpoints := calc.NewEndpoints(svc)

	// Goa HTTP Mux and handler setup
	mux := goahttp.NewMuxer()
	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder
	eh := func(ctx context.Context, w http.ResponseWriter, err error) {
		log.Printf("internal error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
	}
	st := func(ctx context.Context, err error) goahttp.Statuser {
		return statusCode(http.StatusInternalServerError)
	}

	// Mount service
	calcsvr.Mount(mux, calcsvr.New(endpoints, mux, dec, enc, eh, st))

	// Start HTTP server
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
