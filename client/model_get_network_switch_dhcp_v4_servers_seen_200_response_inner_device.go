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

// checks if the GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice{}

// GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice Attributes of the server when it's a device.
type GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice struct {
	// Device serial.
	Serial *string `json:"serial,omitempty"`
	// Device name.
	Name *string `json:"name,omitempty"`
	// Url link to device.
	Url *string `json:"url,omitempty"`
	Interface *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface `json:"interface,omitempty"`
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice{}
	return &this
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceWithDefaults instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceWithDefaults() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) SetSerial(v string) {
	o.Serial = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) SetUrl(v string) {
	o.Url = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) GetInterface() GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface {
	if o == nil || IsNil(o.Interface) {
		var ret GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) GetInterfaceOk() (*GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface, bool) {
	if o == nil || IsNil(o.Interface) {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) HasInterface() bool {
	if o != nil && !IsNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface and assigns it to the Interface field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) SetInterface(v GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDeviceInterface) {
	o.Interface = &v
}

func (o GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice struct {
	value *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice
	isSet bool
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) Get() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice {
	return v.value
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) Set(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice {
	return &NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


