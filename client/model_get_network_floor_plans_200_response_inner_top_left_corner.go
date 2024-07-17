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

// checks if the GetNetworkFloorPlans200ResponseInnerTopLeftCorner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkFloorPlans200ResponseInnerTopLeftCorner{}

// GetNetworkFloorPlans200ResponseInnerTopLeftCorner The longitude and latitude of the top left corner of your floor plan.
type GetNetworkFloorPlans200ResponseInnerTopLeftCorner struct {
	// Latitude
	Lat *float32 `json:"lat,omitempty"`
	// Longitude
	Lng *float32 `json:"lng,omitempty"`
}

// NewGetNetworkFloorPlans200ResponseInnerTopLeftCorner instantiates a new GetNetworkFloorPlans200ResponseInnerTopLeftCorner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFloorPlans200ResponseInnerTopLeftCorner() *GetNetworkFloorPlans200ResponseInnerTopLeftCorner {
	this := GetNetworkFloorPlans200ResponseInnerTopLeftCorner{}
	return &this
}

// NewGetNetworkFloorPlans200ResponseInnerTopLeftCornerWithDefaults instantiates a new GetNetworkFloorPlans200ResponseInnerTopLeftCorner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFloorPlans200ResponseInnerTopLeftCornerWithDefaults() *GetNetworkFloorPlans200ResponseInnerTopLeftCorner {
	this := GetNetworkFloorPlans200ResponseInnerTopLeftCorner{}
	return &this
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInnerTopLeftCorner) GetLat() float32 {
	if o == nil || IsNil(o.Lat) {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInnerTopLeftCorner) GetLatOk() (*float32, bool) {
	if o == nil || IsNil(o.Lat) {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInnerTopLeftCorner) HasLat() bool {
	if o != nil && !IsNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *GetNetworkFloorPlans200ResponseInnerTopLeftCorner) SetLat(v float32) {
	o.Lat = &v
}

// GetLng returns the Lng field value if set, zero value otherwise.
func (o *GetNetworkFloorPlans200ResponseInnerTopLeftCorner) GetLng() float32 {
	if o == nil || IsNil(o.Lng) {
		var ret float32
		return ret
	}
	return *o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFloorPlans200ResponseInnerTopLeftCorner) GetLngOk() (*float32, bool) {
	if o == nil || IsNil(o.Lng) {
		return nil, false
	}
	return o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *GetNetworkFloorPlans200ResponseInnerTopLeftCorner) HasLng() bool {
	if o != nil && !IsNil(o.Lng) {
		return true
	}

	return false
}

// SetLng gets a reference to the given float32 and assigns it to the Lng field.
func (o *GetNetworkFloorPlans200ResponseInnerTopLeftCorner) SetLng(v float32) {
	o.Lng = &v
}

func (o GetNetworkFloorPlans200ResponseInnerTopLeftCorner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkFloorPlans200ResponseInnerTopLeftCorner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lat) {
		toSerialize["lat"] = o.Lat
	}
	if !IsNil(o.Lng) {
		toSerialize["lng"] = o.Lng
	}
	return toSerialize, nil
}

type NullableGetNetworkFloorPlans200ResponseInnerTopLeftCorner struct {
	value *GetNetworkFloorPlans200ResponseInnerTopLeftCorner
	isSet bool
}

func (v NullableGetNetworkFloorPlans200ResponseInnerTopLeftCorner) Get() *GetNetworkFloorPlans200ResponseInnerTopLeftCorner {
	return v.value
}

func (v *NullableGetNetworkFloorPlans200ResponseInnerTopLeftCorner) Set(val *GetNetworkFloorPlans200ResponseInnerTopLeftCorner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFloorPlans200ResponseInnerTopLeftCorner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFloorPlans200ResponseInnerTopLeftCorner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFloorPlans200ResponseInnerTopLeftCorner(val *GetNetworkFloorPlans200ResponseInnerTopLeftCorner) *NullableGetNetworkFloorPlans200ResponseInnerTopLeftCorner {
	return &NullableGetNetworkFloorPlans200ResponseInnerTopLeftCorner{value: val, isSet: true}
}

func (v NullableGetNetworkFloorPlans200ResponseInnerTopLeftCorner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFloorPlans200ResponseInnerTopLeftCorner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


