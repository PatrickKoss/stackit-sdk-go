/*
STACKIT Secrets Manager API

This API provides endpoints for managing the Secrets-Manager.

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package secretsmanager

// UpdateACLPayload struct for UpdateACLPayload
type UpdateACLPayload struct {
	// The given IP/IP Range that is permitted to access.
	// REQUIRED
	Cidr *string `json:"cidr"`
}
