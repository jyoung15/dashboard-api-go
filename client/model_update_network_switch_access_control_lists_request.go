/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateNetworkSwitchAccessControlListsRequest struct for UpdateNetworkSwitchAccessControlListsRequest
type UpdateNetworkSwitchAccessControlListsRequest struct {
	// An ordered array of the access control list rules (not including the default rule). An empty array will clear the rules.
	Rules []UpdateNetworkSwitchAccessControlListsRequestRulesInner `json:"rules"`
}

// NewUpdateNetworkSwitchAccessControlListsRequest instantiates a new UpdateNetworkSwitchAccessControlListsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchAccessControlListsRequest(rules []UpdateNetworkSwitchAccessControlListsRequestRulesInner) *UpdateNetworkSwitchAccessControlListsRequest {
	this := UpdateNetworkSwitchAccessControlListsRequest{}
	this.Rules = rules
	return &this
}

// NewUpdateNetworkSwitchAccessControlListsRequestWithDefaults instantiates a new UpdateNetworkSwitchAccessControlListsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchAccessControlListsRequestWithDefaults() *UpdateNetworkSwitchAccessControlListsRequest {
	this := UpdateNetworkSwitchAccessControlListsRequest{}
	return &this
}

// GetRules returns the Rules field value
func (o *UpdateNetworkSwitchAccessControlListsRequest) GetRules() []UpdateNetworkSwitchAccessControlListsRequestRulesInner {
	if o == nil {
		var ret []UpdateNetworkSwitchAccessControlListsRequestRulesInner
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchAccessControlListsRequest) GetRulesOk() ([]UpdateNetworkSwitchAccessControlListsRequestRulesInner, bool) {
	if o == nil {
    return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *UpdateNetworkSwitchAccessControlListsRequest) SetRules(v []UpdateNetworkSwitchAccessControlListsRequestRulesInner) {
	o.Rules = v
}

func (o UpdateNetworkSwitchAccessControlListsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkSwitchAccessControlListsRequest struct {
	value *UpdateNetworkSwitchAccessControlListsRequest
	isSet bool
}

func (v NullableUpdateNetworkSwitchAccessControlListsRequest) Get() *UpdateNetworkSwitchAccessControlListsRequest {
	return v.value
}

func (v *NullableUpdateNetworkSwitchAccessControlListsRequest) Set(val *UpdateNetworkSwitchAccessControlListsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchAccessControlListsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchAccessControlListsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchAccessControlListsRequest(val *UpdateNetworkSwitchAccessControlListsRequest) *NullableUpdateNetworkSwitchAccessControlListsRequest {
	return &NullableUpdateNetworkSwitchAccessControlListsRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchAccessControlListsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchAccessControlListsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


