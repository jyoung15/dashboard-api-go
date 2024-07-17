/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetAdministeredLicensingSubscriptionEntitlements200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAdministeredLicensingSubscriptionEntitlements200Response{}

// GetAdministeredLicensingSubscriptionEntitlements200Response struct for GetAdministeredLicensingSubscriptionEntitlements200Response
type GetAdministeredLicensingSubscriptionEntitlements200Response struct {
	// The SKU identifier of the entitlement
	Sku *string `json:"sku,omitempty"`
	// The user-facing name of the entitlement
	Name *string `json:"name,omitempty"`
	// The product type of the entitlement
	ProductType *string `json:"productType,omitempty"`
	// The product class associated with the entitlement
	ProductClass *string `json:"productClass,omitempty"`
	// The feature tier associated with the entitlement (null for add-ons)
	FeatureTier *string `json:"featureTier,omitempty"`
	// Whether or not the entitlement is an add-on
	IsAddOn *bool `json:"isAddOn,omitempty"`
}

// NewGetAdministeredLicensingSubscriptionEntitlements200Response instantiates a new GetAdministeredLicensingSubscriptionEntitlements200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAdministeredLicensingSubscriptionEntitlements200Response() *GetAdministeredLicensingSubscriptionEntitlements200Response {
	this := GetAdministeredLicensingSubscriptionEntitlements200Response{}
	return &this
}

// NewGetAdministeredLicensingSubscriptionEntitlements200ResponseWithDefaults instantiates a new GetAdministeredLicensingSubscriptionEntitlements200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAdministeredLicensingSubscriptionEntitlements200ResponseWithDefaults() *GetAdministeredLicensingSubscriptionEntitlements200Response {
	this := GetAdministeredLicensingSubscriptionEntitlements200Response{}
	return &this
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) SetSku(v string) {
	o.Sku = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) SetName(v string) {
	o.Name = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) GetProductType() string {
	if o == nil || IsNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) GetProductTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductType) {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) HasProductType() bool {
	if o != nil && !IsNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) SetProductType(v string) {
	o.ProductType = &v
}

// GetProductClass returns the ProductClass field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) GetProductClass() string {
	if o == nil || IsNil(o.ProductClass) {
		var ret string
		return ret
	}
	return *o.ProductClass
}

// GetProductClassOk returns a tuple with the ProductClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) GetProductClassOk() (*string, bool) {
	if o == nil || IsNil(o.ProductClass) {
		return nil, false
	}
	return o.ProductClass, true
}

// HasProductClass returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) HasProductClass() bool {
	if o != nil && !IsNil(o.ProductClass) {
		return true
	}

	return false
}

// SetProductClass gets a reference to the given string and assigns it to the ProductClass field.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) SetProductClass(v string) {
	o.ProductClass = &v
}

// GetFeatureTier returns the FeatureTier field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) GetFeatureTier() string {
	if o == nil || IsNil(o.FeatureTier) {
		var ret string
		return ret
	}
	return *o.FeatureTier
}

// GetFeatureTierOk returns a tuple with the FeatureTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) GetFeatureTierOk() (*string, bool) {
	if o == nil || IsNil(o.FeatureTier) {
		return nil, false
	}
	return o.FeatureTier, true
}

// HasFeatureTier returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) HasFeatureTier() bool {
	if o != nil && !IsNil(o.FeatureTier) {
		return true
	}

	return false
}

// SetFeatureTier gets a reference to the given string and assigns it to the FeatureTier field.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) SetFeatureTier(v string) {
	o.FeatureTier = &v
}

// GetIsAddOn returns the IsAddOn field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) GetIsAddOn() bool {
	if o == nil || IsNil(o.IsAddOn) {
		var ret bool
		return ret
	}
	return *o.IsAddOn
}

// GetIsAddOnOk returns a tuple with the IsAddOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) GetIsAddOnOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAddOn) {
		return nil, false
	}
	return o.IsAddOn, true
}

// HasIsAddOn returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) HasIsAddOn() bool {
	if o != nil && !IsNil(o.IsAddOn) {
		return true
	}

	return false
}

// SetIsAddOn gets a reference to the given bool and assigns it to the IsAddOn field.
func (o *GetAdministeredLicensingSubscriptionEntitlements200Response) SetIsAddOn(v bool) {
	o.IsAddOn = &v
}

func (o GetAdministeredLicensingSubscriptionEntitlements200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAdministeredLicensingSubscriptionEntitlements200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !IsNil(o.ProductClass) {
		toSerialize["productClass"] = o.ProductClass
	}
	if !IsNil(o.FeatureTier) {
		toSerialize["featureTier"] = o.FeatureTier
	}
	if !IsNil(o.IsAddOn) {
		toSerialize["isAddOn"] = o.IsAddOn
	}
	return toSerialize, nil
}

type NullableGetAdministeredLicensingSubscriptionEntitlements200Response struct {
	value *GetAdministeredLicensingSubscriptionEntitlements200Response
	isSet bool
}

func (v NullableGetAdministeredLicensingSubscriptionEntitlements200Response) Get() *GetAdministeredLicensingSubscriptionEntitlements200Response {
	return v.value
}

func (v *NullableGetAdministeredLicensingSubscriptionEntitlements200Response) Set(val *GetAdministeredLicensingSubscriptionEntitlements200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAdministeredLicensingSubscriptionEntitlements200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAdministeredLicensingSubscriptionEntitlements200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAdministeredLicensingSubscriptionEntitlements200Response(val *GetAdministeredLicensingSubscriptionEntitlements200Response) *NullableGetAdministeredLicensingSubscriptionEntitlements200Response {
	return &NullableGetAdministeredLicensingSubscriptionEntitlements200Response{value: val, isSet: true}
}

func (v NullableGetAdministeredLicensingSubscriptionEntitlements200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAdministeredLicensingSubscriptionEntitlements200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

