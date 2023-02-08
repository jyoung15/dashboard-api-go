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

// CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories Settings for blocked URL categories
type CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories struct {
	// How URL categories are applied. Can be 'network default', 'append' or 'override'.
	Settings *string `json:"settings,omitempty"`
	// A list of URL categories to block
	Categories []string `json:"categories,omitempty"`
}

// NewCreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories instantiates a new CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories() *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories {
	this := CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories{}
	return &this
}

// NewCreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategoriesWithDefaults instantiates a new CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategoriesWithDefaults() *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories {
	this := CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories) GetSettings() string {
	if o == nil || isNil(o.Settings) {
		var ret string
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories) GetSettingsOk() (*string, bool) {
	if o == nil || isNil(o.Settings) {
    return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories) HasSettings() bool {
	if o != nil && !isNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given string and assigns it to the Settings field.
func (o *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories) SetSettings(v string) {
	o.Settings = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories) GetCategories() []string {
	if o == nil || isNil(o.Categories) {
		var ret []string
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories) GetCategoriesOk() ([]string, bool) {
	if o == nil || isNil(o.Categories) {
    return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories) HasCategories() bool {
	if o != nil && !isNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []string and assigns it to the Categories field.
func (o *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories) SetCategories(v []string) {
	o.Categories = v
}

func (o CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !isNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories struct {
	value *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories
	isSet bool
}

func (v NullableCreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories) Get() *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories {
	return v.value
}

func (v *NullableCreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories) Set(val *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories(val *CreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories) *NullableCreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories {
	return &NullableCreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories{value: val, isSet: true}
}

func (v NullableCreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkGroupPolicyRequestContentFilteringBlockedUrlCategories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


