/*
SKE-API

The SKE API provides endpoints to create, update, delete clusters within STACKIT portal projects and to trigger further cluster management tasks.

API version: 1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ske

import (
	"encoding/json"
	"time"
)

// checks if the MachineImageVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MachineImageVersion{}

// MachineImageVersion struct for MachineImageVersion
type MachineImageVersion struct {
	Cri            *[]CRI     `json:"cri,omitempty"`
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
	State          *string    `json:"state,omitempty"`
	Version        *string    `json:"version,omitempty"`
}

// NewMachineImageVersion instantiates a new MachineImageVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMachineImageVersion() *MachineImageVersion {
	this := MachineImageVersion{}
	return &this
}

// NewMachineImageVersionWithDefaults instantiates a new MachineImageVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineImageVersionWithDefaults() *MachineImageVersion {
	this := MachineImageVersion{}
	return &this
}

// GetCri returns the Cri field value if set, zero value otherwise.
func (o *MachineImageVersion) GetCri() *[]CRI {
	if o == nil || IsNil(o.Cri) {
		var ret *[]CRI
		return ret
	}
	return o.Cri
}

// GetCriOk returns a tuple with the Cri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineImageVersion) GetCriOk() (*[]CRI, bool) {
	if o == nil || IsNil(o.Cri) {
		return nil, false
	}
	return o.Cri, true
}

// HasCri returns a boolean if a field has been set.
func (o *MachineImageVersion) HasCri() bool {
	if o != nil && !IsNil(o.Cri) {
		return true
	}

	return false
}

// SetCri gets a reference to the given []CRI and assigns it to the Cri field.
func (o *MachineImageVersion) SetCri(v *[]CRI) {
	o.Cri = v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *MachineImageVersion) GetExpirationDate() *time.Time {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret *time.Time
		return ret
	}
	return o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineImageVersion) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *MachineImageVersion) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given time.Time and assigns it to the ExpirationDate field.
func (o *MachineImageVersion) SetExpirationDate(v *time.Time) {
	o.ExpirationDate = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *MachineImageVersion) GetState() *string {
	if o == nil || IsNil(o.State) {
		var ret *string
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineImageVersion) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MachineImageVersion) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *MachineImageVersion) SetState(v *string) {
	o.State = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *MachineImageVersion) GetVersion() *string {
	if o == nil || IsNil(o.Version) {
		var ret *string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineImageVersion) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *MachineImageVersion) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *MachineImageVersion) SetVersion(v *string) {
	o.Version = v
}

func (o MachineImageVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cri) {
		toSerialize["cri"] = o.Cri
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableMachineImageVersion struct {
	value *MachineImageVersion
	isSet bool
}

func (v NullableMachineImageVersion) Get() *MachineImageVersion {
	return v.value
}

func (v *NullableMachineImageVersion) Set(val *MachineImageVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineImageVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineImageVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineImageVersion(val *MachineImageVersion) *NullableMachineImageVersion {
	return &NullableMachineImageVersion{value: val, isSet: true}
}

func (v NullableMachineImageVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineImageVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
