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

// UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey MD5 authentication credentials. This param is only relevant if md5AuthenticationEnabled is true
type UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey struct {
	// MD5 authentication key index. Key index must be between 1 to 255
	Id *int32 `json:"id,omitempty"`
	// MD5 authentication passphrase
	Passphrase *string `json:"passphrase,omitempty"`
}

// NewUpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey instantiates a new UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey() *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey {
	this := UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey{}
	return &this
}

// NewUpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKeyWithDefaults instantiates a new UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKeyWithDefaults() *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey {
	this := UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) SetId(v int32) {
	o.Id = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) GetPassphrase() string {
	if o == nil || isNil(o.Passphrase) {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) GetPassphraseOk() (*string, bool) {
	if o == nil || isNil(o.Passphrase) {
    return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) HasPassphrase() bool {
	if o != nil && !isNil(o.Passphrase) {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) SetPassphrase(v string) {
	o.Passphrase = &v
}

func (o UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Passphrase) {
		toSerialize["passphrase"] = o.Passphrase
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey struct {
	value *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey
	isSet bool
}

func (v NullableUpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) Get() *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey {
	return v.value
}

func (v *NullableUpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) Set(val *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey(val *UpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) *NullableUpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey {
	return &NullableUpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchRoutingOspfRequestMd5AuthenticationKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


