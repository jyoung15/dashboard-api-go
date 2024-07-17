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

// checks if the GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner{}

// GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner struct for GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner
type GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner struct {
	// The start time of the channel utilization interval.
	StartTime *time.Time `json:"startTime,omitempty"`
	// The end time of the channel utilization interval.
	EndTime *time.Time `json:"endTime,omitempty"`
	// Percentage of total channel utiliation for the given radio.
	UtilizationTotal *float32 `json:"utilizationTotal,omitempty"`
	// Percentage of wifi channel utiliation for the given radio.
	Utilization80211 *float32 `json:"utilization80211,omitempty"`
	// Percentage of non-wifi channel utiliation for the given radio.
	UtilizationNon80211 *float32 `json:"utilizationNon80211,omitempty"`
}

// NewGetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner instantiates a new GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner() *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner {
	this := GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner{}
	return &this
}

// NewGetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0InnerWithDefaults instantiates a new GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0InnerWithDefaults() *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner {
	this := GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetUtilizationTotal returns the UtilizationTotal field value if set, zero value otherwise.
func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) GetUtilizationTotal() float32 {
	if o == nil || IsNil(o.UtilizationTotal) {
		var ret float32
		return ret
	}
	return *o.UtilizationTotal
}

// GetUtilizationTotalOk returns a tuple with the UtilizationTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) GetUtilizationTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.UtilizationTotal) {
		return nil, false
	}
	return o.UtilizationTotal, true
}

// HasUtilizationTotal returns a boolean if a field has been set.
func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) HasUtilizationTotal() bool {
	if o != nil && !IsNil(o.UtilizationTotal) {
		return true
	}

	return false
}

// SetUtilizationTotal gets a reference to the given float32 and assigns it to the UtilizationTotal field.
func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) SetUtilizationTotal(v float32) {
	o.UtilizationTotal = &v
}

// GetUtilization80211 returns the Utilization80211 field value if set, zero value otherwise.
func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) GetUtilization80211() float32 {
	if o == nil || IsNil(o.Utilization80211) {
		var ret float32
		return ret
	}
	return *o.Utilization80211
}

// GetUtilization80211Ok returns a tuple with the Utilization80211 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) GetUtilization80211Ok() (*float32, bool) {
	if o == nil || IsNil(o.Utilization80211) {
		return nil, false
	}
	return o.Utilization80211, true
}

// HasUtilization80211 returns a boolean if a field has been set.
func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) HasUtilization80211() bool {
	if o != nil && !IsNil(o.Utilization80211) {
		return true
	}

	return false
}

// SetUtilization80211 gets a reference to the given float32 and assigns it to the Utilization80211 field.
func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) SetUtilization80211(v float32) {
	o.Utilization80211 = &v
}

// GetUtilizationNon80211 returns the UtilizationNon80211 field value if set, zero value otherwise.
func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) GetUtilizationNon80211() float32 {
	if o == nil || IsNil(o.UtilizationNon80211) {
		var ret float32
		return ret
	}
	return *o.UtilizationNon80211
}

// GetUtilizationNon80211Ok returns a tuple with the UtilizationNon80211 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) GetUtilizationNon80211Ok() (*float32, bool) {
	if o == nil || IsNil(o.UtilizationNon80211) {
		return nil, false
	}
	return o.UtilizationNon80211, true
}

// HasUtilizationNon80211 returns a boolean if a field has been set.
func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) HasUtilizationNon80211() bool {
	if o != nil && !IsNil(o.UtilizationNon80211) {
		return true
	}

	return false
}

// SetUtilizationNon80211 gets a reference to the given float32 and assigns it to the UtilizationNon80211 field.
func (o *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) SetUtilizationNon80211(v float32) {
	o.UtilizationNon80211 = &v
}

func (o GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.UtilizationTotal) {
		toSerialize["utilizationTotal"] = o.UtilizationTotal
	}
	if !IsNil(o.Utilization80211) {
		toSerialize["utilization80211"] = o.Utilization80211
	}
	if !IsNil(o.UtilizationNon80211) {
		toSerialize["utilizationNon80211"] = o.UtilizationNon80211
	}
	return toSerialize, nil
}

type NullableGetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner struct {
	value *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner
	isSet bool
}

func (v NullableGetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) Get() *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner {
	return v.value
}

func (v *NullableGetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) Set(val *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner(val *GetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) *NullableGetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner {
	return &NullableGetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner{value: val, isSet: true}
}

func (v NullableGetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkNetworkHealthChannelUtilization200ResponseInnerWifi0Inner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

