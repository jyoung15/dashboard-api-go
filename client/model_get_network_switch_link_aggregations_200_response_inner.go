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

// checks if the GetNetworkSwitchLinkAggregations200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchLinkAggregations200ResponseInner{}

// GetNetworkSwitchLinkAggregations200ResponseInner struct for GetNetworkSwitchLinkAggregations200ResponseInner
type GetNetworkSwitchLinkAggregations200ResponseInner struct {
	// The ID for the link aggregation.
	Id *string `json:"id,omitempty"`
	// The ID for the link aggregation.
	SwitchPorts []GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner `json:"switchPorts,omitempty"`
}

// NewGetNetworkSwitchLinkAggregations200ResponseInner instantiates a new GetNetworkSwitchLinkAggregations200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchLinkAggregations200ResponseInner() *GetNetworkSwitchLinkAggregations200ResponseInner {
	this := GetNetworkSwitchLinkAggregations200ResponseInner{}
	return &this
}

// NewGetNetworkSwitchLinkAggregations200ResponseInnerWithDefaults instantiates a new GetNetworkSwitchLinkAggregations200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchLinkAggregations200ResponseInnerWithDefaults() *GetNetworkSwitchLinkAggregations200ResponseInner {
	this := GetNetworkSwitchLinkAggregations200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkSwitchLinkAggregations200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchLinkAggregations200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkSwitchLinkAggregations200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkSwitchLinkAggregations200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetSwitchPorts returns the SwitchPorts field value if set, zero value otherwise.
func (o *GetNetworkSwitchLinkAggregations200ResponseInner) GetSwitchPorts() []GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner {
	if o == nil || IsNil(o.SwitchPorts) {
		var ret []GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner
		return ret
	}
	return o.SwitchPorts
}

// GetSwitchPortsOk returns a tuple with the SwitchPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchLinkAggregations200ResponseInner) GetSwitchPortsOk() ([]GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner, bool) {
	if o == nil || IsNil(o.SwitchPorts) {
		return nil, false
	}
	return o.SwitchPorts, true
}

// HasSwitchPorts returns a boolean if a field has been set.
func (o *GetNetworkSwitchLinkAggregations200ResponseInner) HasSwitchPorts() bool {
	if o != nil && !IsNil(o.SwitchPorts) {
		return true
	}

	return false
}

// SetSwitchPorts gets a reference to the given []GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner and assigns it to the SwitchPorts field.
func (o *GetNetworkSwitchLinkAggregations200ResponseInner) SetSwitchPorts(v []GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner) {
	o.SwitchPorts = v
}

func (o GetNetworkSwitchLinkAggregations200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchLinkAggregations200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SwitchPorts) {
		toSerialize["switchPorts"] = o.SwitchPorts
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchLinkAggregations200ResponseInner struct {
	value *GetNetworkSwitchLinkAggregations200ResponseInner
	isSet bool
}

func (v NullableGetNetworkSwitchLinkAggregations200ResponseInner) Get() *GetNetworkSwitchLinkAggregations200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkSwitchLinkAggregations200ResponseInner) Set(val *GetNetworkSwitchLinkAggregations200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchLinkAggregations200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchLinkAggregations200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchLinkAggregations200ResponseInner(val *GetNetworkSwitchLinkAggregations200ResponseInner) *NullableGetNetworkSwitchLinkAggregations200ResponseInner {
	return &NullableGetNetworkSwitchLinkAggregations200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchLinkAggregations200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchLinkAggregations200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


