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

// checks if the CreateNetworkCameraWirelessProfileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkCameraWirelessProfileRequest{}

// CreateNetworkCameraWirelessProfileRequest struct for CreateNetworkCameraWirelessProfileRequest
type CreateNetworkCameraWirelessProfileRequest struct {
	// The name of the camera wireless profile. This parameter is required.
	Name string `json:"name"`
	Ssid CreateNetworkCameraWirelessProfileRequestSsid `json:"ssid"`
	Identity *CreateNetworkCameraWirelessProfileRequestIdentity `json:"identity,omitempty"`
}

// NewCreateNetworkCameraWirelessProfileRequest instantiates a new CreateNetworkCameraWirelessProfileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkCameraWirelessProfileRequest(name string, ssid CreateNetworkCameraWirelessProfileRequestSsid) *CreateNetworkCameraWirelessProfileRequest {
	this := CreateNetworkCameraWirelessProfileRequest{}
	this.Name = name
	this.Ssid = ssid
	return &this
}

// NewCreateNetworkCameraWirelessProfileRequestWithDefaults instantiates a new CreateNetworkCameraWirelessProfileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkCameraWirelessProfileRequestWithDefaults() *CreateNetworkCameraWirelessProfileRequest {
	this := CreateNetworkCameraWirelessProfileRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateNetworkCameraWirelessProfileRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraWirelessProfileRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNetworkCameraWirelessProfileRequest) SetName(v string) {
	o.Name = v
}

// GetSsid returns the Ssid field value
func (o *CreateNetworkCameraWirelessProfileRequest) GetSsid() CreateNetworkCameraWirelessProfileRequestSsid {
	if o == nil {
		var ret CreateNetworkCameraWirelessProfileRequestSsid
		return ret
	}

	return o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraWirelessProfileRequest) GetSsidOk() (*CreateNetworkCameraWirelessProfileRequestSsid, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ssid, true
}

// SetSsid sets field value
func (o *CreateNetworkCameraWirelessProfileRequest) SetSsid(v CreateNetworkCameraWirelessProfileRequestSsid) {
	o.Ssid = v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *CreateNetworkCameraWirelessProfileRequest) GetIdentity() CreateNetworkCameraWirelessProfileRequestIdentity {
	if o == nil || IsNil(o.Identity) {
		var ret CreateNetworkCameraWirelessProfileRequestIdentity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkCameraWirelessProfileRequest) GetIdentityOk() (*CreateNetworkCameraWirelessProfileRequestIdentity, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *CreateNetworkCameraWirelessProfileRequest) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given CreateNetworkCameraWirelessProfileRequestIdentity and assigns it to the Identity field.
func (o *CreateNetworkCameraWirelessProfileRequest) SetIdentity(v CreateNetworkCameraWirelessProfileRequestIdentity) {
	o.Identity = &v
}

func (o CreateNetworkCameraWirelessProfileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkCameraWirelessProfileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["ssid"] = o.Ssid
	if !IsNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	return toSerialize, nil
}

type NullableCreateNetworkCameraWirelessProfileRequest struct {
	value *CreateNetworkCameraWirelessProfileRequest
	isSet bool
}

func (v NullableCreateNetworkCameraWirelessProfileRequest) Get() *CreateNetworkCameraWirelessProfileRequest {
	return v.value
}

func (v *NullableCreateNetworkCameraWirelessProfileRequest) Set(val *CreateNetworkCameraWirelessProfileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkCameraWirelessProfileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkCameraWirelessProfileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkCameraWirelessProfileRequest(val *CreateNetworkCameraWirelessProfileRequest) *NullableCreateNetworkCameraWirelessProfileRequest {
	return &NullableCreateNetworkCameraWirelessProfileRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkCameraWirelessProfileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkCameraWirelessProfileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


