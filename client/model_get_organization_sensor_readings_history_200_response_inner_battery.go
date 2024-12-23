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

// checks if the GetOrganizationSensorReadingsHistory200ResponseInnerBattery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSensorReadingsHistory200ResponseInnerBattery{}

// GetOrganizationSensorReadingsHistory200ResponseInnerBattery Reading for the 'battery' metric. This will only be present if the 'metric' property equals 'battery'.
type GetOrganizationSensorReadingsHistory200ResponseInnerBattery struct {
	// Remaining battery life.
	Percentage *int32 `json:"percentage,omitempty"`
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerBattery instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerBattery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSensorReadingsHistory200ResponseInnerBattery() *GetOrganizationSensorReadingsHistory200ResponseInnerBattery {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerBattery{}
	return &this
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerBatteryWithDefaults instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerBattery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSensorReadingsHistory200ResponseInnerBatteryWithDefaults() *GetOrganizationSensorReadingsHistory200ResponseInnerBattery {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerBattery{}
	return &this
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerBattery) GetPercentage() int32 {
	if o == nil || IsNil(o.Percentage) {
		var ret int32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerBattery) GetPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.Percentage) {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerBattery) HasPercentage() bool {
	if o != nil && !IsNil(o.Percentage) {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given int32 and assigns it to the Percentage field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerBattery) SetPercentage(v int32) {
	o.Percentage = &v
}

func (o GetOrganizationSensorReadingsHistory200ResponseInnerBattery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSensorReadingsHistory200ResponseInnerBattery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Percentage) {
		toSerialize["percentage"] = o.Percentage
	}
	return toSerialize, nil
}

type NullableGetOrganizationSensorReadingsHistory200ResponseInnerBattery struct {
	value *GetOrganizationSensorReadingsHistory200ResponseInnerBattery
	isSet bool
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerBattery) Get() *GetOrganizationSensorReadingsHistory200ResponseInnerBattery {
	return v.value
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerBattery) Set(val *GetOrganizationSensorReadingsHistory200ResponseInnerBattery) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerBattery) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerBattery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSensorReadingsHistory200ResponseInnerBattery(val *GetOrganizationSensorReadingsHistory200ResponseInnerBattery) *NullableGetOrganizationSensorReadingsHistory200ResponseInnerBattery {
	return &NullableGetOrganizationSensorReadingsHistory200ResponseInnerBattery{value: val, isSet: true}
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerBattery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerBattery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


