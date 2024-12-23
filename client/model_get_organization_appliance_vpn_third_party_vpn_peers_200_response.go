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

// checks if the GetOrganizationApplianceVpnThirdPartyVPNPeers200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationApplianceVpnThirdPartyVPNPeers200Response{}

// GetOrganizationApplianceVpnThirdPartyVPNPeers200Response struct for GetOrganizationApplianceVpnThirdPartyVPNPeers200Response
type GetOrganizationApplianceVpnThirdPartyVPNPeers200Response struct {
	// The list of VPN peers
	Peers []GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner `json:"peers,omitempty"`
}

// NewGetOrganizationApplianceVpnThirdPartyVPNPeers200Response instantiates a new GetOrganizationApplianceVpnThirdPartyVPNPeers200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationApplianceVpnThirdPartyVPNPeers200Response() *GetOrganizationApplianceVpnThirdPartyVPNPeers200Response {
	this := GetOrganizationApplianceVpnThirdPartyVPNPeers200Response{}
	return &this
}

// NewGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponseWithDefaults instantiates a new GetOrganizationApplianceVpnThirdPartyVPNPeers200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationApplianceVpnThirdPartyVPNPeers200ResponseWithDefaults() *GetOrganizationApplianceVpnThirdPartyVPNPeers200Response {
	this := GetOrganizationApplianceVpnThirdPartyVPNPeers200Response{}
	return &this
}

// GetPeers returns the Peers field value if set, zero value otherwise.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200Response) GetPeers() []GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner {
	if o == nil || IsNil(o.Peers) {
		var ret []GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner
		return ret
	}
	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200Response) GetPeersOk() ([]GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner, bool) {
	if o == nil || IsNil(o.Peers) {
		return nil, false
	}
	return o.Peers, true
}

// HasPeers returns a boolean if a field has been set.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200Response) HasPeers() bool {
	if o != nil && !IsNil(o.Peers) {
		return true
	}

	return false
}

// SetPeers gets a reference to the given []GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner and assigns it to the Peers field.
func (o *GetOrganizationApplianceVpnThirdPartyVPNPeers200Response) SetPeers(v []GetOrganizationApplianceVpnThirdPartyVPNPeers200ResponsePeersInner) {
	o.Peers = v
}

func (o GetOrganizationApplianceVpnThirdPartyVPNPeers200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationApplianceVpnThirdPartyVPNPeers200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Peers) {
		toSerialize["peers"] = o.Peers
	}
	return toSerialize, nil
}

type NullableGetOrganizationApplianceVpnThirdPartyVPNPeers200Response struct {
	value *GetOrganizationApplianceVpnThirdPartyVPNPeers200Response
	isSet bool
}

func (v NullableGetOrganizationApplianceVpnThirdPartyVPNPeers200Response) Get() *GetOrganizationApplianceVpnThirdPartyVPNPeers200Response {
	return v.value
}

func (v *NullableGetOrganizationApplianceVpnThirdPartyVPNPeers200Response) Set(val *GetOrganizationApplianceVpnThirdPartyVPNPeers200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationApplianceVpnThirdPartyVPNPeers200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationApplianceVpnThirdPartyVPNPeers200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationApplianceVpnThirdPartyVPNPeers200Response(val *GetOrganizationApplianceVpnThirdPartyVPNPeers200Response) *NullableGetOrganizationApplianceVpnThirdPartyVPNPeers200Response {
	return &NullableGetOrganizationApplianceVpnThirdPartyVPNPeers200Response{value: val, isSet: true}
}

func (v NullableGetOrganizationApplianceVpnThirdPartyVPNPeers200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationApplianceVpnThirdPartyVPNPeers200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


