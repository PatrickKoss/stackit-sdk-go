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

// checks if the UpdateMetricsStorageRetentionPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMetricsStorageRetentionPayload{}

// UpdateMetricsStorageRetentionPayload struct for UpdateMetricsStorageRetentionPayload
type UpdateMetricsStorageRetentionPayload struct {
	// Retention time of longtime storage of 1h sampled data. After that time the data will be deleted permanently. `Additional Validators:` * Should be a valid time string * Should not be bigger than metricsRetentionTime5m
	// REQUIRED
	MetricsRetentionTime1h *string `json:"metricsRetentionTime1h"`
	// Retention time of longtime storage of 5m sampled data. After that time the data will be down sampled to 1h. `Additional Validators:` * Should be a valid time string * Should not be bigger than metricsRetentionTimeRaw
	// REQUIRED
	MetricsRetentionTime5m *string `json:"metricsRetentionTime5m"`
	// Retention time of longtime storage of raw sampled data. After that time the data will be down sampled to 5m. Keep in mind, that the initial goal of downsampling is not saving disk or object storage space. In fact, downsampling doesn't save you any space but instead, it adds 2 more blocks for each raw block which are only slightly smaller or relatively similar size to raw block. This is done by internal downsampling implementation which to be mathematically correct holds various aggregations. This means that downsampling can increase the size of your storage a bit (~3x), if you choose to store all resolutions (recommended). The goal of downsampling is to provide an opportunity to get fast results for range queries of big time intervals like months or years. `Additional Validators:` * Should be a valid time string * Should not be bigger than 13 months
	// REQUIRED
	MetricsRetentionTimeRaw *string `json:"metricsRetentionTimeRaw"`
}

type _UpdateMetricsStorageRetentionPayload UpdateMetricsStorageRetentionPayload

// NewUpdateMetricsStorageRetentionPayload instantiates a new UpdateMetricsStorageRetentionPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMetricsStorageRetentionPayload(metricsRetentionTime1h *string, metricsRetentionTime5m *string, metricsRetentionTimeRaw *string) *UpdateMetricsStorageRetentionPayload {
	this := UpdateMetricsStorageRetentionPayload{}
	this.MetricsRetentionTime1h = metricsRetentionTime1h
	this.MetricsRetentionTime5m = metricsRetentionTime5m
	this.MetricsRetentionTimeRaw = metricsRetentionTimeRaw
	return &this
}

// NewUpdateMetricsStorageRetentionPayloadWithDefaults instantiates a new UpdateMetricsStorageRetentionPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMetricsStorageRetentionPayloadWithDefaults() *UpdateMetricsStorageRetentionPayload {
	this := UpdateMetricsStorageRetentionPayload{}
	return &this
}

// GetMetricsRetentionTime1h returns the MetricsRetentionTime1h field value
func (o *UpdateMetricsStorageRetentionPayload) GetMetricsRetentionTime1h() *string {
	if o == nil || IsNil(o.MetricsRetentionTime1h) {
		var ret *string
		return ret
	}

	return o.MetricsRetentionTime1h
}

// GetMetricsRetentionTime1hOk returns a tuple with the MetricsRetentionTime1h field value
// and a boolean to check if the value has been set.
func (o *UpdateMetricsStorageRetentionPayload) GetMetricsRetentionTime1hOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetricsRetentionTime1h, true
}

// SetMetricsRetentionTime1h sets field value
func (o *UpdateMetricsStorageRetentionPayload) SetMetricsRetentionTime1h(v *string) {
	o.MetricsRetentionTime1h = v
}

// GetMetricsRetentionTime5m returns the MetricsRetentionTime5m field value
func (o *UpdateMetricsStorageRetentionPayload) GetMetricsRetentionTime5m() *string {
	if o == nil || IsNil(o.MetricsRetentionTime5m) {
		var ret *string
		return ret
	}

	return o.MetricsRetentionTime5m
}

// GetMetricsRetentionTime5mOk returns a tuple with the MetricsRetentionTime5m field value
// and a boolean to check if the value has been set.
func (o *UpdateMetricsStorageRetentionPayload) GetMetricsRetentionTime5mOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetricsRetentionTime5m, true
}

// SetMetricsRetentionTime5m sets field value
func (o *UpdateMetricsStorageRetentionPayload) SetMetricsRetentionTime5m(v *string) {
	o.MetricsRetentionTime5m = v
}

// GetMetricsRetentionTimeRaw returns the MetricsRetentionTimeRaw field value
func (o *UpdateMetricsStorageRetentionPayload) GetMetricsRetentionTimeRaw() *string {
	if o == nil || IsNil(o.MetricsRetentionTimeRaw) {
		var ret *string
		return ret
	}

	return o.MetricsRetentionTimeRaw
}

// GetMetricsRetentionTimeRawOk returns a tuple with the MetricsRetentionTimeRaw field value
// and a boolean to check if the value has been set.
func (o *UpdateMetricsStorageRetentionPayload) GetMetricsRetentionTimeRawOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetricsRetentionTimeRaw, true
}

// SetMetricsRetentionTimeRaw sets field value
func (o *UpdateMetricsStorageRetentionPayload) SetMetricsRetentionTimeRaw(v *string) {
	o.MetricsRetentionTimeRaw = v
}

func (o UpdateMetricsStorageRetentionPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metricsRetentionTime1h"] = o.MetricsRetentionTime1h
	toSerialize["metricsRetentionTime5m"] = o.MetricsRetentionTime5m
	toSerialize["metricsRetentionTimeRaw"] = o.MetricsRetentionTimeRaw
	return toSerialize, nil
}

type NullableUpdateMetricsStorageRetentionPayload struct {
	value *UpdateMetricsStorageRetentionPayload
	isSet bool
}

func (v NullableUpdateMetricsStorageRetentionPayload) Get() *UpdateMetricsStorageRetentionPayload {
	return v.value
}

func (v *NullableUpdateMetricsStorageRetentionPayload) Set(val *UpdateMetricsStorageRetentionPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMetricsStorageRetentionPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMetricsStorageRetentionPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMetricsStorageRetentionPayload(val *UpdateMetricsStorageRetentionPayload) *NullableUpdateMetricsStorageRetentionPayload {
	return &NullableUpdateMetricsStorageRetentionPayload{value: val, isSet: true}
}

func (v NullableUpdateMetricsStorageRetentionPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMetricsStorageRetentionPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
