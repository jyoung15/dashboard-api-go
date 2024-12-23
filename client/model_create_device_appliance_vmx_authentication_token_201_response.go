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

// checks if the CreateDeviceApplianceVmxAuthenticationToken201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDeviceApplianceVmxAuthenticationToken201Response{}

// CreateDeviceApplianceVmxAuthenticationToken201Response struct for CreateDeviceApplianceVmxAuthenticationToken201Response
type CreateDeviceApplianceVmxAuthenticationToken201Response struct {
	// The newly generated authentication token for the vMX instance
	Token *string `json:"token,omitempty"`
	// The expiration time for the token, in ISO 8601 format
	ExpiresAt *string `json:"expiresAt,omitempty"`
}

// NewCreateDeviceApplianceVmxAuthenticationToken201Response instantiates a new CreateDeviceApplianceVmxAuthenticationToken201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDeviceApplianceVmxAuthenticationToken201Response() *CreateDeviceApplianceVmxAuthenticationToken201Response {
	this := CreateDeviceApplianceVmxAuthenticationToken201Response{}
	return &this
}

// NewCreateDeviceApplianceVmxAuthenticationToken201ResponseWithDefaults instantiates a new CreateDeviceApplianceVmxAuthenticationToken201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDeviceApplianceVmxAuthenticationToken201ResponseWithDefaults() *CreateDeviceApplianceVmxAuthenticationToken201Response {
	this := CreateDeviceApplianceVmxAuthenticationToken201Response{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateDeviceApplianceVmxAuthenticationToken201Response) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceApplianceVmxAuthenticationToken201Response) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateDeviceApplianceVmxAuthenticationToken201Response) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateDeviceApplianceVmxAuthenticationToken201Response) SetToken(v string) {
	o.Token = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *CreateDeviceApplianceVmxAuthenticationToken201Response) GetExpiresAt() string {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDeviceApplianceVmxAuthenticationToken201Response) GetExpiresAtOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CreateDeviceApplianceVmxAuthenticationToken201Response) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *CreateDeviceApplianceVmxAuthenticationToken201Response) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

func (o CreateDeviceApplianceVmxAuthenticationToken201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDeviceApplianceVmxAuthenticationToken201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	return toSerialize, nil
}

type NullableCreateDeviceApplianceVmxAuthenticationToken201Response struct {
	value *CreateDeviceApplianceVmxAuthenticationToken201Response
	isSet bool
}

func (v NullableCreateDeviceApplianceVmxAuthenticationToken201Response) Get() *CreateDeviceApplianceVmxAuthenticationToken201Response {
	return v.value
}

func (v *NullableCreateDeviceApplianceVmxAuthenticationToken201Response) Set(val *CreateDeviceApplianceVmxAuthenticationToken201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDeviceApplianceVmxAuthenticationToken201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDeviceApplianceVmxAuthenticationToken201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDeviceApplianceVmxAuthenticationToken201Response(val *CreateDeviceApplianceVmxAuthenticationToken201Response) *NullableCreateDeviceApplianceVmxAuthenticationToken201Response {
	return &NullableCreateDeviceApplianceVmxAuthenticationToken201Response{value: val, isSet: true}
}

func (v NullableCreateDeviceApplianceVmxAuthenticationToken201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDeviceApplianceVmxAuthenticationToken201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


