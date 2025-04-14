package design

import (
	"goa.design/goa/v3/dsl"
)

var _ = dsl.API("calc", func() {
	dsl.Title("Calculator Service")
	dsl.Description("A simple calculator service")
	dsl.Server("calc", func() {
		dsl.Host("localhost", func() {
			dsl.URI("http://localhost:8080")
		})
	})
})

var _ = dsl.Service("calc", func() {
	dsl.Description("The calc service provides addition.")

	dsl.Method("add", func() {
		dsl.Description("Adds two integers and returns the result.")
		dsl.Payload(func() {
			dsl.Attribute("a", dsl.Int64, "First Operand")
			dsl.Attribute("b", dsl.Int64, "Second Operand")
			dsl.Required("a", "b")
		})
		dsl.Result(dsl.Int64)
		dsl.HTTP(func() {
			dsl.GET("/add/{a}/{b}")
			dsl.Response(dsl.StatusOK)
		})
	})
	dsl.Method("subtract", func() {
		dsl.Description("Subtracts two integers and returns the result.")
		dsl.Payload(func() {
			dsl.Attribute("a", dsl.Int64)
			dsl.Attribute("b", dsl.Int64)
			dsl.Required("a", "b")
		})
		dsl.Result(dsl.Int64)
		dsl.HTTP(func() {
			dsl.GET("/sub/{a}/{b}")
			dsl.Response(dsl.StatusOK)
		})
	})
})
