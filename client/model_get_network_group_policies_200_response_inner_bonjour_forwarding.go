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

// checks if the GetNetworkGroupPolicies200ResponseInnerBonjourForwarding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkGroupPolicies200ResponseInnerBonjourForwarding{}

// GetNetworkGroupPolicies200ResponseInnerBonjourForwarding The Bonjour settings for your group policy. Only valid if your network has a wireless configuration.
type GetNetworkGroupPolicies200ResponseInnerBonjourForwarding struct {
	// How Bonjour rules are applied. Can be 'network default', 'ignore' or 'custom'.
	Settings *string `json:"settings,omitempty"`
	// A list of the Bonjour forwarding rules for your group policy. If 'settings' is set to 'custom', at least one rule must be specified.
	Rules []GetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner `json:"rules,omitempty"`
}

// NewGetNetworkGroupPolicies200ResponseInnerBonjourForwarding instantiates a new GetNetworkGroupPolicies200ResponseInnerBonjourForwarding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkGroupPolicies200ResponseInnerBonjourForwarding() *GetNetworkGroupPolicies200ResponseInnerBonjourForwarding {
	this := GetNetworkGroupPolicies200ResponseInnerBonjourForwarding{}
	return &this
}

// NewGetNetworkGroupPolicies200ResponseInnerBonjourForwardingWithDefaults instantiates a new GetNetworkGroupPolicies200ResponseInnerBonjourForwarding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkGroupPolicies200ResponseInnerBonjourForwardingWithDefaults() *GetNetworkGroupPolicies200ResponseInnerBonjourForwarding {
	this := GetNetworkGroupPolicies200ResponseInnerBonjourForwarding{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *GetNetworkGroupPolicies200ResponseInnerBonjourForwarding) GetSettings() string {
	if o == nil || IsNil(o.Settings) {
		var ret string
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerBonjourForwarding) GetSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerBonjourForwarding) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given string and assigns it to the Settings field.
func (o *GetNetworkGroupPolicies200ResponseInnerBonjourForwarding) SetSettings(v string) {
	o.Settings = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *GetNetworkGroupPolicies200ResponseInnerBonjourForwarding) GetRules() []GetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner {
	if o == nil || IsNil(o.Rules) {
		var ret []GetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerBonjourForwarding) GetRulesOk() ([]GetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerBonjourForwarding) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []GetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner and assigns it to the Rules field.
func (o *GetNetworkGroupPolicies200ResponseInnerBonjourForwarding) SetRules(v []GetNetworkGroupPolicies200ResponseInnerBonjourForwardingRulesInner) {
	o.Rules = v
}

func (o GetNetworkGroupPolicies200ResponseInnerBonjourForwarding) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkGroupPolicies200ResponseInnerBonjourForwarding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableGetNetworkGroupPolicies200ResponseInnerBonjourForwarding struct {
	value *GetNetworkGroupPolicies200ResponseInnerBonjourForwarding
	isSet bool
}

func (v NullableGetNetworkGroupPolicies200ResponseInnerBonjourForwarding) Get() *GetNetworkGroupPolicies200ResponseInnerBonjourForwarding {
	return v.value
}

func (v *NullableGetNetworkGroupPolicies200ResponseInnerBonjourForwarding) Set(val *GetNetworkGroupPolicies200ResponseInnerBonjourForwarding) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkGroupPolicies200ResponseInnerBonjourForwarding) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkGroupPolicies200ResponseInnerBonjourForwarding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkGroupPolicies200ResponseInnerBonjourForwarding(val *GetNetworkGroupPolicies200ResponseInnerBonjourForwarding) *NullableGetNetworkGroupPolicies200ResponseInnerBonjourForwarding {
	return &NullableGetNetworkGroupPolicies200ResponseInnerBonjourForwarding{value: val, isSet: true}
}

func (v NullableGetNetworkGroupPolicies200ResponseInnerBonjourForwarding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkGroupPolicies200ResponseInnerBonjourForwarding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


