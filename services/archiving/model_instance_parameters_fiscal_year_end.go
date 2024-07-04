/*
STACKIT Archiving Service API

The STACKIT Archiving Service (SAS) offers archiving endpoints for SAP Archive Link, SAP-ILM and non-SAP archiving scenarios.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package archiving

type InstanceParametersFiscalYearEnd struct {
	// Day that marks the end of fiscal year.
	// Can be cast to int32 without loss of precision.
	Day *int64 `json:"day,omitempty"`
	// Month that marks the end of fiscal year.
	// Can be cast to int32 without loss of precision.
	Month *int64 `json:"month,omitempty"`
}
