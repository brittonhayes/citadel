// Code generated by goa v3.2.6, DO NOT EDIT.
//
// HTTP request path constructors for the incidents service.
//
// Command:
// $ goa gen github.com/brittonhayes/citadel/design

package client

import (
	"fmt"
)

// FindIncidentsPath returns the URL path to the incidents service find HTTP endpoint.
func FindIncidentsPath(id uint64) string {
	return fmt.Sprintf("/incidents/%v", id)
}

// ListAllIncidentsPath returns the URL path to the incidents service list all HTTP endpoint.
func ListAllIncidentsPath() string {
	return "/incidents"
}
