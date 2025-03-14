/*
STACKIT Server Update Management API

API endpoints for Server Update Operations on STACKIT Servers.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serverupdate

import (
	"encoding/json"
)

// checks if the UpdateScheduleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateScheduleAllOf{}

// UpdateScheduleAllOf struct for UpdateScheduleAllOf
type UpdateScheduleAllOf struct {
	// Can be cast to int32 without loss of precision.
	// REQUIRED
	Id *int64 `json:"id"`
}

type _UpdateScheduleAllOf UpdateScheduleAllOf

// NewUpdateScheduleAllOf instantiates a new UpdateScheduleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateScheduleAllOf(id *int64) *UpdateScheduleAllOf {
	this := UpdateScheduleAllOf{}
	this.Id = id
	return &this
}

// NewUpdateScheduleAllOfWithDefaults instantiates a new UpdateScheduleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateScheduleAllOfWithDefaults() *UpdateScheduleAllOf {
	this := UpdateScheduleAllOf{}
	return &this
}

// GetId returns the Id field value
func (o *UpdateScheduleAllOf) GetId() *int64 {
	if o == nil || IsNil(o.Id) {
		var ret *int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateScheduleAllOf) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id, true
}

// SetId sets field value
func (o *UpdateScheduleAllOf) SetId(v *int64) {
	o.Id = v
}

func (o UpdateScheduleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableUpdateScheduleAllOf struct {
	value *UpdateScheduleAllOf
	isSet bool
}

func (v NullableUpdateScheduleAllOf) Get() *UpdateScheduleAllOf {
	return v.value
}

func (v *NullableUpdateScheduleAllOf) Set(val *UpdateScheduleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateScheduleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateScheduleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateScheduleAllOf(val *UpdateScheduleAllOf) *NullableUpdateScheduleAllOf {
	return &NullableUpdateScheduleAllOf{value: val, isSet: true}
}

func (v NullableUpdateScheduleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateScheduleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
