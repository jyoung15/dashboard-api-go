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

// checks if the UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner{}

// UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner struct for UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner
type UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner struct {
	// Network tag
	Tag *string `json:"tag,omitempty"`
	// Network id
	Id *string `json:"id,omitempty"`
	// Permission scope id
	PermissionScopeId string `json:"permissionScopeId"`
}

type _UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner

// NewUpdateOrganizationCameraRoleRequestAppliedOnNetworksInner instantiates a new UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationCameraRoleRequestAppliedOnNetworksInner(permissionScopeId string) *UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner {
	this := UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner{}
	this.PermissionScopeId = permissionScopeId
	return &this
}

// NewUpdateOrganizationCameraRoleRequestAppliedOnNetworksInnerWithDefaults instantiates a new UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationCameraRoleRequestAppliedOnNetworksInnerWithDefaults() *UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner {
	this := UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner) SetTag(v string) {
	o.Tag = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner) SetId(v string) {
	o.Id = &v
}

// GetPermissionScopeId returns the PermissionScopeId field value
func (o *UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner) GetPermissionScopeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PermissionScopeId
}

// GetPermissionScopeIdOk returns a tuple with the PermissionScopeId field value
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner) GetPermissionScopeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PermissionScopeId, true
}

// SetPermissionScopeId sets field value
func (o *UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner) SetPermissionScopeId(v string) {
	o.PermissionScopeId = v
}

func (o UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["permissionScopeId"] = o.PermissionScopeId
	return toSerialize, nil
}

func (o *UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner) UnmarshalJSON(data []byte) (err error) {
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

	varUpdateOrganizationCameraRoleRequestAppliedOnNetworksInner := _UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateOrganizationCameraRoleRequestAppliedOnNetworksInner)

	if err != nil {
		return err
	}

	*o = UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner(varUpdateOrganizationCameraRoleRequestAppliedOnNetworksInner)

	return err
}

type NullableUpdateOrganizationCameraRoleRequestAppliedOnNetworksInner struct {
	value *UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner
	isSet bool
}

func (v NullableUpdateOrganizationCameraRoleRequestAppliedOnNetworksInner) Get() *UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner {
	return v.value
}

func (v *NullableUpdateOrganizationCameraRoleRequestAppliedOnNetworksInner) Set(val *UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationCameraRoleRequestAppliedOnNetworksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationCameraRoleRequestAppliedOnNetworksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationCameraRoleRequestAppliedOnNetworksInner(val *UpdateOrganizationCameraRoleRequestAppliedOnNetworksInner) *NullableUpdateOrganizationCameraRoleRequestAppliedOnNetworksInner {
	return &NullableUpdateOrganizationCameraRoleRequestAppliedOnNetworksInner{value: val, isSet: true}
}

func (v NullableUpdateOrganizationCameraRoleRequestAppliedOnNetworksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationCameraRoleRequestAppliedOnNetworksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

