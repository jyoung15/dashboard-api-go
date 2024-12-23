/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetDeviceLossAndLatencyHistory200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceLossAndLatencyHistory200ResponseInner{}

// GetDeviceLossAndLatencyHistory200ResponseInner struct for GetDeviceLossAndLatencyHistory200ResponseInner
type GetDeviceLossAndLatencyHistory200ResponseInner struct {
	// Start time of the sample
	StartTime *time.Time `json:"startTime,omitempty"`
	// End time of the sample
	EndTime *time.Time `json:"endTime,omitempty"`
	// Percentage of packets lost
	LossPercent *float32 `json:"lossPercent,omitempty"`
	// Latency in milliseconds
	LatencyMs *float32 `json:"latencyMs,omitempty"`
	// Number of useful information bits delivered
	Goodput *int32 `json:"goodput,omitempty"`
	// Jitter, in milliseconds
	Jitter *float32 `json:"jitter,omitempty"`
}

// NewGetDeviceLossAndLatencyHistory200ResponseInner instantiates a new GetDeviceLossAndLatencyHistory200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceLossAndLatencyHistory200ResponseInner() *GetDeviceLossAndLatencyHistory200ResponseInner {
	this := GetDeviceLossAndLatencyHistory200ResponseInner{}
	return &this
}

// NewGetDeviceLossAndLatencyHistory200ResponseInnerWithDefaults instantiates a new GetDeviceLossAndLatencyHistory200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceLossAndLatencyHistory200ResponseInnerWithDefaults() *GetDeviceLossAndLatencyHistory200ResponseInner {
	this := GetDeviceLossAndLatencyHistory200ResponseInner{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetLossPercent returns the LossPercent field value if set, zero value otherwise.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetLossPercent() float32 {
	if o == nil || IsNil(o.LossPercent) {
		var ret float32
		return ret
	}
	return *o.LossPercent
}

// GetLossPercentOk returns a tuple with the LossPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetLossPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.LossPercent) {
		return nil, false
	}
	return o.LossPercent, true
}

// HasLossPercent returns a boolean if a field has been set.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) HasLossPercent() bool {
	if o != nil && !IsNil(o.LossPercent) {
		return true
	}

	return false
}

// SetLossPercent gets a reference to the given float32 and assigns it to the LossPercent field.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) SetLossPercent(v float32) {
	o.LossPercent = &v
}

// GetLatencyMs returns the LatencyMs field value if set, zero value otherwise.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetLatencyMs() float32 {
	if o == nil || IsNil(o.LatencyMs) {
		var ret float32
		return ret
	}
	return *o.LatencyMs
}

// GetLatencyMsOk returns a tuple with the LatencyMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetLatencyMsOk() (*float32, bool) {
	if o == nil || IsNil(o.LatencyMs) {
		return nil, false
	}
	return o.LatencyMs, true
}

// HasLatencyMs returns a boolean if a field has been set.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) HasLatencyMs() bool {
	if o != nil && !IsNil(o.LatencyMs) {
		return true
	}

	return false
}

// SetLatencyMs gets a reference to the given float32 and assigns it to the LatencyMs field.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) SetLatencyMs(v float32) {
	o.LatencyMs = &v
}

// GetGoodput returns the Goodput field value if set, zero value otherwise.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetGoodput() int32 {
	if o == nil || IsNil(o.Goodput) {
		var ret int32
		return ret
	}
	return *o.Goodput
}

// GetGoodputOk returns a tuple with the Goodput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetGoodputOk() (*int32, bool) {
	if o == nil || IsNil(o.Goodput) {
		return nil, false
	}
	return o.Goodput, true
}

// HasGoodput returns a boolean if a field has been set.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) HasGoodput() bool {
	if o != nil && !IsNil(o.Goodput) {
		return true
	}

	return false
}

// SetGoodput gets a reference to the given int32 and assigns it to the Goodput field.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) SetGoodput(v int32) {
	o.Goodput = &v
}

// GetJitter returns the Jitter field value if set, zero value otherwise.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetJitter() float32 {
	if o == nil || IsNil(o.Jitter) {
		var ret float32
		return ret
	}
	return *o.Jitter
}

// GetJitterOk returns a tuple with the Jitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) GetJitterOk() (*float32, bool) {
	if o == nil || IsNil(o.Jitter) {
		return nil, false
	}
	return o.Jitter, true
}

// HasJitter returns a boolean if a field has been set.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) HasJitter() bool {
	if o != nil && !IsNil(o.Jitter) {
		return true
	}

	return false
}

// SetJitter gets a reference to the given float32 and assigns it to the Jitter field.
func (o *GetDeviceLossAndLatencyHistory200ResponseInner) SetJitter(v float32) {
	o.Jitter = &v
}

func (o GetDeviceLossAndLatencyHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceLossAndLatencyHistory200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.LossPercent) {
		toSerialize["lossPercent"] = o.LossPercent
	}
	if !IsNil(o.LatencyMs) {
		toSerialize["latencyMs"] = o.LatencyMs
	}
	if !IsNil(o.Goodput) {
		toSerialize["goodput"] = o.Goodput
	}
	if !IsNil(o.Jitter) {
		toSerialize["jitter"] = o.Jitter
	}
	return toSerialize, nil
}

type NullableGetDeviceLossAndLatencyHistory200ResponseInner struct {
	value *GetDeviceLossAndLatencyHistory200ResponseInner
	isSet bool
}

func (v NullableGetDeviceLossAndLatencyHistory200ResponseInner) Get() *GetDeviceLossAndLatencyHistory200ResponseInner {
	return v.value
}

func (v *NullableGetDeviceLossAndLatencyHistory200ResponseInner) Set(val *GetDeviceLossAndLatencyHistory200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceLossAndLatencyHistory200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceLossAndLatencyHistory200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceLossAndLatencyHistory200ResponseInner(val *GetDeviceLossAndLatencyHistory200ResponseInner) *NullableGetDeviceLossAndLatencyHistory200ResponseInner {
	return &NullableGetDeviceLossAndLatencyHistory200ResponseInner{value: val, isSet: true}
}

func (v NullableGetDeviceLossAndLatencyHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceLossAndLatencyHistory200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


