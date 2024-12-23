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

// checks if the GetNetworkApplianceSsids200ResponseInnerRadiusServersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceSsids200ResponseInnerRadiusServersInner{}

// GetNetworkApplianceSsids200ResponseInnerRadiusServersInner struct for GetNetworkApplianceSsids200ResponseInnerRadiusServersInner
type GetNetworkApplianceSsids200ResponseInnerRadiusServersInner struct {
	// The IP address of your RADIUS server.
	Host *string `json:"host,omitempty"`
	// The UDP port your RADIUS servers listens on for Access-requests.
	Port *int32 `json:"port,omitempty"`
}

// NewGetNetworkApplianceSsids200ResponseInnerRadiusServersInner instantiates a new GetNetworkApplianceSsids200ResponseInnerRadiusServersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceSsids200ResponseInnerRadiusServersInner() *GetNetworkApplianceSsids200ResponseInnerRadiusServersInner {
	this := GetNetworkApplianceSsids200ResponseInnerRadiusServersInner{}
	return &this
}

// NewGetNetworkApplianceSsids200ResponseInnerRadiusServersInnerWithDefaults instantiates a new GetNetworkApplianceSsids200ResponseInnerRadiusServersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceSsids200ResponseInnerRadiusServersInnerWithDefaults() *GetNetworkApplianceSsids200ResponseInnerRadiusServersInner {
	this := GetNetworkApplianceSsids200ResponseInnerRadiusServersInner{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *GetNetworkApplianceSsids200ResponseInnerRadiusServersInner) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSsids200ResponseInnerRadiusServersInner) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *GetNetworkApplianceSsids200ResponseInnerRadiusServersInner) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *GetNetworkApplianceSsids200ResponseInnerRadiusServersInner) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *GetNetworkApplianceSsids200ResponseInnerRadiusServersInner) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSsids200ResponseInnerRadiusServersInner) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *GetNetworkApplianceSsids200ResponseInnerRadiusServersInner) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *GetNetworkApplianceSsids200ResponseInnerRadiusServersInner) SetPort(v int32) {
	o.Port = &v
}

func (o GetNetworkApplianceSsids200ResponseInnerRadiusServersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceSsids200ResponseInnerRadiusServersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceSsids200ResponseInnerRadiusServersInner struct {
	value *GetNetworkApplianceSsids200ResponseInnerRadiusServersInner
	isSet bool
}

func (v NullableGetNetworkApplianceSsids200ResponseInnerRadiusServersInner) Get() *GetNetworkApplianceSsids200ResponseInnerRadiusServersInner {
	return v.value
}

func (v *NullableGetNetworkApplianceSsids200ResponseInnerRadiusServersInner) Set(val *GetNetworkApplianceSsids200ResponseInnerRadiusServersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceSsids200ResponseInnerRadiusServersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceSsids200ResponseInnerRadiusServersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceSsids200ResponseInnerRadiusServersInner(val *GetNetworkApplianceSsids200ResponseInnerRadiusServersInner) *NullableGetNetworkApplianceSsids200ResponseInnerRadiusServersInner {
	return &NullableGetNetworkApplianceSsids200ResponseInnerRadiusServersInner{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceSsids200ResponseInnerRadiusServersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceSsids200ResponseInnerRadiusServersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


