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

// checks if the Receivers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Receivers{}

// Receivers struct for Receivers
type Receivers struct {
	EmailConfigs *[]EmailConfig `json:"emailConfigs,omitempty"`
	// REQUIRED
	Name            *string           `json:"name"`
	OpsgenieConfigs *[]OpsgenieConfig `json:"opsgenieConfigs,omitempty"`
	WebHookConfigs  *[]WebHook        `json:"webHookConfigs,omitempty"`
}

type _Receivers Receivers

// NewReceivers instantiates a new Receivers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceivers(name *string) *Receivers {
	this := Receivers{}
	this.Name = name
	return &this
}

// NewReceiversWithDefaults instantiates a new Receivers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceiversWithDefaults() *Receivers {
	this := Receivers{}
	return &this
}

// GetEmailConfigs returns the EmailConfigs field value if set, zero value otherwise.
func (o *Receivers) GetEmailConfigs() *[]EmailConfig {
	if o == nil || IsNil(o.EmailConfigs) {
		var ret *[]EmailConfig
		return ret
	}
	return o.EmailConfigs
}

// GetEmailConfigsOk returns a tuple with the EmailConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receivers) GetEmailConfigsOk() (*[]EmailConfig, bool) {
	if o == nil || IsNil(o.EmailConfigs) {
		return nil, false
	}
	return o.EmailConfigs, true
}

// HasEmailConfigs returns a boolean if a field has been set.
func (o *Receivers) HasEmailConfigs() bool {
	if o != nil && !IsNil(o.EmailConfigs) {
		return true
	}

	return false
}

// SetEmailConfigs gets a reference to the given []EmailConfig and assigns it to the EmailConfigs field.
func (o *Receivers) SetEmailConfigs(v *[]EmailConfig) {
	o.EmailConfigs = v
}

// GetName returns the Name field value
func (o *Receivers) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Receivers) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *Receivers) SetName(v *string) {
	o.Name = v
}

// GetOpsgenieConfigs returns the OpsgenieConfigs field value if set, zero value otherwise.
func (o *Receivers) GetOpsgenieConfigs() *[]OpsgenieConfig {
	if o == nil || IsNil(o.OpsgenieConfigs) {
		var ret *[]OpsgenieConfig
		return ret
	}
	return o.OpsgenieConfigs
}

// GetOpsgenieConfigsOk returns a tuple with the OpsgenieConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receivers) GetOpsgenieConfigsOk() (*[]OpsgenieConfig, bool) {
	if o == nil || IsNil(o.OpsgenieConfigs) {
		return nil, false
	}
	return o.OpsgenieConfigs, true
}

// HasOpsgenieConfigs returns a boolean if a field has been set.
func (o *Receivers) HasOpsgenieConfigs() bool {
	if o != nil && !IsNil(o.OpsgenieConfigs) {
		return true
	}

	return false
}

// SetOpsgenieConfigs gets a reference to the given []OpsgenieConfig and assigns it to the OpsgenieConfigs field.
func (o *Receivers) SetOpsgenieConfigs(v *[]OpsgenieConfig) {
	o.OpsgenieConfigs = v
}

// GetWebHookConfigs returns the WebHookConfigs field value if set, zero value otherwise.
func (o *Receivers) GetWebHookConfigs() *[]WebHook {
	if o == nil || IsNil(o.WebHookConfigs) {
		var ret *[]WebHook
		return ret
	}
	return o.WebHookConfigs
}

// GetWebHookConfigsOk returns a tuple with the WebHookConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Receivers) GetWebHookConfigsOk() (*[]WebHook, bool) {
	if o == nil || IsNil(o.WebHookConfigs) {
		return nil, false
	}
	return o.WebHookConfigs, true
}

// HasWebHookConfigs returns a boolean if a field has been set.
func (o *Receivers) HasWebHookConfigs() bool {
	if o != nil && !IsNil(o.WebHookConfigs) {
		return true
	}

	return false
}

// SetWebHookConfigs gets a reference to the given []WebHook and assigns it to the WebHookConfigs field.
func (o *Receivers) SetWebHookConfigs(v *[]WebHook) {
	o.WebHookConfigs = v
}

func (o Receivers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailConfigs) {
		toSerialize["emailConfigs"] = o.EmailConfigs
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.OpsgenieConfigs) {
		toSerialize["opsgenieConfigs"] = o.OpsgenieConfigs
	}
	if !IsNil(o.WebHookConfigs) {
		toSerialize["webHookConfigs"] = o.WebHookConfigs
	}
	return toSerialize, nil
}

type NullableReceivers struct {
	value *Receivers
	isSet bool
}

func (v NullableReceivers) Get() *Receivers {
	return v.value
}

func (v *NullableReceivers) Set(val *Receivers) {
	v.value = val
	v.isSet = true
}

func (v NullableReceivers) IsSet() bool {
	return v.isSet
}

func (v *NullableReceivers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceivers(val *Receivers) *NullableReceivers {
	return &NullableReceivers{value: val, isSet: true}
}

func (v NullableReceivers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceivers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
