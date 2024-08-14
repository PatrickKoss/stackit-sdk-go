/*
STACKIT Membership API

The Membership API is used to manage memberships, roles and permissions of STACKIT resources, like projects, folders, organizations and other resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorization

// Role struct for Role
type Role struct {
	// REQUIRED
	Description *string `json:"description"`
	Id          *string `json:"id,omitempty"`
	// REQUIRED
	Name *string `json:"name"`
	// REQUIRED
	Permissions *[]Permission `json:"permissions"`
}
