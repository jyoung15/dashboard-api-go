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

// checks if the DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner{}

// DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner struct for DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner
type DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner struct {
	// The index of the twisted pair tested.
	Index *int32 `json:"index,omitempty"`
	// The test result of the twisted pair tested.
	Status *string `json:"status,omitempty"`
	// The detected length of the twisted pair.
	LengthMeters *int32 `json:"lengthMeters,omitempty"`
}

// NewDevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner instantiates a new DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner() *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner {
	this := DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner{}
	return &this
}

// NewDevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInnerWithDefaults instantiates a new DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInnerWithDefaults() *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner {
	this := DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner) GetIndex() int32 {
	if o == nil || IsNil(o.Index) {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner) GetIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner) SetIndex(v int32) {
	o.Index = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner) SetStatus(v string) {
	o.Status = &v
}

// GetLengthMeters returns the LengthMeters field value if set, zero value otherwise.
func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner) GetLengthMeters() int32 {
	if o == nil || IsNil(o.LengthMeters) {
		var ret int32
		return ret
	}
	return *o.LengthMeters
}

// GetLengthMetersOk returns a tuple with the LengthMeters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner) GetLengthMetersOk() (*int32, bool) {
	if o == nil || IsNil(o.LengthMeters) {
		return nil, false
	}
	return o.LengthMeters, true
}

// HasLengthMeters returns a boolean if a field has been set.
func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner) HasLengthMeters() bool {
	if o != nil && !IsNil(o.LengthMeters) {
		return true
	}

	return false
}

// SetLengthMeters gets a reference to the given int32 and assigns it to the LengthMeters field.
func (o *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner) SetLengthMeters(v int32) {
	o.LengthMeters = &v
}

func (o DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.LengthMeters) {
		toSerialize["lengthMeters"] = o.LengthMeters
	}
	return toSerialize, nil
}

type NullableDevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner struct {
	value *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner
	isSet bool
}

func (v NullableDevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner) Get() *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner {
	return v.value
}

func (v *NullableDevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner) Set(val *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner(val *DevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner) *NullableDevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner {
	return &NullableDevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner{value: val, isSet: true}
}

func (v NullableDevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevicesSerialLiveToolsCableTestPostRequestMessageResultsInnerPairsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

