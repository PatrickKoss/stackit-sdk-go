/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

// UpdateAlertConfigReceiverPayload Receivers
type UpdateAlertConfigReceiverPayload struct {
	// Email configurations
	EmailConfigs *[]CreateAlertConfigReceiverPayloadEmailConfigsInner `json:"emailConfigs,omitempty"`
	// `Additional Validators:` * must be unique * should only include the characters: a-zA-Z0-9-
	// REQUIRED
	Name *string `json:"name"`
	// Configuration for ops genie.
	OpsgenieConfigs *[]CreateAlertConfigReceiverPayloadOpsgenieConfigsInner `json:"opsgenieConfigs,omitempty"`
	WebHookConfigs  *[]CreateAlertConfigReceiverPayloadWebHookConfigsInner  `json:"webHookConfigs,omitempty"`
}
