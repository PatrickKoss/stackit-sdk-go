/*
Application Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each application load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1beta.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lbapplication

import (
	"encoding/json"
)

// checks if the Listener type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Listener{}

// Listener struct for Listener
type Listener struct {
	DisplayName *string                 `json:"displayName,omitempty"`
	Http        *map[string]interface{} `json:"http,omitempty"`
	Https       *ProtocolOptionsHTTPS   `json:"https,omitempty"`
	// Will be used to reference a listener and will replace display name in the future. Currently uses <protocol>-<port> as the name if no display name is given.
	Name *string `json:"name,omitempty"`
	// Port number where we listen for traffic
	Port *interface{} `json:"port,omitempty"`
	// Protocol is the highest network protocol we understand to load balance. Currently PROTOCOL_HTTP and PROTOCOL_HTTPS are supported.
	Protocol *string `json:"protocol,omitempty"`
	// Rules define the routing parameters for the HTTP and HTTPS listeners.
	Rules *[]Rule `json:"rules,omitempty"`
}

// NewListener instantiates a new Listener object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListener() *Listener {
	this := Listener{}
	return &this
}

// NewListenerWithDefaults instantiates a new Listener object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListenerWithDefaults() *Listener {
	this := Listener{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Listener) GetDisplayName() *string {
	if o == nil || IsNil(o.DisplayName) {
		var ret *string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Listener) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Listener) SetDisplayName(v *string) {
	o.DisplayName = v
}

// GetHttp returns the Http field value if set, zero value otherwise.
func (o *Listener) GetHttp() *map[string]interface{} {
	if o == nil || IsNil(o.Http) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Http
}

// GetHttpOk returns a tuple with the Http field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetHttpOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Http) {
		return &map[string]interface{}{}, false
	}
	return o.Http, true
}

// HasHttp returns a boolean if a field has been set.
func (o *Listener) HasHttp() bool {
	if o != nil && !IsNil(o.Http) {
		return true
	}

	return false
}

// SetHttp gets a reference to the given map[string]interface{} and assigns it to the Http field.
func (o *Listener) SetHttp(v *map[string]interface{}) {
	o.Http = v
}

// GetHttps returns the Https field value if set, zero value otherwise.
func (o *Listener) GetHttps() *ProtocolOptionsHTTPS {
	if o == nil || IsNil(o.Https) {
		var ret *ProtocolOptionsHTTPS
		return ret
	}
	return o.Https
}

// GetHttpsOk returns a tuple with the Https field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetHttpsOk() (*ProtocolOptionsHTTPS, bool) {
	if o == nil || IsNil(o.Https) {
		return nil, false
	}
	return o.Https, true
}

// HasHttps returns a boolean if a field has been set.
func (o *Listener) HasHttps() bool {
	if o != nil && !IsNil(o.Https) {
		return true
	}

	return false
}

// SetHttps gets a reference to the given ProtocolOptionsHTTPS and assigns it to the Https field.
func (o *Listener) SetHttps(v *ProtocolOptionsHTTPS) {
	o.Https = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Listener) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Listener) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Listener) SetName(v *string) {
	o.Name = v
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Listener) GetPort() *interface{} {
	if o == nil || IsNil(o.Port) {
		var ret *interface{}
		return ret
	}
	return o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Listener) GetPortOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *Listener) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given interface{} and assigns it to the Port field.
func (o *Listener) SetPort(v *interface{}) {
	o.Port = v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *Listener) GetProtocol() *string {
	if o == nil || IsNil(o.Protocol) {
		var ret *string
		return ret
	}
	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *Listener) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *Listener) SetProtocol(v *string) {
	o.Protocol = v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *Listener) GetRules() *[]Rule {
	if o == nil || IsNil(o.Rules) {
		var ret *[]Rule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Listener) GetRulesOk() (*[]Rule, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *Listener) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []Rule and assigns it to the Rules field.
func (o *Listener) SetRules(v *[]Rule) {
	o.Rules = v
}

func (o Listener) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Http) {
		toSerialize["http"] = o.Http
	}
	if !IsNil(o.Https) {
		toSerialize["https"] = o.Https
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableListener struct {
	value *Listener
	isSet bool
}

func (v NullableListener) Get() *Listener {
	return v.value
}

func (v *NullableListener) Set(val *Listener) {
	v.value = val
	v.isSet = true
}

func (v NullableListener) IsSet() bool {
	return v.isSet
}

func (v *NullableListener) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListener(val *Listener) *NullableListener {
	return &NullableListener{value: val, isSet: true}
}

func (v NullableListener) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListener) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
