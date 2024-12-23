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

// checks if the UpdateNetworkWirelessSettingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSettingsRequest{}

// UpdateNetworkWirelessSettingsRequest struct for UpdateNetworkWirelessSettingsRequest
type UpdateNetworkWirelessSettingsRequest struct {
	// Toggle for enabling or disabling meshing in a network
	MeshingEnabled *bool `json:"meshingEnabled,omitempty"`
	// Toggle for enabling or disabling IPv6 bridging in a network (Note: if enabled, SSIDs must also be configured to use bridge mode)
	Ipv6BridgeEnabled *bool `json:"ipv6BridgeEnabled,omitempty"`
	// Toggle for enabling or disabling location analytics for your network
	LocationAnalyticsEnabled *bool `json:"locationAnalyticsEnabled,omitempty"`
	// The default strategy that network devices will use to perform an upgrade. Requires firmware version MR 26.8 or higher.
	UpgradeStrategy *string `json:"upgradeStrategy,omitempty"`
	// Toggle for enabling or disabling LED lights on all APs in the network (making them run dark)
	LedLightsOn *bool `json:"ledLightsOn,omitempty"`
	NamedVlans *UpdateNetworkWirelessSettingsRequestNamedVlans `json:"namedVlans,omitempty"`
}

// NewUpdateNetworkWirelessSettingsRequest instantiates a new UpdateNetworkWirelessSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSettingsRequest() *UpdateNetworkWirelessSettingsRequest {
	this := UpdateNetworkWirelessSettingsRequest{}
	return &this
}

// NewUpdateNetworkWirelessSettingsRequestWithDefaults instantiates a new UpdateNetworkWirelessSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSettingsRequestWithDefaults() *UpdateNetworkWirelessSettingsRequest {
	this := UpdateNetworkWirelessSettingsRequest{}
	return &this
}

// GetMeshingEnabled returns the MeshingEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSettingsRequest) GetMeshingEnabled() bool {
	if o == nil || IsNil(o.MeshingEnabled) {
		var ret bool
		return ret
	}
	return *o.MeshingEnabled
}

// GetMeshingEnabledOk returns a tuple with the MeshingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSettingsRequest) GetMeshingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.MeshingEnabled) {
		return nil, false
	}
	return o.MeshingEnabled, true
}

// HasMeshingEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSettingsRequest) HasMeshingEnabled() bool {
	if o != nil && !IsNil(o.MeshingEnabled) {
		return true
	}

	return false
}

// SetMeshingEnabled gets a reference to the given bool and assigns it to the MeshingEnabled field.
func (o *UpdateNetworkWirelessSettingsRequest) SetMeshingEnabled(v bool) {
	o.MeshingEnabled = &v
}

// GetIpv6BridgeEnabled returns the Ipv6BridgeEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSettingsRequest) GetIpv6BridgeEnabled() bool {
	if o == nil || IsNil(o.Ipv6BridgeEnabled) {
		var ret bool
		return ret
	}
	return *o.Ipv6BridgeEnabled
}

// GetIpv6BridgeEnabledOk returns a tuple with the Ipv6BridgeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSettingsRequest) GetIpv6BridgeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Ipv6BridgeEnabled) {
		return nil, false
	}
	return o.Ipv6BridgeEnabled, true
}

// HasIpv6BridgeEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSettingsRequest) HasIpv6BridgeEnabled() bool {
	if o != nil && !IsNil(o.Ipv6BridgeEnabled) {
		return true
	}

	return false
}

// SetIpv6BridgeEnabled gets a reference to the given bool and assigns it to the Ipv6BridgeEnabled field.
func (o *UpdateNetworkWirelessSettingsRequest) SetIpv6BridgeEnabled(v bool) {
	o.Ipv6BridgeEnabled = &v
}

// GetLocationAnalyticsEnabled returns the LocationAnalyticsEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSettingsRequest) GetLocationAnalyticsEnabled() bool {
	if o == nil || IsNil(o.LocationAnalyticsEnabled) {
		var ret bool
		return ret
	}
	return *o.LocationAnalyticsEnabled
}

// GetLocationAnalyticsEnabledOk returns a tuple with the LocationAnalyticsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSettingsRequest) GetLocationAnalyticsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.LocationAnalyticsEnabled) {
		return nil, false
	}
	return o.LocationAnalyticsEnabled, true
}

// HasLocationAnalyticsEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSettingsRequest) HasLocationAnalyticsEnabled() bool {
	if o != nil && !IsNil(o.LocationAnalyticsEnabled) {
		return true
	}

	return false
}

// SetLocationAnalyticsEnabled gets a reference to the given bool and assigns it to the LocationAnalyticsEnabled field.
func (o *UpdateNetworkWirelessSettingsRequest) SetLocationAnalyticsEnabled(v bool) {
	o.LocationAnalyticsEnabled = &v
}

// GetUpgradeStrategy returns the UpgradeStrategy field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSettingsRequest) GetUpgradeStrategy() string {
	if o == nil || IsNil(o.UpgradeStrategy) {
		var ret string
		return ret
	}
	return *o.UpgradeStrategy
}

// GetUpgradeStrategyOk returns a tuple with the UpgradeStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSettingsRequest) GetUpgradeStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.UpgradeStrategy) {
		return nil, false
	}
	return o.UpgradeStrategy, true
}

// HasUpgradeStrategy returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSettingsRequest) HasUpgradeStrategy() bool {
	if o != nil && !IsNil(o.UpgradeStrategy) {
		return true
	}

	return false
}

// SetUpgradeStrategy gets a reference to the given string and assigns it to the UpgradeStrategy field.
func (o *UpdateNetworkWirelessSettingsRequest) SetUpgradeStrategy(v string) {
	o.UpgradeStrategy = &v
}

// GetLedLightsOn returns the LedLightsOn field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSettingsRequest) GetLedLightsOn() bool {
	if o == nil || IsNil(o.LedLightsOn) {
		var ret bool
		return ret
	}
	return *o.LedLightsOn
}

// GetLedLightsOnOk returns a tuple with the LedLightsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSettingsRequest) GetLedLightsOnOk() (*bool, bool) {
	if o == nil || IsNil(o.LedLightsOn) {
		return nil, false
	}
	return o.LedLightsOn, true
}

// HasLedLightsOn returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSettingsRequest) HasLedLightsOn() bool {
	if o != nil && !IsNil(o.LedLightsOn) {
		return true
	}

	return false
}

// SetLedLightsOn gets a reference to the given bool and assigns it to the LedLightsOn field.
func (o *UpdateNetworkWirelessSettingsRequest) SetLedLightsOn(v bool) {
	o.LedLightsOn = &v
}

// GetNamedVlans returns the NamedVlans field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSettingsRequest) GetNamedVlans() UpdateNetworkWirelessSettingsRequestNamedVlans {
	if o == nil || IsNil(o.NamedVlans) {
		var ret UpdateNetworkWirelessSettingsRequestNamedVlans
		return ret
	}
	return *o.NamedVlans
}

// GetNamedVlansOk returns a tuple with the NamedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSettingsRequest) GetNamedVlansOk() (*UpdateNetworkWirelessSettingsRequestNamedVlans, bool) {
	if o == nil || IsNil(o.NamedVlans) {
		return nil, false
	}
	return o.NamedVlans, true
}

// HasNamedVlans returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSettingsRequest) HasNamedVlans() bool {
	if o != nil && !IsNil(o.NamedVlans) {
		return true
	}

	return false
}

// SetNamedVlans gets a reference to the given UpdateNetworkWirelessSettingsRequestNamedVlans and assigns it to the NamedVlans field.
func (o *UpdateNetworkWirelessSettingsRequest) SetNamedVlans(v UpdateNetworkWirelessSettingsRequestNamedVlans) {
	o.NamedVlans = &v
}

func (o UpdateNetworkWirelessSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSettingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MeshingEnabled) {
		toSerialize["meshingEnabled"] = o.MeshingEnabled
	}
	if !IsNil(o.Ipv6BridgeEnabled) {
		toSerialize["ipv6BridgeEnabled"] = o.Ipv6BridgeEnabled
	}
	if !IsNil(o.LocationAnalyticsEnabled) {
		toSerialize["locationAnalyticsEnabled"] = o.LocationAnalyticsEnabled
	}
	if !IsNil(o.UpgradeStrategy) {
		toSerialize["upgradeStrategy"] = o.UpgradeStrategy
	}
	if !IsNil(o.LedLightsOn) {
		toSerialize["ledLightsOn"] = o.LedLightsOn
	}
	if !IsNil(o.NamedVlans) {
		toSerialize["namedVlans"] = o.NamedVlans
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSettingsRequest struct {
	value *UpdateNetworkWirelessSettingsRequest
	isSet bool
}

func (v NullableUpdateNetworkWirelessSettingsRequest) Get() *UpdateNetworkWirelessSettingsRequest {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSettingsRequest) Set(val *UpdateNetworkWirelessSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSettingsRequest(val *UpdateNetworkWirelessSettingsRequest) *NullableUpdateNetworkWirelessSettingsRequest {
	return &NullableUpdateNetworkWirelessSettingsRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


