package design

import (
	. "goa.design/goa/v3/dsl"
)

// API describes the global properties of the API server.
var _ = API("citadel", func() {
	Title("Citadel service")
	Description("HTTP service for working with Security Operations Center resources")
	Server("citadel", func() {
		Description("citadel contains the vulnerability, risk, incident, and intel services.")
		Services("vulnerabilities", "incidents")
		Host("development", func() {
			Description("Development hosts")
			URI("http://localhost:8000/{version}/citadel")
			// Variable describes a URI variable.
			Variable("version", String, "API version", func() {
				Default("v1")
			})
		})

		Host("production", func() {
			Description("Production hosts")
			URI("https://localhost:8000/{version}/citadel")

			// Variable describes a URI variable.
			Variable("version", String, "API version", func() {
				Default("v1")
			})
		})
	})
})

var LimitPayload = Type("LimitPayload", func() {
	Attribute("limit", Int32, "Limit the number of results", func() {
		Meta("rpc:tag", "1")
		Minimum(0)
		Maximum(1000)
	})
})
