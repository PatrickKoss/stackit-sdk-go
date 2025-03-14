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

// checks if the CatalogProductOverview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogProductOverview{}

// CatalogProductOverview struct for CatalogProductOverview
type CatalogProductOverview struct {
	// The product type. For reference: SAAS - Software as a Service, SAI - STACKIT Application Image
	// REQUIRED
	DeliveryMethod *string `json:"deliveryMethod"`
	// The lifecycle state of the product.
	// REQUIRED
	LifecycleState *string `json:"lifecycleState"`
	// The logo base64 encoded.
	Logo *string `json:"logo,omitempty"`
	// The product name.
	// REQUIRED
	Name *string `json:"name"`
	// The product ID.
	// REQUIRED
	ProductId *string `json:"productId"`
	// The short summary of the product.
	// REQUIRED
	Summary *string `json:"summary"`
	// REQUIRED
	Vendor *CatalogProductOverviewVendor `json:"vendor"`
}

type _CatalogProductOverview CatalogProductOverview

// NewCatalogProductOverview instantiates a new CatalogProductOverview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogProductOverview(deliveryMethod *string, lifecycleState *string, name *string, productId *string, summary *string, vendor *CatalogProductOverviewVendor) *CatalogProductOverview {
	this := CatalogProductOverview{}
	this.DeliveryMethod = deliveryMethod
	this.LifecycleState = lifecycleState
	this.Name = name
	this.ProductId = productId
	this.Summary = summary
	this.Vendor = vendor
	return &this
}

// NewCatalogProductOverviewWithDefaults instantiates a new CatalogProductOverview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogProductOverviewWithDefaults() *CatalogProductOverview {
	this := CatalogProductOverview{}
	return &this
}

// GetDeliveryMethod returns the DeliveryMethod field value
func (o *CatalogProductOverview) GetDeliveryMethod() *string {
	if o == nil || IsNil(o.DeliveryMethod) {
		var ret *string
		return ret
	}

	return o.DeliveryMethod
}

// GetDeliveryMethodOk returns a tuple with the DeliveryMethod field value
// and a boolean to check if the value has been set.
func (o *CatalogProductOverview) GetDeliveryMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeliveryMethod, true
}

// SetDeliveryMethod sets field value
func (o *CatalogProductOverview) SetDeliveryMethod(v *string) {
	o.DeliveryMethod = v
}

// GetLifecycleState returns the LifecycleState field value
func (o *CatalogProductOverview) GetLifecycleState() *string {
	if o == nil || IsNil(o.LifecycleState) {
		var ret *string
		return ret
	}

	return o.LifecycleState
}

// GetLifecycleStateOk returns a tuple with the LifecycleState field value
// and a boolean to check if the value has been set.
func (o *CatalogProductOverview) GetLifecycleStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LifecycleState, true
}

// SetLifecycleState sets field value
func (o *CatalogProductOverview) SetLifecycleState(v *string) {
	o.LifecycleState = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *CatalogProductOverview) GetLogo() *string {
	if o == nil || IsNil(o.Logo) {
		var ret *string
		return ret
	}
	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogProductOverview) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *CatalogProductOverview) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *CatalogProductOverview) SetLogo(v *string) {
	o.Logo = v
}

// GetName returns the Name field value
func (o *CatalogProductOverview) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CatalogProductOverview) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *CatalogProductOverview) SetName(v *string) {
	o.Name = v
}

// GetProductId returns the ProductId field value
func (o *CatalogProductOverview) GetProductId() *string {
	if o == nil || IsNil(o.ProductId) {
		var ret *string
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *CatalogProductOverview) GetProductIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductId, true
}

// SetProductId sets field value
func (o *CatalogProductOverview) SetProductId(v *string) {
	o.ProductId = v
}

// GetSummary returns the Summary field value
func (o *CatalogProductOverview) GetSummary() *string {
	if o == nil || IsNil(o.Summary) {
		var ret *string
		return ret
	}

	return o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value
// and a boolean to check if the value has been set.
func (o *CatalogProductOverview) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Summary, true
}

// SetSummary sets field value
func (o *CatalogProductOverview) SetSummary(v *string) {
	o.Summary = v
}

// GetVendor returns the Vendor field value
func (o *CatalogProductOverview) GetVendor() *CatalogProductOverviewVendor {
	if o == nil || IsNil(o.Vendor) {
		var ret *CatalogProductOverviewVendor
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *CatalogProductOverview) GetVendorOk() (*CatalogProductOverviewVendor, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vendor, true
}

// SetVendor sets field value
func (o *CatalogProductOverview) SetVendor(v *CatalogProductOverviewVendor) {
	o.Vendor = v
}

func (o CatalogProductOverview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deliveryMethod"] = o.DeliveryMethod
	toSerialize["lifecycleState"] = o.LifecycleState
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	toSerialize["name"] = o.Name
	toSerialize["productId"] = o.ProductId
	toSerialize["summary"] = o.Summary
	toSerialize["vendor"] = o.Vendor
	return toSerialize, nil
}

type NullableCatalogProductOverview struct {
	value *CatalogProductOverview
	isSet bool
}

func (v NullableCatalogProductOverview) Get() *CatalogProductOverview {
	return v.value
}

func (v *NullableCatalogProductOverview) Set(val *CatalogProductOverview) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogProductOverview) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogProductOverview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogProductOverview(val *CatalogProductOverview) *NullableCatalogProductOverview {
	return &NullableCatalogProductOverview{value: val, isSet: true}
}

func (v NullableCatalogProductOverview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogProductOverview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
