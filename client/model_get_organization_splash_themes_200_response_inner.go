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

// checks if the GetOrganizationSplashThemes200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSplashThemes200ResponseInner{}

// GetOrganizationSplashThemes200ResponseInner struct for GetOrganizationSplashThemes200ResponseInner
type GetOrganizationSplashThemes200ResponseInner struct {
	// theme id
	Id *string `json:"id,omitempty"`
	// theme name
	Name *string `json:"name,omitempty"`
	// list of theme assets
	ThemeAssets []GetOrganizationSplashThemes200ResponseInnerThemeAssetsInner `json:"themeAssets,omitempty"`
}

// NewGetOrganizationSplashThemes200ResponseInner instantiates a new GetOrganizationSplashThemes200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSplashThemes200ResponseInner() *GetOrganizationSplashThemes200ResponseInner {
	this := GetOrganizationSplashThemes200ResponseInner{}
	return &this
}

// NewGetOrganizationSplashThemes200ResponseInnerWithDefaults instantiates a new GetOrganizationSplashThemes200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSplashThemes200ResponseInnerWithDefaults() *GetOrganizationSplashThemes200ResponseInner {
	this := GetOrganizationSplashThemes200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationSplashThemes200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSplashThemes200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationSplashThemes200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationSplashThemes200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationSplashThemes200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSplashThemes200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationSplashThemes200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationSplashThemes200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetThemeAssets returns the ThemeAssets field value if set, zero value otherwise.
func (o *GetOrganizationSplashThemes200ResponseInner) GetThemeAssets() []GetOrganizationSplashThemes200ResponseInnerThemeAssetsInner {
	if o == nil || IsNil(o.ThemeAssets) {
		var ret []GetOrganizationSplashThemes200ResponseInnerThemeAssetsInner
		return ret
	}
	return o.ThemeAssets
}

// GetThemeAssetsOk returns a tuple with the ThemeAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSplashThemes200ResponseInner) GetThemeAssetsOk() ([]GetOrganizationSplashThemes200ResponseInnerThemeAssetsInner, bool) {
	if o == nil || IsNil(o.ThemeAssets) {
		return nil, false
	}
	return o.ThemeAssets, true
}

// HasThemeAssets returns a boolean if a field has been set.
func (o *GetOrganizationSplashThemes200ResponseInner) HasThemeAssets() bool {
	if o != nil && !IsNil(o.ThemeAssets) {
		return true
	}

	return false
}

// SetThemeAssets gets a reference to the given []GetOrganizationSplashThemes200ResponseInnerThemeAssetsInner and assigns it to the ThemeAssets field.
func (o *GetOrganizationSplashThemes200ResponseInner) SetThemeAssets(v []GetOrganizationSplashThemes200ResponseInnerThemeAssetsInner) {
	o.ThemeAssets = v
}

func (o GetOrganizationSplashThemes200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSplashThemes200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ThemeAssets) {
		toSerialize["themeAssets"] = o.ThemeAssets
	}
	return toSerialize, nil
}

type NullableGetOrganizationSplashThemes200ResponseInner struct {
	value *GetOrganizationSplashThemes200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationSplashThemes200ResponseInner) Get() *GetOrganizationSplashThemes200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationSplashThemes200ResponseInner) Set(val *GetOrganizationSplashThemes200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSplashThemes200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSplashThemes200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSplashThemes200ResponseInner(val *GetOrganizationSplashThemes200ResponseInner) *NullableGetOrganizationSplashThemes200ResponseInner {
	return &NullableGetOrganizationSplashThemes200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationSplashThemes200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSplashThemes200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


