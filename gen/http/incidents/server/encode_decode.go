// Code generated by goa v3.2.6, DO NOT EDIT.
//
// incidents HTTP server encoders and decoders
//
// Command:
// $ goa gen citadel/design

package server

import (
	incidents "citadel/gen/incidents"
	"context"
	"net/http"
	"strconv"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeFindResponse returns an encoder for responses returned by the
// incidents find endpoint.
func EncodeFindResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*incidents.Incident)
		enc := encoder(ctx, w)
		body := NewFindResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeFindRequest returns a decoder for requests sent to the incidents find
// endpoint.
func DecodeFindRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  uint64
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = v
		}
		if err != nil {
			return nil, err
		}
		payload := NewFindPayload(id)

		return payload, nil
	}
}

// EncodeListAllResponse returns an encoder for responses returned by the
// incidents list all endpoint.
func EncodeListAllResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.([]*incidents.Incident)
		enc := encoder(ctx, w)
		body := NewListAllResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListAllRequest returns a decoder for requests sent to the incidents
// list all endpoint.
func DecodeListAllRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			limit *int32
			err   error
		)
		{
			limitRaw := r.URL.Query().Get("limit")
			if limitRaw != "" {
				v, err2 := strconv.ParseInt(limitRaw, 10, 32)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("limit", limitRaw, "integer"))
				}
				pv := int32(v)
				limit = &pv
			}
		}
		if limit != nil {
			if *limit < 0 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", *limit, 0, true))
			}
		}
		if limit != nil {
			if *limit > 1000 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", *limit, 1000, false))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewListAllLimitPayload(limit)

		return payload, nil
	}
}

// marshalIncidentsIncidentToIncidentResponse builds a value of type
// *IncidentResponse from a value of type *incidents.Incident.
func marshalIncidentsIncidentToIncidentResponse(v *incidents.Incident) *IncidentResponse {
	res := &IncidentResponse{
		ID:               v.ID,
		Date:             v.Date,
		DateClosed:       v.DateClosed,
		Permissions:      v.Permissions,
		Severity:         v.Severity,
		Title:            v.Title,
		Summary:          v.Summary,
		Scope:            v.Scope,
		ResponsibleParty: v.ResponsibleParty,
		RootCause:        v.RootCause,
		SlackChannel:     v.SlackChannel,
		CreatedAt:        v.CreatedAt,
		UpdatedAt:        v.UpdatedAt,
	}
	if v.AffectedCustomers != nil {
		res.AffectedCustomers = make([]string, len(v.AffectedCustomers))
		for i, val := range v.AffectedCustomers {
			res.AffectedCustomers[i] = val
		}
	}

	return res
}