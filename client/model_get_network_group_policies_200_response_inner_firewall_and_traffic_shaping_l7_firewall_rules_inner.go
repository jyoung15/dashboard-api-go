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

// checks if the GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner{}

// GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner struct for GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner
type GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner struct {
	// The policy applied to matching traffic. Must be 'deny'.
	Policy *string `json:"policy,omitempty"`
	// Type of the L7 Rule. Must be 'application', 'applicationCategory', 'host', 'port' or 'ipRange'
	Type *string `json:"type,omitempty"`
	// The 'value' of what you want to block. If 'type' is 'host', 'port' or 'ipRange', 'value' must be a string matching either a hostname (e.g. somewhere.com), a port (e.g. 8080), or an IP range (e.g. 192.1.0.0/16). If 'type' is 'application' or 'applicationCategory', then 'value' must be an object with an ID for the application.
	Value *string `json:"value,omitempty"`
}

// NewGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner instantiates a new GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner() *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner {
	this := GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner{}
	return &this
}

// NewGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInnerWithDefaults instantiates a new GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInnerWithDefaults() *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner {
	this := GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner) GetPolicy() string {
	if o == nil || IsNil(o.Policy) {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner) GetPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner) SetPolicy(v string) {
	o.Policy = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner) SetValue(v string) {
	o.Value = &v
}

func (o GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner struct {
	value *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner
	isSet bool
}

func (v NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner) Get() *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner {
	return v.value
}

func (v *NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner) Set(val *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner(val *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner) *NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner {
	return &NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner{value: val, isSet: true}
}

func (v NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


