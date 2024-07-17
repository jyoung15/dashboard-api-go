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

// checks if the GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping{}

// GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping     The firewall and traffic shaping rules and settings for your policy. 
type GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping struct {
	// How firewall and traffic shaping rules are enforced. Can be 'network default', 'ignore' or 'custom'.
	Settings *string `json:"settings,omitempty"`
	//     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules. 
	TrafficShapingRules []GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner `json:"trafficShapingRules,omitempty"`
	// An ordered array of the L3 firewall rules
	L3FirewallRules []GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL3FirewallRulesInner `json:"l3FirewallRules,omitempty"`
	// An ordered array of L7 firewall rules
	L7FirewallRules []GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner `json:"l7FirewallRules,omitempty"`
}

// NewGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping instantiates a new GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping() *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping {
	this := GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping{}
	return &this
}

// NewGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingWithDefaults instantiates a new GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingWithDefaults() *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping {
	this := GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) GetSettings() string {
	if o == nil || IsNil(o.Settings) {
		var ret string
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) GetSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given string and assigns it to the Settings field.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) SetSettings(v string) {
	o.Settings = &v
}

// GetTrafficShapingRules returns the TrafficShapingRules field value if set, zero value otherwise.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) GetTrafficShapingRules() []GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner {
	if o == nil || IsNil(o.TrafficShapingRules) {
		var ret []GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner
		return ret
	}
	return o.TrafficShapingRules
}

// GetTrafficShapingRulesOk returns a tuple with the TrafficShapingRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) GetTrafficShapingRulesOk() ([]GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner, bool) {
	if o == nil || IsNil(o.TrafficShapingRules) {
		return nil, false
	}
	return o.TrafficShapingRules, true
}

// HasTrafficShapingRules returns a boolean if a field has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) HasTrafficShapingRules() bool {
	if o != nil && !IsNil(o.TrafficShapingRules) {
		return true
	}

	return false
}

// SetTrafficShapingRules gets a reference to the given []GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner and assigns it to the TrafficShapingRules field.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) SetTrafficShapingRules(v []GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) {
	o.TrafficShapingRules = v
}

// GetL3FirewallRules returns the L3FirewallRules field value if set, zero value otherwise.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) GetL3FirewallRules() []GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL3FirewallRulesInner {
	if o == nil || IsNil(o.L3FirewallRules) {
		var ret []GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL3FirewallRulesInner
		return ret
	}
	return o.L3FirewallRules
}

// GetL3FirewallRulesOk returns a tuple with the L3FirewallRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) GetL3FirewallRulesOk() ([]GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL3FirewallRulesInner, bool) {
	if o == nil || IsNil(o.L3FirewallRules) {
		return nil, false
	}
	return o.L3FirewallRules, true
}

// HasL3FirewallRules returns a boolean if a field has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) HasL3FirewallRules() bool {
	if o != nil && !IsNil(o.L3FirewallRules) {
		return true
	}

	return false
}

// SetL3FirewallRules gets a reference to the given []GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL3FirewallRulesInner and assigns it to the L3FirewallRules field.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) SetL3FirewallRules(v []GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL3FirewallRulesInner) {
	o.L3FirewallRules = v
}

// GetL7FirewallRules returns the L7FirewallRules field value if set, zero value otherwise.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) GetL7FirewallRules() []GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner {
	if o == nil || IsNil(o.L7FirewallRules) {
		var ret []GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner
		return ret
	}
	return o.L7FirewallRules
}

// GetL7FirewallRulesOk returns a tuple with the L7FirewallRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) GetL7FirewallRulesOk() ([]GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner, bool) {
	if o == nil || IsNil(o.L7FirewallRules) {
		return nil, false
	}
	return o.L7FirewallRules, true
}

// HasL7FirewallRules returns a boolean if a field has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) HasL7FirewallRules() bool {
	if o != nil && !IsNil(o.L7FirewallRules) {
		return true
	}

	return false
}

// SetL7FirewallRules gets a reference to the given []GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner and assigns it to the L7FirewallRules field.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) SetL7FirewallRules(v []GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner) {
	o.L7FirewallRules = v
}

func (o GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.TrafficShapingRules) {
		toSerialize["trafficShapingRules"] = o.TrafficShapingRules
	}
	if !IsNil(o.L3FirewallRules) {
		toSerialize["l3FirewallRules"] = o.L3FirewallRules
	}
	if !IsNil(o.L7FirewallRules) {
		toSerialize["l7FirewallRules"] = o.L7FirewallRules
	}
	return toSerialize, nil
}

type NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping struct {
	value *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping
	isSet bool
}

func (v NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) Get() *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping {
	return v.value
}

func (v *NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) Set(val *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping(val *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) *NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping {
	return &NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping{value: val, isSet: true}
}

func (v NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

