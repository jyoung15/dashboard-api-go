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

// checks if the GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC{}

// GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC An object containing current disk usage details.
type GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC struct {
	// The used disk space.
	Used *int32 `json:"used,omitempty"`
	// The available disk space.
	Space *int32 `json:"space,omitempty"`
}

// NewGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC instantiates a new GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC() *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC {
	this := GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC{}
	return &this
}

// NewGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageCWithDefaults instantiates a new GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageCWithDefaults() *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC {
	this := GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC{}
	return &this
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC) GetUsed() int32 {
	if o == nil || IsNil(o.Used) {
		var ret int32
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC) GetUsedOk() (*int32, bool) {
	if o == nil || IsNil(o.Used) {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC) HasUsed() bool {
	if o != nil && !IsNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given int32 and assigns it to the Used field.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC) SetUsed(v int32) {
	o.Used = &v
}

// GetSpace returns the Space field value if set, zero value otherwise.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC) GetSpace() int32 {
	if o == nil || IsNil(o.Space) {
		var ret int32
		return ret
	}
	return *o.Space
}

// GetSpaceOk returns a tuple with the Space field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC) GetSpaceOk() (*int32, bool) {
	if o == nil || IsNil(o.Space) {
		return nil, false
	}
	return o.Space, true
}

// HasSpace returns a boolean if a field has been set.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC) HasSpace() bool {
	if o != nil && !IsNil(o.Space) {
		return true
	}

	return false
}

// SetSpace gets a reference to the given int32 and assigns it to the Space field.
func (o *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC) SetSpace(v int32) {
	o.Space = &v
}

func (o GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Used) {
		toSerialize["used"] = o.Used
	}
	if !IsNil(o.Space) {
		toSerialize["space"] = o.Space
	}
	return toSerialize, nil
}

type NullableGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC struct {
	value *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC
	isSet bool
}

func (v NullableGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC) Get() *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC {
	return v.value
}

func (v *NullableGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC) Set(val *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC(val *GetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC) *NullableGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC {
	return &NullableGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC{value: val, isSet: true}
}

func (v NullableGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSmDevicePerformanceHistory200ResponseInnerDiskUsageC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


