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

// checks if the ListCredentialsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCredentialsResponse{}

// ListCredentialsResponse struct for ListCredentialsResponse
type ListCredentialsResponse struct {
	// REQUIRED
	Credentials *[]ServiceKeysList `json:"credentials"`
	// REQUIRED
	Message *string `json:"message"`
}

type _ListCredentialsResponse ListCredentialsResponse

// NewListCredentialsResponse instantiates a new ListCredentialsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCredentialsResponse(credentials *[]ServiceKeysList, message *string) *ListCredentialsResponse {
	this := ListCredentialsResponse{}
	this.Credentials = credentials
	this.Message = message
	return &this
}

// NewListCredentialsResponseWithDefaults instantiates a new ListCredentialsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCredentialsResponseWithDefaults() *ListCredentialsResponse {
	this := ListCredentialsResponse{}
	return &this
}

// GetCredentials returns the Credentials field value
func (o *ListCredentialsResponse) GetCredentials() *[]ServiceKeysList {
	if o == nil || IsNil(o.Credentials) {
		var ret *[]ServiceKeysList
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *ListCredentialsResponse) GetCredentialsOk() (*[]ServiceKeysList, bool) {
	if o == nil {
		return nil, false
	}
	return o.Credentials, true
}

// SetCredentials sets field value
func (o *ListCredentialsResponse) SetCredentials(v *[]ServiceKeysList) {
	o.Credentials = v
}

// GetMessage returns the Message field value
func (o *ListCredentialsResponse) GetMessage() *string {
	if o == nil || IsNil(o.Message) {
		var ret *string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ListCredentialsResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message, true
}

// SetMessage sets field value
func (o *ListCredentialsResponse) SetMessage(v *string) {
	o.Message = v
}

func (o ListCredentialsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["credentials"] = o.Credentials
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableListCredentialsResponse struct {
	value *ListCredentialsResponse
	isSet bool
}

func (v NullableListCredentialsResponse) Get() *ListCredentialsResponse {
	return v.value
}

func (v *NullableListCredentialsResponse) Set(val *ListCredentialsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListCredentialsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListCredentialsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCredentialsResponse(val *ListCredentialsResponse) *NullableListCredentialsResponse {
	return &NullableListCredentialsResponse{value: val, isSet: true}
}

func (v NullableListCredentialsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCredentialsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
