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

// checks if the CreateNetworkSwitchStackRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkSwitchStackRequest{}

// CreateNetworkSwitchStackRequest struct for CreateNetworkSwitchStackRequest
type CreateNetworkSwitchStackRequest struct {
	// The name of the new stack
	Name string `json:"name"`
	// An array of switch serials to be added into the new stack
	Serials []string `json:"serials"`
}

// NewCreateNetworkSwitchStackRequest instantiates a new CreateNetworkSwitchStackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkSwitchStackRequest(name string, serials []string) *CreateNetworkSwitchStackRequest {
	this := CreateNetworkSwitchStackRequest{}
	this.Name = name
	this.Serials = serials
	return &this
}

// NewCreateNetworkSwitchStackRequestWithDefaults instantiates a new CreateNetworkSwitchStackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkSwitchStackRequestWithDefaults() *CreateNetworkSwitchStackRequest {
	this := CreateNetworkSwitchStackRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateNetworkSwitchStackRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchStackRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNetworkSwitchStackRequest) SetName(v string) {
	o.Name = v
}

// GetSerials returns the Serials field value
func (o *CreateNetworkSwitchStackRequest) GetSerials() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchStackRequest) GetSerialsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Serials, true
}

// SetSerials sets field value
func (o *CreateNetworkSwitchStackRequest) SetSerials(v []string) {
	o.Serials = v
}

func (o CreateNetworkSwitchStackRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkSwitchStackRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["serials"] = o.Serials
	return toSerialize, nil
}

type NullableCreateNetworkSwitchStackRequest struct {
	value *CreateNetworkSwitchStackRequest
	isSet bool
}

func (v NullableCreateNetworkSwitchStackRequest) Get() *CreateNetworkSwitchStackRequest {
	return v.value
}

func (v *NullableCreateNetworkSwitchStackRequest) Set(val *CreateNetworkSwitchStackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkSwitchStackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkSwitchStackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkSwitchStackRequest(val *CreateNetworkSwitchStackRequest) *NullableCreateNetworkSwitchStackRequest {
	return &NullableCreateNetworkSwitchStackRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkSwitchStackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkSwitchStackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


