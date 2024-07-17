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

// checks if the GetDeviceCellularSims200ResponseSimsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceCellularSims200ResponseSimsInner{}

// GetDeviceCellularSims200ResponseSimsInner struct for GetDeviceCellularSims200ResponseSimsInner
type GetDeviceCellularSims200ResponseSimsInner struct {
	// SIM slot being configured. Must be 'sim1' on single-sim devices. Use 'sim3' for eSIM.
	Slot *string `json:"slot,omitempty"`
	// If true, this SIM is activated on platform bootup. It must be true on single-SIM devices and is a required field for dual-SIM MGs unless it is being configured using 'simOrdering'.
	IsPrimary *bool `json:"isPrimary,omitempty"`
	// APN configurations. If empty, the default APN will be used.
	Apns []GetDeviceCellularSims200ResponseSimsInnerApnsInner `json:"apns,omitempty"`
}

// NewGetDeviceCellularSims200ResponseSimsInner instantiates a new GetDeviceCellularSims200ResponseSimsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceCellularSims200ResponseSimsInner() *GetDeviceCellularSims200ResponseSimsInner {
	this := GetDeviceCellularSims200ResponseSimsInner{}
	var isPrimary bool = false
	this.IsPrimary = &isPrimary
	return &this
}

// NewGetDeviceCellularSims200ResponseSimsInnerWithDefaults instantiates a new GetDeviceCellularSims200ResponseSimsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceCellularSims200ResponseSimsInnerWithDefaults() *GetDeviceCellularSims200ResponseSimsInner {
	this := GetDeviceCellularSims200ResponseSimsInner{}
	var isPrimary bool = false
	this.IsPrimary = &isPrimary
	return &this
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *GetDeviceCellularSims200ResponseSimsInner) GetSlot() string {
	if o == nil || IsNil(o.Slot) {
		var ret string
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceCellularSims200ResponseSimsInner) GetSlotOk() (*string, bool) {
	if o == nil || IsNil(o.Slot) {
		return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *GetDeviceCellularSims200ResponseSimsInner) HasSlot() bool {
	if o != nil && !IsNil(o.Slot) {
		return true
	}

	return false
}

// SetSlot gets a reference to the given string and assigns it to the Slot field.
func (o *GetDeviceCellularSims200ResponseSimsInner) SetSlot(v string) {
	o.Slot = &v
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *GetDeviceCellularSims200ResponseSimsInner) GetIsPrimary() bool {
	if o == nil || IsNil(o.IsPrimary) {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceCellularSims200ResponseSimsInner) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrimary) {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *GetDeviceCellularSims200ResponseSimsInner) HasIsPrimary() bool {
	if o != nil && !IsNil(o.IsPrimary) {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *GetDeviceCellularSims200ResponseSimsInner) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetApns returns the Apns field value if set, zero value otherwise.
func (o *GetDeviceCellularSims200ResponseSimsInner) GetApns() []GetDeviceCellularSims200ResponseSimsInnerApnsInner {
	if o == nil || IsNil(o.Apns) {
		var ret []GetDeviceCellularSims200ResponseSimsInnerApnsInner
		return ret
	}
	return o.Apns
}

// GetApnsOk returns a tuple with the Apns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceCellularSims200ResponseSimsInner) GetApnsOk() ([]GetDeviceCellularSims200ResponseSimsInnerApnsInner, bool) {
	if o == nil || IsNil(o.Apns) {
		return nil, false
	}
	return o.Apns, true
}

// HasApns returns a boolean if a field has been set.
func (o *GetDeviceCellularSims200ResponseSimsInner) HasApns() bool {
	if o != nil && !IsNil(o.Apns) {
		return true
	}

	return false
}

// SetApns gets a reference to the given []GetDeviceCellularSims200ResponseSimsInnerApnsInner and assigns it to the Apns field.
func (o *GetDeviceCellularSims200ResponseSimsInner) SetApns(v []GetDeviceCellularSims200ResponseSimsInnerApnsInner) {
	o.Apns = v
}

func (o GetDeviceCellularSims200ResponseSimsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceCellularSims200ResponseSimsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Slot) {
		toSerialize["slot"] = o.Slot
	}
	if !IsNil(o.IsPrimary) {
		toSerialize["isPrimary"] = o.IsPrimary
	}
	if !IsNil(o.Apns) {
		toSerialize["apns"] = o.Apns
	}
	return toSerialize, nil
}

type NullableGetDeviceCellularSims200ResponseSimsInner struct {
	value *GetDeviceCellularSims200ResponseSimsInner
	isSet bool
}

func (v NullableGetDeviceCellularSims200ResponseSimsInner) Get() *GetDeviceCellularSims200ResponseSimsInner {
	return v.value
}

func (v *NullableGetDeviceCellularSims200ResponseSimsInner) Set(val *GetDeviceCellularSims200ResponseSimsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceCellularSims200ResponseSimsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceCellularSims200ResponseSimsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceCellularSims200ResponseSimsInner(val *GetDeviceCellularSims200ResponseSimsInner) *NullableGetDeviceCellularSims200ResponseSimsInner {
	return &NullableGetDeviceCellularSims200ResponseSimsInner{value: val, isSet: true}
}

func (v NullableGetDeviceCellularSims200ResponseSimsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceCellularSims200ResponseSimsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

