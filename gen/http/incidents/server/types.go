// Code generated by goa v3.2.6, DO NOT EDIT.
//
// incidents HTTP server types
//
// Command:
// $ goa gen github.com/brittonhayes/citadel/design

package server

import (
	incidents "github.com/brittonhayes/citadel/gen/incidents"
)

// FindResponseBody is the type of the "incidents" service "find" endpoint HTTP
// response body.
type FindResponseBody struct {
	// Unique ID of the incident
	ID *uint64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Date the incident occurred
	Date *string `form:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// Date the incident occurred
	DateClosed *string `form:"date_closed,omitempty" json:"date_closed,omitempty" xml:"date_closed,omitempty"`
	// Permissions associated with incident
	Permissions *string `form:"Permissions,omitempty" json:"Permissions,omitempty" xml:"Permissions,omitempty"`
	// The severity of the incident
	Severity *int `form:"severity,omitempty" json:"severity,omitempty" xml:"severity,omitempty"`
	// The short title of the incident
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// The detailed description of the incident
	Summary *string `form:"summary,omitempty" json:"summary,omitempty" xml:"summary,omitempty"`
	// The scope of impact of this incident
	Scope *string `form:"scope,omitempty" json:"scope,omitempty" xml:"scope,omitempty"`
	// What group or individual caused the initial incident
	ResponsibleParty *string `form:"responsible_party,omitempty" json:"responsible_party,omitempty" xml:"responsible_party,omitempty"`
	// A list of the affected customers
	AffectedCustomers []string `form:"affected_customers,omitempty" json:"affected_customers,omitempty" xml:"affected_customers,omitempty"`
	// The original cause of the incident
	RootCause *string `form:"root_cause,omitempty" json:"root_cause,omitempty" xml:"root_cause,omitempty"`
	// The slack channel for incident discussions
	SlackChannel *string `form:"slack_channel,omitempty" json:"slack_channel,omitempty" xml:"slack_channel,omitempty"`
	// When the incident was submitted
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// When the incident was last updated
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// ListAllResponseBody is the type of the "incidents" service "list all"
// endpoint HTTP response body.
type ListAllResponseBody []*IncidentResponse

// IncidentResponse is used to define fields on response body types.
type IncidentResponse struct {
	// Unique ID of the incident
	ID *uint64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Date the incident occurred
	Date *string `form:"date,omitempty" json:"date,omitempty" xml:"date,omitempty"`
	// Date the incident occurred
	DateClosed *string `form:"date_closed,omitempty" json:"date_closed,omitempty" xml:"date_closed,omitempty"`
	// Permissions associated with incident
	Permissions *string `form:"Permissions,omitempty" json:"Permissions,omitempty" xml:"Permissions,omitempty"`
	// The severity of the incident
	Severity *int `form:"severity,omitempty" json:"severity,omitempty" xml:"severity,omitempty"`
	// The short title of the incident
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// The detailed description of the incident
	Summary *string `form:"summary,omitempty" json:"summary,omitempty" xml:"summary,omitempty"`
	// The scope of impact of this incident
	Scope *string `form:"scope,omitempty" json:"scope,omitempty" xml:"scope,omitempty"`
	// What group or individual caused the initial incident
	ResponsibleParty *string `form:"responsible_party,omitempty" json:"responsible_party,omitempty" xml:"responsible_party,omitempty"`
	// A list of the affected customers
	AffectedCustomers []string `form:"affected_customers,omitempty" json:"affected_customers,omitempty" xml:"affected_customers,omitempty"`
	// The original cause of the incident
	RootCause *string `form:"root_cause,omitempty" json:"root_cause,omitempty" xml:"root_cause,omitempty"`
	// The slack channel for incident discussions
	SlackChannel *string `form:"slack_channel,omitempty" json:"slack_channel,omitempty" xml:"slack_channel,omitempty"`
	// When the incident was submitted
	CreatedAt *string `form:"created_at,omitempty" json:"created_at,omitempty" xml:"created_at,omitempty"`
	// When the incident was last updated
	UpdatedAt *string `form:"updated_at,omitempty" json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

// NewFindResponseBody builds the HTTP response body from the result of the
// "find" endpoint of the "incidents" service.
func NewFindResponseBody(res *incidents.Incident) *FindResponseBody {
	body := &FindResponseBody{
		ID:               res.ID,
		Date:             res.Date,
		DateClosed:       res.DateClosed,
		Permissions:      res.Permissions,
		Severity:         res.Severity,
		Title:            res.Title,
		Summary:          res.Summary,
		Scope:            res.Scope,
		ResponsibleParty: res.ResponsibleParty,
		RootCause:        res.RootCause,
		SlackChannel:     res.SlackChannel,
		CreatedAt:        res.CreatedAt,
		UpdatedAt:        res.UpdatedAt,
	}
	if res.AffectedCustomers != nil {
		body.AffectedCustomers = make([]string, len(res.AffectedCustomers))
		for i, val := range res.AffectedCustomers {
			body.AffectedCustomers[i] = val
		}
	}
	return body
}

// NewListAllResponseBody builds the HTTP response body from the result of the
// "list all" endpoint of the "incidents" service.
func NewListAllResponseBody(res []*incidents.Incident) ListAllResponseBody {
	body := make([]*IncidentResponse, len(res))
	for i, val := range res {
		body[i] = marshalIncidentsIncidentToIncidentResponse(val)
	}
	return body
}

// NewFindPayload builds a incidents service find endpoint payload.
func NewFindPayload(id uint64) *incidents.FindPayload {
	v := &incidents.FindPayload{}
	v.ID = id

	return v
}

// NewListAllLimitPayload builds a incidents service list all endpoint payload.
func NewListAllLimitPayload(limit *int32) *incidents.LimitPayload {
	v := &incidents.LimitPayload{}
	v.Limit = limit

	return v
}
