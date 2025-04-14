package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("calc", func() {
	Title("Calculator Service")
	Description("A simple calculator service")
	Server("calc", func() {
		Host("localhost", func() {
			URI("http://localhost:8080")
		})
	})
})

var _ = Service("calc", func() {
	Description("The calc service provides addition.")
	
	Method("add", func() {
		Description("Adds two integes and returns the result.")
		Payload(func() {
			Attribute("a", Int)
			Attribute("b", Int)
			Required("a", "b")
		})
		Result(Int)
		HTTP(func(){
			GET("/add/{a}/{b}")
			Response(StatusOK)
		})
	})
})