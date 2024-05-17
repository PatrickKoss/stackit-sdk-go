/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

type Request struct {
	Details *string `json:"details,omitempty"`
	// Resource action.
	// REQUIRED
	RequestAction *string `json:"requestAction"`
	// ID representing a single API request.
	// REQUIRED
	RequestId *string `json:"requestId"`
	// Resource type.
	// REQUIRED
	RequestType *string `json:"requestType"`
	// REQUIRED
	Resources *[]RequestResource `json:"resources"`
	// The state of a resource object.
	// REQUIRED
	Status *string `json:"status"`
}
