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

// checks if the GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface{}

// GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface Interface attributes of the server. Only for configured servers.
type GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface struct {
	// Interface name.
	Name *string `json:"name,omitempty"`
	// Url link to interface.
	Url *string `json:"url,omitempty"`
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface{}
	return &this
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterfaceWithDefaults instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterfaceWithDefaults() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface) SetUrl(v string) {
	o.Url = &v
}

func (o GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface struct {
	value *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface
	isSet bool
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface) Get() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface {
	return v.value
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface) Set(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface) *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface {
	return &NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


