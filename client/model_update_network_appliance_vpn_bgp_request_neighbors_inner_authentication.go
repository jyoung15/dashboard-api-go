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

// checks if the UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication{}

// UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication Authentication settings between BGP peers.
type UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication struct {
	// Password to configure MD5 authentication between BGP peers.
	Password *string `json:"password,omitempty"`
}

// NewUpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication instantiates a new UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication() *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication {
	this := UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication{}
	return &this
}

// NewUpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthenticationWithDefaults instantiates a new UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthenticationWithDefaults() *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication {
	this := UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication) SetPassword(v string) {
	o.Password = &v
}

func (o UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication struct {
	value *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication
	isSet bool
}

func (v NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication) Get() *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication {
	return v.value
}

func (v *NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication) Set(val *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication(val *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication) *NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication {
	return &NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


