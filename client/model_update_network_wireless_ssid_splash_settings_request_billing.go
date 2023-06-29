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

// checks if the UpdateNetworkWirelessSsidSplashSettingsRequestBilling type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidSplashSettingsRequestBilling{}

// UpdateNetworkWirelessSsidSplashSettingsRequestBilling Details associated with billing splash.
type UpdateNetworkWirelessSsidSplashSettingsRequestBilling struct {
	FreeAccess *UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess `json:"freeAccess,omitempty"`
	// Whether or not billing uses the fast login prepaid access option.
	PrepaidAccessFastLoginEnabled *bool `json:"prepaidAccessFastLoginEnabled,omitempty"`
	// The email address that receives replies from clients.
	ReplyToEmailAddress *string `json:"replyToEmailAddress,omitempty"`
}

// NewUpdateNetworkWirelessSsidSplashSettingsRequestBilling instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequestBilling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidSplashSettingsRequestBilling() *UpdateNetworkWirelessSsidSplashSettingsRequestBilling {
	this := UpdateNetworkWirelessSsidSplashSettingsRequestBilling{}
	return &this
}

// NewUpdateNetworkWirelessSsidSplashSettingsRequestBillingWithDefaults instantiates a new UpdateNetworkWirelessSsidSplashSettingsRequestBilling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidSplashSettingsRequestBillingWithDefaults() *UpdateNetworkWirelessSsidSplashSettingsRequestBilling {
	this := UpdateNetworkWirelessSsidSplashSettingsRequestBilling{}
	return &this
}

// GetFreeAccess returns the FreeAccess field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestBilling) GetFreeAccess() UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess {
	if o == nil || IsNil(o.FreeAccess) {
		var ret UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess
		return ret
	}
	return *o.FreeAccess
}

// GetFreeAccessOk returns a tuple with the FreeAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestBilling) GetFreeAccessOk() (*UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess, bool) {
	if o == nil || IsNil(o.FreeAccess) {
		return nil, false
	}
	return o.FreeAccess, true
}

// HasFreeAccess returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestBilling) HasFreeAccess() bool {
	if o != nil && !IsNil(o.FreeAccess) {
		return true
	}

	return false
}

// SetFreeAccess gets a reference to the given UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess and assigns it to the FreeAccess field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestBilling) SetFreeAccess(v UpdateNetworkWirelessSsidSplashSettingsRequestBillingFreeAccess) {
	o.FreeAccess = &v
}

// GetPrepaidAccessFastLoginEnabled returns the PrepaidAccessFastLoginEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestBilling) GetPrepaidAccessFastLoginEnabled() bool {
	if o == nil || IsNil(o.PrepaidAccessFastLoginEnabled) {
		var ret bool
		return ret
	}
	return *o.PrepaidAccessFastLoginEnabled
}

// GetPrepaidAccessFastLoginEnabledOk returns a tuple with the PrepaidAccessFastLoginEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestBilling) GetPrepaidAccessFastLoginEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PrepaidAccessFastLoginEnabled) {
		return nil, false
	}
	return o.PrepaidAccessFastLoginEnabled, true
}

// HasPrepaidAccessFastLoginEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestBilling) HasPrepaidAccessFastLoginEnabled() bool {
	if o != nil && !IsNil(o.PrepaidAccessFastLoginEnabled) {
		return true
	}

	return false
}

// SetPrepaidAccessFastLoginEnabled gets a reference to the given bool and assigns it to the PrepaidAccessFastLoginEnabled field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestBilling) SetPrepaidAccessFastLoginEnabled(v bool) {
	o.PrepaidAccessFastLoginEnabled = &v
}

// GetReplyToEmailAddress returns the ReplyToEmailAddress field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestBilling) GetReplyToEmailAddress() string {
	if o == nil || IsNil(o.ReplyToEmailAddress) {
		var ret string
		return ret
	}
	return *o.ReplyToEmailAddress
}

// GetReplyToEmailAddressOk returns a tuple with the ReplyToEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestBilling) GetReplyToEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ReplyToEmailAddress) {
		return nil, false
	}
	return o.ReplyToEmailAddress, true
}

// HasReplyToEmailAddress returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestBilling) HasReplyToEmailAddress() bool {
	if o != nil && !IsNil(o.ReplyToEmailAddress) {
		return true
	}

	return false
}

// SetReplyToEmailAddress gets a reference to the given string and assigns it to the ReplyToEmailAddress field.
func (o *UpdateNetworkWirelessSsidSplashSettingsRequestBilling) SetReplyToEmailAddress(v string) {
	o.ReplyToEmailAddress = &v
}

func (o UpdateNetworkWirelessSsidSplashSettingsRequestBilling) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidSplashSettingsRequestBilling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FreeAccess) {
		toSerialize["freeAccess"] = o.FreeAccess
	}
	if !IsNil(o.PrepaidAccessFastLoginEnabled) {
		toSerialize["prepaidAccessFastLoginEnabled"] = o.PrepaidAccessFastLoginEnabled
	}
	if !IsNil(o.ReplyToEmailAddress) {
		toSerialize["replyToEmailAddress"] = o.ReplyToEmailAddress
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidSplashSettingsRequestBilling struct {
	value *UpdateNetworkWirelessSsidSplashSettingsRequestBilling
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestBilling) Get() *UpdateNetworkWirelessSsidSplashSettingsRequestBilling {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestBilling) Set(val *UpdateNetworkWirelessSsidSplashSettingsRequestBilling) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestBilling) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestBilling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidSplashSettingsRequestBilling(val *UpdateNetworkWirelessSsidSplashSettingsRequestBilling) *NullableUpdateNetworkWirelessSsidSplashSettingsRequestBilling {
	return &NullableUpdateNetworkWirelessSsidSplashSettingsRequestBilling{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidSplashSettingsRequestBilling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidSplashSettingsRequestBilling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


