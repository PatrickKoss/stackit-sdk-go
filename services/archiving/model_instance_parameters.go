/*
STACKIT Archiving Service API

The STACKIT Archiving Service (SAS) offers archiving endpoints for SAP Archive Link, SAP-ILM and non-SAP archiving scenarios.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package archiving

type InstanceParameters struct {
	// Billing address for Service
	// REQUIRED
	BillingAddress *string `json:"billingAddress"`
	// Date of instance creation
	CreatedAt     *string                          `json:"createdAt,omitempty"`
	FiscalYearEnd *InstanceParametersFiscalYearEnd `json:"fiscalYearEnd,omitempty"`
	// REQUIRED
	Kind *string `json:"kind"`
	// Service name
	// REQUIRED
	ServiceName *string `json:"serviceName"`
	// REQUIRED
	System *InstanceParametersSystem `json:"system"`
	// Version of operator kind
	Version *string `json:"version,omitempty"`
}
