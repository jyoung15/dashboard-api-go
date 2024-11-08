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

// checks if the AddNetworkSwitchStackRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddNetworkSwitchStackRequest{}

// AddNetworkSwitchStackRequest struct for AddNetworkSwitchStackRequest
type AddNetworkSwitchStackRequest struct {
	// The serial of the switch to be added
	Serial string `json:"serial"`
}

type _AddNetworkSwitchStackRequest AddNetworkSwitchStackRequest

// NewAddNetworkSwitchStackRequest instantiates a new AddNetworkSwitchStackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddNetworkSwitchStackRequest(serial string) *AddNetworkSwitchStackRequest {
	this := AddNetworkSwitchStackRequest{}
	this.Serial = serial
	return &this
}

// NewAddNetworkSwitchStackRequestWithDefaults instantiates a new AddNetworkSwitchStackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddNetworkSwitchStackRequestWithDefaults() *AddNetworkSwitchStackRequest {
	this := AddNetworkSwitchStackRequest{}
	return &this
}

// GetSerial returns the Serial field value
func (o *AddNetworkSwitchStackRequest) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *AddNetworkSwitchStackRequest) GetSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *AddNetworkSwitchStackRequest) SetSerial(v string) {
	o.Serial = v
}

func (o AddNetworkSwitchStackRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddNetworkSwitchStackRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serial"] = o.Serial
	return toSerialize, nil
}

func (o *AddNetworkSwitchStackRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"serial",
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

	varAddNetworkSwitchStackRequest := _AddNetworkSwitchStackRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddNetworkSwitchStackRequest)

	if err != nil {
		return err
	}

	*o = AddNetworkSwitchStackRequest(varAddNetworkSwitchStackRequest)

	return err
}

type NullableAddNetworkSwitchStackRequest struct {
	value *AddNetworkSwitchStackRequest
	isSet bool
}

func (v NullableAddNetworkSwitchStackRequest) Get() *AddNetworkSwitchStackRequest {
	return v.value
}

func (v *NullableAddNetworkSwitchStackRequest) Set(val *AddNetworkSwitchStackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddNetworkSwitchStackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddNetworkSwitchStackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddNetworkSwitchStackRequest(val *AddNetworkSwitchStackRequest) *NullableAddNetworkSwitchStackRequest {
	return &NullableAddNetworkSwitchStackRequest{value: val, isSet: true}
}

func (v NullableAddNetworkSwitchStackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddNetworkSwitchStackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


