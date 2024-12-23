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

// checks if the GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner{}

// GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner struct for GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner
type GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner struct {
	// Type of address for the device uplink. Available options are: ipv4, ipv6.
	Protocol *string `json:"protocol,omitempty"`
	// Indicates how the device uplink address is assigned. Available options are: static, dynamic.
	AssignmentMode *string `json:"assignmentMode,omitempty"`
	// Device uplink address.
	Address *string `json:"address,omitempty"`
	// Device uplink gateway address.
	Gateway *string `json:"gateway,omitempty"`
	Public *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic `json:"public,omitempty"`
	Vlan *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerVlan `json:"vlan,omitempty"`
}

// NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner instantiates a new GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner() *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner {
	this := GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner{}
	return &this
}

// NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerWithDefaults instantiates a new GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerWithDefaults() *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner {
	this := GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) SetProtocol(v string) {
	o.Protocol = &v
}

// GetAssignmentMode returns the AssignmentMode field value if set, zero value otherwise.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) GetAssignmentMode() string {
	if o == nil || IsNil(o.AssignmentMode) {
		var ret string
		return ret
	}
	return *o.AssignmentMode
}

// GetAssignmentModeOk returns a tuple with the AssignmentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) GetAssignmentModeOk() (*string, bool) {
	if o == nil || IsNil(o.AssignmentMode) {
		return nil, false
	}
	return o.AssignmentMode, true
}

// HasAssignmentMode returns a boolean if a field has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) HasAssignmentMode() bool {
	if o != nil && !IsNil(o.AssignmentMode) {
		return true
	}

	return false
}

// SetAssignmentMode gets a reference to the given string and assigns it to the AssignmentMode field.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) SetAssignmentMode(v string) {
	o.AssignmentMode = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) SetAddress(v string) {
	o.Address = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) GetGateway() string {
	if o == nil || IsNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) SetGateway(v string) {
	o.Gateway = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) GetPublic() GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic {
	if o == nil || IsNil(o.Public) {
		var ret GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) GetPublicOk() (*GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic, bool) {
	if o == nil || IsNil(o.Public) {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) HasPublic() bool {
	if o != nil && !IsNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic and assigns it to the Public field.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) SetPublic(v GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerPublic) {
	o.Public = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) GetVlan() GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerVlan {
	if o == nil || IsNil(o.Vlan) {
		var ret GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerVlan
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) GetVlanOk() (*GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerVlan, bool) {
	if o == nil || IsNil(o.Vlan) {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) HasVlan() bool {
	if o != nil && !IsNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerVlan and assigns it to the Vlan field.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) SetVlan(v GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInnerVlan) {
	o.Vlan = &v
}

func (o GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Public) {
		toSerialize["public"] = o.Public
	}
	if !IsNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	return toSerialize, nil
}

type NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner struct {
	value *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner
	isSet bool
}

func (v NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) Get() *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner {
	return v.value
}

func (v *NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) Set(val *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner(val *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) *NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner {
	return &NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner{value: val, isSet: true}
}

func (v NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInnerAddressesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


