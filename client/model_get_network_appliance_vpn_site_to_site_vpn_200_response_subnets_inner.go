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

// checks if the GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner{}

// GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner struct for GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner
type GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner struct {
	// The CIDR notation subnet used within the VPN
	LocalSubnet *string `json:"localSubnet,omitempty"`
	// Indicates the presence of the subnet in the VPN
	UseVpn *bool `json:"useVpn,omitempty"`
}

// NewGetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner instantiates a new GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner() *GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner {
	this := GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner{}
	return &this
}

// NewGetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInnerWithDefaults instantiates a new GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInnerWithDefaults() *GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner {
	this := GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner{}
	return &this
}

// GetLocalSubnet returns the LocalSubnet field value if set, zero value otherwise.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner) GetLocalSubnet() string {
	if o == nil || IsNil(o.LocalSubnet) {
		var ret string
		return ret
	}
	return *o.LocalSubnet
}

// GetLocalSubnetOk returns a tuple with the LocalSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner) GetLocalSubnetOk() (*string, bool) {
	if o == nil || IsNil(o.LocalSubnet) {
		return nil, false
	}
	return o.LocalSubnet, true
}

// HasLocalSubnet returns a boolean if a field has been set.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner) HasLocalSubnet() bool {
	if o != nil && !IsNil(o.LocalSubnet) {
		return true
	}

	return false
}

// SetLocalSubnet gets a reference to the given string and assigns it to the LocalSubnet field.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner) SetLocalSubnet(v string) {
	o.LocalSubnet = &v
}

// GetUseVpn returns the UseVpn field value if set, zero value otherwise.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner) GetUseVpn() bool {
	if o == nil || IsNil(o.UseVpn) {
		var ret bool
		return ret
	}
	return *o.UseVpn
}

// GetUseVpnOk returns a tuple with the UseVpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner) GetUseVpnOk() (*bool, bool) {
	if o == nil || IsNil(o.UseVpn) {
		return nil, false
	}
	return o.UseVpn, true
}

// HasUseVpn returns a boolean if a field has been set.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner) HasUseVpn() bool {
	if o != nil && !IsNil(o.UseVpn) {
		return true
	}

	return false
}

// SetUseVpn gets a reference to the given bool and assigns it to the UseVpn field.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner) SetUseVpn(v bool) {
	o.UseVpn = &v
}

func (o GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LocalSubnet) {
		toSerialize["localSubnet"] = o.LocalSubnet
	}
	if !IsNil(o.UseVpn) {
		toSerialize["useVpn"] = o.UseVpn
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner struct {
	value *GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner
	isSet bool
}

func (v NullableGetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner) Get() *GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner {
	return v.value
}

func (v *NullableGetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner) Set(val *GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner(val *GetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner) *NullableGetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner {
	return &NullableGetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceVpnSiteToSiteVpn200ResponseSubnetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


