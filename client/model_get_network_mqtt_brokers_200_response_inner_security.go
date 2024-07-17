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

// checks if the GetNetworkMqttBrokers200ResponseInnerSecurity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkMqttBrokers200ResponseInnerSecurity{}

// GetNetworkMqttBrokers200ResponseInnerSecurity Security settings of the MQTT broker.
type GetNetworkMqttBrokers200ResponseInnerSecurity struct {
	// Security protocol of the MQTT broker.
	Mode *string `json:"mode,omitempty"`
	Tls *GetNetworkMqttBrokers200ResponseInnerSecurityTls `json:"tls,omitempty"`
}

// NewGetNetworkMqttBrokers200ResponseInnerSecurity instantiates a new GetNetworkMqttBrokers200ResponseInnerSecurity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkMqttBrokers200ResponseInnerSecurity() *GetNetworkMqttBrokers200ResponseInnerSecurity {
	this := GetNetworkMqttBrokers200ResponseInnerSecurity{}
	return &this
}

// NewGetNetworkMqttBrokers200ResponseInnerSecurityWithDefaults instantiates a new GetNetworkMqttBrokers200ResponseInnerSecurity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkMqttBrokers200ResponseInnerSecurityWithDefaults() *GetNetworkMqttBrokers200ResponseInnerSecurity {
	this := GetNetworkMqttBrokers200ResponseInnerSecurity{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *GetNetworkMqttBrokers200ResponseInnerSecurity) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkMqttBrokers200ResponseInnerSecurity) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *GetNetworkMqttBrokers200ResponseInnerSecurity) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *GetNetworkMqttBrokers200ResponseInnerSecurity) SetMode(v string) {
	o.Mode = &v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *GetNetworkMqttBrokers200ResponseInnerSecurity) GetTls() GetNetworkMqttBrokers200ResponseInnerSecurityTls {
	if o == nil || IsNil(o.Tls) {
		var ret GetNetworkMqttBrokers200ResponseInnerSecurityTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkMqttBrokers200ResponseInnerSecurity) GetTlsOk() (*GetNetworkMqttBrokers200ResponseInnerSecurityTls, bool) {
	if o == nil || IsNil(o.Tls) {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *GetNetworkMqttBrokers200ResponseInnerSecurity) HasTls() bool {
	if o != nil && !IsNil(o.Tls) {
		return true
	}

	return false
}

// SetTls gets a reference to the given GetNetworkMqttBrokers200ResponseInnerSecurityTls and assigns it to the Tls field.
func (o *GetNetworkMqttBrokers200ResponseInnerSecurity) SetTls(v GetNetworkMqttBrokers200ResponseInnerSecurityTls) {
	o.Tls = &v
}

func (o GetNetworkMqttBrokers200ResponseInnerSecurity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkMqttBrokers200ResponseInnerSecurity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Tls) {
		toSerialize["tls"] = o.Tls
	}
	return toSerialize, nil
}

type NullableGetNetworkMqttBrokers200ResponseInnerSecurity struct {
	value *GetNetworkMqttBrokers200ResponseInnerSecurity
	isSet bool
}

func (v NullableGetNetworkMqttBrokers200ResponseInnerSecurity) Get() *GetNetworkMqttBrokers200ResponseInnerSecurity {
	return v.value
}

func (v *NullableGetNetworkMqttBrokers200ResponseInnerSecurity) Set(val *GetNetworkMqttBrokers200ResponseInnerSecurity) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkMqttBrokers200ResponseInnerSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkMqttBrokers200ResponseInnerSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkMqttBrokers200ResponseInnerSecurity(val *GetNetworkMqttBrokers200ResponseInnerSecurity) *NullableGetNetworkMqttBrokers200ResponseInnerSecurity {
	return &NullableGetNetworkMqttBrokers200ResponseInnerSecurity{value: val, isSet: true}
}

func (v NullableGetNetworkMqttBrokers200ResponseInnerSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkMqttBrokers200ResponseInnerSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

