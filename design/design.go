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
		// dsl.Description("Adds two integers and returns the result.")
		dsl.Payload(func() {
			dsl.Attribute("a", dsl.String, "First Operand")
			dsl.Attribute("b", dsl.String, "Second Operand")
			dsl.Required("a", "b")
		})
		dsl.Result(dsl.String)
		dsl.HTTP(func() {
			dsl.GET("/add/{a}/{b}")
			dsl.Response(dsl.StatusOK, func() {
				dsl.ContentType("text/plain")
			})
		})
	})
	dsl.Method("subtract", func() {
		dsl.Description("Subtracts two integers and returns the result.")
		dsl.Payload(func() {
			dsl.Attribute("a", dsl.String, "First Operand")
			dsl.Attribute("b", dsl.String, "Second Operand")
			dsl.Required("a", "b")
		})
		dsl.Result(dsl.String)
		dsl.HTTP(func() {
			dsl.GET("/sub/{a}/{b}")
			dsl.Response(dsl.StatusOK, func() {
				dsl.ContentType("text/plain")
			})
		})
	})
})
