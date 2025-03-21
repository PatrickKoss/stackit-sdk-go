/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"encoding/json"
)

// checks if the PublicIp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicIp{}

// PublicIp Object that represents a public IP.
type PublicIp struct {
	// Universally Unique Identifier (UUID).
	Id *string `json:"id,omitempty"`
	// Object that represents an IP address.
	Ip *string `json:"ip,omitempty"`
	// Object that represents the labels of an object. Regex for keys: `^[a-z]((-|_|[a-z0-9])){0,62}$`. Regex for values: `^(-|_|[a-z0-9]){0,63}$`.
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// Universally Unique Identifier (UUID).
	NetworkInterface *NullableString `json:"networkInterface,omitempty"`
}

// NewPublicIp instantiates a new PublicIp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicIp() *PublicIp {
	this := PublicIp{}
	return &this
}

// NewPublicIpWithDefaults instantiates a new PublicIp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicIpWithDefaults() *PublicIp {
	this := PublicIp{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PublicIp) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIp) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PublicIp) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PublicIp) SetId(v *string) {
	o.Id = v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *PublicIp) GetIp() *string {
	if o == nil || IsNil(o.Ip) {
		var ret *string
		return ret
	}
	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIp) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *PublicIp) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *PublicIp) SetIp(v *string) {
	o.Ip = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *PublicIp) GetLabels() *map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIp) GetLabelsOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return &map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *PublicIp) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *PublicIp) SetLabels(v *map[string]interface{}) {
	o.Labels = v
}

// GetNetworkInterface returns the NetworkInterface field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicIp) GetNetworkInterface() *string {
	if o == nil || IsNil(o.NetworkInterface) || IsNil(o.NetworkInterface.Get()) {
		var ret *string
		return ret
	}
	return o.NetworkInterface.Get()
}

// GetNetworkInterfaceOk returns a tuple with the NetworkInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicIp) GetNetworkInterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkInterface) {
		return nil, false
	}
	return o.NetworkInterface.Get(), o.NetworkInterface.IsSet()
}

// HasNetworkInterface returns a boolean if a field has been set.
func (o *PublicIp) HasNetworkInterface() bool {
	if o != nil && !IsNil(o.NetworkInterface) && o.NetworkInterface.IsSet() {
		return true
	}

	return false
}

// SetNetworkInterface gets a reference to the given string and assigns it to the NetworkInterface field.
func (o *PublicIp) SetNetworkInterface(v *string) {
	if IsNil(o.NetworkInterface) {
		o.NetworkInterface = new(NullableString)
	}
	o.NetworkInterface.Set(v)
}

// SetNetworkInterfaceNil sets the value for NetworkInterface to be an explicit nil
func (o *PublicIp) SetNetworkInterfaceNil() {
	if IsNil(o.NetworkInterface) {
		o.NetworkInterface = new(NullableString)
	}
	o.NetworkInterface.Set(nil)
}

// UnsetNetworkInterface ensures that no value is present for NetworkInterface, not even an explicit nil
func (o *PublicIp) UnsetNetworkInterface() {
	if IsNil(o.NetworkInterface) {
		o.NetworkInterface = new(NullableString)
	}
	o.NetworkInterface.Unset()
}

func (o PublicIp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if o.NetworkInterface.IsSet() {
		toSerialize["networkInterface"] = o.NetworkInterface.Get()
	}
	return toSerialize, nil
}

type NullablePublicIp struct {
	value *PublicIp
	isSet bool
}

func (v NullablePublicIp) Get() *PublicIp {
	return v.value
}

func (v *NullablePublicIp) Set(val *PublicIp) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicIp) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicIp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicIp(val *PublicIp) *NullablePublicIp {
	return &NullablePublicIp{value: val, isSet: true}
}

func (v NullablePublicIp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicIp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
