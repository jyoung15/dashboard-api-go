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

// checks if the UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner{}

// UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner struct for UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner
type UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner struct {
	// The format for the realm ('1' or '0')
	Format *string `json:"format,omitempty"`
	// The name of the realm
	Realm *string `json:"realm,omitempty"`
	// An array of EAP methods for the realm.
	Methods []UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner `json:"methods,omitempty"`
}

// NewUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner instantiates a new UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner() *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner {
	this := UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner{}
	return &this
}

// NewUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerWithDefaults instantiates a new UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerWithDefaults() *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner {
	this := UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner) SetFormat(v string) {
	o.Format = &v
}

// GetRealm returns the Realm field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner) GetRealm() string {
	if o == nil || IsNil(o.Realm) {
		var ret string
		return ret
	}
	return *o.Realm
}

// GetRealmOk returns a tuple with the Realm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner) GetRealmOk() (*string, bool) {
	if o == nil || IsNil(o.Realm) {
		return nil, false
	}
	return o.Realm, true
}

// HasRealm returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner) HasRealm() bool {
	if o != nil && !IsNil(o.Realm) {
		return true
	}

	return false
}

// SetRealm gets a reference to the given string and assigns it to the Realm field.
func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner) SetRealm(v string) {
	o.Realm = &v
}

// GetMethods returns the Methods field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner) GetMethods() []UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner {
	if o == nil || IsNil(o.Methods) {
		var ret []UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner
		return ret
	}
	return o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner) GetMethodsOk() ([]UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner, bool) {
	if o == nil || IsNil(o.Methods) {
		return nil, false
	}
	return o.Methods, true
}

// HasMethods returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner) HasMethods() bool {
	if o != nil && !IsNil(o.Methods) {
		return true
	}

	return false
}

// SetMethods gets a reference to the given []UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner and assigns it to the Methods field.
func (o *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner) SetMethods(v []UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInnerMethodsInner) {
	o.Methods = v
}

func (o UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.Realm) {
		toSerialize["realm"] = o.Realm
	}
	if !IsNil(o.Methods) {
		toSerialize["methods"] = o.Methods
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner struct {
	value *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner) Get() *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner) Set(val *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner(val *UpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner) *NullableUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner {
	return &NullableUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidHotspot20RequestNaiRealmsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


