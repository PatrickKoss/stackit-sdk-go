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

// checks if the ListFlavorsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFlavorsResponse{}

// ListFlavorsResponse struct for ListFlavorsResponse
type ListFlavorsResponse struct {
	Flavors *[]InstanceFlavorEntry `json:"flavors,omitempty"`
}

// NewListFlavorsResponse instantiates a new ListFlavorsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFlavorsResponse() *ListFlavorsResponse {
	this := ListFlavorsResponse{}
	return &this
}

// NewListFlavorsResponseWithDefaults instantiates a new ListFlavorsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFlavorsResponseWithDefaults() *ListFlavorsResponse {
	this := ListFlavorsResponse{}
	return &this
}

// GetFlavors returns the Flavors field value if set, zero value otherwise.
func (o *ListFlavorsResponse) GetFlavors() *[]InstanceFlavorEntry {
	if o == nil || IsNil(o.Flavors) {
		var ret *[]InstanceFlavorEntry
		return ret
	}
	return o.Flavors
}

// GetFlavorsOk returns a tuple with the Flavors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFlavorsResponse) GetFlavorsOk() (*[]InstanceFlavorEntry, bool) {
	if o == nil || IsNil(o.Flavors) {
		return nil, false
	}
	return o.Flavors, true
}

// HasFlavors returns a boolean if a field has been set.
func (o *ListFlavorsResponse) HasFlavors() bool {
	if o != nil && !IsNil(o.Flavors) {
		return true
	}

	return false
}

// SetFlavors gets a reference to the given []InstanceFlavorEntry and assigns it to the Flavors field.
func (o *ListFlavorsResponse) SetFlavors(v *[]InstanceFlavorEntry) {
	o.Flavors = v
}

func (o ListFlavorsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Flavors) {
		toSerialize["flavors"] = o.Flavors
	}
	return toSerialize, nil
}

type NullableListFlavorsResponse struct {
	value *ListFlavorsResponse
	isSet bool
}

func (v NullableListFlavorsResponse) Get() *ListFlavorsResponse {
	return v.value
}

func (v *NullableListFlavorsResponse) Set(val *ListFlavorsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListFlavorsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListFlavorsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFlavorsResponse(val *ListFlavorsResponse) *NullableListFlavorsResponse {
	return &NullableListFlavorsResponse{value: val, isSet: true}
}

func (v NullableListFlavorsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFlavorsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
