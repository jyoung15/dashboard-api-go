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

// checks if the UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship{}

// UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship Details associated with guest sponsored splash.
type UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship struct {
	// Duration in minutes of sponsored guest authorization. Must be between 1 and 60480 (6 weeks)
	DurationInMinutes *int32 `json:"durationInMinutes,omitempty"`
	// Whether or not guests can specify how much time they are requesting.
	GuestCanRequestTimeframe *bool `json:"guestCanRequestTimeframe,omitempty"`
}

// NewUpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship() *UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship {
	this := UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship{}
	return &this
}

// NewUpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorshipWithDefaults instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorshipWithDefaults() *UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship {
	this := UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship{}
	return &this
}

// GetDurationInMinutes returns the DurationInMinutes field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship) GetDurationInMinutes() int32 {
	if o == nil || IsNil(o.DurationInMinutes) {
		var ret int32
		return ret
	}
	return *o.DurationInMinutes
}

// GetDurationInMinutesOk returns a tuple with the DurationInMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship) GetDurationInMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.DurationInMinutes) {
		return nil, false
	}
	return o.DurationInMinutes, true
}

// HasDurationInMinutes returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship) HasDurationInMinutes() bool {
	if o != nil && !IsNil(o.DurationInMinutes) {
		return true
	}

	return false
}

// SetDurationInMinutes gets a reference to the given int32 and assigns it to the DurationInMinutes field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship) SetDurationInMinutes(v int32) {
	o.DurationInMinutes = &v
}

// GetGuestCanRequestTimeframe returns the GuestCanRequestTimeframe field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship) GetGuestCanRequestTimeframe() bool {
	if o == nil || IsNil(o.GuestCanRequestTimeframe) {
		var ret bool
		return ret
	}
	return *o.GuestCanRequestTimeframe
}

// GetGuestCanRequestTimeframeOk returns a tuple with the GuestCanRequestTimeframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship) GetGuestCanRequestTimeframeOk() (*bool, bool) {
	if o == nil || IsNil(o.GuestCanRequestTimeframe) {
		return nil, false
	}
	return o.GuestCanRequestTimeframe, true
}

// HasGuestCanRequestTimeframe returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship) HasGuestCanRequestTimeframe() bool {
	if o != nil && !IsNil(o.GuestCanRequestTimeframe) {
		return true
	}

	return false
}

// SetGuestCanRequestTimeframe gets a reference to the given bool and assigns it to the GuestCanRequestTimeframe field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship) SetGuestCanRequestTimeframe(v bool) {
	o.GuestCanRequestTimeframe = &v
}

func (o UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DurationInMinutes) {
		toSerialize["durationInMinutes"] = o.DurationInMinutes
	}
	if !IsNil(o.GuestCanRequestTimeframe) {
		toSerialize["guestCanRequestTimeframe"] = o.GuestCanRequestTimeframe
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship struct {
	value *UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship) Get() *UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship) Set(val *UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship(val *UpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship) *NullableUpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship {
	return &NullableUpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestGuestSponsorship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


