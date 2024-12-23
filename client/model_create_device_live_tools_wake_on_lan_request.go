/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateDeviceLiveToolsWakeOnLanRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDeviceLiveToolsWakeOnLanRequest{}

// CreateDeviceLiveToolsWakeOnLanRequest struct for CreateDeviceLiveToolsWakeOnLanRequest
type CreateDeviceLiveToolsWakeOnLanRequest struct {
	// The target's VLAN (1 to 4094)
	VlanId int32 `json:"vlanId"`
	// The target's MAC address
	Mac string `json:"mac"`
	Callback *CreateDeviceLiveToolsArpTableRequestCallback `json:"callback,omitempty"`
}

type _CreateDeviceLiveToolsWakeOnLanRequest CreateDeviceLiveToolsWakeOnLanRequest

// NewCreateDeviceLiveToolsWakeOnLanRequest instantiates a new CreateDeviceLiveToolsWakeOnLanRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeviceLiveToolsWakeOnLanRequest(vlanId int32, mac string) *CreateDeviceLiveToolsWakeOnLanRequest {
	this := CreateDeviceLiveToolsWakeOnLanRequest{}
	this.VlanId = vlanId
	this.Mac = mac
	return &this
}

// NewCreateDeviceLiveToolsWakeOnLanRequestWithDefaults instantiates a new CreateDeviceLiveToolsWakeOnLanRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeviceLiveToolsWakeOnLanRequestWithDefaults() *CreateDeviceLiveToolsWakeOnLanRequest {
	this := CreateDeviceLiveToolsWakeOnLanRequest{}
	return &this
}

// GetVlanId returns the VlanId field value
func (o *CreateDeviceLiveToolsWakeOnLanRequest) GetVlanId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsWakeOnLanRequest) GetVlanIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VlanId, true
}

// SetVlanId sets field value
func (o *CreateDeviceLiveToolsWakeOnLanRequest) SetVlanId(v int32) {
	o.VlanId = v
}

// GetMac returns the Mac field value
func (o *CreateDeviceLiveToolsWakeOnLanRequest) GetMac() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mac
}

// GetMacOk returns a tuple with the Mac field value
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsWakeOnLanRequest) GetMacOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mac, true
}

// SetMac sets field value
func (o *CreateDeviceLiveToolsWakeOnLanRequest) SetMac(v string) {
	o.Mac = v
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *CreateDeviceLiveToolsWakeOnLanRequest) GetCallback() CreateDeviceLiveToolsArpTableRequestCallback {
	if o == nil || IsNil(o.Callback) {
		var ret CreateDeviceLiveToolsArpTableRequestCallback
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceLiveToolsWakeOnLanRequest) GetCallbackOk() (*CreateDeviceLiveToolsArpTableRequestCallback, bool) {
	if o == nil || IsNil(o.Callback) {
		return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *CreateDeviceLiveToolsWakeOnLanRequest) HasCallback() bool {
	if o != nil && !IsNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given CreateDeviceLiveToolsArpTableRequestCallback and assigns it to the Callback field.
func (o *CreateDeviceLiveToolsWakeOnLanRequest) SetCallback(v CreateDeviceLiveToolsArpTableRequestCallback) {
	o.Callback = &v
}

func (o CreateDeviceLiveToolsWakeOnLanRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDeviceLiveToolsWakeOnLanRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vlanId"] = o.VlanId
	toSerialize["mac"] = o.Mac
	if !IsNil(o.Callback) {
		toSerialize["callback"] = o.Callback
	}
	return toSerialize, nil
}

func (o *CreateDeviceLiveToolsWakeOnLanRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"vlanId",
		"mac",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateDeviceLiveToolsWakeOnLanRequest := _CreateDeviceLiveToolsWakeOnLanRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateDeviceLiveToolsWakeOnLanRequest)

	if err != nil {
		return err
	}

	*o = CreateDeviceLiveToolsWakeOnLanRequest(varCreateDeviceLiveToolsWakeOnLanRequest)

	return err
}

type NullableCreateDeviceLiveToolsWakeOnLanRequest struct {
	value *CreateDeviceLiveToolsWakeOnLanRequest
	isSet bool
}

func (v NullableCreateDeviceLiveToolsWakeOnLanRequest) Get() *CreateDeviceLiveToolsWakeOnLanRequest {
	return v.value
}

func (v *NullableCreateDeviceLiveToolsWakeOnLanRequest) Set(val *CreateDeviceLiveToolsWakeOnLanRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeviceLiveToolsWakeOnLanRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeviceLiveToolsWakeOnLanRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeviceLiveToolsWakeOnLanRequest(val *CreateDeviceLiveToolsWakeOnLanRequest) *NullableCreateDeviceLiveToolsWakeOnLanRequest {
	return &NullableCreateDeviceLiveToolsWakeOnLanRequest{value: val, isSet: true}
}

func (v NullableCreateDeviceLiveToolsWakeOnLanRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeviceLiveToolsWakeOnLanRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


