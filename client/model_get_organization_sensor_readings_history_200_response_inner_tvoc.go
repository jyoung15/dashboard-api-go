/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GetOrganizationSensorReadingsHistory200ResponseInnerTvoc Reading for the 'tvoc' metric. This will only be present if the 'metric' property equals 'tvoc'.
type GetOrganizationSensorReadingsHistory200ResponseInnerTvoc struct {
	// TVOC reading in micrograms per cubic meter.
	Concentration *int32 `json:"concentration,omitempty"`
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerTvoc instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerTvoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSensorReadingsHistory200ResponseInnerTvoc() *GetOrganizationSensorReadingsHistory200ResponseInnerTvoc {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerTvoc{}
	return &this
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerTvocWithDefaults instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerTvoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSensorReadingsHistory200ResponseInnerTvocWithDefaults() *GetOrganizationSensorReadingsHistory200ResponseInnerTvoc {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerTvoc{}
	return &this
}

// GetConcentration returns the Concentration field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerTvoc) GetConcentration() int32 {
	if o == nil || isNil(o.Concentration) {
		var ret int32
		return ret
	}
	return *o.Concentration
}

// GetConcentrationOk returns a tuple with the Concentration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerTvoc) GetConcentrationOk() (*int32, bool) {
	if o == nil || isNil(o.Concentration) {
    return nil, false
	}
	return o.Concentration, true
}

// HasConcentration returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerTvoc) HasConcentration() bool {
	if o != nil && !isNil(o.Concentration) {
		return true
	}

	return false
}

// SetConcentration gets a reference to the given int32 and assigns it to the Concentration field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerTvoc) SetConcentration(v int32) {
	o.Concentration = &v
}

func (o GetOrganizationSensorReadingsHistory200ResponseInnerTvoc) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Concentration) {
		toSerialize["concentration"] = o.Concentration
	}
	return json.Marshal(toSerialize)
}

type NullableGetOrganizationSensorReadingsHistory200ResponseInnerTvoc struct {
	value *GetOrganizationSensorReadingsHistory200ResponseInnerTvoc
	isSet bool
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerTvoc) Get() *GetOrganizationSensorReadingsHistory200ResponseInnerTvoc {
	return v.value
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerTvoc) Set(val *GetOrganizationSensorReadingsHistory200ResponseInnerTvoc) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerTvoc) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerTvoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSensorReadingsHistory200ResponseInnerTvoc(val *GetOrganizationSensorReadingsHistory200ResponseInnerTvoc) *NullableGetOrganizationSensorReadingsHistory200ResponseInnerTvoc {
	return &NullableGetOrganizationSensorReadingsHistory200ResponseInnerTvoc{value: val, isSet: true}
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerTvoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerTvoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


