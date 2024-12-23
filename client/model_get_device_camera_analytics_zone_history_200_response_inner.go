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

// checks if the GetDeviceCameraAnalyticsZoneHistory200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceCameraAnalyticsZoneHistory200ResponseInner{}

// GetDeviceCameraAnalyticsZoneHistory200ResponseInner struct for GetDeviceCameraAnalyticsZoneHistory200ResponseInner
type GetDeviceCameraAnalyticsZoneHistory200ResponseInner struct {
	// The start time
	StartTs *time.Time `json:"startTs,omitempty"`
	// The end time
	EndTs *time.Time `json:"endTs,omitempty"`
	// The number of entrances
	Entrances *int32 `json:"entrances,omitempty"`
	// The average count
	AverageCount *float32 `json:"averageCount,omitempty"`
}

// NewGetDeviceCameraAnalyticsZoneHistory200ResponseInner instantiates a new GetDeviceCameraAnalyticsZoneHistory200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceCameraAnalyticsZoneHistory200ResponseInner() *GetDeviceCameraAnalyticsZoneHistory200ResponseInner {
	this := GetDeviceCameraAnalyticsZoneHistory200ResponseInner{}
	return &this
}

// NewGetDeviceCameraAnalyticsZoneHistory200ResponseInnerWithDefaults instantiates a new GetDeviceCameraAnalyticsZoneHistory200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceCameraAnalyticsZoneHistory200ResponseInnerWithDefaults() *GetDeviceCameraAnalyticsZoneHistory200ResponseInner {
	this := GetDeviceCameraAnalyticsZoneHistory200ResponseInner{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *GetDeviceCameraAnalyticsZoneHistory200ResponseInner) GetStartTs() time.Time {
	if o == nil || IsNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceCameraAnalyticsZoneHistory200ResponseInner) GetStartTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTs) {
		return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *GetDeviceCameraAnalyticsZoneHistory200ResponseInner) HasStartTs() bool {
	if o != nil && !IsNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *GetDeviceCameraAnalyticsZoneHistory200ResponseInner) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *GetDeviceCameraAnalyticsZoneHistory200ResponseInner) GetEndTs() time.Time {
	if o == nil || IsNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceCameraAnalyticsZoneHistory200ResponseInner) GetEndTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTs) {
		return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *GetDeviceCameraAnalyticsZoneHistory200ResponseInner) HasEndTs() bool {
	if o != nil && !IsNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *GetDeviceCameraAnalyticsZoneHistory200ResponseInner) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetEntrances returns the Entrances field value if set, zero value otherwise.
func (o *GetDeviceCameraAnalyticsZoneHistory200ResponseInner) GetEntrances() int32 {
	if o == nil || IsNil(o.Entrances) {
		var ret int32
		return ret
	}
	return *o.Entrances
}

// GetEntrancesOk returns a tuple with the Entrances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceCameraAnalyticsZoneHistory200ResponseInner) GetEntrancesOk() (*int32, bool) {
	if o == nil || IsNil(o.Entrances) {
		return nil, false
	}
	return o.Entrances, true
}

// HasEntrances returns a boolean if a field has been set.
func (o *GetDeviceCameraAnalyticsZoneHistory200ResponseInner) HasEntrances() bool {
	if o != nil && !IsNil(o.Entrances) {
		return true
	}

	return false
}

// SetEntrances gets a reference to the given int32 and assigns it to the Entrances field.
func (o *GetDeviceCameraAnalyticsZoneHistory200ResponseInner) SetEntrances(v int32) {
	o.Entrances = &v
}

// GetAverageCount returns the AverageCount field value if set, zero value otherwise.
func (o *GetDeviceCameraAnalyticsZoneHistory200ResponseInner) GetAverageCount() float32 {
	if o == nil || IsNil(o.AverageCount) {
		var ret float32
		return ret
	}
	return *o.AverageCount
}

// GetAverageCountOk returns a tuple with the AverageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceCameraAnalyticsZoneHistory200ResponseInner) GetAverageCountOk() (*float32, bool) {
	if o == nil || IsNil(o.AverageCount) {
		return nil, false
	}
	return o.AverageCount, true
}

// HasAverageCount returns a boolean if a field has been set.
func (o *GetDeviceCameraAnalyticsZoneHistory200ResponseInner) HasAverageCount() bool {
	if o != nil && !IsNil(o.AverageCount) {
		return true
	}

	return false
}

// SetAverageCount gets a reference to the given float32 and assigns it to the AverageCount field.
func (o *GetDeviceCameraAnalyticsZoneHistory200ResponseInner) SetAverageCount(v float32) {
	o.AverageCount = &v
}

func (o GetDeviceCameraAnalyticsZoneHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceCameraAnalyticsZoneHistory200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !IsNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if !IsNil(o.Entrances) {
		toSerialize["entrances"] = o.Entrances
	}
	if !IsNil(o.AverageCount) {
		toSerialize["averageCount"] = o.AverageCount
	}
	return toSerialize, nil
}

type NullableGetDeviceCameraAnalyticsZoneHistory200ResponseInner struct {
	value *GetDeviceCameraAnalyticsZoneHistory200ResponseInner
	isSet bool
}

func (v NullableGetDeviceCameraAnalyticsZoneHistory200ResponseInner) Get() *GetDeviceCameraAnalyticsZoneHistory200ResponseInner {
	return v.value
}

func (v *NullableGetDeviceCameraAnalyticsZoneHistory200ResponseInner) Set(val *GetDeviceCameraAnalyticsZoneHistory200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceCameraAnalyticsZoneHistory200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceCameraAnalyticsZoneHistory200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceCameraAnalyticsZoneHistory200ResponseInner(val *GetDeviceCameraAnalyticsZoneHistory200ResponseInner) *NullableGetDeviceCameraAnalyticsZoneHistory200ResponseInner {
	return &NullableGetDeviceCameraAnalyticsZoneHistory200ResponseInner{value: val, isSet: true}
}

func (v NullableGetDeviceCameraAnalyticsZoneHistory200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceCameraAnalyticsZoneHistory200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


