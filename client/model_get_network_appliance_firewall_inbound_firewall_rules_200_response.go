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

// checks if the GetNetworkApplianceFirewallInboundFirewallRules200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceFirewallInboundFirewallRules200Response{}

// GetNetworkApplianceFirewallInboundFirewallRules200Response struct for GetNetworkApplianceFirewallInboundFirewallRules200Response
type GetNetworkApplianceFirewallInboundFirewallRules200Response struct {
	// An ordered array of the firewall rules (not including the default rule)
	Rules []GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner `json:"rules,omitempty"`
	// Log the special default rule (boolean value - enable only if you've configured a syslog server) (optional)
	SyslogDefaultRule *bool `json:"syslogDefaultRule,omitempty"`
}

// NewGetNetworkApplianceFirewallInboundFirewallRules200Response instantiates a new GetNetworkApplianceFirewallInboundFirewallRules200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceFirewallInboundFirewallRules200Response() *GetNetworkApplianceFirewallInboundFirewallRules200Response {
	this := GetNetworkApplianceFirewallInboundFirewallRules200Response{}
	return &this
}

// NewGetNetworkApplianceFirewallInboundFirewallRules200ResponseWithDefaults instantiates a new GetNetworkApplianceFirewallInboundFirewallRules200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceFirewallInboundFirewallRules200ResponseWithDefaults() *GetNetworkApplianceFirewallInboundFirewallRules200Response {
	this := GetNetworkApplianceFirewallInboundFirewallRules200Response{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200Response) GetRules() []GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner {
	if o == nil || IsNil(o.Rules) {
		var ret []GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200Response) GetRulesOk() ([]GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200Response) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner and assigns it to the Rules field.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200Response) SetRules(v []GetNetworkApplianceFirewallInboundCellularFirewallRules200ResponseRulesInner) {
	o.Rules = v
}

// GetSyslogDefaultRule returns the SyslogDefaultRule field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200Response) GetSyslogDefaultRule() bool {
	if o == nil || IsNil(o.SyslogDefaultRule) {
		var ret bool
		return ret
	}
	return *o.SyslogDefaultRule
}

// GetSyslogDefaultRuleOk returns a tuple with the SyslogDefaultRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200Response) GetSyslogDefaultRuleOk() (*bool, bool) {
	if o == nil || IsNil(o.SyslogDefaultRule) {
		return nil, false
	}
	return o.SyslogDefaultRule, true
}

// HasSyslogDefaultRule returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200Response) HasSyslogDefaultRule() bool {
	if o != nil && !IsNil(o.SyslogDefaultRule) {
		return true
	}

	return false
}

// SetSyslogDefaultRule gets a reference to the given bool and assigns it to the SyslogDefaultRule field.
func (o *GetNetworkApplianceFirewallInboundFirewallRules200Response) SetSyslogDefaultRule(v bool) {
	o.SyslogDefaultRule = &v
}

func (o GetNetworkApplianceFirewallInboundFirewallRules200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceFirewallInboundFirewallRules200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !IsNil(o.SyslogDefaultRule) {
		toSerialize["syslogDefaultRule"] = o.SyslogDefaultRule
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceFirewallInboundFirewallRules200Response struct {
	value *GetNetworkApplianceFirewallInboundFirewallRules200Response
	isSet bool
}

func (v NullableGetNetworkApplianceFirewallInboundFirewallRules200Response) Get() *GetNetworkApplianceFirewallInboundFirewallRules200Response {
	return v.value
}

func (v *NullableGetNetworkApplianceFirewallInboundFirewallRules200Response) Set(val *GetNetworkApplianceFirewallInboundFirewallRules200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceFirewallInboundFirewallRules200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceFirewallInboundFirewallRules200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceFirewallInboundFirewallRules200Response(val *GetNetworkApplianceFirewallInboundFirewallRules200Response) *NullableGetNetworkApplianceFirewallInboundFirewallRules200Response {
	return &NullableGetNetworkApplianceFirewallInboundFirewallRules200Response{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceFirewallInboundFirewallRules200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceFirewallInboundFirewallRules200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


