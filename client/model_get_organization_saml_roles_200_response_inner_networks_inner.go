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

// checks if the GetOrganizationSamlRoles200ResponseInnerNetworksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSamlRoles200ResponseInnerNetworksInner{}

// GetOrganizationSamlRoles200ResponseInnerNetworksInner struct for GetOrganizationSamlRoles200ResponseInnerNetworksInner
type GetOrganizationSamlRoles200ResponseInnerNetworksInner struct {
	// The network ID
	Id *string `json:"id,omitempty"`
	// The privilege of the SAML administrator on the network
	Access *string `json:"access,omitempty"`
}

// NewGetOrganizationSamlRoles200ResponseInnerNetworksInner instantiates a new GetOrganizationSamlRoles200ResponseInnerNetworksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSamlRoles200ResponseInnerNetworksInner() *GetOrganizationSamlRoles200ResponseInnerNetworksInner {
	this := GetOrganizationSamlRoles200ResponseInnerNetworksInner{}
	return &this
}

// NewGetOrganizationSamlRoles200ResponseInnerNetworksInnerWithDefaults instantiates a new GetOrganizationSamlRoles200ResponseInnerNetworksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSamlRoles200ResponseInnerNetworksInnerWithDefaults() *GetOrganizationSamlRoles200ResponseInnerNetworksInner {
	this := GetOrganizationSamlRoles200ResponseInnerNetworksInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationSamlRoles200ResponseInnerNetworksInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSamlRoles200ResponseInnerNetworksInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationSamlRoles200ResponseInnerNetworksInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationSamlRoles200ResponseInnerNetworksInner) SetId(v string) {
	o.Id = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *GetOrganizationSamlRoles200ResponseInnerNetworksInner) GetAccess() string {
	if o == nil || IsNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSamlRoles200ResponseInnerNetworksInner) GetAccessOk() (*string, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *GetOrganizationSamlRoles200ResponseInnerNetworksInner) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *GetOrganizationSamlRoles200ResponseInnerNetworksInner) SetAccess(v string) {
	o.Access = &v
}

func (o GetOrganizationSamlRoles200ResponseInnerNetworksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSamlRoles200ResponseInnerNetworksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	return toSerialize, nil
}

type NullableGetOrganizationSamlRoles200ResponseInnerNetworksInner struct {
	value *GetOrganizationSamlRoles200ResponseInnerNetworksInner
	isSet bool
}

func (v NullableGetOrganizationSamlRoles200ResponseInnerNetworksInner) Get() *GetOrganizationSamlRoles200ResponseInnerNetworksInner {
	return v.value
}

func (v *NullableGetOrganizationSamlRoles200ResponseInnerNetworksInner) Set(val *GetOrganizationSamlRoles200ResponseInnerNetworksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSamlRoles200ResponseInnerNetworksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSamlRoles200ResponseInnerNetworksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSamlRoles200ResponseInnerNetworksInner(val *GetOrganizationSamlRoles200ResponseInnerNetworksInner) *NullableGetOrganizationSamlRoles200ResponseInnerNetworksInner {
	return &NullableGetOrganizationSamlRoles200ResponseInnerNetworksInner{value: val, isSet: true}
}

func (v NullableGetOrganizationSamlRoles200ResponseInnerNetworksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSamlRoles200ResponseInnerNetworksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


