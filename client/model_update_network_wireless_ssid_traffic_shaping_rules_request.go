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

// checks if the UpdateNetworkWirelessSsidTrafficShapingRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidTrafficShapingRulesRequest{}

// UpdateNetworkWirelessSsidTrafficShapingRulesRequest struct for UpdateNetworkWirelessSsidTrafficShapingRulesRequest
type UpdateNetworkWirelessSsidTrafficShapingRulesRequest struct {
	// Whether traffic shaping rules are applied to clients on your SSID.
	TrafficShapingEnabled *bool `json:"trafficShapingEnabled,omitempty"`
	// Whether default traffic shaping rules are enabled (true) or disabled (false). There are 4 default rules, which can be seen on your network's traffic shaping page. Note that default rules count against the rule limit of 8.
	DefaultRulesEnabled *bool `json:"defaultRulesEnabled,omitempty"`
	//     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules. 
	Rules []GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner `json:"rules,omitempty"`
}

// NewUpdateNetworkWirelessSsidTrafficShapingRulesRequest instantiates a new UpdateNetworkWirelessSsidTrafficShapingRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidTrafficShapingRulesRequest() *UpdateNetworkWirelessSsidTrafficShapingRulesRequest {
	this := UpdateNetworkWirelessSsidTrafficShapingRulesRequest{}
	return &this
}

// NewUpdateNetworkWirelessSsidTrafficShapingRulesRequestWithDefaults instantiates a new UpdateNetworkWirelessSsidTrafficShapingRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidTrafficShapingRulesRequestWithDefaults() *UpdateNetworkWirelessSsidTrafficShapingRulesRequest {
	this := UpdateNetworkWirelessSsidTrafficShapingRulesRequest{}
	return &this
}

// GetTrafficShapingEnabled returns the TrafficShapingEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequest) GetTrafficShapingEnabled() bool {
	if o == nil || IsNil(o.TrafficShapingEnabled) {
		var ret bool
		return ret
	}
	return *o.TrafficShapingEnabled
}

// GetTrafficShapingEnabledOk returns a tuple with the TrafficShapingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequest) GetTrafficShapingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TrafficShapingEnabled) {
		return nil, false
	}
	return o.TrafficShapingEnabled, true
}

// HasTrafficShapingEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequest) HasTrafficShapingEnabled() bool {
	if o != nil && !IsNil(o.TrafficShapingEnabled) {
		return true
	}

	return false
}

// SetTrafficShapingEnabled gets a reference to the given bool and assigns it to the TrafficShapingEnabled field.
func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequest) SetTrafficShapingEnabled(v bool) {
	o.TrafficShapingEnabled = &v
}

// GetDefaultRulesEnabled returns the DefaultRulesEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequest) GetDefaultRulesEnabled() bool {
	if o == nil || IsNil(o.DefaultRulesEnabled) {
		var ret bool
		return ret
	}
	return *o.DefaultRulesEnabled
}

// GetDefaultRulesEnabledOk returns a tuple with the DefaultRulesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequest) GetDefaultRulesEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultRulesEnabled) {
		return nil, false
	}
	return o.DefaultRulesEnabled, true
}

// HasDefaultRulesEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequest) HasDefaultRulesEnabled() bool {
	if o != nil && !IsNil(o.DefaultRulesEnabled) {
		return true
	}

	return false
}

// SetDefaultRulesEnabled gets a reference to the given bool and assigns it to the DefaultRulesEnabled field.
func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequest) SetDefaultRulesEnabled(v bool) {
	o.DefaultRulesEnabled = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequest) GetRules() []GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner {
	if o == nil || IsNil(o.Rules) {
		var ret []GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequest) GetRulesOk() ([]GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequest) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner and assigns it to the Rules field.
func (o *UpdateNetworkWirelessSsidTrafficShapingRulesRequest) SetRules(v []GetNetworkWirelessSsidTrafficShapingRules200ResponseRulesInner) {
	o.Rules = v
}

func (o UpdateNetworkWirelessSsidTrafficShapingRulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidTrafficShapingRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TrafficShapingEnabled) {
		toSerialize["trafficShapingEnabled"] = o.TrafficShapingEnabled
	}
	if !IsNil(o.DefaultRulesEnabled) {
		toSerialize["defaultRulesEnabled"] = o.DefaultRulesEnabled
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidTrafficShapingRulesRequest struct {
	value *UpdateNetworkWirelessSsidTrafficShapingRulesRequest
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidTrafficShapingRulesRequest) Get() *UpdateNetworkWirelessSsidTrafficShapingRulesRequest {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidTrafficShapingRulesRequest) Set(val *UpdateNetworkWirelessSsidTrafficShapingRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidTrafficShapingRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidTrafficShapingRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidTrafficShapingRulesRequest(val *UpdateNetworkWirelessSsidTrafficShapingRulesRequest) *NullableUpdateNetworkWirelessSsidTrafficShapingRulesRequest {
	return &NullableUpdateNetworkWirelessSsidTrafficShapingRulesRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidTrafficShapingRulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidTrafficShapingRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


