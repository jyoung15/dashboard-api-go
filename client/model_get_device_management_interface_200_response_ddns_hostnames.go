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

// checks if the GetDeviceManagementInterface200ResponseDdnsHostnames type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceManagementInterface200ResponseDdnsHostnames{}

// GetDeviceManagementInterface200ResponseDdnsHostnames Dynamic DNS hostnames.
type GetDeviceManagementInterface200ResponseDdnsHostnames struct {
	// Active dynamic DNS hostname.
	ActiveDdnsHostname *string `json:"activeDdnsHostname,omitempty"`
	// WAN 1 dynamic DNS hostname.
	DdnsHostnameWan1 *string `json:"ddnsHostnameWan1,omitempty"`
	// WAN 2 dynamic DNS hostname.
	DdnsHostnameWan2 *string `json:"ddnsHostnameWan2,omitempty"`
}

// NewGetDeviceManagementInterface200ResponseDdnsHostnames instantiates a new GetDeviceManagementInterface200ResponseDdnsHostnames object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceManagementInterface200ResponseDdnsHostnames() *GetDeviceManagementInterface200ResponseDdnsHostnames {
	this := GetDeviceManagementInterface200ResponseDdnsHostnames{}
	return &this
}

// NewGetDeviceManagementInterface200ResponseDdnsHostnamesWithDefaults instantiates a new GetDeviceManagementInterface200ResponseDdnsHostnames object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceManagementInterface200ResponseDdnsHostnamesWithDefaults() *GetDeviceManagementInterface200ResponseDdnsHostnames {
	this := GetDeviceManagementInterface200ResponseDdnsHostnames{}
	return &this
}

// GetActiveDdnsHostname returns the ActiveDdnsHostname field value if set, zero value otherwise.
func (o *GetDeviceManagementInterface200ResponseDdnsHostnames) GetActiveDdnsHostname() string {
	if o == nil || IsNil(o.ActiveDdnsHostname) {
		var ret string
		return ret
	}
	return *o.ActiveDdnsHostname
}

// GetActiveDdnsHostnameOk returns a tuple with the ActiveDdnsHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceManagementInterface200ResponseDdnsHostnames) GetActiveDdnsHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.ActiveDdnsHostname) {
		return nil, false
	}
	return o.ActiveDdnsHostname, true
}

// HasActiveDdnsHostname returns a boolean if a field has been set.
func (o *GetDeviceManagementInterface200ResponseDdnsHostnames) HasActiveDdnsHostname() bool {
	if o != nil && !IsNil(o.ActiveDdnsHostname) {
		return true
	}

	return false
}

// SetActiveDdnsHostname gets a reference to the given string and assigns it to the ActiveDdnsHostname field.
func (o *GetDeviceManagementInterface200ResponseDdnsHostnames) SetActiveDdnsHostname(v string) {
	o.ActiveDdnsHostname = &v
}

// GetDdnsHostnameWan1 returns the DdnsHostnameWan1 field value if set, zero value otherwise.
func (o *GetDeviceManagementInterface200ResponseDdnsHostnames) GetDdnsHostnameWan1() string {
	if o == nil || IsNil(o.DdnsHostnameWan1) {
		var ret string
		return ret
	}
	return *o.DdnsHostnameWan1
}

// GetDdnsHostnameWan1Ok returns a tuple with the DdnsHostnameWan1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceManagementInterface200ResponseDdnsHostnames) GetDdnsHostnameWan1Ok() (*string, bool) {
	if o == nil || IsNil(o.DdnsHostnameWan1) {
		return nil, false
	}
	return o.DdnsHostnameWan1, true
}

// HasDdnsHostnameWan1 returns a boolean if a field has been set.
func (o *GetDeviceManagementInterface200ResponseDdnsHostnames) HasDdnsHostnameWan1() bool {
	if o != nil && !IsNil(o.DdnsHostnameWan1) {
		return true
	}

	return false
}

// SetDdnsHostnameWan1 gets a reference to the given string and assigns it to the DdnsHostnameWan1 field.
func (o *GetDeviceManagementInterface200ResponseDdnsHostnames) SetDdnsHostnameWan1(v string) {
	o.DdnsHostnameWan1 = &v
}

// GetDdnsHostnameWan2 returns the DdnsHostnameWan2 field value if set, zero value otherwise.
func (o *GetDeviceManagementInterface200ResponseDdnsHostnames) GetDdnsHostnameWan2() string {
	if o == nil || IsNil(o.DdnsHostnameWan2) {
		var ret string
		return ret
	}
	return *o.DdnsHostnameWan2
}

// GetDdnsHostnameWan2Ok returns a tuple with the DdnsHostnameWan2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceManagementInterface200ResponseDdnsHostnames) GetDdnsHostnameWan2Ok() (*string, bool) {
	if o == nil || IsNil(o.DdnsHostnameWan2) {
		return nil, false
	}
	return o.DdnsHostnameWan2, true
}

// HasDdnsHostnameWan2 returns a boolean if a field has been set.
func (o *GetDeviceManagementInterface200ResponseDdnsHostnames) HasDdnsHostnameWan2() bool {
	if o != nil && !IsNil(o.DdnsHostnameWan2) {
		return true
	}

	return false
}

// SetDdnsHostnameWan2 gets a reference to the given string and assigns it to the DdnsHostnameWan2 field.
func (o *GetDeviceManagementInterface200ResponseDdnsHostnames) SetDdnsHostnameWan2(v string) {
	o.DdnsHostnameWan2 = &v
}

func (o GetDeviceManagementInterface200ResponseDdnsHostnames) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceManagementInterface200ResponseDdnsHostnames) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActiveDdnsHostname) {
		toSerialize["activeDdnsHostname"] = o.ActiveDdnsHostname
	}
	if !IsNil(o.DdnsHostnameWan1) {
		toSerialize["ddnsHostnameWan1"] = o.DdnsHostnameWan1
	}
	if !IsNil(o.DdnsHostnameWan2) {
		toSerialize["ddnsHostnameWan2"] = o.DdnsHostnameWan2
	}
	return toSerialize, nil
}

type NullableGetDeviceManagementInterface200ResponseDdnsHostnames struct {
	value *GetDeviceManagementInterface200ResponseDdnsHostnames
	isSet bool
}

func (v NullableGetDeviceManagementInterface200ResponseDdnsHostnames) Get() *GetDeviceManagementInterface200ResponseDdnsHostnames {
	return v.value
}

func (v *NullableGetDeviceManagementInterface200ResponseDdnsHostnames) Set(val *GetDeviceManagementInterface200ResponseDdnsHostnames) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceManagementInterface200ResponseDdnsHostnames) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceManagementInterface200ResponseDdnsHostnames) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceManagementInterface200ResponseDdnsHostnames(val *GetDeviceManagementInterface200ResponseDdnsHostnames) *NullableGetDeviceManagementInterface200ResponseDdnsHostnames {
	return &NullableGetDeviceManagementInterface200ResponseDdnsHostnames{value: val, isSet: true}
}

func (v NullableGetDeviceManagementInterface200ResponseDdnsHostnames) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceManagementInterface200ResponseDdnsHostnames) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


