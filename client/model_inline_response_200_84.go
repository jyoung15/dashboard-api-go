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

// InlineResponse20084 struct for InlineResponse20084
type InlineResponse20084 struct {
	// The list of VPN peers
	Peers []InlineResponse20084Peers `json:"peers,omitempty"`
}

// NewInlineResponse20084 instantiates a new InlineResponse20084 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20084() *InlineResponse20084 {
	this := InlineResponse20084{}
	return &this
}

// NewInlineResponse20084WithDefaults instantiates a new InlineResponse20084 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20084WithDefaults() *InlineResponse20084 {
	this := InlineResponse20084{}
	return &this
}

// GetPeers returns the Peers field value if set, zero value otherwise.
func (o *InlineResponse20084) GetPeers() []InlineResponse20084Peers {
	if o == nil || isNil(o.Peers) {
		var ret []InlineResponse20084Peers
		return ret
	}
	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20084) GetPeersOk() ([]InlineResponse20084Peers, bool) {
	if o == nil || isNil(o.Peers) {
    return nil, false
	}
	return o.Peers, true
}

// HasPeers returns a boolean if a field has been set.
func (o *InlineResponse20084) HasPeers() bool {
	if o != nil && !isNil(o.Peers) {
		return true
	}

	return false
}

// SetPeers gets a reference to the given []InlineResponse20084Peers and assigns it to the Peers field.
func (o *InlineResponse20084) SetPeers(v []InlineResponse20084Peers) {
	o.Peers = v
}

func (o InlineResponse20084) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Peers) {
		toSerialize["peers"] = o.Peers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20084 struct {
	value *InlineResponse20084
	isSet bool
}

func (v NullableInlineResponse20084) Get() *InlineResponse20084 {
	return v.value
}

func (v *NullableInlineResponse20084) Set(val *InlineResponse20084) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20084) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20084) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20084(val *InlineResponse20084) *NullableInlineResponse20084 {
	return &NullableInlineResponse20084{value: val, isSet: true}
}

func (v NullableInlineResponse20084) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20084) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

