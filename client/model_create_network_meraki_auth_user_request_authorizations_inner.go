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

// checks if the CreateNetworkMerakiAuthUserRequestAuthorizationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkMerakiAuthUserRequestAuthorizationsInner{}

// CreateNetworkMerakiAuthUserRequestAuthorizationsInner struct for CreateNetworkMerakiAuthUserRequestAuthorizationsInner
type CreateNetworkMerakiAuthUserRequestAuthorizationsInner struct {
	// Required for wireless networks. The SSID for which the user is being authorized, which must be configured for the user's given accountType.
	SsidNumber *int32 `json:"ssidNumber,omitempty"`
	// Date for authorization to expire. Set to 'Never' for the authorization to not expire, which is the default.
	ExpiresAt *string `json:"expiresAt,omitempty"`
}

// NewCreateNetworkMerakiAuthUserRequestAuthorizationsInner instantiates a new CreateNetworkMerakiAuthUserRequestAuthorizationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkMerakiAuthUserRequestAuthorizationsInner() *CreateNetworkMerakiAuthUserRequestAuthorizationsInner {
	this := CreateNetworkMerakiAuthUserRequestAuthorizationsInner{}
	var expiresAt string = "Never"
	this.ExpiresAt = &expiresAt
	return &this
}

// NewCreateNetworkMerakiAuthUserRequestAuthorizationsInnerWithDefaults instantiates a new CreateNetworkMerakiAuthUserRequestAuthorizationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkMerakiAuthUserRequestAuthorizationsInnerWithDefaults() *CreateNetworkMerakiAuthUserRequestAuthorizationsInner {
	this := CreateNetworkMerakiAuthUserRequestAuthorizationsInner{}
	var expiresAt string = "Never"
	this.ExpiresAt = &expiresAt
	return &this
}

// GetSsidNumber returns the SsidNumber field value if set, zero value otherwise.
func (o *CreateNetworkMerakiAuthUserRequestAuthorizationsInner) GetSsidNumber() int32 {
	if o == nil || IsNil(o.SsidNumber) {
		var ret int32
		return ret
	}
	return *o.SsidNumber
}

// GetSsidNumberOk returns a tuple with the SsidNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkMerakiAuthUserRequestAuthorizationsInner) GetSsidNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.SsidNumber) {
		return nil, false
	}
	return o.SsidNumber, true
}

// HasSsidNumber returns a boolean if a field has been set.
func (o *CreateNetworkMerakiAuthUserRequestAuthorizationsInner) HasSsidNumber() bool {
	if o != nil && !IsNil(o.SsidNumber) {
		return true
	}

	return false
}

// SetSsidNumber gets a reference to the given int32 and assigns it to the SsidNumber field.
func (o *CreateNetworkMerakiAuthUserRequestAuthorizationsInner) SetSsidNumber(v int32) {
	o.SsidNumber = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *CreateNetworkMerakiAuthUserRequestAuthorizationsInner) GetExpiresAt() string {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkMerakiAuthUserRequestAuthorizationsInner) GetExpiresAtOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CreateNetworkMerakiAuthUserRequestAuthorizationsInner) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *CreateNetworkMerakiAuthUserRequestAuthorizationsInner) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

func (o CreateNetworkMerakiAuthUserRequestAuthorizationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkMerakiAuthUserRequestAuthorizationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SsidNumber) {
		toSerialize["ssidNumber"] = o.SsidNumber
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	return toSerialize, nil
}

type NullableCreateNetworkMerakiAuthUserRequestAuthorizationsInner struct {
	value *CreateNetworkMerakiAuthUserRequestAuthorizationsInner
	isSet bool
}

func (v NullableCreateNetworkMerakiAuthUserRequestAuthorizationsInner) Get() *CreateNetworkMerakiAuthUserRequestAuthorizationsInner {
	return v.value
}

func (v *NullableCreateNetworkMerakiAuthUserRequestAuthorizationsInner) Set(val *CreateNetworkMerakiAuthUserRequestAuthorizationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkMerakiAuthUserRequestAuthorizationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkMerakiAuthUserRequestAuthorizationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkMerakiAuthUserRequestAuthorizationsInner(val *CreateNetworkMerakiAuthUserRequestAuthorizationsInner) *NullableCreateNetworkMerakiAuthUserRequestAuthorizationsInner {
	return &NullableCreateNetworkMerakiAuthUserRequestAuthorizationsInner{value: val, isSet: true}
}

func (v NullableCreateNetworkMerakiAuthUserRequestAuthorizationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkMerakiAuthUserRequestAuthorizationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


