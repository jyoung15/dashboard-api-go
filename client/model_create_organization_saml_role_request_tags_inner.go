/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateOrganizationSamlRoleRequestTagsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationSamlRoleRequestTagsInner{}

// CreateOrganizationSamlRoleRequestTagsInner struct for CreateOrganizationSamlRoleRequestTagsInner
type CreateOrganizationSamlRoleRequestTagsInner struct {
	// The name of the tag
	Tag string `json:"tag"`
	// The privilege of the SAML administrator on the tag. Can be one of 'full', 'read-only', 'guest-ambassador' or 'monitor-only'
	Access string `json:"access"`
}

type _CreateOrganizationSamlRoleRequestTagsInner CreateOrganizationSamlRoleRequestTagsInner

// NewCreateOrganizationSamlRoleRequestTagsInner instantiates a new CreateOrganizationSamlRoleRequestTagsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationSamlRoleRequestTagsInner(tag string, access string) *CreateOrganizationSamlRoleRequestTagsInner {
	this := CreateOrganizationSamlRoleRequestTagsInner{}
	this.Tag = tag
	this.Access = access
	return &this
}

// NewCreateOrganizationSamlRoleRequestTagsInnerWithDefaults instantiates a new CreateOrganizationSamlRoleRequestTagsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationSamlRoleRequestTagsInnerWithDefaults() *CreateOrganizationSamlRoleRequestTagsInner {
	this := CreateOrganizationSamlRoleRequestTagsInner{}
	return &this
}

// GetTag returns the Tag field value
func (o *CreateOrganizationSamlRoleRequestTagsInner) GetTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationSamlRoleRequestTagsInner) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tag, true
}

// SetTag sets field value
func (o *CreateOrganizationSamlRoleRequestTagsInner) SetTag(v string) {
	o.Tag = v
}

// GetAccess returns the Access field value
func (o *CreateOrganizationSamlRoleRequestTagsInner) GetAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationSamlRoleRequestTagsInner) GetAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *CreateOrganizationSamlRoleRequestTagsInner) SetAccess(v string) {
	o.Access = v
}

func (o CreateOrganizationSamlRoleRequestTagsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationSamlRoleRequestTagsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tag"] = o.Tag
	toSerialize["access"] = o.Access
	return toSerialize, nil
}

func (o *CreateOrganizationSamlRoleRequestTagsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tag",
		"access",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateOrganizationSamlRoleRequestTagsInner := _CreateOrganizationSamlRoleRequestTagsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrganizationSamlRoleRequestTagsInner)

	if err != nil {
		return err
	}

	*o = CreateOrganizationSamlRoleRequestTagsInner(varCreateOrganizationSamlRoleRequestTagsInner)

	return err
}

type NullableCreateOrganizationSamlRoleRequestTagsInner struct {
	value *CreateOrganizationSamlRoleRequestTagsInner
	isSet bool
}

func (v NullableCreateOrganizationSamlRoleRequestTagsInner) Get() *CreateOrganizationSamlRoleRequestTagsInner {
	return v.value
}

func (v *NullableCreateOrganizationSamlRoleRequestTagsInner) Set(val *CreateOrganizationSamlRoleRequestTagsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationSamlRoleRequestTagsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationSamlRoleRequestTagsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationSamlRoleRequestTagsInner(val *CreateOrganizationSamlRoleRequestTagsInner) *NullableCreateOrganizationSamlRoleRequestTagsInner {
	return &NullableCreateOrganizationSamlRoleRequestTagsInner{value: val, isSet: true}
}

func (v NullableCreateOrganizationSamlRoleRequestTagsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationSamlRoleRequestTagsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


