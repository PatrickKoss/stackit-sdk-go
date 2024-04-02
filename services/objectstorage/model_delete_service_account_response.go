/*
STACKIT Object Storage API

STACKIT API to manage the Object Storage

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstorage

type DeleteServiceAccountResponse struct {
	// Project ID
	// REQUIRED
	Project *string `json:"project"`
	// The ID of the service account
	// REQUIRED
	ServiceAccountId *string `json:"serviceAccountId"`
}
