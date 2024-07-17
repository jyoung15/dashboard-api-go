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

// checks if the GetNetworkWirelessElectronicShelfLabel200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessElectronicShelfLabel200Response{}

// GetNetworkWirelessElectronicShelfLabel200Response struct for GetNetworkWirelessElectronicShelfLabel200Response
type GetNetworkWirelessElectronicShelfLabel200Response struct {
	// Desired ESL hostname of the network
	Hostname *string `json:"hostname,omitempty"`
	// Turn ESL features on and off for this network
	Enabled *bool `json:"enabled,omitempty"`
}

// NewGetNetworkWirelessElectronicShelfLabel200Response instantiates a new GetNetworkWirelessElectronicShelfLabel200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessElectronicShelfLabel200Response() *GetNetworkWirelessElectronicShelfLabel200Response {
	this := GetNetworkWirelessElectronicShelfLabel200Response{}
	return &this
}

// NewGetNetworkWirelessElectronicShelfLabel200ResponseWithDefaults instantiates a new GetNetworkWirelessElectronicShelfLabel200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessElectronicShelfLabel200ResponseWithDefaults() *GetNetworkWirelessElectronicShelfLabel200Response {
	this := GetNetworkWirelessElectronicShelfLabel200Response{}
	return &this
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *GetNetworkWirelessElectronicShelfLabel200Response) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessElectronicShelfLabel200Response) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *GetNetworkWirelessElectronicShelfLabel200Response) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *GetNetworkWirelessElectronicShelfLabel200Response) SetHostname(v string) {
	o.Hostname = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetNetworkWirelessElectronicShelfLabel200Response) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessElectronicShelfLabel200Response) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetNetworkWirelessElectronicShelfLabel200Response) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetNetworkWirelessElectronicShelfLabel200Response) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o GetNetworkWirelessElectronicShelfLabel200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessElectronicShelfLabel200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessElectronicShelfLabel200Response struct {
	value *GetNetworkWirelessElectronicShelfLabel200Response
	isSet bool
}

func (v NullableGetNetworkWirelessElectronicShelfLabel200Response) Get() *GetNetworkWirelessElectronicShelfLabel200Response {
	return v.value
}

func (v *NullableGetNetworkWirelessElectronicShelfLabel200Response) Set(val *GetNetworkWirelessElectronicShelfLabel200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessElectronicShelfLabel200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessElectronicShelfLabel200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessElectronicShelfLabel200Response(val *GetNetworkWirelessElectronicShelfLabel200Response) *NullableGetNetworkWirelessElectronicShelfLabel200Response {
	return &NullableGetNetworkWirelessElectronicShelfLabel200Response{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessElectronicShelfLabel200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessElectronicShelfLabel200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

