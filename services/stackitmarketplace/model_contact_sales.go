/*
STACKIT Marketplace API

API to manage STACKIT Marketplace.

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackitmarketplace

import (
	"encoding/json"
)

// checks if the ContactSales type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactSales{}

// ContactSales Contact sales.
type ContactSales struct {
	// REQUIRED
	ContactSales *ContactSalesContactSales `json:"contactSales"`
	// The form type.
	// REQUIRED
	Type *string `json:"type"`
}

type _ContactSales ContactSales

// NewContactSales instantiates a new ContactSales object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactSales(contactSales *ContactSalesContactSales, type_ *string) *ContactSales {
	this := ContactSales{}
	this.ContactSales = contactSales
	this.Type = type_
	return &this
}

// NewContactSalesWithDefaults instantiates a new ContactSales object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactSalesWithDefaults() *ContactSales {
	this := ContactSales{}
	return &this
}

// GetContactSales returns the ContactSales field value
func (o *ContactSales) GetContactSales() *ContactSalesContactSales {
	if o == nil || IsNil(o.ContactSales) {
		var ret *ContactSalesContactSales
		return ret
	}

	return o.ContactSales
}

// GetContactSalesOk returns a tuple with the ContactSales field value
// and a boolean to check if the value has been set.
func (o *ContactSales) GetContactSalesOk() (*ContactSalesContactSales, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContactSales, true
}

// SetContactSales sets field value
func (o *ContactSales) SetContactSales(v *ContactSalesContactSales) {
	o.ContactSales = v
}

// GetType returns the Type field value
func (o *ContactSales) GetType() *string {
	if o == nil || IsNil(o.Type) {
		var ret *string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContactSales) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *ContactSales) SetType(v *string) {
	o.Type = v
}

func (o ContactSales) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contactSales"] = o.ContactSales
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableContactSales struct {
	value *ContactSales
	isSet bool
}

func (v NullableContactSales) Get() *ContactSales {
	return v.value
}

func (v *NullableContactSales) Set(val *ContactSales) {
	v.value = val
	v.isSet = true
}

func (v NullableContactSales) IsSet() bool {
	return v.isSet
}

func (v *NullableContactSales) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactSales(val *ContactSales) *NullableContactSales {
	return &NullableContactSales{value: val, isSet: true}
}

func (v NullableContactSales) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactSales) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
