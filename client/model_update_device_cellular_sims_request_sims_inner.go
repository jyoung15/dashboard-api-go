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

// checks if the UpdateDeviceCellularSimsRequestSimsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceCellularSimsRequestSimsInner{}

// UpdateDeviceCellularSimsRequestSimsInner struct for UpdateDeviceCellularSimsRequestSimsInner
type UpdateDeviceCellularSimsRequestSimsInner struct {
	// SIM slot being configured. Must be 'sim1' on single-sim devices.
	Slot *string `json:"slot,omitempty"`
	// If true, this SIM is used for boot. Must be true on single-sim devices.
	IsPrimary *bool `json:"isPrimary,omitempty"`
	// APN configurations. If empty, the default APN will be used.
	Apns []UpdateDeviceCellularSimsRequestSimsInnerApnsInner `json:"apns,omitempty"`
}

// NewUpdateDeviceCellularSimsRequestSimsInner instantiates a new UpdateDeviceCellularSimsRequestSimsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceCellularSimsRequestSimsInner() *UpdateDeviceCellularSimsRequestSimsInner {
	this := UpdateDeviceCellularSimsRequestSimsInner{}
	var isPrimary bool = false
	this.IsPrimary = &isPrimary
	return &this
}

// NewUpdateDeviceCellularSimsRequestSimsInnerWithDefaults instantiates a new UpdateDeviceCellularSimsRequestSimsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceCellularSimsRequestSimsInnerWithDefaults() *UpdateDeviceCellularSimsRequestSimsInner {
	this := UpdateDeviceCellularSimsRequestSimsInner{}
	var isPrimary bool = false
	this.IsPrimary = &isPrimary
	return &this
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *UpdateDeviceCellularSimsRequestSimsInner) GetSlot() string {
	if o == nil || IsNil(o.Slot) {
		var ret string
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCellularSimsRequestSimsInner) GetSlotOk() (*string, bool) {
	if o == nil || IsNil(o.Slot) {
		return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *UpdateDeviceCellularSimsRequestSimsInner) HasSlot() bool {
	if o != nil && !IsNil(o.Slot) {
		return true
	}

	return false
}

// SetSlot gets a reference to the given string and assigns it to the Slot field.
func (o *UpdateDeviceCellularSimsRequestSimsInner) SetSlot(v string) {
	o.Slot = &v
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *UpdateDeviceCellularSimsRequestSimsInner) GetIsPrimary() bool {
	if o == nil || IsNil(o.IsPrimary) {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCellularSimsRequestSimsInner) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrimary) {
		return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *UpdateDeviceCellularSimsRequestSimsInner) HasIsPrimary() bool {
	if o != nil && !IsNil(o.IsPrimary) {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *UpdateDeviceCellularSimsRequestSimsInner) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetApns returns the Apns field value if set, zero value otherwise.
func (o *UpdateDeviceCellularSimsRequestSimsInner) GetApns() []UpdateDeviceCellularSimsRequestSimsInnerApnsInner {
	if o == nil || IsNil(o.Apns) {
		var ret []UpdateDeviceCellularSimsRequestSimsInnerApnsInner
		return ret
	}
	return o.Apns
}

// GetApnsOk returns a tuple with the Apns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCellularSimsRequestSimsInner) GetApnsOk() ([]UpdateDeviceCellularSimsRequestSimsInnerApnsInner, bool) {
	if o == nil || IsNil(o.Apns) {
		return nil, false
	}
	return o.Apns, true
}

// HasApns returns a boolean if a field has been set.
func (o *UpdateDeviceCellularSimsRequestSimsInner) HasApns() bool {
	if o != nil && !IsNil(o.Apns) {
		return true
	}

	return false
}

// SetApns gets a reference to the given []UpdateDeviceCellularSimsRequestSimsInnerApnsInner and assigns it to the Apns field.
func (o *UpdateDeviceCellularSimsRequestSimsInner) SetApns(v []UpdateDeviceCellularSimsRequestSimsInnerApnsInner) {
	o.Apns = v
}

func (o UpdateDeviceCellularSimsRequestSimsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceCellularSimsRequestSimsInner) ToMap() (map[string]interface{}, error) {
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

type NullableUpdateDeviceCellularSimsRequestSimsInner struct {
	value *UpdateDeviceCellularSimsRequestSimsInner
	isSet bool
}

func (v NullableUpdateDeviceCellularSimsRequestSimsInner) Get() *UpdateDeviceCellularSimsRequestSimsInner {
	return v.value
}

func (v *NullableUpdateDeviceCellularSimsRequestSimsInner) Set(val *UpdateDeviceCellularSimsRequestSimsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceCellularSimsRequestSimsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceCellularSimsRequestSimsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceCellularSimsRequestSimsInner(val *UpdateDeviceCellularSimsRequestSimsInner) *NullableUpdateDeviceCellularSimsRequestSimsInner {
	return &NullableUpdateDeviceCellularSimsRequestSimsInner{value: val, isSet: true}
}

func (v NullableUpdateDeviceCellularSimsRequestSimsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceCellularSimsRequestSimsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


