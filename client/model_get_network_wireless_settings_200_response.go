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

// checks if the GetNetworkWirelessSettings200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessSettings200Response{}

// GetNetworkWirelessSettings200Response struct for GetNetworkWirelessSettings200Response
type GetNetworkWirelessSettings200Response struct {
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
	NamedVlans *GetNetworkWirelessSettings200ResponseNamedVlans `json:"namedVlans,omitempty"`
	RegulatoryDomain *GetNetworkWirelessSettings200ResponseRegulatoryDomain `json:"regulatoryDomain,omitempty"`
}

// NewGetNetworkWirelessSettings200Response instantiates a new GetNetworkWirelessSettings200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessSettings200Response() *GetNetworkWirelessSettings200Response {
	this := GetNetworkWirelessSettings200Response{}
	return &this
}

// NewGetNetworkWirelessSettings200ResponseWithDefaults instantiates a new GetNetworkWirelessSettings200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessSettings200ResponseWithDefaults() *GetNetworkWirelessSettings200Response {
	this := GetNetworkWirelessSettings200Response{}
	return &this
}

// GetMeshingEnabled returns the MeshingEnabled field value if set, zero value otherwise.
func (o *GetNetworkWirelessSettings200Response) GetMeshingEnabled() bool {
	if o == nil || IsNil(o.MeshingEnabled) {
		var ret bool
		return ret
	}
	return *o.MeshingEnabled
}

// GetMeshingEnabledOk returns a tuple with the MeshingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSettings200Response) GetMeshingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.MeshingEnabled) {
		return nil, false
	}
	return o.MeshingEnabled, true
}

// HasMeshingEnabled returns a boolean if a field has been set.
func (o *GetNetworkWirelessSettings200Response) HasMeshingEnabled() bool {
	if o != nil && !IsNil(o.MeshingEnabled) {
		return true
	}

	return false
}

// SetMeshingEnabled gets a reference to the given bool and assigns it to the MeshingEnabled field.
func (o *GetNetworkWirelessSettings200Response) SetMeshingEnabled(v bool) {
	o.MeshingEnabled = &v
}

// GetIpv6BridgeEnabled returns the Ipv6BridgeEnabled field value if set, zero value otherwise.
func (o *GetNetworkWirelessSettings200Response) GetIpv6BridgeEnabled() bool {
	if o == nil || IsNil(o.Ipv6BridgeEnabled) {
		var ret bool
		return ret
	}
	return *o.Ipv6BridgeEnabled
}

// GetIpv6BridgeEnabledOk returns a tuple with the Ipv6BridgeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSettings200Response) GetIpv6BridgeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Ipv6BridgeEnabled) {
		return nil, false
	}
	return o.Ipv6BridgeEnabled, true
}

// HasIpv6BridgeEnabled returns a boolean if a field has been set.
func (o *GetNetworkWirelessSettings200Response) HasIpv6BridgeEnabled() bool {
	if o != nil && !IsNil(o.Ipv6BridgeEnabled) {
		return true
	}

	return false
}

// SetIpv6BridgeEnabled gets a reference to the given bool and assigns it to the Ipv6BridgeEnabled field.
func (o *GetNetworkWirelessSettings200Response) SetIpv6BridgeEnabled(v bool) {
	o.Ipv6BridgeEnabled = &v
}

// GetLocationAnalyticsEnabled returns the LocationAnalyticsEnabled field value if set, zero value otherwise.
func (o *GetNetworkWirelessSettings200Response) GetLocationAnalyticsEnabled() bool {
	if o == nil || IsNil(o.LocationAnalyticsEnabled) {
		var ret bool
		return ret
	}
	return *o.LocationAnalyticsEnabled
}

// GetLocationAnalyticsEnabledOk returns a tuple with the LocationAnalyticsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSettings200Response) GetLocationAnalyticsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.LocationAnalyticsEnabled) {
		return nil, false
	}
	return o.LocationAnalyticsEnabled, true
}

// HasLocationAnalyticsEnabled returns a boolean if a field has been set.
func (o *GetNetworkWirelessSettings200Response) HasLocationAnalyticsEnabled() bool {
	if o != nil && !IsNil(o.LocationAnalyticsEnabled) {
		return true
	}

	return false
}

// SetLocationAnalyticsEnabled gets a reference to the given bool and assigns it to the LocationAnalyticsEnabled field.
func (o *GetNetworkWirelessSettings200Response) SetLocationAnalyticsEnabled(v bool) {
	o.LocationAnalyticsEnabled = &v
}

// GetUpgradeStrategy returns the UpgradeStrategy field value if set, zero value otherwise.
func (o *GetNetworkWirelessSettings200Response) GetUpgradeStrategy() string {
	if o == nil || IsNil(o.UpgradeStrategy) {
		var ret string
		return ret
	}
	return *o.UpgradeStrategy
}

// GetUpgradeStrategyOk returns a tuple with the UpgradeStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSettings200Response) GetUpgradeStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.UpgradeStrategy) {
		return nil, false
	}
	return o.UpgradeStrategy, true
}

// HasUpgradeStrategy returns a boolean if a field has been set.
func (o *GetNetworkWirelessSettings200Response) HasUpgradeStrategy() bool {
	if o != nil && !IsNil(o.UpgradeStrategy) {
		return true
	}

	return false
}

// SetUpgradeStrategy gets a reference to the given string and assigns it to the UpgradeStrategy field.
func (o *GetNetworkWirelessSettings200Response) SetUpgradeStrategy(v string) {
	o.UpgradeStrategy = &v
}

// GetLedLightsOn returns the LedLightsOn field value if set, zero value otherwise.
func (o *GetNetworkWirelessSettings200Response) GetLedLightsOn() bool {
	if o == nil || IsNil(o.LedLightsOn) {
		var ret bool
		return ret
	}
	return *o.LedLightsOn
}

// GetLedLightsOnOk returns a tuple with the LedLightsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSettings200Response) GetLedLightsOnOk() (*bool, bool) {
	if o == nil || IsNil(o.LedLightsOn) {
		return nil, false
	}
	return o.LedLightsOn, true
}

// HasLedLightsOn returns a boolean if a field has been set.
func (o *GetNetworkWirelessSettings200Response) HasLedLightsOn() bool {
	if o != nil && !IsNil(o.LedLightsOn) {
		return true
	}

	return false
}

// SetLedLightsOn gets a reference to the given bool and assigns it to the LedLightsOn field.
func (o *GetNetworkWirelessSettings200Response) SetLedLightsOn(v bool) {
	o.LedLightsOn = &v
}

// GetNamedVlans returns the NamedVlans field value if set, zero value otherwise.
func (o *GetNetworkWirelessSettings200Response) GetNamedVlans() GetNetworkWirelessSettings200ResponseNamedVlans {
	if o == nil || IsNil(o.NamedVlans) {
		var ret GetNetworkWirelessSettings200ResponseNamedVlans
		return ret
	}
	return *o.NamedVlans
}

// GetNamedVlansOk returns a tuple with the NamedVlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSettings200Response) GetNamedVlansOk() (*GetNetworkWirelessSettings200ResponseNamedVlans, bool) {
	if o == nil || IsNil(o.NamedVlans) {
		return nil, false
	}
	return o.NamedVlans, true
}

// HasNamedVlans returns a boolean if a field has been set.
func (o *GetNetworkWirelessSettings200Response) HasNamedVlans() bool {
	if o != nil && !IsNil(o.NamedVlans) {
		return true
	}

	return false
}

// SetNamedVlans gets a reference to the given GetNetworkWirelessSettings200ResponseNamedVlans and assigns it to the NamedVlans field.
func (o *GetNetworkWirelessSettings200Response) SetNamedVlans(v GetNetworkWirelessSettings200ResponseNamedVlans) {
	o.NamedVlans = &v
}

// GetRegulatoryDomain returns the RegulatoryDomain field value if set, zero value otherwise.
func (o *GetNetworkWirelessSettings200Response) GetRegulatoryDomain() GetNetworkWirelessSettings200ResponseRegulatoryDomain {
	if o == nil || IsNil(o.RegulatoryDomain) {
		var ret GetNetworkWirelessSettings200ResponseRegulatoryDomain
		return ret
	}
	return *o.RegulatoryDomain
}

// GetRegulatoryDomainOk returns a tuple with the RegulatoryDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSettings200Response) GetRegulatoryDomainOk() (*GetNetworkWirelessSettings200ResponseRegulatoryDomain, bool) {
	if o == nil || IsNil(o.RegulatoryDomain) {
		return nil, false
	}
	return o.RegulatoryDomain, true
}

// HasRegulatoryDomain returns a boolean if a field has been set.
func (o *GetNetworkWirelessSettings200Response) HasRegulatoryDomain() bool {
	if o != nil && !IsNil(o.RegulatoryDomain) {
		return true
	}

	return false
}

// SetRegulatoryDomain gets a reference to the given GetNetworkWirelessSettings200ResponseRegulatoryDomain and assigns it to the RegulatoryDomain field.
func (o *GetNetworkWirelessSettings200Response) SetRegulatoryDomain(v GetNetworkWirelessSettings200ResponseRegulatoryDomain) {
	o.RegulatoryDomain = &v
}

func (o GetNetworkWirelessSettings200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessSettings200Response) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.RegulatoryDomain) {
		toSerialize["regulatoryDomain"] = o.RegulatoryDomain
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessSettings200Response struct {
	value *GetNetworkWirelessSettings200Response
	isSet bool
}

func (v NullableGetNetworkWirelessSettings200Response) Get() *GetNetworkWirelessSettings200Response {
	return v.value
}

func (v *NullableGetNetworkWirelessSettings200Response) Set(val *GetNetworkWirelessSettings200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessSettings200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessSettings200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessSettings200Response(val *GetNetworkWirelessSettings200Response) *NullableGetNetworkWirelessSettings200Response {
	return &NullableGetNetworkWirelessSettings200Response{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessSettings200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessSettings200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


