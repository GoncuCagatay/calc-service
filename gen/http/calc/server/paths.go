// Code generated by goa v3.20.1, DO NOT EDIT.
//
// HTTP request path constructors for the calc service.
//
// Command:
// $ goa gen github.com/GoncuCagatay/calc-service/design

package server

import (
	"fmt"
)

// AddCalcPath returns the URL path to the calc service add HTTP endpoint.
func AddCalcPath(a int, b int) string {
	return fmt.Sprintf("/add/%v/%v", a, b)
}

// SubtractCalcPath returns the URL path to the calc service subtract HTTP endpoint.
func SubtractCalcPath(a int, b int) string {
	return fmt.Sprintf("/sub/%v/%v", a, b)
}
