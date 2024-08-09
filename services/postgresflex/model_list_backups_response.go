/*
STACKIT PostgreSQL Flex API

This is the documentation for the STACKIT postgres service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package postgresflex

import (
	"encoding/json"
)

// checks if the ListBackupsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListBackupsResponse{}

// ListBackupsResponse struct for ListBackupsResponse
type ListBackupsResponse struct {
	Count *int64    `json:"count,omitempty"`
	Items *[]Backup `json:"items,omitempty"`
}

// NewListBackupsResponse instantiates a new ListBackupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListBackupsResponse() *ListBackupsResponse {
	this := ListBackupsResponse{}
	return &this
}

// NewListBackupsResponseWithDefaults instantiates a new ListBackupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListBackupsResponseWithDefaults() *ListBackupsResponse {
	this := ListBackupsResponse{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ListBackupsResponse) GetCount() *int64 {
	if o == nil || IsNil(o.Count) {
		var ret *int64
		return ret
	}
	return o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBackupsResponse) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ListBackupsResponse) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ListBackupsResponse) SetCount(v *int64) {
	o.Count = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListBackupsResponse) GetItems() *[]Backup {
	if o == nil || IsNil(o.Items) {
		var ret *[]Backup
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBackupsResponse) GetItemsOk() (*[]Backup, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListBackupsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Backup and assigns it to the Items field.
func (o *ListBackupsResponse) SetItems(v *[]Backup) {
	o.Items = v
}

func (o ListBackupsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableListBackupsResponse struct {
	value *ListBackupsResponse
	isSet bool
}

func (v NullableListBackupsResponse) Get() *ListBackupsResponse {
	return v.value
}

func (v *NullableListBackupsResponse) Set(val *ListBackupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListBackupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListBackupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListBackupsResponse(val *ListBackupsResponse) *NullableListBackupsResponse {
	return &NullableListBackupsResponse{value: val, isSet: true}
}

func (v NullableListBackupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListBackupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
