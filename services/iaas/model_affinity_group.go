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

// checks if the AffinityGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AffinityGroup{}

// AffinityGroup Definition of an affinity group.
type AffinityGroup struct {
	// Universally Unique Identifier (UUID).
	Id *string `json:"id,omitempty"`
	// The servers that are part of the affinity group.
	Members *[]string `json:"members,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	// REQUIRED
	Name *string `json:"name"`
	// The affinity group policy. `hard-affinity`: All servers in this group will be hosted on the same compute node. `soft-affinity`: All servers in this group will be hosted on as few compute nodes as possible. `hard-anti-affinity`: All servers in this group will be hosted on different compute nodes. `soft-anti-affinity`: All servers in this group will be hosted on as many compute nodes as possible. Possible values: `hard-anti-affinity`, `hard-affinity`, `soft-anti-affinity`, `soft-affinity`.
	// REQUIRED
	Policy *string `json:"policy"`
}

type _AffinityGroup AffinityGroup

// NewAffinityGroup instantiates a new AffinityGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffinityGroup(name *string, policy *string) *AffinityGroup {
	this := AffinityGroup{}
	this.Name = name
	this.Policy = policy
	return &this
}

// NewAffinityGroupWithDefaults instantiates a new AffinityGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffinityGroupWithDefaults() *AffinityGroup {
	this := AffinityGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AffinityGroup) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffinityGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AffinityGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AffinityGroup) SetId(v *string) {
	o.Id = v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *AffinityGroup) GetMembers() *[]string {
	if o == nil || IsNil(o.Members) {
		var ret *[]string
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffinityGroup) GetMembersOk() (*[]string, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *AffinityGroup) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []string and assigns it to the Members field.
func (o *AffinityGroup) SetMembers(v *[]string) {
	o.Members = v
}

// GetName returns the Name field value
func (o *AffinityGroup) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AffinityGroup) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *AffinityGroup) SetName(v *string) {
	o.Name = v
}

// GetPolicy returns the Policy field value
func (o *AffinityGroup) GetPolicy() *string {
	if o == nil || IsNil(o.Policy) {
		var ret *string
		return ret
	}

	return o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value
// and a boolean to check if the value has been set.
func (o *AffinityGroup) GetPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Policy, true
}

// SetPolicy sets field value
func (o *AffinityGroup) SetPolicy(v *string) {
	o.Policy = v
}

func (o AffinityGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	toSerialize["name"] = o.Name
	toSerialize["policy"] = o.Policy
	return toSerialize, nil
}

type NullableAffinityGroup struct {
	value *AffinityGroup
	isSet bool
}

func (v NullableAffinityGroup) Get() *AffinityGroup {
	return v.value
}

func (v *NullableAffinityGroup) Set(val *AffinityGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableAffinityGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableAffinityGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffinityGroup(val *AffinityGroup) *NullableAffinityGroup {
	return &NullableAffinityGroup{value: val, isSet: true}
}

func (v NullableAffinityGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffinityGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
