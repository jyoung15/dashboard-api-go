/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetDeviceSwitchRoutingInterfaces200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceSwitchRoutingInterfaces200ResponseInner{}

// GetDeviceSwitchRoutingInterfaces200ResponseInner struct for GetDeviceSwitchRoutingInterfaces200ResponseInner
type GetDeviceSwitchRoutingInterfaces200ResponseInner struct {
	// The id
	InterfaceId *string `json:"interfaceId,omitempty"`
	// The name
	Name *string `json:"name,omitempty"`
	// IPv4 subnet
	Subnet *string `json:"subnet,omitempty"`
	// IPv4 address
	InterfaceIp *string `json:"interfaceIp,omitempty"`
	// Multicast routing status
	MulticastRouting *string `json:"multicastRouting,omitempty"`
	// VLAN id
	VlanId *int32 `json:"vlanId,omitempty"`
	// IPv4 default gateway
	DefaultGateway *string `json:"defaultGateway,omitempty"`
	OspfSettings *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings `json:"ospfSettings,omitempty"`
	OspfV3 *GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3 `json:"ospfV3,omitempty"`
	Ipv6 *GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6 `json:"ipv6,omitempty"`
}

// NewGetDeviceSwitchRoutingInterfaces200ResponseInner instantiates a new GetDeviceSwitchRoutingInterfaces200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceSwitchRoutingInterfaces200ResponseInner() *GetDeviceSwitchRoutingInterfaces200ResponseInner {
	this := GetDeviceSwitchRoutingInterfaces200ResponseInner{}
	return &this
}

// NewGetDeviceSwitchRoutingInterfaces200ResponseInnerWithDefaults instantiates a new GetDeviceSwitchRoutingInterfaces200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceSwitchRoutingInterfaces200ResponseInnerWithDefaults() *GetDeviceSwitchRoutingInterfaces200ResponseInner {
	this := GetDeviceSwitchRoutingInterfaces200ResponseInner{}
	return &this
}

// GetInterfaceId returns the InterfaceId field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetInterfaceId() string {
	if o == nil || IsNil(o.InterfaceId) {
		var ret string
		return ret
	}
	return *o.InterfaceId
}

// GetInterfaceIdOk returns a tuple with the InterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetInterfaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceId) {
		return nil, false
	}
	return o.InterfaceId, true
}

// HasInterfaceId returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) HasInterfaceId() bool {
	if o != nil && !IsNil(o.InterfaceId) {
		return true
	}

	return false
}

// SetInterfaceId gets a reference to the given string and assigns it to the InterfaceId field.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) SetInterfaceId(v string) {
	o.InterfaceId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetSubnet() string {
	if o == nil || IsNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetSubnetOk() (*string, bool) {
	if o == nil || IsNil(o.Subnet) {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) HasSubnet() bool {
	if o != nil && !IsNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) SetSubnet(v string) {
	o.Subnet = &v
}

// GetInterfaceIp returns the InterfaceIp field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetInterfaceIp() string {
	if o == nil || IsNil(o.InterfaceIp) {
		var ret string
		return ret
	}
	return *o.InterfaceIp
}

// GetInterfaceIpOk returns a tuple with the InterfaceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetInterfaceIpOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceIp) {
		return nil, false
	}
	return o.InterfaceIp, true
}

// HasInterfaceIp returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) HasInterfaceIp() bool {
	if o != nil && !IsNil(o.InterfaceIp) {
		return true
	}

	return false
}

// SetInterfaceIp gets a reference to the given string and assigns it to the InterfaceIp field.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) SetInterfaceIp(v string) {
	o.InterfaceIp = &v
}

// GetMulticastRouting returns the MulticastRouting field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetMulticastRouting() string {
	if o == nil || IsNil(o.MulticastRouting) {
		var ret string
		return ret
	}
	return *o.MulticastRouting
}

// GetMulticastRoutingOk returns a tuple with the MulticastRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetMulticastRoutingOk() (*string, bool) {
	if o == nil || IsNil(o.MulticastRouting) {
		return nil, false
	}
	return o.MulticastRouting, true
}

// HasMulticastRouting returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) HasMulticastRouting() bool {
	if o != nil && !IsNil(o.MulticastRouting) {
		return true
	}

	return false
}

// SetMulticastRouting gets a reference to the given string and assigns it to the MulticastRouting field.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) SetMulticastRouting(v string) {
	o.MulticastRouting = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetVlanId() int32 {
	if o == nil || IsNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) SetVlanId(v int32) {
	o.VlanId = &v
}

// GetDefaultGateway returns the DefaultGateway field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetDefaultGateway() string {
	if o == nil || IsNil(o.DefaultGateway) {
		var ret string
		return ret
	}
	return *o.DefaultGateway
}

// GetDefaultGatewayOk returns a tuple with the DefaultGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetDefaultGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultGateway) {
		return nil, false
	}
	return o.DefaultGateway, true
}

// HasDefaultGateway returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) HasDefaultGateway() bool {
	if o != nil && !IsNil(o.DefaultGateway) {
		return true
	}

	return false
}

// SetDefaultGateway gets a reference to the given string and assigns it to the DefaultGateway field.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) SetDefaultGateway(v string) {
	o.DefaultGateway = &v
}

// GetOspfSettings returns the OspfSettings field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetOspfSettings() GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings {
	if o == nil || IsNil(o.OspfSettings) {
		var ret GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings
		return ret
	}
	return *o.OspfSettings
}

// GetOspfSettingsOk returns a tuple with the OspfSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetOspfSettingsOk() (*GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings, bool) {
	if o == nil || IsNil(o.OspfSettings) {
		return nil, false
	}
	return o.OspfSettings, true
}

// HasOspfSettings returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) HasOspfSettings() bool {
	if o != nil && !IsNil(o.OspfSettings) {
		return true
	}

	return false
}

// SetOspfSettings gets a reference to the given GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings and assigns it to the OspfSettings field.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) SetOspfSettings(v GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings) {
	o.OspfSettings = &v
}

// GetOspfV3 returns the OspfV3 field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetOspfV3() GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3 {
	if o == nil || IsNil(o.OspfV3) {
		var ret GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3
		return ret
	}
	return *o.OspfV3
}

// GetOspfV3Ok returns a tuple with the OspfV3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetOspfV3Ok() (*GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3, bool) {
	if o == nil || IsNil(o.OspfV3) {
		return nil, false
	}
	return o.OspfV3, true
}

// HasOspfV3 returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) HasOspfV3() bool {
	if o != nil && !IsNil(o.OspfV3) {
		return true
	}

	return false
}

// SetOspfV3 gets a reference to the given GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3 and assigns it to the OspfV3 field.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) SetOspfV3(v GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3) {
	o.OspfV3 = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetIpv6() GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6 {
	if o == nil || IsNil(o.Ipv6) {
		var ret GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) GetIpv6Ok() (*GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6, bool) {
	if o == nil || IsNil(o.Ipv6) {
		return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) HasIpv6() bool {
	if o != nil && !IsNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6 and assigns it to the Ipv6 field.
func (o *GetDeviceSwitchRoutingInterfaces200ResponseInner) SetIpv6(v GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6) {
	o.Ipv6 = &v
}

func (o GetDeviceSwitchRoutingInterfaces200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceSwitchRoutingInterfaces200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InterfaceId) {
		toSerialize["interfaceId"] = o.InterfaceId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !IsNil(o.InterfaceIp) {
		toSerialize["interfaceIp"] = o.InterfaceIp
	}
	if !IsNil(o.MulticastRouting) {
		toSerialize["multicastRouting"] = o.MulticastRouting
	}
	if !IsNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	if !IsNil(o.DefaultGateway) {
		toSerialize["defaultGateway"] = o.DefaultGateway
	}
	if !IsNil(o.OspfSettings) {
		toSerialize["ospfSettings"] = o.OspfSettings
	}
	if !IsNil(o.OspfV3) {
		toSerialize["ospfV3"] = o.OspfV3
	}
	if !IsNil(o.Ipv6) {
		toSerialize["ipv6"] = o.Ipv6
	}
	return toSerialize, nil
}

type NullableGetDeviceSwitchRoutingInterfaces200ResponseInner struct {
	value *GetDeviceSwitchRoutingInterfaces200ResponseInner
	isSet bool
}

func (v NullableGetDeviceSwitchRoutingInterfaces200ResponseInner) Get() *GetDeviceSwitchRoutingInterfaces200ResponseInner {
	return v.value
}

func (v *NullableGetDeviceSwitchRoutingInterfaces200ResponseInner) Set(val *GetDeviceSwitchRoutingInterfaces200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceSwitchRoutingInterfaces200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceSwitchRoutingInterfaces200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceSwitchRoutingInterfaces200ResponseInner(val *GetDeviceSwitchRoutingInterfaces200ResponseInner) *NullableGetDeviceSwitchRoutingInterfaces200ResponseInner {
	return &NullableGetDeviceSwitchRoutingInterfaces200ResponseInner{value: val, isSet: true}
}

func (v NullableGetDeviceSwitchRoutingInterfaces200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceSwitchRoutingInterfaces200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


