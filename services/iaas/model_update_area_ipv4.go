/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the UpdateAreaIPv4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAreaIPv4{}

// UpdateAreaIPv4 The update object for a IPv4 network area.
type UpdateAreaIPv4 struct {
	DefaultNameservers *[]string `json:"defaultNameservers,omitempty"`
	// The default prefix length for networks in the network area.
	DefaultPrefixLen *int64 `json:"defaultPrefixLen,omitempty"`
	// The maximal prefix length for networks in the network area.
	MaxPrefixLen *int64 `json:"maxPrefixLen,omitempty"`
	// The minimal prefix length for networks in the network area.
	MinPrefixLen *int64 `json:"minPrefixLen,omitempty"`
}

// NewUpdateAreaIPv4 instantiates a new UpdateAreaIPv4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAreaIPv4() *UpdateAreaIPv4 {
	this := UpdateAreaIPv4{}
	return &this
}

// NewUpdateAreaIPv4WithDefaults instantiates a new UpdateAreaIPv4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAreaIPv4WithDefaults() *UpdateAreaIPv4 {
	this := UpdateAreaIPv4{}
	return &this
}

// GetDefaultNameservers returns the DefaultNameservers field value if set, zero value otherwise.
func (o *UpdateAreaIPv4) GetDefaultNameservers() *[]string {
	if o == nil || IsNil(o.DefaultNameservers) {
		var ret *[]string
		return ret
	}
	return o.DefaultNameservers
}

// GetDefaultNameserversOk returns a tuple with the DefaultNameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAreaIPv4) GetDefaultNameserversOk() (*[]string, bool) {
	if o == nil || IsNil(o.DefaultNameservers) {
		return nil, false
	}
	return o.DefaultNameservers, true
}

// HasDefaultNameservers returns a boolean if a field has been set.
func (o *UpdateAreaIPv4) HasDefaultNameservers() bool {
	if o != nil && !IsNil(o.DefaultNameservers) {
		return true
	}

	return false
}

// SetDefaultNameservers gets a reference to the given []string and assigns it to the DefaultNameservers field.
func (o *UpdateAreaIPv4) SetDefaultNameservers(v *[]string) {
	o.DefaultNameservers = v
}

// GetDefaultPrefixLen returns the DefaultPrefixLen field value if set, zero value otherwise.
func (o *UpdateAreaIPv4) GetDefaultPrefixLen() *int64 {
	if o == nil || IsNil(o.DefaultPrefixLen) {
		var ret *int64
		return ret
	}
	return o.DefaultPrefixLen
}

// GetDefaultPrefixLenOk returns a tuple with the DefaultPrefixLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAreaIPv4) GetDefaultPrefixLenOk() (*int64, bool) {
	if o == nil || IsNil(o.DefaultPrefixLen) {
		return nil, false
	}
	return o.DefaultPrefixLen, true
}

// HasDefaultPrefixLen returns a boolean if a field has been set.
func (o *UpdateAreaIPv4) HasDefaultPrefixLen() bool {
	if o != nil && !IsNil(o.DefaultPrefixLen) {
		return true
	}

	return false
}

// SetDefaultPrefixLen gets a reference to the given int64 and assigns it to the DefaultPrefixLen field.
func (o *UpdateAreaIPv4) SetDefaultPrefixLen(v *int64) {
	o.DefaultPrefixLen = v
}

// GetMaxPrefixLen returns the MaxPrefixLen field value if set, zero value otherwise.
func (o *UpdateAreaIPv4) GetMaxPrefixLen() *int64 {
	if o == nil || IsNil(o.MaxPrefixLen) {
		var ret *int64
		return ret
	}
	return o.MaxPrefixLen
}

// GetMaxPrefixLenOk returns a tuple with the MaxPrefixLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAreaIPv4) GetMaxPrefixLenOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxPrefixLen) {
		return nil, false
	}
	return o.MaxPrefixLen, true
}

// HasMaxPrefixLen returns a boolean if a field has been set.
func (o *UpdateAreaIPv4) HasMaxPrefixLen() bool {
	if o != nil && !IsNil(o.MaxPrefixLen) {
		return true
	}

	return false
}

// SetMaxPrefixLen gets a reference to the given int64 and assigns it to the MaxPrefixLen field.
func (o *UpdateAreaIPv4) SetMaxPrefixLen(v *int64) {
	o.MaxPrefixLen = v
}

// GetMinPrefixLen returns the MinPrefixLen field value if set, zero value otherwise.
func (o *UpdateAreaIPv4) GetMinPrefixLen() *int64 {
	if o == nil || IsNil(o.MinPrefixLen) {
		var ret *int64
		return ret
	}
	return o.MinPrefixLen
}

// GetMinPrefixLenOk returns a tuple with the MinPrefixLen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAreaIPv4) GetMinPrefixLenOk() (*int64, bool) {
	if o == nil || IsNil(o.MinPrefixLen) {
		return nil, false
	}
	return o.MinPrefixLen, true
}

// HasMinPrefixLen returns a boolean if a field has been set.
func (o *UpdateAreaIPv4) HasMinPrefixLen() bool {
	if o != nil && !IsNil(o.MinPrefixLen) {
		return true
	}

	return false
}

// SetMinPrefixLen gets a reference to the given int64 and assigns it to the MinPrefixLen field.
func (o *UpdateAreaIPv4) SetMinPrefixLen(v *int64) {
	o.MinPrefixLen = v
}

func (o UpdateAreaIPv4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultNameservers) {
		toSerialize["defaultNameservers"] = o.DefaultNameservers
	}
	if !IsNil(o.DefaultPrefixLen) {
		toSerialize["defaultPrefixLen"] = o.DefaultPrefixLen
	}
	if !IsNil(o.MaxPrefixLen) {
		toSerialize["maxPrefixLen"] = o.MaxPrefixLen
	}
	if !IsNil(o.MinPrefixLen) {
		toSerialize["minPrefixLen"] = o.MinPrefixLen
	}
	return toSerialize, nil
}

type NullableUpdateAreaIPv4 struct {
	value *UpdateAreaIPv4
	isSet bool
}

func (v NullableUpdateAreaIPv4) Get() *UpdateAreaIPv4 {
	return v.value
}

func (v *NullableUpdateAreaIPv4) Set(val *UpdateAreaIPv4) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAreaIPv4) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAreaIPv4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAreaIPv4(val *UpdateAreaIPv4) *NullableUpdateAreaIPv4 {
	return &NullableUpdateAreaIPv4{value: val, isSet: true}
}

func (v NullableUpdateAreaIPv4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAreaIPv4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
