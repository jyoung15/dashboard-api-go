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

// checks if the GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia{}

// GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia The inactive count data, indexed by media type (RJ45 or SFP)
type GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia struct {
	Rj45 *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMediaRj45 `json:"rj45,omitempty"`
	Sfp *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMediaSfp `json:"sfp,omitempty"`
}

// NewGetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia instantiates a new GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia() *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia {
	this := GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia{}
	return &this
}

// NewGetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMediaWithDefaults instantiates a new GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMediaWithDefaults() *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia {
	this := GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia{}
	return &this
}

// GetRj45 returns the Rj45 field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia) GetRj45() GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMediaRj45 {
	if o == nil || IsNil(o.Rj45) {
		var ret GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMediaRj45
		return ret
	}
	return *o.Rj45
}

// GetRj45Ok returns a tuple with the Rj45 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia) GetRj45Ok() (*GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMediaRj45, bool) {
	if o == nil || IsNil(o.Rj45) {
		return nil, false
	}
	return o.Rj45, true
}

// HasRj45 returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia) HasRj45() bool {
	if o != nil && !IsNil(o.Rj45) {
		return true
	}

	return false
}

// SetRj45 gets a reference to the given GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMediaRj45 and assigns it to the Rj45 field.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia) SetRj45(v GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMediaRj45) {
	o.Rj45 = &v
}

// GetSfp returns the Sfp field value if set, zero value otherwise.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia) GetSfp() GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMediaSfp {
	if o == nil || IsNil(o.Sfp) {
		var ret GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMediaSfp
		return ret
	}
	return *o.Sfp
}

// GetSfpOk returns a tuple with the Sfp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia) GetSfpOk() (*GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMediaSfp, bool) {
	if o == nil || IsNil(o.Sfp) {
		return nil, false
	}
	return o.Sfp, true
}

// HasSfp returns a boolean if a field has been set.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia) HasSfp() bool {
	if o != nil && !IsNil(o.Sfp) {
		return true
	}

	return false
}

// SetSfp gets a reference to the given GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMediaSfp and assigns it to the Sfp field.
func (o *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia) SetSfp(v GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMediaSfp) {
	o.Sfp = &v
}

func (o GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rj45) {
		toSerialize["rj45"] = o.Rj45
	}
	if !IsNil(o.Sfp) {
		toSerialize["sfp"] = o.Sfp
	}
	return toSerialize, nil
}

type NullableGetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia struct {
	value *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia
	isSet bool
}

func (v NullableGetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia) Get() *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia {
	return v.value
}

func (v *NullableGetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia) Set(val *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia(val *GetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia) *NullableGetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia {
	return &NullableGetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia{value: val, isSet: true}
}

func (v NullableGetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSwitchPortsOverview200ResponseCountsByStatusInactiveByMedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


