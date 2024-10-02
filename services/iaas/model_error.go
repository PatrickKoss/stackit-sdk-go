/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

// Error Error with HTTP error code and an error message.
type Error struct {
	// REQUIRED
	Code *int64 `json:"code"`
	// An error message.
	// REQUIRED
	Msg *string `json:"msg"`
}
