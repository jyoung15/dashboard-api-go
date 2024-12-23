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

// checks if the CreateOrganizationSplashThemeAssetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationSplashThemeAssetRequest{}

// CreateOrganizationSplashThemeAssetRequest struct for CreateOrganizationSplashThemeAssetRequest
type CreateOrganizationSplashThemeAssetRequest struct {
	// File name. Will overwrite files with same name.
	Name *string `json:"name,omitempty"`
	// a file containing the asset content
	Content *string `json:"content,omitempty"`
}

// NewCreateOrganizationSplashThemeAssetRequest instantiates a new CreateOrganizationSplashThemeAssetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationSplashThemeAssetRequest() *CreateOrganizationSplashThemeAssetRequest {
	this := CreateOrganizationSplashThemeAssetRequest{}
	return &this
}

// NewCreateOrganizationSplashThemeAssetRequestWithDefaults instantiates a new CreateOrganizationSplashThemeAssetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationSplashThemeAssetRequestWithDefaults() *CreateOrganizationSplashThemeAssetRequest {
	this := CreateOrganizationSplashThemeAssetRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateOrganizationSplashThemeAssetRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationSplashThemeAssetRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateOrganizationSplashThemeAssetRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateOrganizationSplashThemeAssetRequest) SetName(v string) {
	o.Name = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *CreateOrganizationSplashThemeAssetRequest) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationSplashThemeAssetRequest) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *CreateOrganizationSplashThemeAssetRequest) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *CreateOrganizationSplashThemeAssetRequest) SetContent(v string) {
	o.Content = &v
}

func (o CreateOrganizationSplashThemeAssetRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationSplashThemeAssetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	return toSerialize, nil
}

type NullableCreateOrganizationSplashThemeAssetRequest struct {
	value *CreateOrganizationSplashThemeAssetRequest
	isSet bool
}

func (v NullableCreateOrganizationSplashThemeAssetRequest) Get() *CreateOrganizationSplashThemeAssetRequest {
	return v.value
}

func (v *NullableCreateOrganizationSplashThemeAssetRequest) Set(val *CreateOrganizationSplashThemeAssetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationSplashThemeAssetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationSplashThemeAssetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationSplashThemeAssetRequest(val *CreateOrganizationSplashThemeAssetRequest) *NullableCreateOrganizationSplashThemeAssetRequest {
	return &NullableCreateOrganizationSplashThemeAssetRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationSplashThemeAssetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationSplashThemeAssetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


