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

// checks if the GetOrganizationSensorReadingsHistory200ResponseInnerCo2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSensorReadingsHistory200ResponseInnerCo2{}

// GetOrganizationSensorReadingsHistory200ResponseInnerCo2 Reading for the 'co2' metric. This will only be present if the 'metric' property equals 'co2'.
type GetOrganizationSensorReadingsHistory200ResponseInnerCo2 struct {
	// CO2 reading in parts per million.
	Concentration *int32 `json:"concentration,omitempty"`
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerCo2 instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerCo2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSensorReadingsHistory200ResponseInnerCo2() *GetOrganizationSensorReadingsHistory200ResponseInnerCo2 {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerCo2{}
	return &this
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerCo2WithDefaults instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerCo2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSensorReadingsHistory200ResponseInnerCo2WithDefaults() *GetOrganizationSensorReadingsHistory200ResponseInnerCo2 {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerCo2{}
	return &this
}

// GetConcentration returns the Concentration field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerCo2) GetConcentration() int32 {
	if o == nil || IsNil(o.Concentration) {
		var ret int32
		return ret
	}
	return *o.Concentration
}

// GetConcentrationOk returns a tuple with the Concentration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerCo2) GetConcentrationOk() (*int32, bool) {
	if o == nil || IsNil(o.Concentration) {
		return nil, false
	}
	return o.Concentration, true
}

// HasConcentration returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerCo2) HasConcentration() bool {
	if o != nil && !IsNil(o.Concentration) {
		return true
	}

	return false
}

// SetConcentration gets a reference to the given int32 and assigns it to the Concentration field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerCo2) SetConcentration(v int32) {
	o.Concentration = &v
}

func (o GetOrganizationSensorReadingsHistory200ResponseInnerCo2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSensorReadingsHistory200ResponseInnerCo2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Concentration) {
		toSerialize["concentration"] = o.Concentration
	}
	return toSerialize, nil
}

type NullableGetOrganizationSensorReadingsHistory200ResponseInnerCo2 struct {
	value *GetOrganizationSensorReadingsHistory200ResponseInnerCo2
	isSet bool
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerCo2) Get() *GetOrganizationSensorReadingsHistory200ResponseInnerCo2 {
	return v.value
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerCo2) Set(val *GetOrganizationSensorReadingsHistory200ResponseInnerCo2) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerCo2) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerCo2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSensorReadingsHistory200ResponseInnerCo2(val *GetOrganizationSensorReadingsHistory200ResponseInnerCo2) *NullableGetOrganizationSensorReadingsHistory200ResponseInnerCo2 {
	return &NullableGetOrganizationSensorReadingsHistory200ResponseInnerCo2{value: val, isSet: true}
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerCo2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerCo2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

