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

// checks if the UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess{}

// UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess Details associated with a free access plan with limits.
type UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess struct {
	// Whether or not free access is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// How long a device can use a network for free.
	DurationInMinutes *int32 `json:"durationInMinutes,omitempty"`
}

// NewUpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess() *UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess {
	this := UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess{}
	return &this
}

// NewUpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccessWithDefaults instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccessWithDefaults() *UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess {
	this := UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDurationInMinutes returns the DurationInMinutes field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess) GetDurationInMinutes() int32 {
	if o == nil || IsNil(o.DurationInMinutes) {
		var ret int32
		return ret
	}
	return *o.DurationInMinutes
}

// GetDurationInMinutesOk returns a tuple with the DurationInMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess) GetDurationInMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.DurationInMinutes) {
		return nil, false
	}
	return o.DurationInMinutes, true
}

// HasDurationInMinutes returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess) HasDurationInMinutes() bool {
	if o != nil && !IsNil(o.DurationInMinutes) {
		return true
	}

	return false
}

// SetDurationInMinutes gets a reference to the given int32 and assigns it to the DurationInMinutes field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess) SetDurationInMinutes(v int32) {
	o.DurationInMinutes = &v
}

func (o UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.DurationInMinutes) {
		toSerialize["durationInMinutes"] = o.DurationInMinutes
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess struct {
	value *UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess) Get() *UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess) Set(val *UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess(val *UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess) *NullableUpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess {
	return &NullableUpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


