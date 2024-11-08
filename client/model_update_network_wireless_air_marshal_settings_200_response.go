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

// checks if the UpdateNetworkWirelessAirMarshalSettings200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessAirMarshalSettings200Response{}

// UpdateNetworkWirelessAirMarshalSettings200Response struct for UpdateNetworkWirelessAirMarshalSettings200Response
type UpdateNetworkWirelessAirMarshalSettings200Response struct {
	// The network ID
	NetworkId *string `json:"networkId,omitempty"`
	// Indicates whether or not clients are allowed to       connect to rogue SSIDs. (blocked by default)
	DefaultPolicy *string `json:"defaultPolicy,omitempty"`
}

// NewUpdateNetworkWirelessAirMarshalSettings200Response instantiates a new UpdateNetworkWirelessAirMarshalSettings200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessAirMarshalSettings200Response() *UpdateNetworkWirelessAirMarshalSettings200Response {
	this := UpdateNetworkWirelessAirMarshalSettings200Response{}
	return &this
}

// NewUpdateNetworkWirelessAirMarshalSettings200ResponseWithDefaults instantiates a new UpdateNetworkWirelessAirMarshalSettings200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessAirMarshalSettings200ResponseWithDefaults() *UpdateNetworkWirelessAirMarshalSettings200Response {
	this := UpdateNetworkWirelessAirMarshalSettings200Response{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessAirMarshalSettings200Response) GetNetworkId() string {
	if o == nil || IsNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessAirMarshalSettings200Response) GetNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessAirMarshalSettings200Response) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *UpdateNetworkWirelessAirMarshalSettings200Response) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetDefaultPolicy returns the DefaultPolicy field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessAirMarshalSettings200Response) GetDefaultPolicy() string {
	if o == nil || IsNil(o.DefaultPolicy) {
		var ret string
		return ret
	}
	return *o.DefaultPolicy
}

// GetDefaultPolicyOk returns a tuple with the DefaultPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessAirMarshalSettings200Response) GetDefaultPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultPolicy) {
		return nil, false
	}
	return o.DefaultPolicy, true
}

// HasDefaultPolicy returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessAirMarshalSettings200Response) HasDefaultPolicy() bool {
	if o != nil && !IsNil(o.DefaultPolicy) {
		return true
	}

	return false
}

// SetDefaultPolicy gets a reference to the given string and assigns it to the DefaultPolicy field.
func (o *UpdateNetworkWirelessAirMarshalSettings200Response) SetDefaultPolicy(v string) {
	o.DefaultPolicy = &v
}

func (o UpdateNetworkWirelessAirMarshalSettings200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessAirMarshalSettings200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !IsNil(o.DefaultPolicy) {
		toSerialize["defaultPolicy"] = o.DefaultPolicy
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessAirMarshalSettings200Response struct {
	value *UpdateNetworkWirelessAirMarshalSettings200Response
	isSet bool
}

func (v NullableUpdateNetworkWirelessAirMarshalSettings200Response) Get() *UpdateNetworkWirelessAirMarshalSettings200Response {
	return v.value
}

func (v *NullableUpdateNetworkWirelessAirMarshalSettings200Response) Set(val *UpdateNetworkWirelessAirMarshalSettings200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessAirMarshalSettings200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessAirMarshalSettings200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessAirMarshalSettings200Response(val *UpdateNetworkWirelessAirMarshalSettings200Response) *NullableUpdateNetworkWirelessAirMarshalSettings200Response {
	return &NullableUpdateNetworkWirelessAirMarshalSettings200Response{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessAirMarshalSettings200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessAirMarshalSettings200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


