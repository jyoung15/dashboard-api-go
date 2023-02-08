/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront The prepaid front image used in the splash page.
type UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront struct {
	// The MD5 value of the prepaid front image file. Setting this to null will remove the prepaid front from the splash page.
	Md5 *string `json:"md5,omitempty"`
	// The extension of the prepaid front image file.
	Extension *string `json:"extension,omitempty"`
	Image *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage `json:"image,omitempty"`
}

// NewUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront() *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront {
	this := UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront{}
	return &this
}

// NewUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontWithDefaults instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontWithDefaults() *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront {
	this := UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront{}
	return &this
}

// GetMd5 returns the Md5 field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront) GetMd5() string {
	if o == nil || isNil(o.Md5) {
		var ret string
		return ret
	}
	return *o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront) GetMd5Ok() (*string, bool) {
	if o == nil || isNil(o.Md5) {
    return nil, false
	}
	return o.Md5, true
}

// HasMd5 returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront) HasMd5() bool {
	if o != nil && !isNil(o.Md5) {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given string and assigns it to the Md5 field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront) SetMd5(v string) {
	o.Md5 = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront) GetExtension() string {
	if o == nil || isNil(o.Extension) {
		var ret string
		return ret
	}
	return *o.Extension
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront) GetExtensionOk() (*string, bool) {
	if o == nil || isNil(o.Extension) {
    return nil, false
	}
	return o.Extension, true
}

// HasExtension returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront) HasExtension() bool {
	if o != nil && !isNil(o.Extension) {
		return true
	}

	return false
}

// SetExtension gets a reference to the given string and assigns it to the Extension field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront) SetExtension(v string) {
	o.Extension = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront) GetImage() UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage {
	if o == nil || isNil(o.Image) {
		var ret UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront) GetImageOk() (*UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage, bool) {
	if o == nil || isNil(o.Image) {
    return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront) HasImage() bool {
	if o != nil && !isNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage and assigns it to the Image field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront) SetImage(v UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFrontImage) {
	o.Image = &v
}

func (o UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Md5) {
		toSerialize["md5"] = o.Md5
	}
	if !isNil(o.Extension) {
		toSerialize["extension"] = o.Extension
	}
	if !isNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront struct {
	value *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront) Get() *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront) Set(val *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront(val *UpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront) *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront {
	return &NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestSplashPrepaidFront) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


