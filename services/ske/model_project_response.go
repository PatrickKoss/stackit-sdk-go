/*
SKE-API

The SKE API provides endpoints to create, update, delete clusters within STACKIT portal projects and to trigger further cluster management tasks.

API version: 1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ske

import (
	"encoding/json"
)

// checks if the ProjectResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectResponse{}

// ProjectResponse struct for ProjectResponse
type ProjectResponse struct {
	ProjectId *string       `json:"projectId,omitempty"`
	State     *ProjectState `json:"state,omitempty"`
}

// NewProjectResponse instantiates a new ProjectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectResponse() *ProjectResponse {
	this := ProjectResponse{}
	var state ProjectState = PROJECTSTATE_UNSPECIFIED
	this.State = &state
	return &this
}

// NewProjectResponseWithDefaults instantiates a new ProjectResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectResponseWithDefaults() *ProjectResponse {
	this := ProjectResponse{}
	var state ProjectState = PROJECTSTATE_UNSPECIFIED
	this.State = &state
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ProjectResponse) GetProjectId() *string {
	if o == nil || IsNil(o.ProjectId) {
		var ret *string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectResponse) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ProjectResponse) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *ProjectResponse) SetProjectId(v *string) {
	o.ProjectId = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ProjectResponse) GetState() *ProjectState {
	if o == nil || IsNil(o.State) {
		var ret *ProjectState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectResponse) GetStateOk() (*ProjectState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ProjectResponse) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given ProjectState and assigns it to the State field.
func (o *ProjectResponse) SetState(v *ProjectState) {
	o.State = v
}

func (o ProjectResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableProjectResponse struct {
	value *ProjectResponse
	isSet bool
}

func (v NullableProjectResponse) Get() *ProjectResponse {
	return v.value
}

func (v *NullableProjectResponse) Set(val *ProjectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectResponse(val *ProjectResponse) *NullableProjectResponse {
	return &NullableProjectResponse{value: val, isSet: true}
}

func (v NullableProjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
