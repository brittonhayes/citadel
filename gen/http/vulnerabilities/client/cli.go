// Code generated by goa v3.2.6, DO NOT EDIT.
//
// vulnerabilities HTTP client CLI support package
//
// Command:
// $ goa gen github.com/brittonhayes/citadel/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	vulnerabilities "github.com/brittonhayes/citadel/gen/vulnerabilities"
	goa "goa.design/goa/v3/pkg"
)

// BuildFindPayload builds the payload for the vulnerabilities find endpoint
// from CLI flags.
func BuildFindPayload(vulnerabilitiesFindID string) (*vulnerabilities.FindPayload, error) {
	var err error
	var id uint64
	{
		id, err = strconv.ParseUint(vulnerabilitiesFindID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT64")
		}
	}
	v := &vulnerabilities.FindPayload{}
	v.ID = id

	return v, nil
}

// BuildListPayload builds the payload for the vulnerabilities list endpoint
// from CLI flags.
func BuildListPayload(vulnerabilitiesListLimit string) (*vulnerabilities.LimitPayload, error) {
	var err error
	var limit *int32
	{
		if vulnerabilitiesListLimit != "" {
			var v int64
			v, err = strconv.ParseInt(vulnerabilitiesListLimit, 10, 32)
			val := int32(v)
			limit = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for limit, must be INT32")
			}
			if limit != nil {
				if *limit < 0 {
					err = goa.MergeErrors(err, goa.InvalidRangeError("limit", *limit, 0, true))
				}
			}
			if limit != nil {
				if *limit > 500 {
					err = goa.MergeErrors(err, goa.InvalidRangeError("limit", *limit, 500, false))
				}
			}
			if err != nil {
				return nil, err
			}
		}
	}
	v := &vulnerabilities.LimitPayload{}
	v.Limit = limit

	return v, nil
}

// BuildSubmitPayload builds the payload for the vulnerabilities submit
// endpoint from CLI flags.
func BuildSubmitPayload(vulnerabilitiesSubmitBody string) (*vulnerabilities.SubmitPayload, error) {
	var err error
	var body SubmitRequestBody
	{
		err = json.Unmarshal([]byte(vulnerabilitiesSubmitBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"cvss_score\": 0.7719438,\n      \"description\": \"Veniam iste velit illo maxime impedit pariatur.\",\n      \"exploitable\": true,\n      \"is_patchable\": true,\n      \"is_upgradeable\": true,\n      \"title\": \"Voluptatum ut eveniet.\"\n   }'")
		}
	}
	v := &vulnerabilities.SubmitPayload{
		Title:         body.Title,
		Description:   body.Description,
		Exploitable:   body.Exploitable,
		CvssScore:     body.CvssScore,
		IsPatchable:   body.IsPatchable,
		IsUpgradeable: body.IsUpgradeable,
	}

	return v, nil
}
