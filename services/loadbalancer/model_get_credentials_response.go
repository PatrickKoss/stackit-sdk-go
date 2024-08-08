/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1.7.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

import (
	"encoding/json"
)

// checks if the GetCredentialsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCredentialsResponse{}

// GetCredentialsResponse struct for GetCredentialsResponse
type GetCredentialsResponse struct {
	Credential *CredentialsResponse `json:"credential,omitempty"`
}

// NewGetCredentialsResponse instantiates a new GetCredentialsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCredentialsResponse() *GetCredentialsResponse {
	this := GetCredentialsResponse{}
	return &this
}

// NewGetCredentialsResponseWithDefaults instantiates a new GetCredentialsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCredentialsResponseWithDefaults() *GetCredentialsResponse {
	this := GetCredentialsResponse{}
	return &this
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *GetCredentialsResponse) GetCredential() *CredentialsResponse {
	if o == nil || IsNil(o.Credential) {
		var ret *CredentialsResponse
		return ret
	}
	return o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCredentialsResponse) GetCredentialOk() (*CredentialsResponse, bool) {
	if o == nil || IsNil(o.Credential) {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *GetCredentialsResponse) HasCredential() bool {
	if o != nil && !IsNil(o.Credential) {
		return true
	}

	return false
}

// SetCredential gets a reference to the given CredentialsResponse and assigns it to the Credential field.
func (o *GetCredentialsResponse) SetCredential(v *CredentialsResponse) {
	o.Credential = v
}

func (o GetCredentialsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Credential) {
		toSerialize["credential"] = o.Credential
	}
	return toSerialize, nil
}

type NullableGetCredentialsResponse struct {
	value *GetCredentialsResponse
	isSet bool
}

func (v NullableGetCredentialsResponse) Get() *GetCredentialsResponse {
	return v.value
}

func (v *NullableGetCredentialsResponse) Set(val *GetCredentialsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCredentialsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCredentialsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCredentialsResponse(val *GetCredentialsResponse) *NullableGetCredentialsResponse {
	return &NullableGetCredentialsResponse{value: val, isSet: true}
}

func (v NullableGetCredentialsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCredentialsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
