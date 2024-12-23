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

// checks if the ClaimIntoOrganization200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClaimIntoOrganization200Response{}

// ClaimIntoOrganization200Response struct for ClaimIntoOrganization200Response
type ClaimIntoOrganization200Response struct {
	// The numbers of the orders claimed
	Orders []string `json:"orders,omitempty"`
	// The serials of the devices claimed
	Serials []string `json:"serials,omitempty"`
	// The licenses claimed
	Licenses []ClaimIntoOrganization200ResponseLicensesInner `json:"licenses,omitempty"`
}

// NewClaimIntoOrganization200Response instantiates a new ClaimIntoOrganization200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClaimIntoOrganization200Response() *ClaimIntoOrganization200Response {
	this := ClaimIntoOrganization200Response{}
	return &this
}

// NewClaimIntoOrganization200ResponseWithDefaults instantiates a new ClaimIntoOrganization200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClaimIntoOrganization200ResponseWithDefaults() *ClaimIntoOrganization200Response {
	this := ClaimIntoOrganization200Response{}
	return &this
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *ClaimIntoOrganization200Response) GetOrders() []string {
	if o == nil || IsNil(o.Orders) {
		var ret []string
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimIntoOrganization200Response) GetOrdersOk() ([]string, bool) {
	if o == nil || IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *ClaimIntoOrganization200Response) HasOrders() bool {
	if o != nil && !IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []string and assigns it to the Orders field.
func (o *ClaimIntoOrganization200Response) SetOrders(v []string) {
	o.Orders = v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *ClaimIntoOrganization200Response) GetSerials() []string {
	if o == nil || IsNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimIntoOrganization200Response) GetSerialsOk() ([]string, bool) {
	if o == nil || IsNil(o.Serials) {
		return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *ClaimIntoOrganization200Response) HasSerials() bool {
	if o != nil && !IsNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *ClaimIntoOrganization200Response) SetSerials(v []string) {
	o.Serials = v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *ClaimIntoOrganization200Response) GetLicenses() []ClaimIntoOrganization200ResponseLicensesInner {
	if o == nil || IsNil(o.Licenses) {
		var ret []ClaimIntoOrganization200ResponseLicensesInner
		return ret
	}
	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClaimIntoOrganization200Response) GetLicensesOk() ([]ClaimIntoOrganization200ResponseLicensesInner, bool) {
	if o == nil || IsNil(o.Licenses) {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *ClaimIntoOrganization200Response) HasLicenses() bool {
	if o != nil && !IsNil(o.Licenses) {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given []ClaimIntoOrganization200ResponseLicensesInner and assigns it to the Licenses field.
func (o *ClaimIntoOrganization200Response) SetLicenses(v []ClaimIntoOrganization200ResponseLicensesInner) {
	o.Licenses = v
}

func (o ClaimIntoOrganization200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClaimIntoOrganization200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Orders) {
		toSerialize["orders"] = o.Orders
	}
	if !IsNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	if !IsNil(o.Licenses) {
		toSerialize["licenses"] = o.Licenses
	}
	return toSerialize, nil
}

type NullableClaimIntoOrganization200Response struct {
	value *ClaimIntoOrganization200Response
	isSet bool
}

func (v NullableClaimIntoOrganization200Response) Get() *ClaimIntoOrganization200Response {
	return v.value
}

func (v *NullableClaimIntoOrganization200Response) Set(val *ClaimIntoOrganization200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableClaimIntoOrganization200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableClaimIntoOrganization200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClaimIntoOrganization200Response(val *ClaimIntoOrganization200Response) *NullableClaimIntoOrganization200Response {
	return &NullableClaimIntoOrganization200Response{value: val, isSet: true}
}

func (v NullableClaimIntoOrganization200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClaimIntoOrganization200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


