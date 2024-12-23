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

// checks if the UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner{}

// UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner struct for UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner
type UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner struct {
	// 'Deny' traffic specified by this rule
	Policy *string `json:"policy,omitempty"`
	// Type of the L7 rule. One of: 'application', 'applicationCategory', 'host', 'port', 'ipRange'
	Type *string `json:"type,omitempty"`
	// The 'value' of what you want to block. Format of 'value' varies depending on type of the rule. The application categories and application ids can be retrieved from the the 'MX L7 application categories' endpoint. The countries follow the two-letter ISO 3166-1 alpha-2 format.
	Value *string `json:"value,omitempty"`
}

// NewUpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner instantiates a new UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner() *UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner {
	this := UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner{}
	return &this
}

// NewUpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInnerWithDefaults instantiates a new UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInnerWithDefaults() *UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner {
	this := UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner) GetPolicy() string {
	if o == nil || IsNil(o.Policy) {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner) GetPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner) SetPolicy(v string) {
	o.Policy = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner) SetValue(v string) {
	o.Value = &v
}

func (o UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner) ToMap() (map[string]interface{}, error) {
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

type NullableUpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner struct {
	value *UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner
	isSet bool
}

func (v NullableUpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner) Get() *UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner {
	return v.value
}

func (v *NullableUpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner) Set(val *UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner(val *UpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner) *NullableUpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner {
	return &NullableUpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceFirewallL7FirewallRulesRequestRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


