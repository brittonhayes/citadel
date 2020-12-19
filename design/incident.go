package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("incidents", func() {
	Description("The incident service handles the creation and querying of security incidents")
	HTTP(func() {
		Path("/incidents")
	})
	Method("find", func() {
		Payload(func() {
			Field(1, "id", UInt64, "Unique ID of the incident")
			Required("id")
		})
		Result(Incident)
		Error("no_match", String, "No incident matched given criteria")
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})

	Method("list all", func() {
		Payload(LimitPayload)
		Result(ArrayOf(Incident))
		Error("no_match", String, "No incident matched given criteria")
		HTTP(func() {
			GET("/")
			Response(StatusOK)
			Params(func() {
				Param("limit:limit")
			})
		})
	})
})

var Incident = Type("Incident", func() {
	Attribute("id", UInt64, "Unique ID of the incident", func() {
		Minimum(0)
	})
	Attribute("date", String, "Date the incident occurred", func() {
		Format(FormatDate)
	})
	Attribute("date_closed", String, "Date the incident occurred", func() {
		Format(FormatDate)
	})
	Attribute("Permissions", String, "Permissions associated with incident", func() {
		Enum("admin", "analyst", "guest")
	})
	Attribute("severity", Int, "The severity of the incident", func() {
		Enum(0, 1, 2, 3, 4, 5)
	})
	Attribute("title", String, "The short title of the incident")
	Attribute("summary", String, "The detailed description of the incident")
	Attribute("scope", String, "The scope of impact of this incident")
	Attribute("responsible_party", String, "What group or individual caused the initial incident")
	Attribute("affected_customers", ArrayOf(String), "A list of the affected customers")
	Attribute("root_cause", String, "The original cause of the incident")
	Attribute("slack_channel", String, "The slack channel for incident discussions")
	Attribute("created_at", String, "When the incident was submitted", func() {
		Format(FormatDateTime)
	})
	Attribute("updated_at", String, "When the incident was last updated", func() {
		Format(FormatDateTime)
	})
})
