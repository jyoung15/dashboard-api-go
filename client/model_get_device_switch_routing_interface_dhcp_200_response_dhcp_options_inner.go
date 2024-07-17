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

// checks if the GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner{}

// GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner struct for GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner
type GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner struct {
	// The code for DHCP option which should be from 2 to 254
	Code *string `json:"code,omitempty"`
	// The type of the DHCP option which should be one of ('text', 'ip', 'integer' or 'hex')
	Type *string `json:"type,omitempty"`
	// The value of the DHCP option
	Value *string `json:"value,omitempty"`
}

// NewGetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner instantiates a new GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner() *GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner {
	this := GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner{}
	return &this
}

// NewGetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInnerWithDefaults instantiates a new GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInnerWithDefaults() *GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner {
	this := GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner) SetCode(v string) {
	o.Code = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner) SetValue(v string) {
	o.Value = &v
}

func (o GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableGetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner struct {
	value *GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner
	isSet bool
}

func (v NullableGetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner) Get() *GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner {
	return v.value
}

func (v *NullableGetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner) Set(val *GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner(val *GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner) *NullableGetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner {
	return &NullableGetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner{value: val, isSet: true}
}

func (v NullableGetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


