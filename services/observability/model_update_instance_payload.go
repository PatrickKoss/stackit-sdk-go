/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

// UpdateInstancePayload Create update instance body.
type UpdateInstancePayload struct {
	// Name of the service
	Name *string `json:"name,omitempty"`
	// additional parameters
	Parameter *map[string]interface{} `json:"parameter,omitempty"`
	// uuid of the plan to create/update
	// REQUIRED
	PlanId *string `json:"planId"`
}
