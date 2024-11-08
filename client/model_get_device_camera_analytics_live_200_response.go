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

// checks if the GetDeviceCameraAnalyticsLive200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceCameraAnalyticsLive200Response{}

// GetDeviceCameraAnalyticsLive200Response struct for GetDeviceCameraAnalyticsLive200Response
type GetDeviceCameraAnalyticsLive200Response struct {
	// The current time
	Ts *string `json:"ts,omitempty"`
	Zones *GetDeviceCameraAnalyticsLive200ResponseZones `json:"zones,omitempty"`
}

// NewGetDeviceCameraAnalyticsLive200Response instantiates a new GetDeviceCameraAnalyticsLive200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceCameraAnalyticsLive200Response() *GetDeviceCameraAnalyticsLive200Response {
	this := GetDeviceCameraAnalyticsLive200Response{}
	return &this
}

// NewGetDeviceCameraAnalyticsLive200ResponseWithDefaults instantiates a new GetDeviceCameraAnalyticsLive200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceCameraAnalyticsLive200ResponseWithDefaults() *GetDeviceCameraAnalyticsLive200Response {
	this := GetDeviceCameraAnalyticsLive200Response{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *GetDeviceCameraAnalyticsLive200Response) GetTs() string {
	if o == nil || IsNil(o.Ts) {
		var ret string
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceCameraAnalyticsLive200Response) GetTsOk() (*string, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *GetDeviceCameraAnalyticsLive200Response) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given string and assigns it to the Ts field.
func (o *GetDeviceCameraAnalyticsLive200Response) SetTs(v string) {
	o.Ts = &v
}

// GetZones returns the Zones field value if set, zero value otherwise.
func (o *GetDeviceCameraAnalyticsLive200Response) GetZones() GetDeviceCameraAnalyticsLive200ResponseZones {
	if o == nil || IsNil(o.Zones) {
		var ret GetDeviceCameraAnalyticsLive200ResponseZones
		return ret
	}
	return *o.Zones
}

// GetZonesOk returns a tuple with the Zones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceCameraAnalyticsLive200Response) GetZonesOk() (*GetDeviceCameraAnalyticsLive200ResponseZones, bool) {
	if o == nil || IsNil(o.Zones) {
		return nil, false
	}
	return o.Zones, true
}

// HasZones returns a boolean if a field has been set.
func (o *GetDeviceCameraAnalyticsLive200Response) HasZones() bool {
	if o != nil && !IsNil(o.Zones) {
		return true
	}

	return false
}

// SetZones gets a reference to the given GetDeviceCameraAnalyticsLive200ResponseZones and assigns it to the Zones field.
func (o *GetDeviceCameraAnalyticsLive200Response) SetZones(v GetDeviceCameraAnalyticsLive200ResponseZones) {
	o.Zones = &v
}

func (o GetDeviceCameraAnalyticsLive200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceCameraAnalyticsLive200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.Zones) {
		toSerialize["zones"] = o.Zones
	}
	return toSerialize, nil
}

type NullableGetDeviceCameraAnalyticsLive200Response struct {
	value *GetDeviceCameraAnalyticsLive200Response
	isSet bool
}

func (v NullableGetDeviceCameraAnalyticsLive200Response) Get() *GetDeviceCameraAnalyticsLive200Response {
	return v.value
}

func (v *NullableGetDeviceCameraAnalyticsLive200Response) Set(val *GetDeviceCameraAnalyticsLive200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceCameraAnalyticsLive200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceCameraAnalyticsLive200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceCameraAnalyticsLive200Response(val *GetDeviceCameraAnalyticsLive200Response) *NullableGetDeviceCameraAnalyticsLive200Response {
	return &NullableGetDeviceCameraAnalyticsLive200Response{value: val, isSet: true}
}

func (v NullableGetDeviceCameraAnalyticsLive200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceCameraAnalyticsLive200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


