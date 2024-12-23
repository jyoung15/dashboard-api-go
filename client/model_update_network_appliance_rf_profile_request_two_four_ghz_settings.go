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

// checks if the UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings{}

// UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings Settings related to 2.4Ghz band
type UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings struct {
	// Sets min bitrate (Mbps) of 2.4Ghz band. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	MinBitrate *float32 `json:"minBitrate,omitempty"`
	// Determines whether ax radio on 2.4Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering.
	AxEnabled *bool `json:"axEnabled,omitempty"`
}

// NewUpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings instantiates a new UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings() *UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings {
	this := UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings{}
	return &this
}

// NewUpdateNetworkApplianceRfProfileRequestTwoFourGhzSettingsWithDefaults instantiates a new UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceRfProfileRequestTwoFourGhzSettingsWithDefaults() *UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings {
	this := UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings{}
	return &this
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings) GetMinBitrate() float32 {
	if o == nil || IsNil(o.MinBitrate) {
		var ret float32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings) GetMinBitrateOk() (*float32, bool) {
	if o == nil || IsNil(o.MinBitrate) {
		return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings) HasMinBitrate() bool {
	if o != nil && !IsNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given float32 and assigns it to the MinBitrate field.
func (o *UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings) SetMinBitrate(v float32) {
	o.MinBitrate = &v
}

// GetAxEnabled returns the AxEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings) GetAxEnabled() bool {
	if o == nil || IsNil(o.AxEnabled) {
		var ret bool
		return ret
	}
	return *o.AxEnabled
}

// GetAxEnabledOk returns a tuple with the AxEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings) GetAxEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AxEnabled) {
		return nil, false
	}
	return o.AxEnabled, true
}

// HasAxEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings) HasAxEnabled() bool {
	if o != nil && !IsNil(o.AxEnabled) {
		return true
	}

	return false
}

// SetAxEnabled gets a reference to the given bool and assigns it to the AxEnabled field.
func (o *UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings) SetAxEnabled(v bool) {
	o.AxEnabled = &v
}

func (o UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MinBitrate) {
		toSerialize["minBitrate"] = o.MinBitrate
	}
	if !IsNil(o.AxEnabled) {
		toSerialize["axEnabled"] = o.AxEnabled
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings struct {
	value *UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings
	isSet bool
}

func (v NullableUpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings) Get() *UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings {
	return v.value
}

func (v *NullableUpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings) Set(val *UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings(val *UpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings) *NullableUpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings {
	return &NullableUpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceRfProfileRequestTwoFourGhzSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


