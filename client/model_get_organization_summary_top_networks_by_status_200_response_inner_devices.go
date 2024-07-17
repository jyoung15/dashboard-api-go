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

// checks if the GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices{}

// GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices Network device information
type GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices struct {
	// URLs by product type
	ByProductType []GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevicesByProductTypeInner `json:"byProductType,omitempty"`
}

// NewGetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices instantiates a new GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices() *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices {
	this := GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices{}
	return &this
}

// NewGetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevicesWithDefaults instantiates a new GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevicesWithDefaults() *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices {
	this := GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices{}
	return &this
}

// GetByProductType returns the ByProductType field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices) GetByProductType() []GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevicesByProductTypeInner {
	if o == nil || IsNil(o.ByProductType) {
		var ret []GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevicesByProductTypeInner
		return ret
	}
	return o.ByProductType
}

// GetByProductTypeOk returns a tuple with the ByProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices) GetByProductTypeOk() ([]GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevicesByProductTypeInner, bool) {
	if o == nil || IsNil(o.ByProductType) {
		return nil, false
	}
	return o.ByProductType, true
}

// HasByProductType returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices) HasByProductType() bool {
	if o != nil && !IsNil(o.ByProductType) {
		return true
	}

	return false
}

// SetByProductType gets a reference to the given []GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevicesByProductTypeInner and assigns it to the ByProductType field.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices) SetByProductType(v []GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevicesByProductTypeInner) {
	o.ByProductType = v
}

func (o GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ByProductType) {
		toSerialize["byProductType"] = o.ByProductType
	}
	return toSerialize, nil
}

type NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices struct {
	value *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices
	isSet bool
}

func (v NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices) Get() *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices {
	return v.value
}

func (v *NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices) Set(val *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices(val *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices) *NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices {
	return &NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices{value: val, isSet: true}
}

func (v NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

