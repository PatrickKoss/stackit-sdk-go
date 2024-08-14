/*
STACKIT MSSQL Service API

This is the documentation for the STACKIT MSSQL service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlserverflex

// InstanceFlavorEntry struct for InstanceFlavorEntry
type InstanceFlavorEntry struct {
	Categories  *string `json:"categories,omitempty"`
	Cpu         *int64  `json:"cpu,omitempty"`
	Description *string `json:"description,omitempty"`
	Id          *string `json:"id,omitempty"`
	Memory      *int64  `json:"memory,omitempty"`
}
