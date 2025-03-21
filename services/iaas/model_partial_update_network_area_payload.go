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

// checks if the PartialUpdateNetworkAreaPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartialUpdateNetworkAreaPayload{}

// PartialUpdateNetworkAreaPayload struct for PartialUpdateNetworkAreaPayload
type PartialUpdateNetworkAreaPayload struct {
	AddressFamily *UpdateAreaAddressFamily `json:"addressFamily,omitempty"`
	// Object that represents the labels of an object. Regex for keys: `^[a-z]((-|_|[a-z0-9])){0,62}$`. Regex for values: `^(-|_|[a-z0-9]){0,63}$`.
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	Name *string `json:"name,omitempty"`
}

// NewPartialUpdateNetworkAreaPayload instantiates a new PartialUpdateNetworkAreaPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartialUpdateNetworkAreaPayload() *PartialUpdateNetworkAreaPayload {
	this := PartialUpdateNetworkAreaPayload{}
	return &this
}

// NewPartialUpdateNetworkAreaPayloadWithDefaults instantiates a new PartialUpdateNetworkAreaPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartialUpdateNetworkAreaPayloadWithDefaults() *PartialUpdateNetworkAreaPayload {
	this := PartialUpdateNetworkAreaPayload{}
	return &this
}

// GetAddressFamily returns the AddressFamily field value if set, zero value otherwise.
func (o *PartialUpdateNetworkAreaPayload) GetAddressFamily() *UpdateAreaAddressFamily {
	if o == nil || IsNil(o.AddressFamily) {
		var ret *UpdateAreaAddressFamily
		return ret
	}
	return o.AddressFamily
}

// GetAddressFamilyOk returns a tuple with the AddressFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateNetworkAreaPayload) GetAddressFamilyOk() (*UpdateAreaAddressFamily, bool) {
	if o == nil || IsNil(o.AddressFamily) {
		return nil, false
	}
	return o.AddressFamily, true
}

// HasAddressFamily returns a boolean if a field has been set.
func (o *PartialUpdateNetworkAreaPayload) HasAddressFamily() bool {
	if o != nil && !IsNil(o.AddressFamily) {
		return true
	}

	return false
}

// SetAddressFamily gets a reference to the given UpdateAreaAddressFamily and assigns it to the AddressFamily field.
func (o *PartialUpdateNetworkAreaPayload) SetAddressFamily(v *UpdateAreaAddressFamily) {
	o.AddressFamily = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *PartialUpdateNetworkAreaPayload) GetLabels() *map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateNetworkAreaPayload) GetLabelsOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return &map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *PartialUpdateNetworkAreaPayload) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *PartialUpdateNetworkAreaPayload) SetLabels(v *map[string]interface{}) {
	o.Labels = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PartialUpdateNetworkAreaPayload) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialUpdateNetworkAreaPayload) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PartialUpdateNetworkAreaPayload) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PartialUpdateNetworkAreaPayload) SetName(v *string) {
	o.Name = v
}

func (o PartialUpdateNetworkAreaPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressFamily) {
		toSerialize["addressFamily"] = o.AddressFamily
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullablePartialUpdateNetworkAreaPayload struct {
	value *PartialUpdateNetworkAreaPayload
	isSet bool
}

func (v NullablePartialUpdateNetworkAreaPayload) Get() *PartialUpdateNetworkAreaPayload {
	return v.value
}

func (v *NullablePartialUpdateNetworkAreaPayload) Set(val *PartialUpdateNetworkAreaPayload) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialUpdateNetworkAreaPayload) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialUpdateNetworkAreaPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialUpdateNetworkAreaPayload(val *PartialUpdateNetworkAreaPayload) *NullablePartialUpdateNetworkAreaPayload {
	return &NullablePartialUpdateNetworkAreaPayload{value: val, isSet: true}
}

func (v NullablePartialUpdateNetworkAreaPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialUpdateNetworkAreaPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
