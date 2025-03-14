/*
STACKIT Object Storage API

STACKIT API to manage the Object Storage

API version: 2.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objectstorage

import (
	"encoding/json"
)

// checks if the Bucket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Bucket{}

// Bucket struct for Bucket
type Bucket struct {
	// REQUIRED
	Name *string `json:"name"`
	// REQUIRED
	Region *string `json:"region"`
	// URL in path style
	// REQUIRED
	UrlPathStyle *string `json:"urlPathStyle"`
	// URL in virtual hosted style
	// REQUIRED
	UrlVirtualHostedStyle *string `json:"urlVirtualHostedStyle"`
}

type _Bucket Bucket

// NewBucket instantiates a new Bucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBucket(name *string, region *string, urlPathStyle *string, urlVirtualHostedStyle *string) *Bucket {
	this := Bucket{}
	this.Name = name
	this.Region = region
	this.UrlPathStyle = urlPathStyle
	this.UrlVirtualHostedStyle = urlVirtualHostedStyle
	return &this
}

// NewBucketWithDefaults instantiates a new Bucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBucketWithDefaults() *Bucket {
	this := Bucket{}
	return &this
}

// GetName returns the Name field value
func (o *Bucket) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Bucket) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *Bucket) SetName(v *string) {
	o.Name = v
}

// GetRegion returns the Region field value
func (o *Bucket) GetRegion() *string {
	if o == nil || IsNil(o.Region) {
		var ret *string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *Bucket) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region, true
}

// SetRegion sets field value
func (o *Bucket) SetRegion(v *string) {
	o.Region = v
}

// GetUrlPathStyle returns the UrlPathStyle field value
func (o *Bucket) GetUrlPathStyle() *string {
	if o == nil || IsNil(o.UrlPathStyle) {
		var ret *string
		return ret
	}

	return o.UrlPathStyle
}

// GetUrlPathStyleOk returns a tuple with the UrlPathStyle field value
// and a boolean to check if the value has been set.
func (o *Bucket) GetUrlPathStyleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UrlPathStyle, true
}

// SetUrlPathStyle sets field value
func (o *Bucket) SetUrlPathStyle(v *string) {
	o.UrlPathStyle = v
}

// GetUrlVirtualHostedStyle returns the UrlVirtualHostedStyle field value
func (o *Bucket) GetUrlVirtualHostedStyle() *string {
	if o == nil || IsNil(o.UrlVirtualHostedStyle) {
		var ret *string
		return ret
	}

	return o.UrlVirtualHostedStyle
}

// GetUrlVirtualHostedStyleOk returns a tuple with the UrlVirtualHostedStyle field value
// and a boolean to check if the value has been set.
func (o *Bucket) GetUrlVirtualHostedStyleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UrlVirtualHostedStyle, true
}

// SetUrlVirtualHostedStyle sets field value
func (o *Bucket) SetUrlVirtualHostedStyle(v *string) {
	o.UrlVirtualHostedStyle = v
}

func (o Bucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["region"] = o.Region
	toSerialize["urlPathStyle"] = o.UrlPathStyle
	toSerialize["urlVirtualHostedStyle"] = o.UrlVirtualHostedStyle
	return toSerialize, nil
}

type NullableBucket struct {
	value *Bucket
	isSet bool
}

func (v NullableBucket) Get() *Bucket {
	return v.value
}

func (v *NullableBucket) Set(val *Bucket) {
	v.value = val
	v.isSet = true
}

func (v NullableBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucket(val *Bucket) *NullableBucket {
	return &NullableBucket{value: val, isSet: true}
}

func (v NullableBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
