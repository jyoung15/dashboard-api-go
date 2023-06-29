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

// checks if the CreateOrganizationAdminRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationAdminRequest{}

// CreateOrganizationAdminRequest struct for CreateOrganizationAdminRequest
type CreateOrganizationAdminRequest struct {
	// The email of the dashboard administrator. This attribute can not be updated.
	Email string `json:"email"`
	// The name of the dashboard administrator
	Name string `json:"name"`
	// The privilege of the dashboard administrator on the organization. Can be one of 'full', 'read-only', 'enterprise' or 'none'
	OrgAccess string `json:"orgAccess"`
	// The list of tags that the dashboard administrator has privileges on
	Tags []CreateOrganizationAdminRequestTagsInner `json:"tags,omitempty"`
	// The list of networks that the dashboard administrator has privileges on
	Networks []CreateOrganizationAdminRequestNetworksInner `json:"networks,omitempty"`
	// The method of authentication the user will use to sign in to the Meraki dashboard. Can be one of 'Email' or 'Cisco SecureX Sign-On'. The default is Email authentication
	AuthenticationMethod *string `json:"authenticationMethod,omitempty"`
}

// NewCreateOrganizationAdminRequest instantiates a new CreateOrganizationAdminRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationAdminRequest(email string, name string, orgAccess string) *CreateOrganizationAdminRequest {
	this := CreateOrganizationAdminRequest{}
	this.Email = email
	this.Name = name
	this.OrgAccess = orgAccess
	return &this
}

// NewCreateOrganizationAdminRequestWithDefaults instantiates a new CreateOrganizationAdminRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationAdminRequestWithDefaults() *CreateOrganizationAdminRequest {
	this := CreateOrganizationAdminRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *CreateOrganizationAdminRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdminRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CreateOrganizationAdminRequest) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value
func (o *CreateOrganizationAdminRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdminRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrganizationAdminRequest) SetName(v string) {
	o.Name = v
}

// GetOrgAccess returns the OrgAccess field value
func (o *CreateOrganizationAdminRequest) GetOrgAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgAccess
}

// GetOrgAccessOk returns a tuple with the OrgAccess field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdminRequest) GetOrgAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgAccess, true
}

// SetOrgAccess sets field value
func (o *CreateOrganizationAdminRequest) SetOrgAccess(v string) {
	o.OrgAccess = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateOrganizationAdminRequest) GetTags() []CreateOrganizationAdminRequestTagsInner {
	if o == nil || IsNil(o.Tags) {
		var ret []CreateOrganizationAdminRequestTagsInner
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdminRequest) GetTagsOk() ([]CreateOrganizationAdminRequestTagsInner, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateOrganizationAdminRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []CreateOrganizationAdminRequestTagsInner and assigns it to the Tags field.
func (o *CreateOrganizationAdminRequest) SetTags(v []CreateOrganizationAdminRequestTagsInner) {
	o.Tags = v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *CreateOrganizationAdminRequest) GetNetworks() []CreateOrganizationAdminRequestNetworksInner {
	if o == nil || IsNil(o.Networks) {
		var ret []CreateOrganizationAdminRequestNetworksInner
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdminRequest) GetNetworksOk() ([]CreateOrganizationAdminRequestNetworksInner, bool) {
	if o == nil || IsNil(o.Networks) {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *CreateOrganizationAdminRequest) HasNetworks() bool {
	if o != nil && !IsNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []CreateOrganizationAdminRequestNetworksInner and assigns it to the Networks field.
func (o *CreateOrganizationAdminRequest) SetNetworks(v []CreateOrganizationAdminRequestNetworksInner) {
	o.Networks = v
}

// GetAuthenticationMethod returns the AuthenticationMethod field value if set, zero value otherwise.
func (o *CreateOrganizationAdminRequest) GetAuthenticationMethod() string {
	if o == nil || IsNil(o.AuthenticationMethod) {
		var ret string
		return ret
	}
	return *o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrganizationAdminRequest) GetAuthenticationMethodOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationMethod) {
		return nil, false
	}
	return o.AuthenticationMethod, true
}

// HasAuthenticationMethod returns a boolean if a field has been set.
func (o *CreateOrganizationAdminRequest) HasAuthenticationMethod() bool {
	if o != nil && !IsNil(o.AuthenticationMethod) {
		return true
	}

	return false
}

// SetAuthenticationMethod gets a reference to the given string and assigns it to the AuthenticationMethod field.
func (o *CreateOrganizationAdminRequest) SetAuthenticationMethod(v string) {
	o.AuthenticationMethod = &v
}

func (o CreateOrganizationAdminRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationAdminRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["name"] = o.Name
	toSerialize["orgAccess"] = o.OrgAccess
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Networks) {
		toSerialize["networks"] = o.Networks
	}
	if !IsNil(o.AuthenticationMethod) {
		toSerialize["authenticationMethod"] = o.AuthenticationMethod
	}
	return toSerialize, nil
}

type NullableCreateOrganizationAdminRequest struct {
	value *CreateOrganizationAdminRequest
	isSet bool
}

func (v NullableCreateOrganizationAdminRequest) Get() *CreateOrganizationAdminRequest {
	return v.value
}

func (v *NullableCreateOrganizationAdminRequest) Set(val *CreateOrganizationAdminRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationAdminRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationAdminRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationAdminRequest(val *CreateOrganizationAdminRequest) *NullableCreateOrganizationAdminRequest {
	return &NullableCreateOrganizationAdminRequest{value: val, isSet: true}
}

func (v NullableCreateOrganizationAdminRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationAdminRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


