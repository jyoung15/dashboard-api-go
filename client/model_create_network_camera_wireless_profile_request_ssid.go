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

// checks if the CreateNetworkCameraWirelessProfileRequestSsid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkCameraWirelessProfileRequestSsid{}

// CreateNetworkCameraWirelessProfileRequestSsid The details of the SSID config.
type CreateNetworkCameraWirelessProfileRequestSsid struct {
	// The name of the SSID.
	Name *string `json:"name,omitempty"`
	// The auth mode of the SSID. It can be set to ('psk', '8021x-radius').
	AuthMode *string `json:"authMode,omitempty"`
	// The encryption mode of the SSID. It can be set to ('wpa', 'wpa-eap'). With 'wpa' mode, the authMode should be 'psk' and with 'wpa-eap' the authMode should be '8021x-radius'
	EncryptionMode *string `json:"encryptionMode,omitempty"`
	// The pre-shared key of the SSID.
	Psk *string `json:"psk,omitempty"`
}

// NewCreateNetworkCameraWirelessProfileRequestSsid instantiates a new CreateNetworkCameraWirelessProfileRequestSsid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkCameraWirelessProfileRequestSsid() *CreateNetworkCameraWirelessProfileRequestSsid {
	this := CreateNetworkCameraWirelessProfileRequestSsid{}
	return &this
}

// NewCreateNetworkCameraWirelessProfileRequestSsidWithDefaults instantiates a new CreateNetworkCameraWirelessProfileRequestSsid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkCameraWirelessProfileRequestSsidWithDefaults() *CreateNetworkCameraWirelessProfileRequestSsid {
	this := CreateNetworkCameraWirelessProfileRequestSsid{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateNetworkCameraWirelessProfileRequestSsid) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraWirelessProfileRequestSsid) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateNetworkCameraWirelessProfileRequestSsid) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateNetworkCameraWirelessProfileRequestSsid) SetName(v string) {
	o.Name = &v
}

// GetAuthMode returns the AuthMode field value if set, zero value otherwise.
func (o *CreateNetworkCameraWirelessProfileRequestSsid) GetAuthMode() string {
	if o == nil || IsNil(o.AuthMode) {
		var ret string
		return ret
	}
	return *o.AuthMode
}

// GetAuthModeOk returns a tuple with the AuthMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraWirelessProfileRequestSsid) GetAuthModeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthMode) {
		return nil, false
	}
	return o.AuthMode, true
}

// HasAuthMode returns a boolean if a field has been set.
func (o *CreateNetworkCameraWirelessProfileRequestSsid) HasAuthMode() bool {
	if o != nil && !IsNil(o.AuthMode) {
		return true
	}

	return false
}

// SetAuthMode gets a reference to the given string and assigns it to the AuthMode field.
func (o *CreateNetworkCameraWirelessProfileRequestSsid) SetAuthMode(v string) {
	o.AuthMode = &v
}

// GetEncryptionMode returns the EncryptionMode field value if set, zero value otherwise.
func (o *CreateNetworkCameraWirelessProfileRequestSsid) GetEncryptionMode() string {
	if o == nil || IsNil(o.EncryptionMode) {
		var ret string
		return ret
	}
	return *o.EncryptionMode
}

// GetEncryptionModeOk returns a tuple with the EncryptionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraWirelessProfileRequestSsid) GetEncryptionModeOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionMode) {
		return nil, false
	}
	return o.EncryptionMode, true
}

// HasEncryptionMode returns a boolean if a field has been set.
func (o *CreateNetworkCameraWirelessProfileRequestSsid) HasEncryptionMode() bool {
	if o != nil && !IsNil(o.EncryptionMode) {
		return true
	}

	return false
}

// SetEncryptionMode gets a reference to the given string and assigns it to the EncryptionMode field.
func (o *CreateNetworkCameraWirelessProfileRequestSsid) SetEncryptionMode(v string) {
	o.EncryptionMode = &v
}

// GetPsk returns the Psk field value if set, zero value otherwise.
func (o *CreateNetworkCameraWirelessProfileRequestSsid) GetPsk() string {
	if o == nil || IsNil(o.Psk) {
		var ret string
		return ret
	}
	return *o.Psk
}

// GetPskOk returns a tuple with the Psk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraWirelessProfileRequestSsid) GetPskOk() (*string, bool) {
	if o == nil || IsNil(o.Psk) {
		return nil, false
	}
	return o.Psk, true
}

// HasPsk returns a boolean if a field has been set.
func (o *CreateNetworkCameraWirelessProfileRequestSsid) HasPsk() bool {
	if o != nil && !IsNil(o.Psk) {
		return true
	}

	return false
}

// SetPsk gets a reference to the given string and assigns it to the Psk field.
func (o *CreateNetworkCameraWirelessProfileRequestSsid) SetPsk(v string) {
	o.Psk = &v
}

func (o CreateNetworkCameraWirelessProfileRequestSsid) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkCameraWirelessProfileRequestSsid) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.AuthMode) {
		toSerialize["authMode"] = o.AuthMode
	}
	if !IsNil(o.EncryptionMode) {
		toSerialize["encryptionMode"] = o.EncryptionMode
	}
	if !IsNil(o.Psk) {
		toSerialize["psk"] = o.Psk
	}
	return toSerialize, nil
}

type NullableCreateNetworkCameraWirelessProfileRequestSsid struct {
	value *CreateNetworkCameraWirelessProfileRequestSsid
	isSet bool
}

func (v NullableCreateNetworkCameraWirelessProfileRequestSsid) Get() *CreateNetworkCameraWirelessProfileRequestSsid {
	return v.value
}

func (v *NullableCreateNetworkCameraWirelessProfileRequestSsid) Set(val *CreateNetworkCameraWirelessProfileRequestSsid) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkCameraWirelessProfileRequestSsid) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkCameraWirelessProfileRequestSsid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkCameraWirelessProfileRequestSsid(val *CreateNetworkCameraWirelessProfileRequestSsid) *NullableCreateNetworkCameraWirelessProfileRequestSsid {
	return &NullableCreateNetworkCameraWirelessProfileRequestSsid{value: val, isSet: true}
}

func (v NullableCreateNetworkCameraWirelessProfileRequestSsid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkCameraWirelessProfileRequestSsid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


