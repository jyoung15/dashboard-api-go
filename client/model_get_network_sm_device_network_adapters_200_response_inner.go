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

// GetNetworkSmDeviceNetworkAdapters200ResponseInner struct for GetNetworkSmDeviceNetworkAdapters200ResponseInner
type GetNetworkSmDeviceNetworkAdapters200ResponseInner struct {
	// The IP address of the DCHP Server.
	DhcpServer *string `json:"dhcpServer,omitempty"`
	// The IP address of the DNS Server.
	DnsServer *string `json:"dnsServer,omitempty"`
	// The IP address of the Gateway.
	Gateway *string `json:"gateway,omitempty"`
	// The Meraki Id of the network adapter record.
	Id *string `json:"id,omitempty"`
	// The IP address of the network adapter.
	Ip *string `json:"ip,omitempty"`
	// The MAC associated with the network adapter.
	Mac *string `json:"mac,omitempty"`
	// The name of the newtwork adapter.
	Name *string `json:"name,omitempty"`
	// The subnet for the network adapter.
	Subnet *string `json:"subnet,omitempty"`
}

// NewGetNetworkSmDeviceNetworkAdapters200ResponseInner instantiates a new GetNetworkSmDeviceNetworkAdapters200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSmDeviceNetworkAdapters200ResponseInner() *GetNetworkSmDeviceNetworkAdapters200ResponseInner {
	this := GetNetworkSmDeviceNetworkAdapters200ResponseInner{}
	return &this
}

// NewGetNetworkSmDeviceNetworkAdapters200ResponseInnerWithDefaults instantiates a new GetNetworkSmDeviceNetworkAdapters200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSmDeviceNetworkAdapters200ResponseInnerWithDefaults() *GetNetworkSmDeviceNetworkAdapters200ResponseInner {
	this := GetNetworkSmDeviceNetworkAdapters200ResponseInner{}
	return &this
}

// GetDhcpServer returns the DhcpServer field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) GetDhcpServer() string {
	if o == nil || isNil(o.DhcpServer) {
		var ret string
		return ret
	}
	return *o.DhcpServer
}

// GetDhcpServerOk returns a tuple with the DhcpServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) GetDhcpServerOk() (*string, bool) {
	if o == nil || isNil(o.DhcpServer) {
    return nil, false
	}
	return o.DhcpServer, true
}

// HasDhcpServer returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) HasDhcpServer() bool {
	if o != nil && !isNil(o.DhcpServer) {
		return true
	}

	return false
}

// SetDhcpServer gets a reference to the given string and assigns it to the DhcpServer field.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) SetDhcpServer(v string) {
	o.DhcpServer = &v
}

// GetDnsServer returns the DnsServer field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) GetDnsServer() string {
	if o == nil || isNil(o.DnsServer) {
		var ret string
		return ret
	}
	return *o.DnsServer
}

// GetDnsServerOk returns a tuple with the DnsServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) GetDnsServerOk() (*string, bool) {
	if o == nil || isNil(o.DnsServer) {
    return nil, false
	}
	return o.DnsServer, true
}

// HasDnsServer returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) HasDnsServer() bool {
	if o != nil && !isNil(o.DnsServer) {
		return true
	}

	return false
}

// SetDnsServer gets a reference to the given string and assigns it to the DnsServer field.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) SetDnsServer(v string) {
	o.DnsServer = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) GetGateway() string {
	if o == nil || isNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) GetGatewayOk() (*string, bool) {
	if o == nil || isNil(o.Gateway) {
    return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) HasGateway() bool {
	if o != nil && !isNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) SetGateway(v string) {
	o.Gateway = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) SetIp(v string) {
	o.Ip = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) SetMac(v string) {
	o.Mac = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) GetSubnet() string {
	if o == nil || isNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) GetSubnetOk() (*string, bool) {
	if o == nil || isNil(o.Subnet) {
    return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) HasSubnet() bool {
	if o != nil && !isNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *GetNetworkSmDeviceNetworkAdapters200ResponseInner) SetSubnet(v string) {
	o.Subnet = &v
}

func (o GetNetworkSmDeviceNetworkAdapters200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DhcpServer) {
		toSerialize["dhcpServer"] = o.DhcpServer
	}
	if !isNil(o.DnsServer) {
		toSerialize["dnsServer"] = o.DnsServer
	}
	if !isNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkSmDeviceNetworkAdapters200ResponseInner struct {
	value *GetNetworkSmDeviceNetworkAdapters200ResponseInner
	isSet bool
}

func (v NullableGetNetworkSmDeviceNetworkAdapters200ResponseInner) Get() *GetNetworkSmDeviceNetworkAdapters200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkSmDeviceNetworkAdapters200ResponseInner) Set(val *GetNetworkSmDeviceNetworkAdapters200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSmDeviceNetworkAdapters200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSmDeviceNetworkAdapters200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSmDeviceNetworkAdapters200ResponseInner(val *GetNetworkSmDeviceNetworkAdapters200ResponseInner) *NullableGetNetworkSmDeviceNetworkAdapters200ResponseInner {
	return &NullableGetNetworkSmDeviceNetworkAdapters200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkSmDeviceNetworkAdapters200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSmDeviceNetworkAdapters200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


