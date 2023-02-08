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

// UpdateNetworkApplianceSecurityIntrusionRequest struct for UpdateNetworkApplianceSecurityIntrusionRequest
type UpdateNetworkApplianceSecurityIntrusionRequest struct {
	// Set mode to 'disabled'/'detection'/'prevention' (optional - omitting will leave current config unchanged)
	Mode *string `json:"mode,omitempty"`
	// Set the detection ruleset 'connectivity'/'balanced'/'security' (optional - omitting will leave current config unchanged). Default value is 'balanced' if none currently saved
	IdsRulesets *string `json:"idsRulesets,omitempty"`
	ProtectedNetworks *UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks `json:"protectedNetworks,omitempty"`
}

// NewUpdateNetworkApplianceSecurityIntrusionRequest instantiates a new UpdateNetworkApplianceSecurityIntrusionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceSecurityIntrusionRequest() *UpdateNetworkApplianceSecurityIntrusionRequest {
	this := UpdateNetworkApplianceSecurityIntrusionRequest{}
	return &this
}

// NewUpdateNetworkApplianceSecurityIntrusionRequestWithDefaults instantiates a new UpdateNetworkApplianceSecurityIntrusionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceSecurityIntrusionRequestWithDefaults() *UpdateNetworkApplianceSecurityIntrusionRequest {
	this := UpdateNetworkApplianceSecurityIntrusionRequest{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSecurityIntrusionRequest) GetMode() string {
	if o == nil || isNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSecurityIntrusionRequest) GetModeOk() (*string, bool) {
	if o == nil || isNil(o.Mode) {
    return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSecurityIntrusionRequest) HasMode() bool {
	if o != nil && !isNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *UpdateNetworkApplianceSecurityIntrusionRequest) SetMode(v string) {
	o.Mode = &v
}

// GetIdsRulesets returns the IdsRulesets field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSecurityIntrusionRequest) GetIdsRulesets() string {
	if o == nil || isNil(o.IdsRulesets) {
		var ret string
		return ret
	}
	return *o.IdsRulesets
}

// GetIdsRulesetsOk returns a tuple with the IdsRulesets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSecurityIntrusionRequest) GetIdsRulesetsOk() (*string, bool) {
	if o == nil || isNil(o.IdsRulesets) {
    return nil, false
	}
	return o.IdsRulesets, true
}

// HasIdsRulesets returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSecurityIntrusionRequest) HasIdsRulesets() bool {
	if o != nil && !isNil(o.IdsRulesets) {
		return true
	}

	return false
}

// SetIdsRulesets gets a reference to the given string and assigns it to the IdsRulesets field.
func (o *UpdateNetworkApplianceSecurityIntrusionRequest) SetIdsRulesets(v string) {
	o.IdsRulesets = &v
}

// GetProtectedNetworks returns the ProtectedNetworks field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSecurityIntrusionRequest) GetProtectedNetworks() UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks {
	if o == nil || isNil(o.ProtectedNetworks) {
		var ret UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks
		return ret
	}
	return *o.ProtectedNetworks
}

// GetProtectedNetworksOk returns a tuple with the ProtectedNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSecurityIntrusionRequest) GetProtectedNetworksOk() (*UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks, bool) {
	if o == nil || isNil(o.ProtectedNetworks) {
    return nil, false
	}
	return o.ProtectedNetworks, true
}

// HasProtectedNetworks returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSecurityIntrusionRequest) HasProtectedNetworks() bool {
	if o != nil && !isNil(o.ProtectedNetworks) {
		return true
	}

	return false
}

// SetProtectedNetworks gets a reference to the given UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks and assigns it to the ProtectedNetworks field.
func (o *UpdateNetworkApplianceSecurityIntrusionRequest) SetProtectedNetworks(v UpdateNetworkApplianceSecurityIntrusionRequestProtectedNetworks) {
	o.ProtectedNetworks = &v
}

func (o UpdateNetworkApplianceSecurityIntrusionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !isNil(o.IdsRulesets) {
		toSerialize["idsRulesets"] = o.IdsRulesets
	}
	if !isNil(o.ProtectedNetworks) {
		toSerialize["protectedNetworks"] = o.ProtectedNetworks
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkApplianceSecurityIntrusionRequest struct {
	value *UpdateNetworkApplianceSecurityIntrusionRequest
	isSet bool
}

func (v NullableUpdateNetworkApplianceSecurityIntrusionRequest) Get() *UpdateNetworkApplianceSecurityIntrusionRequest {
	return v.value
}

func (v *NullableUpdateNetworkApplianceSecurityIntrusionRequest) Set(val *UpdateNetworkApplianceSecurityIntrusionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceSecurityIntrusionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceSecurityIntrusionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceSecurityIntrusionRequest(val *UpdateNetworkApplianceSecurityIntrusionRequest) *NullableUpdateNetworkApplianceSecurityIntrusionRequest {
	return &NullableUpdateNetworkApplianceSecurityIntrusionRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceSecurityIntrusionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceSecurityIntrusionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


