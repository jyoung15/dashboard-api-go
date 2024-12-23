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

// checks if the GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6{}

// GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6 IPv6 addressing
type GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6 struct {
	// Assignment mode
	AssignmentMode *string `json:"assignmentMode,omitempty"`
	// IPv6 address
	Address *string `json:"address,omitempty"`
	// IPv6 subnet
	Prefix *string `json:"prefix,omitempty"`
	// IPv6 gateway
	Gateway *string `json:"gateway,omitempty"`
}

// NewGetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6 instantiates a new GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6() *GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6 {
	this := GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6{}
	return &this
}

// NewGetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6WithDefaults instantiates a new GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6WithDefaults() *GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6 {
	this := GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6{}
	return &this
}

// GetAssignmentMode returns the AssignmentMode field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) GetAssignmentMode() string {
	if o == nil || IsNil(o.AssignmentMode) {
		var ret string
		return ret
	}
	return *o.AssignmentMode
}

// GetAssignmentModeOk returns a tuple with the AssignmentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) GetAssignmentModeOk() (*string, bool) {
	if o == nil || IsNil(o.AssignmentMode) {
		return nil, false
	}
	return o.AssignmentMode, true
}

// HasAssignmentMode returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) HasAssignmentMode() bool {
	if o != nil && !IsNil(o.AssignmentMode) {
		return true
	}

	return false
}

// SetAssignmentMode gets a reference to the given string and assigns it to the AssignmentMode field.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) SetAssignmentMode(v string) {
	o.AssignmentMode = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) SetAddress(v string) {
	o.Address = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) SetPrefix(v string) {
	o.Prefix = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) GetGateway() string {
	if o == nil || IsNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) SetGateway(v string) {
	o.Gateway = &v
}

func (o GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssignmentMode) {
		toSerialize["assignmentMode"] = o.AssignmentMode
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	return toSerialize, nil
}

type NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6 struct {
	value *GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6
	isSet bool
}

func (v NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) Get() *GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6 {
	return v.value
}

func (v *NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) Set(val *GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6(val *GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) *NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6 {
	return &NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6{value: val, isSet: true}
}

func (v NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


