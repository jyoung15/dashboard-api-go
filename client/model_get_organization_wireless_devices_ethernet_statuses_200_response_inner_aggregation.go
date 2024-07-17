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

// checks if the GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation{}

// GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation Aggregation details object
type GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation struct {
	// Link Aggregation enabled flag
	Enabled *bool `json:"enabled,omitempty"`
	// Link Aggregation speed
	Speed *int32 `json:"speed,omitempty"`
}

// NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation instantiates a new GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation {
	this := GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation{}
	return &this
}

// NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregationWithDefaults instantiates a new GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregationWithDefaults() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation {
	this := GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation) GetSpeed() int32 {
	if o == nil || IsNil(o.Speed) {
		var ret int32
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation) GetSpeedOk() (*int32, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given int32 and assigns it to the Speed field.
func (o *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation) SetSpeed(v int32) {
	o.Speed = &v
}

func (o GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	return toSerialize, nil
}

type NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation struct {
	value *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation
	isSet bool
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation) Get() *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation {
	return v.value
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation) Set(val *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation(val *GetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation) *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation {
	return &NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation{value: val, isSet: true}
}

func (v NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWirelessDevicesEthernetStatuses200ResponseInnerAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


