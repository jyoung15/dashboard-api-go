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

// checks if the GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner{}

// GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner struct for GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner
type GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner struct {
	// The serial number for the switch port.
	Serial *string `json:"serial,omitempty"`
	// The ID for the switch port.
	PortId *string `json:"portId,omitempty"`
}

// NewGetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner instantiates a new GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner() *GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner {
	this := GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner{}
	return &this
}

// NewGetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInnerWithDefaults instantiates a new GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInnerWithDefaults() *GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner {
	this := GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner) SetSerial(v string) {
	o.Serial = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner) GetPortId() string {
	if o == nil || IsNil(o.PortId) {
		var ret string
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner) GetPortIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortId) {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner) HasPortId() bool {
	if o != nil && !IsNil(o.PortId) {
		return true
	}

	return false
}

// SetPortId gets a reference to the given string and assigns it to the PortId field.
func (o *GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner) SetPortId(v string) {
	o.PortId = &v
}

func (o GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.PortId) {
		toSerialize["portId"] = o.PortId
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner struct {
	value *GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner
	isSet bool
}

func (v NullableGetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner) Get() *GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner {
	return v.value
}

func (v *NullableGetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner) Set(val *GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner(val *GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner) *NullableGetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner {
	return &NullableGetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

