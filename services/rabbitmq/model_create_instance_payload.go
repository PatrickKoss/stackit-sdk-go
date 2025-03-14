/*
STACKIT RabbitMQ API

The STACKIT RabbitMQ API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rabbitmq

import (
	"encoding/json"
)

// checks if the CreateInstancePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInstancePayload{}

// CreateInstancePayload struct for CreateInstancePayload
type CreateInstancePayload struct {
	// REQUIRED
	InstanceName *string             `json:"instanceName"`
	Parameters   *InstanceParameters `json:"parameters,omitempty"`
	// REQUIRED
	PlanId *string `json:"planId"`
}

type _CreateInstancePayload CreateInstancePayload

// NewCreateInstancePayload instantiates a new CreateInstancePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInstancePayload(instanceName *string, planId *string) *CreateInstancePayload {
	this := CreateInstancePayload{}
	this.InstanceName = instanceName
	this.PlanId = planId
	return &this
}

// NewCreateInstancePayloadWithDefaults instantiates a new CreateInstancePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInstancePayloadWithDefaults() *CreateInstancePayload {
	this := CreateInstancePayload{}
	return &this
}

// GetInstanceName returns the InstanceName field value
func (o *CreateInstancePayload) GetInstanceName() *string {
	if o == nil || IsNil(o.InstanceName) {
		var ret *string
		return ret
	}

	return o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value
// and a boolean to check if the value has been set.
func (o *CreateInstancePayload) GetInstanceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstanceName, true
}

// SetInstanceName sets field value
func (o *CreateInstancePayload) SetInstanceName(v *string) {
	o.InstanceName = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *CreateInstancePayload) GetParameters() *InstanceParameters {
	if o == nil || IsNil(o.Parameters) {
		var ret *InstanceParameters
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInstancePayload) GetParametersOk() (*InstanceParameters, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *CreateInstancePayload) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given InstanceParameters and assigns it to the Parameters field.
func (o *CreateInstancePayload) SetParameters(v *InstanceParameters) {
	o.Parameters = v
}

// GetPlanId returns the PlanId field value
func (o *CreateInstancePayload) GetPlanId() *string {
	if o == nil || IsNil(o.PlanId) {
		var ret *string
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *CreateInstancePayload) GetPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlanId, true
}

// SetPlanId sets field value
func (o *CreateInstancePayload) SetPlanId(v *string) {
	o.PlanId = v
}

func (o CreateInstancePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instanceName"] = o.InstanceName
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	toSerialize["planId"] = o.PlanId
	return toSerialize, nil
}

type NullableCreateInstancePayload struct {
	value *CreateInstancePayload
	isSet bool
}

func (v NullableCreateInstancePayload) Get() *CreateInstancePayload {
	return v.value
}

func (v *NullableCreateInstancePayload) Set(val *CreateInstancePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInstancePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInstancePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInstancePayload(val *CreateInstancePayload) *NullableCreateInstancePayload {
	return &NullableCreateInstancePayload{value: val, isSet: true}
}

func (v NullableCreateInstancePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInstancePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
