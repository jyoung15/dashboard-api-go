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

// checks if the GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest{}

// GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest The region of interest
type GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest struct {
	// The x0 coordinate
	X0 *string `json:"x0,omitempty"`
	// The y0 coordinate
	Y0 *string `json:"y0,omitempty"`
	// The x1 coordinate
	X1 *string `json:"x1,omitempty"`
	// The y1 coordinate
	Y1 *string `json:"y1,omitempty"`
}

// NewGetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest instantiates a new GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest() *GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest {
	this := GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest{}
	return &this
}

// NewGetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterestWithDefaults instantiates a new GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterestWithDefaults() *GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest {
	this := GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest{}
	return &this
}

// GetX0 returns the X0 field value if set, zero value otherwise.
func (o *GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) GetX0() string {
	if o == nil || IsNil(o.X0) {
		var ret string
		return ret
	}
	return *o.X0
}

// GetX0Ok returns a tuple with the X0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) GetX0Ok() (*string, bool) {
	if o == nil || IsNil(o.X0) {
		return nil, false
	}
	return o.X0, true
}

// HasX0 returns a boolean if a field has been set.
func (o *GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) HasX0() bool {
	if o != nil && !IsNil(o.X0) {
		return true
	}

	return false
}

// SetX0 gets a reference to the given string and assigns it to the X0 field.
func (o *GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) SetX0(v string) {
	o.X0 = &v
}

// GetY0 returns the Y0 field value if set, zero value otherwise.
func (o *GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) GetY0() string {
	if o == nil || IsNil(o.Y0) {
		var ret string
		return ret
	}
	return *o.Y0
}

// GetY0Ok returns a tuple with the Y0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) GetY0Ok() (*string, bool) {
	if o == nil || IsNil(o.Y0) {
		return nil, false
	}
	return o.Y0, true
}

// HasY0 returns a boolean if a field has been set.
func (o *GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) HasY0() bool {
	if o != nil && !IsNil(o.Y0) {
		return true
	}

	return false
}

// SetY0 gets a reference to the given string and assigns it to the Y0 field.
func (o *GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) SetY0(v string) {
	o.Y0 = &v
}

// GetX1 returns the X1 field value if set, zero value otherwise.
func (o *GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) GetX1() string {
	if o == nil || IsNil(o.X1) {
		var ret string
		return ret
	}
	return *o.X1
}

// GetX1Ok returns a tuple with the X1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) GetX1Ok() (*string, bool) {
	if o == nil || IsNil(o.X1) {
		return nil, false
	}
	return o.X1, true
}

// HasX1 returns a boolean if a field has been set.
func (o *GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) HasX1() bool {
	if o != nil && !IsNil(o.X1) {
		return true
	}

	return false
}

// SetX1 gets a reference to the given string and assigns it to the X1 field.
func (o *GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) SetX1(v string) {
	o.X1 = &v
}

// GetY1 returns the Y1 field value if set, zero value otherwise.
func (o *GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) GetY1() string {
	if o == nil || IsNil(o.Y1) {
		var ret string
		return ret
	}
	return *o.Y1
}

// GetY1Ok returns a tuple with the Y1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) GetY1Ok() (*string, bool) {
	if o == nil || IsNil(o.Y1) {
		return nil, false
	}
	return o.Y1, true
}

// HasY1 returns a boolean if a field has been set.
func (o *GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) HasY1() bool {
	if o != nil && !IsNil(o.Y1) {
		return true
	}

	return false
}

// SetY1 gets a reference to the given string and assigns it to the Y1 field.
func (o *GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) SetY1(v string) {
	o.Y1 = &v
}

func (o GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.X0) {
		toSerialize["x0"] = o.X0
	}
	if !IsNil(o.Y0) {
		toSerialize["y0"] = o.Y0
	}
	if !IsNil(o.X1) {
		toSerialize["x1"] = o.X1
	}
	if !IsNil(o.Y1) {
		toSerialize["y1"] = o.Y1
	}
	return toSerialize, nil
}

type NullableGetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest struct {
	value *GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest
	isSet bool
}

func (v NullableGetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) Get() *GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest {
	return v.value
}

func (v *NullableGetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) Set(val *GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest(val *GetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) *NullableGetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest {
	return &NullableGetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest{value: val, isSet: true}
}

func (v NullableGetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceCameraAnalyticsZones200ResponseInnerRegionOfInterest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


