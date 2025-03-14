/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

import (
	"encoding/json"
)

// checks if the CredentialsRemoteWriteDeleteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialsRemoteWriteDeleteResponse{}

// CredentialsRemoteWriteDeleteResponse struct for CredentialsRemoteWriteDeleteResponse
type CredentialsRemoteWriteDeleteResponse struct {
	// REQUIRED
	MaxLimit *int64 `json:"maxLimit"`
	// REQUIRED
	Message *string `json:"message"`
}

type _CredentialsRemoteWriteDeleteResponse CredentialsRemoteWriteDeleteResponse

// NewCredentialsRemoteWriteDeleteResponse instantiates a new CredentialsRemoteWriteDeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsRemoteWriteDeleteResponse(maxLimit *int64, message *string) *CredentialsRemoteWriteDeleteResponse {
	this := CredentialsRemoteWriteDeleteResponse{}
	this.MaxLimit = maxLimit
	this.Message = message
	return &this
}

// NewCredentialsRemoteWriteDeleteResponseWithDefaults instantiates a new CredentialsRemoteWriteDeleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsRemoteWriteDeleteResponseWithDefaults() *CredentialsRemoteWriteDeleteResponse {
	this := CredentialsRemoteWriteDeleteResponse{}
	return &this
}

// GetMaxLimit returns the MaxLimit field value
func (o *CredentialsRemoteWriteDeleteResponse) GetMaxLimit() *int64 {
	if o == nil || IsNil(o.MaxLimit) {
		var ret *int64
		return ret
	}

	return o.MaxLimit
}

// GetMaxLimitOk returns a tuple with the MaxLimit field value
// and a boolean to check if the value has been set.
func (o *CredentialsRemoteWriteDeleteResponse) GetMaxLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxLimit, true
}

// SetMaxLimit sets field value
func (o *CredentialsRemoteWriteDeleteResponse) SetMaxLimit(v *int64) {
	o.MaxLimit = v
}

// GetMessage returns the Message field value
func (o *CredentialsRemoteWriteDeleteResponse) GetMessage() *string {
	if o == nil || IsNil(o.Message) {
		var ret *string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *CredentialsRemoteWriteDeleteResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message, true
}

// SetMessage sets field value
func (o *CredentialsRemoteWriteDeleteResponse) SetMessage(v *string) {
	o.Message = v
}

func (o CredentialsRemoteWriteDeleteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["maxLimit"] = o.MaxLimit
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableCredentialsRemoteWriteDeleteResponse struct {
	value *CredentialsRemoteWriteDeleteResponse
	isSet bool
}

func (v NullableCredentialsRemoteWriteDeleteResponse) Get() *CredentialsRemoteWriteDeleteResponse {
	return v.value
}

func (v *NullableCredentialsRemoteWriteDeleteResponse) Set(val *CredentialsRemoteWriteDeleteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsRemoteWriteDeleteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsRemoteWriteDeleteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsRemoteWriteDeleteResponse(val *CredentialsRemoteWriteDeleteResponse) *NullableCredentialsRemoteWriteDeleteResponse {
	return &NullableCredentialsRemoteWriteDeleteResponse{value: val, isSet: true}
}

func (v NullableCredentialsRemoteWriteDeleteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsRemoteWriteDeleteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
