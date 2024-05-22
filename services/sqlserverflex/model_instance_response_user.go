/*
STACKIT MSSQL Service API

This is the documentation for the STACKIT MSSQL service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlserverflex

type InstanceResponseUser struct {
	Database *string   `json:"database,omitempty"`
	Host     *string   `json:"host,omitempty"`
	Id       *string   `json:"id,omitempty"`
	Port     *int64    `json:"port,omitempty"`
	Roles    *[]string `json:"roles,omitempty"`
	Username *string   `json:"username,omitempty"`
}
