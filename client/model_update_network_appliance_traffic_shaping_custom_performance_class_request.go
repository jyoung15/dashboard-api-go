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

// checks if the UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest{}

// UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct for UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest
type UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct {
	// Name of the custom performance class
	Name *string `json:"name,omitempty"`
	// Maximum latency in milliseconds
	MaxLatency *int32 `json:"maxLatency,omitempty"`
	// Maximum jitter in milliseconds
	MaxJitter *int32 `json:"maxJitter,omitempty"`
	// Maximum percentage of packet loss
	MaxLossPercentage *int32 `json:"maxLossPercentage,omitempty"`
}

// NewUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest instantiates a new UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest() *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest {
	this := UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest{}
	return &this
}

// NewUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequestWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequestWithDefaults() *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest {
	this := UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) SetName(v string) {
	o.Name = &v
}

// GetMaxLatency returns the MaxLatency field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetMaxLatency() int32 {
	if o == nil || IsNil(o.MaxLatency) {
		var ret int32
		return ret
	}
	return *o.MaxLatency
}

// GetMaxLatencyOk returns a tuple with the MaxLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetMaxLatencyOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxLatency) {
		return nil, false
	}
	return o.MaxLatency, true
}

// HasMaxLatency returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) HasMaxLatency() bool {
	if o != nil && !IsNil(o.MaxLatency) {
		return true
	}

	return false
}

// SetMaxLatency gets a reference to the given int32 and assigns it to the MaxLatency field.
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) SetMaxLatency(v int32) {
	o.MaxLatency = &v
}

// GetMaxJitter returns the MaxJitter field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetMaxJitter() int32 {
	if o == nil || IsNil(o.MaxJitter) {
		var ret int32
		return ret
	}
	return *o.MaxJitter
}

// GetMaxJitterOk returns a tuple with the MaxJitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetMaxJitterOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxJitter) {
		return nil, false
	}
	return o.MaxJitter, true
}

// HasMaxJitter returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) HasMaxJitter() bool {
	if o != nil && !IsNil(o.MaxJitter) {
		return true
	}

	return false
}

// SetMaxJitter gets a reference to the given int32 and assigns it to the MaxJitter field.
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) SetMaxJitter(v int32) {
	o.MaxJitter = &v
}

// GetMaxLossPercentage returns the MaxLossPercentage field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetMaxLossPercentage() int32 {
	if o == nil || IsNil(o.MaxLossPercentage) {
		var ret int32
		return ret
	}
	return *o.MaxLossPercentage
}

// GetMaxLossPercentageOk returns a tuple with the MaxLossPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) GetMaxLossPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxLossPercentage) {
		return nil, false
	}
	return o.MaxLossPercentage, true
}

// HasMaxLossPercentage returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) HasMaxLossPercentage() bool {
	if o != nil && !IsNil(o.MaxLossPercentage) {
		return true
	}

	return false
}

// SetMaxLossPercentage gets a reference to the given int32 and assigns it to the MaxLossPercentage field.
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) SetMaxLossPercentage(v int32) {
	o.MaxLossPercentage = &v
}

func (o UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.MaxLatency) {
		toSerialize["maxLatency"] = o.MaxLatency
	}
	if !IsNil(o.MaxJitter) {
		toSerialize["maxJitter"] = o.MaxJitter
	}
	if !IsNil(o.MaxLossPercentage) {
		toSerialize["maxLossPercentage"] = o.MaxLossPercentage
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct {
	value *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest
	isSet bool
}

func (v NullableUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) Get() *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest {
	return v.value
}

func (v *NullableUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) Set(val *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest(val *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) *NullableUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest {
	return &NullableUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


