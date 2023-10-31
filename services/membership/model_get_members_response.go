/*
STACKIT Membership API

The Membership API is used to manage memberships, roles and permissions of STACKIT resources, like projects, folders, organizations and other resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package membership

type GetMembersResponse struct {
	// REQUIRED
	Members *[]Member `json:"members"`
	// REQUIRED
	ResourceId *string `json:"resourceId"`
	// REQUIRED
	ResourceType *string `json:"resourceType"`
}
