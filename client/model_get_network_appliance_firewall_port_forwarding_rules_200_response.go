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

// checks if the GetNetworkApplianceFirewallPortForwardingRules200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceFirewallPortForwardingRules200Response{}

// GetNetworkApplianceFirewallPortForwardingRules200Response struct for GetNetworkApplianceFirewallPortForwardingRules200Response
type GetNetworkApplianceFirewallPortForwardingRules200Response struct {
	// An array of port forwarding rules
	Rules []GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner `json:"rules,omitempty"`
}

// NewGetNetworkApplianceFirewallPortForwardingRules200Response instantiates a new GetNetworkApplianceFirewallPortForwardingRules200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceFirewallPortForwardingRules200Response() *GetNetworkApplianceFirewallPortForwardingRules200Response {
	this := GetNetworkApplianceFirewallPortForwardingRules200Response{}
	return &this
}

// NewGetNetworkApplianceFirewallPortForwardingRules200ResponseWithDefaults instantiates a new GetNetworkApplianceFirewallPortForwardingRules200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceFirewallPortForwardingRules200ResponseWithDefaults() *GetNetworkApplianceFirewallPortForwardingRules200Response {
	this := GetNetworkApplianceFirewallPortForwardingRules200Response{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallPortForwardingRules200Response) GetRules() []GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner {
	if o == nil || IsNil(o.Rules) {
		var ret []GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallPortForwardingRules200Response) GetRulesOk() ([]GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallPortForwardingRules200Response) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner and assigns it to the Rules field.
func (o *GetNetworkApplianceFirewallPortForwardingRules200Response) SetRules(v []GetNetworkApplianceFirewallPortForwardingRules200ResponseRulesInner) {
	o.Rules = v
}

func (o GetNetworkApplianceFirewallPortForwardingRules200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceFirewallPortForwardingRules200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceFirewallPortForwardingRules200Response struct {
	value *GetNetworkApplianceFirewallPortForwardingRules200Response
	isSet bool
}

func (v NullableGetNetworkApplianceFirewallPortForwardingRules200Response) Get() *GetNetworkApplianceFirewallPortForwardingRules200Response {
	return v.value
}

func (v *NullableGetNetworkApplianceFirewallPortForwardingRules200Response) Set(val *GetNetworkApplianceFirewallPortForwardingRules200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceFirewallPortForwardingRules200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceFirewallPortForwardingRules200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceFirewallPortForwardingRules200Response(val *GetNetworkApplianceFirewallPortForwardingRules200Response) *NullableGetNetworkApplianceFirewallPortForwardingRules200Response {
	return &NullableGetNetworkApplianceFirewallPortForwardingRules200Response{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceFirewallPortForwardingRules200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceFirewallPortForwardingRules200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


