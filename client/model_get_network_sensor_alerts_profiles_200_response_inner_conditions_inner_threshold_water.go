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

// GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater Water detection threshold. 'present' must be provided and set to true.
type GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater struct {
	// Alerting threshold for a water detection event. Must be set to true.
	Present bool `json:"present"`
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater(present bool) *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater{}
	this.Present = present
	return &this
}

// NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWaterWithDefaults instantiates a new GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWaterWithDefaults() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater {
	this := GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater{}
	return &this
}

// GetPresent returns the Present field value
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater) GetPresent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Present
}

// GetPresentOk returns a tuple with the Present field value
// and a boolean to check if the value has been set.
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater) GetPresentOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Present, true
}

// SetPresent sets field value
func (o *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater) SetPresent(v bool) {
	o.Present = v
}

func (o GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["present"] = o.Present
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater struct {
	value *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater
	isSet bool
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater) Get() *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater {
	return v.value
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater) Set(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater(val *GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater) *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater {
	return &NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater{value: val, isSet: true}
}

func (v NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThresholdWater) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


