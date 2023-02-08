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

// UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner struct for UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner
type UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner struct {
	// IP address of your Active Directory server.
	Host string `json:"host"`
	// (Optional) UDP port the Active Directory server listens on. By default, uses port 3268.
	Port *int32 `json:"port,omitempty"`
}

// NewUpdateNetworkWirelessSsidRequestActiveDirectoryServersInner instantiates a new UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidRequestActiveDirectoryServersInner(host string) *UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner {
	this := UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner{}
	this.Host = host
	return &this
}

// NewUpdateNetworkWirelessSsidRequestActiveDirectoryServersInnerWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidRequestActiveDirectoryServersInnerWithDefaults() *UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner {
	this := UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner{}
	return &this
}

// GetHost returns the Host field value
func (o *UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) GetHostOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) SetPort(v int32) {
	o.Port = &v
}

func (o UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["host"] = o.Host
	}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkWirelessSsidRequestActiveDirectoryServersInner struct {
	value *UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) Get() *UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) Set(val *UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidRequestActiveDirectoryServersInner(val *UpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) *NullableUpdateNetworkWirelessSsidRequestActiveDirectoryServersInner {
	return &NullableUpdateNetworkWirelessSsidRequestActiveDirectoryServersInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidRequestActiveDirectoryServersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


