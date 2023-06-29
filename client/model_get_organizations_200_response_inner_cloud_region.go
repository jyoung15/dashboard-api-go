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

// checks if the GetOrganizations200ResponseInnerCloudRegion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizations200ResponseInnerCloudRegion{}

// GetOrganizations200ResponseInnerCloudRegion Region info
type GetOrganizations200ResponseInnerCloudRegion struct {
	// Name of region
	Name *string `json:"name,omitempty"`
}

// NewGetOrganizations200ResponseInnerCloudRegion instantiates a new GetOrganizations200ResponseInnerCloudRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizations200ResponseInnerCloudRegion() *GetOrganizations200ResponseInnerCloudRegion {
	this := GetOrganizations200ResponseInnerCloudRegion{}
	return &this
}

// NewGetOrganizations200ResponseInnerCloudRegionWithDefaults instantiates a new GetOrganizations200ResponseInnerCloudRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizations200ResponseInnerCloudRegionWithDefaults() *GetOrganizations200ResponseInnerCloudRegion {
	this := GetOrganizations200ResponseInnerCloudRegion{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizations200ResponseInnerCloudRegion) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizations200ResponseInnerCloudRegion) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizations200ResponseInnerCloudRegion) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizations200ResponseInnerCloudRegion) SetName(v string) {
	o.Name = &v
}

func (o GetOrganizations200ResponseInnerCloudRegion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizations200ResponseInnerCloudRegion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGetOrganizations200ResponseInnerCloudRegion struct {
	value *GetOrganizations200ResponseInnerCloudRegion
	isSet bool
}

func (v NullableGetOrganizations200ResponseInnerCloudRegion) Get() *GetOrganizations200ResponseInnerCloudRegion {
	return v.value
}

func (v *NullableGetOrganizations200ResponseInnerCloudRegion) Set(val *GetOrganizations200ResponseInnerCloudRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizations200ResponseInnerCloudRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizations200ResponseInnerCloudRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizations200ResponseInnerCloudRegion(val *GetOrganizations200ResponseInnerCloudRegion) *NullableGetOrganizations200ResponseInnerCloudRegion {
	return &NullableGetOrganizations200ResponseInnerCloudRegion{value: val, isSet: true}
}

func (v NullableGetOrganizations200ResponseInnerCloudRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizations200ResponseInnerCloudRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


