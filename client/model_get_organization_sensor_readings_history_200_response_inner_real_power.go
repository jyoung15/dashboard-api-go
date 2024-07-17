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

// checks if the GetOrganizationSensorReadingsHistory200ResponseInnerRealPower type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSensorReadingsHistory200ResponseInnerRealPower{}

// GetOrganizationSensorReadingsHistory200ResponseInnerRealPower Reading for the 'realPower' metric. This will only be present if the 'metric' property equals 'realPower'.
type GetOrganizationSensorReadingsHistory200ResponseInnerRealPower struct {
	// Real power reading in watts.
	Draw *float32 `json:"draw,omitempty"`
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerRealPower instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerRealPower object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSensorReadingsHistory200ResponseInnerRealPower() *GetOrganizationSensorReadingsHistory200ResponseInnerRealPower {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerRealPower{}
	return &this
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerRealPowerWithDefaults instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerRealPower object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSensorReadingsHistory200ResponseInnerRealPowerWithDefaults() *GetOrganizationSensorReadingsHistory200ResponseInnerRealPower {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerRealPower{}
	return &this
}

// GetDraw returns the Draw field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerRealPower) GetDraw() float32 {
	if o == nil || IsNil(o.Draw) {
		var ret float32
		return ret
	}
	return *o.Draw
}

// GetDrawOk returns a tuple with the Draw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerRealPower) GetDrawOk() (*float32, bool) {
	if o == nil || IsNil(o.Draw) {
		return nil, false
	}
	return o.Draw, true
}

// HasDraw returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerRealPower) HasDraw() bool {
	if o != nil && !IsNil(o.Draw) {
		return true
	}

	return false
}

// SetDraw gets a reference to the given float32 and assigns it to the Draw field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerRealPower) SetDraw(v float32) {
	o.Draw = &v
}

func (o GetOrganizationSensorReadingsHistory200ResponseInnerRealPower) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSensorReadingsHistory200ResponseInnerRealPower) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Draw) {
		toSerialize["draw"] = o.Draw
	}
	return toSerialize, nil
}

type NullableGetOrganizationSensorReadingsHistory200ResponseInnerRealPower struct {
	value *GetOrganizationSensorReadingsHistory200ResponseInnerRealPower
	isSet bool
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerRealPower) Get() *GetOrganizationSensorReadingsHistory200ResponseInnerRealPower {
	return v.value
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerRealPower) Set(val *GetOrganizationSensorReadingsHistory200ResponseInnerRealPower) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerRealPower) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerRealPower) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSensorReadingsHistory200ResponseInnerRealPower(val *GetOrganizationSensorReadingsHistory200ResponseInnerRealPower) *NullableGetOrganizationSensorReadingsHistory200ResponseInnerRealPower {
	return &NullableGetOrganizationSensorReadingsHistory200ResponseInnerRealPower{value: val, isSet: true}
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerRealPower) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerRealPower) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


