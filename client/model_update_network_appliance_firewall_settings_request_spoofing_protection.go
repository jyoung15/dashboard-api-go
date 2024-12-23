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

// checks if the UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection{}

// UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection Spoofing protection settings
type UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection struct {
	IpSourceGuard *UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard `json:"ipSourceGuard,omitempty"`
}

// NewUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection instantiates a new UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection() *UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection {
	this := UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection{}
	return &this
}

// NewUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionWithDefaults instantiates a new UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionWithDefaults() *UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection {
	this := UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection{}
	return &this
}

// GetIpSourceGuard returns the IpSourceGuard field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection) GetIpSourceGuard() UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard {
	if o == nil || IsNil(o.IpSourceGuard) {
		var ret UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard
		return ret
	}
	return *o.IpSourceGuard
}

// GetIpSourceGuardOk returns a tuple with the IpSourceGuard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection) GetIpSourceGuardOk() (*UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard, bool) {
	if o == nil || IsNil(o.IpSourceGuard) {
		return nil, false
	}
	return o.IpSourceGuard, true
}

// HasIpSourceGuard returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection) HasIpSourceGuard() bool {
	if o != nil && !IsNil(o.IpSourceGuard) {
		return true
	}

	return false
}

// SetIpSourceGuard gets a reference to the given UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard and assigns it to the IpSourceGuard field.
func (o *UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection) SetIpSourceGuard(v UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtectionIpSourceGuard) {
	o.IpSourceGuard = &v
}

func (o UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IpSourceGuard) {
		toSerialize["ipSourceGuard"] = o.IpSourceGuard
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection struct {
	value *UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection
	isSet bool
}

func (v NullableUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection) Get() *UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection {
	return v.value
}

func (v *NullableUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection) Set(val *UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection(val *UpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection) *NullableUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection {
	return &NullableUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceFirewallSettingsRequestSpoofingProtection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


