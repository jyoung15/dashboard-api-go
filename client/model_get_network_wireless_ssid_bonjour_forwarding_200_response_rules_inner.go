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

// checks if the GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner{}

// GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner struct for GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner
type GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner struct {
	// Desctiption of the bonjour forwarding rule
	Description *string `json:"description,omitempty"`
	// The ID of the service VLAN. Required
	VlanId *string `json:"vlanId,omitempty"`
	// A list of Bonjour services. At least one service must be specified. Available services are 'All Services', 'AirPlay', 'AFP', 'BitTorrent', 'FTP', 'iChat', 'iTunes', 'Printers', 'Samba', 'Scanners' and 'SSH'
	Services []string `json:"services,omitempty"`
}

// NewGetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner instantiates a new GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner() *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner {
	this := GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner{}
	return &this
}

// NewGetNetworkWirelessSsidBonjourForwarding200ResponseRulesInnerWithDefaults instantiates a new GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessSsidBonjourForwarding200ResponseRulesInnerWithDefaults() *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner {
	this := GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) SetDescription(v string) {
	o.Description = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) GetVlanId() string {
	if o == nil || IsNil(o.VlanId) {
		var ret string
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) GetVlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given string and assigns it to the VlanId field.
func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) SetVlanId(v string) {
	o.VlanId = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) GetServices() []string {
	if o == nil || IsNil(o.Services) {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) GetServicesOk() ([]string, bool) {
	if o == nil || IsNil(o.Services) {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) HasServices() bool {
	if o != nil && !IsNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) SetServices(v []string) {
	o.Services = v
}

func (o GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	if !IsNil(o.Services) {
		toSerialize["services"] = o.Services
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner struct {
	value *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner
	isSet bool
}

func (v NullableGetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) Get() *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner {
	return v.value
}

func (v *NullableGetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) Set(val *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner(val *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) *NullableGetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner {
	return &NullableGetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


