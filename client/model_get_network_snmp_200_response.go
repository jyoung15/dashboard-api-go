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

// checks if the GetNetworkSnmp200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSnmp200Response{}

// GetNetworkSnmp200Response struct for GetNetworkSnmp200Response
type GetNetworkSnmp200Response struct {
	// The type of SNMP access. Can be one of 'none' (disabled), 'community' (V1/V2c), or 'users' (V3).
	Access *string `json:"access,omitempty"`
	// SNMP community string if access is 'community'.
	CommunityString *string `json:"communityString,omitempty"`
	// SNMP settings if access is 'users'.
	Users []GetNetworkSnmp200ResponseUsersInner `json:"users,omitempty"`
}

// NewGetNetworkSnmp200Response instantiates a new GetNetworkSnmp200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSnmp200Response() *GetNetworkSnmp200Response {
	this := GetNetworkSnmp200Response{}
	return &this
}

// NewGetNetworkSnmp200ResponseWithDefaults instantiates a new GetNetworkSnmp200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSnmp200ResponseWithDefaults() *GetNetworkSnmp200Response {
	this := GetNetworkSnmp200Response{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *GetNetworkSnmp200Response) GetAccess() string {
	if o == nil || IsNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSnmp200Response) GetAccessOk() (*string, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *GetNetworkSnmp200Response) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *GetNetworkSnmp200Response) SetAccess(v string) {
	o.Access = &v
}

// GetCommunityString returns the CommunityString field value if set, zero value otherwise.
func (o *GetNetworkSnmp200Response) GetCommunityString() string {
	if o == nil || IsNil(o.CommunityString) {
		var ret string
		return ret
	}
	return *o.CommunityString
}

// GetCommunityStringOk returns a tuple with the CommunityString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSnmp200Response) GetCommunityStringOk() (*string, bool) {
	if o == nil || IsNil(o.CommunityString) {
		return nil, false
	}
	return o.CommunityString, true
}

// HasCommunityString returns a boolean if a field has been set.
func (o *GetNetworkSnmp200Response) HasCommunityString() bool {
	if o != nil && !IsNil(o.CommunityString) {
		return true
	}

	return false
}

// SetCommunityString gets a reference to the given string and assigns it to the CommunityString field.
func (o *GetNetworkSnmp200Response) SetCommunityString(v string) {
	o.CommunityString = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *GetNetworkSnmp200Response) GetUsers() []GetNetworkSnmp200ResponseUsersInner {
	if o == nil || IsNil(o.Users) {
		var ret []GetNetworkSnmp200ResponseUsersInner
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSnmp200Response) GetUsersOk() ([]GetNetworkSnmp200ResponseUsersInner, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *GetNetworkSnmp200Response) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []GetNetworkSnmp200ResponseUsersInner and assigns it to the Users field.
func (o *GetNetworkSnmp200Response) SetUsers(v []GetNetworkSnmp200ResponseUsersInner) {
	o.Users = v
}

func (o GetNetworkSnmp200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSnmp200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if !IsNil(o.CommunityString) {
		toSerialize["communityString"] = o.CommunityString
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableGetNetworkSnmp200Response struct {
	value *GetNetworkSnmp200Response
	isSet bool
}

func (v NullableGetNetworkSnmp200Response) Get() *GetNetworkSnmp200Response {
	return v.value
}

func (v *NullableGetNetworkSnmp200Response) Set(val *GetNetworkSnmp200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSnmp200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSnmp200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSnmp200Response(val *GetNetworkSnmp200Response) *NullableGetNetworkSnmp200Response {
	return &NullableGetNetworkSnmp200Response{value: val, isSet: true}
}

func (v NullableGetNetworkSnmp200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSnmp200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


