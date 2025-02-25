/*
STACKIT Model Serving API

This API provides endpoints for the model serving api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package modelserving

import (
	"encoding/json"
)

// checks if the ErrorMessageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorMessageResponse{}

// ErrorMessageResponse struct for ErrorMessageResponse
type ErrorMessageResponse struct {
	Error   *string `json:"error,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewErrorMessageResponse instantiates a new ErrorMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorMessageResponse() *ErrorMessageResponse {
	this := ErrorMessageResponse{}
	return &this
}

// NewErrorMessageResponseWithDefaults instantiates a new ErrorMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorMessageResponseWithDefaults() *ErrorMessageResponse {
	this := ErrorMessageResponse{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ErrorMessageResponse) GetError() *string {
	if o == nil || IsNil(o.Error) {
		var ret *string
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorMessageResponse) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ErrorMessageResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ErrorMessageResponse) SetError(v *string) {
	o.Error = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorMessageResponse) GetMessage() *string {
	if o == nil || IsNil(o.Message) {
		var ret *string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorMessageResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorMessageResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorMessageResponse) SetMessage(v *string) {
	o.Message = v
}

func (o ErrorMessageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableErrorMessageResponse struct {
	value *ErrorMessageResponse
	isSet bool
}

func (v NullableErrorMessageResponse) Get() *ErrorMessageResponse {
	return v.value
}

func (v *NullableErrorMessageResponse) Set(val *ErrorMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorMessageResponse(val *ErrorMessageResponse) *NullableErrorMessageResponse {
	return &NullableErrorMessageResponse{value: val, isSet: true}
}

func (v NullableErrorMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
