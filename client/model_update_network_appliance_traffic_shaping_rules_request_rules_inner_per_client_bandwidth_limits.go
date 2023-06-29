/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits{}

// UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits     An object describing the bandwidth settings for your rule. 
type UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits struct {
	// How bandwidth limits are applied by your rule. Can be one of 'network default', 'ignore' or 'custom'.
	Settings *string `json:"settings,omitempty"`
	BandwidthLimits *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits `json:"bandwidthLimits,omitempty"`
}

// NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits instantiates a new UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits() *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits {
	this := UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits{}
	return &this
}

// NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsWithDefaults() *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits {
	this := UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits) GetSettings() string {
	if o == nil || IsNil(o.Settings) {
		var ret string
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits) GetSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given string and assigns it to the Settings field.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits) SetSettings(v string) {
	o.Settings = &v
}

// GetBandwidthLimits returns the BandwidthLimits field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits) GetBandwidthLimits() UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits {
	if o == nil || IsNil(o.BandwidthLimits) {
		var ret UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits
		return ret
	}
	return *o.BandwidthLimits
}

// GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits) GetBandwidthLimitsOk() (*UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits, bool) {
	if o == nil || IsNil(o.BandwidthLimits) {
		return nil, false
	}
	return o.BandwidthLimits, true
}

// HasBandwidthLimits returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits) HasBandwidthLimits() bool {
	if o != nil && !IsNil(o.BandwidthLimits) {
		return true
	}

	return false
}

// SetBandwidthLimits gets a reference to the given UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits and assigns it to the BandwidthLimits field.
func (o *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits) SetBandwidthLimits(v UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimitsBandwidthLimits) {
	o.BandwidthLimits = &v
}

func (o UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.BandwidthLimits) {
		toSerialize["bandwidthLimits"] = o.BandwidthLimits
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits struct {
	value *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits
	isSet bool
}

func (v NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits) Get() *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits {
	return v.value
}

func (v *NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits) Set(val *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits(val *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits) *NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits {
	return &NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


