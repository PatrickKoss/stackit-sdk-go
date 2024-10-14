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

// checks if the UpdateBackupSchedulePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateBackupSchedulePayload{}

// UpdateBackupSchedulePayload struct for UpdateBackupSchedulePayload
type UpdateBackupSchedulePayload struct {
	// REQUIRED
	BackupSchedule *string `json:"backupSchedule"`
}

type _UpdateBackupSchedulePayload UpdateBackupSchedulePayload

// NewUpdateBackupSchedulePayload instantiates a new UpdateBackupSchedulePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateBackupSchedulePayload(backupSchedule *string) *UpdateBackupSchedulePayload {
	this := UpdateBackupSchedulePayload{}
	this.BackupSchedule = backupSchedule
	return &this
}

// NewUpdateBackupSchedulePayloadWithDefaults instantiates a new UpdateBackupSchedulePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateBackupSchedulePayloadWithDefaults() *UpdateBackupSchedulePayload {
	this := UpdateBackupSchedulePayload{}
	return &this
}

// GetBackupSchedule returns the BackupSchedule field value
func (o *UpdateBackupSchedulePayload) GetBackupSchedule() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.BackupSchedule
}

// GetBackupScheduleOk returns a tuple with the BackupSchedule field value
// and a boolean to check if the value has been set.
func (o *UpdateBackupSchedulePayload) GetBackupScheduleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackupSchedule, true
}

// SetBackupSchedule sets field value
func (o *UpdateBackupSchedulePayload) SetBackupSchedule(v *string) {
	o.BackupSchedule = v
}

func (o UpdateBackupSchedulePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["backupSchedule"] = o.BackupSchedule
	return toSerialize, nil
}

type NullableUpdateBackupSchedulePayload struct {
	value *UpdateBackupSchedulePayload
	isSet bool
}

func (v NullableUpdateBackupSchedulePayload) Get() *UpdateBackupSchedulePayload {
	return v.value
}

func (v *NullableUpdateBackupSchedulePayload) Set(val *UpdateBackupSchedulePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateBackupSchedulePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateBackupSchedulePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateBackupSchedulePayload(val *UpdateBackupSchedulePayload) *NullableUpdateBackupSchedulePayload {
	return &NullableUpdateBackupSchedulePayload{value: val, isSet: true}
}

func (v NullableUpdateBackupSchedulePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateBackupSchedulePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
