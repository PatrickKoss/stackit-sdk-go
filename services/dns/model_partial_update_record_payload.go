/*
STACKIT DNS API

This api provides dns

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the PartialUpdateRecordPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartialUpdateRecordPayload{}

// PartialUpdateRecordPayload RecordPatch for record patch in record set.
type PartialUpdateRecordPayload struct {
	// REQUIRED
	Action *string `json:"action"`
	// records
	// REQUIRED
	Records *[]RecordPayload `json:"records"`
}

type _PartialUpdateRecordPayload PartialUpdateRecordPayload

// NewPartialUpdateRecordPayload instantiates a new PartialUpdateRecordPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartialUpdateRecordPayload(action *string, records *[]RecordPayload) *PartialUpdateRecordPayload {
	this := PartialUpdateRecordPayload{}
	this.Action = action
	this.Records = records
	return &this
}

// NewPartialUpdateRecordPayloadWithDefaults instantiates a new PartialUpdateRecordPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartialUpdateRecordPayloadWithDefaults() *PartialUpdateRecordPayload {
	this := PartialUpdateRecordPayload{}
	return &this
}

// GetAction returns the Action field value
func (o *PartialUpdateRecordPayload) GetAction() *string {
	if o == nil || IsNil(o.Action) {
		var ret *string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *PartialUpdateRecordPayload) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action, true
}

// SetAction sets field value
func (o *PartialUpdateRecordPayload) SetAction(v *string) {
	o.Action = v
}

// GetRecords returns the Records field value
func (o *PartialUpdateRecordPayload) GetRecords() *[]RecordPayload {
	if o == nil || IsNil(o.Records) {
		var ret *[]RecordPayload
		return ret
	}

	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value
// and a boolean to check if the value has been set.
func (o *PartialUpdateRecordPayload) GetRecordsOk() (*[]RecordPayload, bool) {
	if o == nil {
		return nil, false
	}
	return o.Records, true
}

// SetRecords sets field value
func (o *PartialUpdateRecordPayload) SetRecords(v *[]RecordPayload) {
	o.Records = v
}

func (o PartialUpdateRecordPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["records"] = o.Records
	return toSerialize, nil
}

type NullablePartialUpdateRecordPayload struct {
	value *PartialUpdateRecordPayload
	isSet bool
}

func (v NullablePartialUpdateRecordPayload) Get() *PartialUpdateRecordPayload {
	return v.value
}

func (v *NullablePartialUpdateRecordPayload) Set(val *PartialUpdateRecordPayload) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialUpdateRecordPayload) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialUpdateRecordPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialUpdateRecordPayload(val *PartialUpdateRecordPayload) *NullablePartialUpdateRecordPayload {
	return &NullablePartialUpdateRecordPayload{value: val, isSet: true}
}

func (v NullablePartialUpdateRecordPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialUpdateRecordPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
