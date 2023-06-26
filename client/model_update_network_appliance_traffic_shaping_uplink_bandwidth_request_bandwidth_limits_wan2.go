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

// UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2 The bandwidth settings for the 'wan2' uplink
type UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2 struct {
	// The maximum upload limit (integer, in Kbps). null indicates no limit
	LimitUp *int32 `json:"limitUp,omitempty"`
	// The maximum download limit (integer, in Kbps). null indicates no limit
	LimitDown *int32 `json:"limitDown,omitempty"`
}

// NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2 instantiates a new UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2() *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2 {
	this := UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2{}
	return &this
}

// NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2WithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2WithDefaults() *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2 {
	this := UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2{}
	return &this
}

// GetLimitUp returns the LimitUp field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2) GetLimitUp() int32 {
	if o == nil || isNil(o.LimitUp) {
		var ret int32
		return ret
	}
	return *o.LimitUp
}

// GetLimitUpOk returns a tuple with the LimitUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2) GetLimitUpOk() (*int32, bool) {
	if o == nil || isNil(o.LimitUp) {
    return nil, false
	}
	return o.LimitUp, true
}

// HasLimitUp returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2) HasLimitUp() bool {
	if o != nil && !isNil(o.LimitUp) {
		return true
	}

	return false
}

// SetLimitUp gets a reference to the given int32 and assigns it to the LimitUp field.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2) SetLimitUp(v int32) {
	o.LimitUp = &v
}

// GetLimitDown returns the LimitDown field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2) GetLimitDown() int32 {
	if o == nil || isNil(o.LimitDown) {
		var ret int32
		return ret
	}
	return *o.LimitDown
}

// GetLimitDownOk returns a tuple with the LimitDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2) GetLimitDownOk() (*int32, bool) {
	if o == nil || isNil(o.LimitDown) {
    return nil, false
	}
	return o.LimitDown, true
}

// HasLimitDown returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2) HasLimitDown() bool {
	if o != nil && !isNil(o.LimitDown) {
		return true
	}

	return false
}

// SetLimitDown gets a reference to the given int32 and assigns it to the LimitDown field.
func (o *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2) SetLimitDown(v int32) {
	o.LimitDown = &v
}

func (o UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LimitUp) {
		toSerialize["limitUp"] = o.LimitUp
	}
	if !isNil(o.LimitDown) {
		toSerialize["limitDown"] = o.LimitDown
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2 struct {
	value *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2
	isSet bool
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2) Get() *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2 {
	return v.value
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2) Set(val *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2(val *UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2) *NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2 {
	return &NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequestBandwidthLimitsWan2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


