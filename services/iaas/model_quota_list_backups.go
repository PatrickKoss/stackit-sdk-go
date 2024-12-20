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

// checks if the QuotaListBackups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuotaListBackups{}

// QuotaListBackups Number of backups.
type QuotaListBackups struct {
	// REQUIRED
	Limit *int64 `json:"limit"`
	// REQUIRED
	Usage *int64 `json:"usage"`
}

type _QuotaListBackups QuotaListBackups

// NewQuotaListBackups instantiates a new QuotaListBackups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotaListBackups(limit *int64, usage *int64) *QuotaListBackups {
	this := QuotaListBackups{}
	this.Limit = limit
	this.Usage = usage
	return &this
}

// NewQuotaListBackupsWithDefaults instantiates a new QuotaListBackups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaListBackupsWithDefaults() *QuotaListBackups {
	this := QuotaListBackups{}
	return &this
}

// GetLimit returns the Limit field value
func (o *QuotaListBackups) GetLimit() *int64 {
	if o == nil || IsNil(o.Limit) {
		var ret *int64
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *QuotaListBackups) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit, true
}

// SetLimit sets field value
func (o *QuotaListBackups) SetLimit(v *int64) {
	o.Limit = v
}

// GetUsage returns the Usage field value
func (o *QuotaListBackups) GetUsage() *int64 {
	if o == nil || IsNil(o.Usage) {
		var ret *int64
		return ret
	}

	return o.Usage
}

// GetUsageOk returns a tuple with the Usage field value
// and a boolean to check if the value has been set.
func (o *QuotaListBackups) GetUsageOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Usage, true
}

// SetUsage sets field value
func (o *QuotaListBackups) SetUsage(v *int64) {
	o.Usage = v
}

func (o QuotaListBackups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["limit"] = o.Limit
	toSerialize["usage"] = o.Usage
	return toSerialize, nil
}

type NullableQuotaListBackups struct {
	value *QuotaListBackups
	isSet bool
}

func (v NullableQuotaListBackups) Get() *QuotaListBackups {
	return v.value
}

func (v *NullableQuotaListBackups) Set(val *QuotaListBackups) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaListBackups) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaListBackups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaListBackups(val *QuotaListBackups) *NullableQuotaListBackups {
	return &NullableQuotaListBackups{value: val, isSet: true}
}

func (v NullableQuotaListBackups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaListBackups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
