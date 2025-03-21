/*
STACKIT MSSQL Service API

This is the documentation for the STACKIT MSSQL service

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sqlserverflex

import (
	"encoding/json"
)

// checks if the CreateDatabasePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDatabasePayload{}

// CreateDatabasePayload struct for CreateDatabasePayload
type CreateDatabasePayload struct {
	// REQUIRED
	Name *string `json:"name"`
	// REQUIRED
	Options *DatabaseDocumentationCreateDatabaseRequestOptions `json:"options"`
}

type _CreateDatabasePayload CreateDatabasePayload

// NewCreateDatabasePayload instantiates a new CreateDatabasePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDatabasePayload(name *string, options *DatabaseDocumentationCreateDatabaseRequestOptions) *CreateDatabasePayload {
	this := CreateDatabasePayload{}
	this.Name = name
	this.Options = options
	return &this
}

// NewCreateDatabasePayloadWithDefaults instantiates a new CreateDatabasePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDatabasePayloadWithDefaults() *CreateDatabasePayload {
	this := CreateDatabasePayload{}
	return &this
}

// GetName returns the Name field value
func (o *CreateDatabasePayload) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateDatabasePayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *CreateDatabasePayload) SetName(v *string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *CreateDatabasePayload) GetOptions() *DatabaseDocumentationCreateDatabaseRequestOptions {
	if o == nil || IsNil(o.Options) {
		var ret *DatabaseDocumentationCreateDatabaseRequestOptions
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *CreateDatabasePayload) GetOptionsOk() (*DatabaseDocumentationCreateDatabaseRequestOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *CreateDatabasePayload) SetOptions(v *DatabaseDocumentationCreateDatabaseRequestOptions) {
	o.Options = v
}

func (o CreateDatabasePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["options"] = o.Options
	return toSerialize, nil
}

type NullableCreateDatabasePayload struct {
	value *CreateDatabasePayload
	isSet bool
}

func (v NullableCreateDatabasePayload) Get() *CreateDatabasePayload {
	return v.value
}

func (v *NullableCreateDatabasePayload) Set(val *CreateDatabasePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDatabasePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDatabasePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDatabasePayload(val *CreateDatabasePayload) *NullableCreateDatabasePayload {
	return &NullableCreateDatabasePayload{value: val, isSet: true}
}

func (v NullableCreateDatabasePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDatabasePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
