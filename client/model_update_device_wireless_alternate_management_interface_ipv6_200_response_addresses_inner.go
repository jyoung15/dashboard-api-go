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

// checks if the UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner{}

// UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner struct for UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner
type UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner struct {
	// The IP protocol used for the address
	Protocol *string `json:"protocol,omitempty"`
	// The type of address assignment. Either static or dynamic.
	AssignmentMode *string `json:"assignmentMode,omitempty"`
	// The IP address configured for the alternate management interface
	Address *string `json:"address,omitempty"`
	// The gateway address configured for the alternate managment interface
	Gateway *string `json:"gateway,omitempty"`
	// The IPv6 prefix of the interface. Required if IPv6 object is included.
	Prefix *string `json:"prefix,omitempty"`
	Nameservers *UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers `json:"nameservers,omitempty"`
}

// NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner instantiates a new UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner() *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner {
	this := UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner{}
	return &this
}

// NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInnerWithDefaults instantiates a new UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInnerWithDefaults() *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner {
	this := UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) SetProtocol(v string) {
	o.Protocol = &v
}

// GetAssignmentMode returns the AssignmentMode field value if set, zero value otherwise.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetAssignmentMode() string {
	if o == nil || IsNil(o.AssignmentMode) {
		var ret string
		return ret
	}
	return *o.AssignmentMode
}

// GetAssignmentModeOk returns a tuple with the AssignmentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetAssignmentModeOk() (*string, bool) {
	if o == nil || IsNil(o.AssignmentMode) {
		return nil, false
	}
	return o.AssignmentMode, true
}

// HasAssignmentMode returns a boolean if a field has been set.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) HasAssignmentMode() bool {
	if o != nil && !IsNil(o.AssignmentMode) {
		return true
	}

	return false
}

// SetAssignmentMode gets a reference to the given string and assigns it to the AssignmentMode field.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) SetAssignmentMode(v string) {
	o.AssignmentMode = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) SetAddress(v string) {
	o.Address = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetGateway() string {
	if o == nil || IsNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) SetGateway(v string) {
	o.Gateway = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) SetPrefix(v string) {
	o.Prefix = &v
}

// GetNameservers returns the Nameservers field value if set, zero value otherwise.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetNameservers() UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers {
	if o == nil || IsNil(o.Nameservers) {
		var ret UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers
		return ret
	}
	return *o.Nameservers
}

// GetNameserversOk returns a tuple with the Nameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) GetNameserversOk() (*UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers, bool) {
	if o == nil || IsNil(o.Nameservers) {
		return nil, false
	}
	return o.Nameservers, true
}

// HasNameservers returns a boolean if a field has been set.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) HasNameservers() bool {
	if o != nil && !IsNil(o.Nameservers) {
		return true
	}

	return false
}

// SetNameservers gets a reference to the given UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers and assigns it to the Nameservers field.
func (o *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) SetNameservers(v UpdateDeviceWirelessAlternateManagementInterfaceIpv6RequestAddressesInnerNameservers) {
	o.Nameservers = &v
}

func (o UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.AssignmentMode) {
		toSerialize["assignmentMode"] = o.AssignmentMode
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.Nameservers) {
		toSerialize["nameservers"] = o.Nameservers
	}
	return toSerialize, nil
}

type NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner struct {
	value *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner
	isSet bool
}

func (v NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) Get() *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner {
	return v.value
}

func (v *NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) Set(val *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner(val *UpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) *NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner {
	return &NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner{value: val, isSet: true}
}

func (v NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceWirelessAlternateManagementInterfaceIpv6200ResponseAddressesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


