// Code generated by goa v3.2.6, DO NOT EDIT.
//
// HTTP request path constructors for the vulnerabilities service.
//
// Command:
// $ goa gen citadel/design

package client

import (
	"fmt"
)

// FindVulnerabilitiesPath returns the URL path to the vulnerabilities service find HTTP endpoint.
func FindVulnerabilitiesPath(id uint64) string {
	return fmt.Sprintf("/vulnerabilities/%v", id)
}

// ListVulnerabilitiesPath returns the URL path to the vulnerabilities service list HTTP endpoint.
func ListVulnerabilitiesPath() string {
	return "/vulnerabilities"
}

// SubmitVulnerabilitiesPath returns the URL path to the vulnerabilities service submit HTTP endpoint.
func SubmitVulnerabilitiesPath() string {
	return "/vulnerabilities"
}