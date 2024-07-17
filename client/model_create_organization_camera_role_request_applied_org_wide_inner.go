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

// checks if the CreateOrganizationCameraRoleRequestAppliedOrgWideInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationCameraRoleRequestAppliedOrgWideInner{}

// CreateOrganizationCameraRoleRequestAppliedOrgWideInner struct for CreateOrganizationCameraRoleRequestAppliedOrgWideInner
type CreateOrganizationCameraRoleRequestAppliedOrgWideInner struct {
	// Permission scope id
	PermissionScopeId string `json:"permissionScopeId"`
}

type _CreateOrganizationCameraRoleRequestAppliedOrgWideInner CreateOrganizationCameraRoleRequestAppliedOrgWideInner

// NewCreateOrganizationCameraRoleRequestAppliedOrgWideInner instantiates a new CreateOrganizationCameraRoleRequestAppliedOrgWideInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationCameraRoleRequestAppliedOrgWideInner(permissionScopeId string) *CreateOrganizationCameraRoleRequestAppliedOrgWideInner {
	this := CreateOrganizationCameraRoleRequestAppliedOrgWideInner{}
	this.PermissionScopeId = permissionScopeId
	return &this
}

// NewCreateOrganizationCameraRoleRequestAppliedOrgWideInnerWithDefaults instantiates a new CreateOrganizationCameraRoleRequestAppliedOrgWideInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationCameraRoleRequestAppliedOrgWideInnerWithDefaults() *CreateOrganizationCameraRoleRequestAppliedOrgWideInner {
	this := CreateOrganizationCameraRoleRequestAppliedOrgWideInner{}
	return &this
}

// GetPermissionScopeId returns the PermissionScopeId field value
func (o *CreateOrganizationCameraRoleRequestAppliedOrgWideInner) GetPermissionScopeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PermissionScopeId
}

// GetPermissionScopeIdOk returns a tuple with the PermissionScopeId field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationCameraRoleRequestAppliedOrgWideInner) GetPermissionScopeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PermissionScopeId, true
}

// SetPermissionScopeId sets field value
func (o *CreateOrganizationCameraRoleRequestAppliedOrgWideInner) SetPermissionScopeId(v string) {
	o.PermissionScopeId = v
}

func (o CreateOrganizationCameraRoleRequestAppliedOrgWideInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationCameraRoleRequestAppliedOrgWideInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permissionScopeId"] = o.PermissionScopeId
	return toSerialize, nil
}

func (o *CreateOrganizationCameraRoleRequestAppliedOrgWideInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"permissionScopeId",
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

	varCreateOrganizationCameraRoleRequestAppliedOrgWideInner := _CreateOrganizationCameraRoleRequestAppliedOrgWideInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrganizationCameraRoleRequestAppliedOrgWideInner)

	if err != nil {
		return err
	}

	*o = CreateOrganizationCameraRoleRequestAppliedOrgWideInner(varCreateOrganizationCameraRoleRequestAppliedOrgWideInner)

	return err
}

type NullableCreateOrganizationCameraRoleRequestAppliedOrgWideInner struct {
	value *CreateOrganizationCameraRoleRequestAppliedOrgWideInner
	isSet bool
}

func (v NullableCreateOrganizationCameraRoleRequestAppliedOrgWideInner) Get() *CreateOrganizationCameraRoleRequestAppliedOrgWideInner {
	return v.value
}

func (v *NullableCreateOrganizationCameraRoleRequestAppliedOrgWideInner) Set(val *CreateOrganizationCameraRoleRequestAppliedOrgWideInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationCameraRoleRequestAppliedOrgWideInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationCameraRoleRequestAppliedOrgWideInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationCameraRoleRequestAppliedOrgWideInner(val *CreateOrganizationCameraRoleRequestAppliedOrgWideInner) *NullableCreateOrganizationCameraRoleRequestAppliedOrgWideInner {
	return &NullableCreateOrganizationCameraRoleRequestAppliedOrgWideInner{value: val, isSet: true}
}

func (v NullableCreateOrganizationCameraRoleRequestAppliedOrgWideInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationCameraRoleRequestAppliedOrgWideInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


