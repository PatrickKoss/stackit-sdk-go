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

// checks if the Matcher type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Matcher{}

// Matcher struct for Matcher
type Matcher struct {
	CookiePersistence *CookiePersistence `json:"cookiePersistence,omitempty"`
	// Headers for the matcher
	Headers *[]Header `json:"headers,omitempty"`
	// Path prefix for the matcher
	PathPrefix *string `json:"pathPrefix,omitempty"`
	// Query Parameters for the matcher
	QueryParameters *[]QueryParameters `json:"queryParameters,omitempty"`
	// Reference target pool by target pool name.
	TargetPool *string `json:"targetPool,omitempty"`
	// If enabled, when client sends an HTTP request with and Upgrade header, indicating the desire to establish a Websocket connection,  if backend server supports WebSocket, it responds with HTTP 101 status code, switching protocols from HTTP to WebSocket. Hence the client and the server can exchange data in real-time using one long-lived TCP connection.
	WebSocket *bool `json:"webSocket,omitempty"`
}

// NewMatcher instantiates a new Matcher object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatcher() *Matcher {
	this := Matcher{}
	return &this
}

// NewMatcherWithDefaults instantiates a new Matcher object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatcherWithDefaults() *Matcher {
	this := Matcher{}
	return &this
}

// GetCookiePersistence returns the CookiePersistence field value if set, zero value otherwise.
func (o *Matcher) GetCookiePersistence() *CookiePersistence {
	if o == nil || IsNil(o.CookiePersistence) {
		var ret *CookiePersistence
		return ret
	}
	return o.CookiePersistence
}

// GetCookiePersistenceOk returns a tuple with the CookiePersistence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Matcher) GetCookiePersistenceOk() (*CookiePersistence, bool) {
	if o == nil || IsNil(o.CookiePersistence) {
		return nil, false
	}
	return o.CookiePersistence, true
}

// HasCookiePersistence returns a boolean if a field has been set.
func (o *Matcher) HasCookiePersistence() bool {
	if o != nil && !IsNil(o.CookiePersistence) {
		return true
	}

	return false
}

// SetCookiePersistence gets a reference to the given CookiePersistence and assigns it to the CookiePersistence field.
func (o *Matcher) SetCookiePersistence(v *CookiePersistence) {
	o.CookiePersistence = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *Matcher) GetHeaders() *[]Header {
	if o == nil || IsNil(o.Headers) {
		var ret *[]Header
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Matcher) GetHeadersOk() (*[]Header, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *Matcher) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []Header and assigns it to the Headers field.
func (o *Matcher) SetHeaders(v *[]Header) {
	o.Headers = v
}

// GetPathPrefix returns the PathPrefix field value if set, zero value otherwise.
func (o *Matcher) GetPathPrefix() *string {
	if o == nil || IsNil(o.PathPrefix) {
		var ret *string
		return ret
	}
	return o.PathPrefix
}

// GetPathPrefixOk returns a tuple with the PathPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Matcher) GetPathPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.PathPrefix) {
		return nil, false
	}
	return o.PathPrefix, true
}

// HasPathPrefix returns a boolean if a field has been set.
func (o *Matcher) HasPathPrefix() bool {
	if o != nil && !IsNil(o.PathPrefix) {
		return true
	}

	return false
}

// SetPathPrefix gets a reference to the given string and assigns it to the PathPrefix field.
func (o *Matcher) SetPathPrefix(v *string) {
	o.PathPrefix = v
}

// GetQueryParameters returns the QueryParameters field value if set, zero value otherwise.
func (o *Matcher) GetQueryParameters() *[]QueryParameters {
	if o == nil || IsNil(o.QueryParameters) {
		var ret *[]QueryParameters
		return ret
	}
	return o.QueryParameters
}

// GetQueryParametersOk returns a tuple with the QueryParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Matcher) GetQueryParametersOk() (*[]QueryParameters, bool) {
	if o == nil || IsNil(o.QueryParameters) {
		return nil, false
	}
	return o.QueryParameters, true
}

// HasQueryParameters returns a boolean if a field has been set.
func (o *Matcher) HasQueryParameters() bool {
	if o != nil && !IsNil(o.QueryParameters) {
		return true
	}

	return false
}

// SetQueryParameters gets a reference to the given []QueryParameters and assigns it to the QueryParameters field.
func (o *Matcher) SetQueryParameters(v *[]QueryParameters) {
	o.QueryParameters = v
}

// GetTargetPool returns the TargetPool field value if set, zero value otherwise.
func (o *Matcher) GetTargetPool() *string {
	if o == nil || IsNil(o.TargetPool) {
		var ret *string
		return ret
	}
	return o.TargetPool
}

// GetTargetPoolOk returns a tuple with the TargetPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Matcher) GetTargetPoolOk() (*string, bool) {
	if o == nil || IsNil(o.TargetPool) {
		return nil, false
	}
	return o.TargetPool, true
}

// HasTargetPool returns a boolean if a field has been set.
func (o *Matcher) HasTargetPool() bool {
	if o != nil && !IsNil(o.TargetPool) {
		return true
	}

	return false
}

// SetTargetPool gets a reference to the given string and assigns it to the TargetPool field.
func (o *Matcher) SetTargetPool(v *string) {
	o.TargetPool = v
}

// GetWebSocket returns the WebSocket field value if set, zero value otherwise.
func (o *Matcher) GetWebSocket() *bool {
	if o == nil || IsNil(o.WebSocket) {
		var ret *bool
		return ret
	}
	return o.WebSocket
}

// GetWebSocketOk returns a tuple with the WebSocket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Matcher) GetWebSocketOk() (*bool, bool) {
	if o == nil || IsNil(o.WebSocket) {
		return nil, false
	}
	return o.WebSocket, true
}

// HasWebSocket returns a boolean if a field has been set.
func (o *Matcher) HasWebSocket() bool {
	if o != nil && !IsNil(o.WebSocket) {
		return true
	}

	return false
}

// SetWebSocket gets a reference to the given bool and assigns it to the WebSocket field.
func (o *Matcher) SetWebSocket(v *bool) {
	o.WebSocket = v
}

func (o Matcher) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CookiePersistence) {
		toSerialize["cookiePersistence"] = o.CookiePersistence
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.PathPrefix) {
		toSerialize["pathPrefix"] = o.PathPrefix
	}
	if !IsNil(o.QueryParameters) {
		toSerialize["queryParameters"] = o.QueryParameters
	}
	if !IsNil(o.TargetPool) {
		toSerialize["targetPool"] = o.TargetPool
	}
	if !IsNil(o.WebSocket) {
		toSerialize["webSocket"] = o.WebSocket
	}
	return toSerialize, nil
}

type NullableMatcher struct {
	value *Matcher
	isSet bool
}

func (v NullableMatcher) Get() *Matcher {
	return v.value
}

func (v *NullableMatcher) Set(val *Matcher) {
	v.value = val
	v.isSet = true
}

func (v NullableMatcher) IsSet() bool {
	return v.isSet
}

func (v *NullableMatcher) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatcher(val *Matcher) *NullableMatcher {
	return &NullableMatcher{value: val, isSet: true}
}

func (v NullableMatcher) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatcher) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
