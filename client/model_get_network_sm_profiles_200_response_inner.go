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

// checks if the GetNetworkSmProfiles200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSmProfiles200ResponseInner{}

// GetNetworkSmProfiles200ResponseInner struct for GetNetworkSmProfiles200ResponseInner
type GetNetworkSmProfiles200ResponseInner struct {
	// ID of a profile.
	Id *string `json:"id,omitempty"`
	// Name of a profile.
	Name *string `json:"name,omitempty"`
	// Description of a profile.
	Description *string `json:"description,omitempty"`
	// Scope of a profile.
	Scope *string `json:"scope,omitempty"`
	// Tags of a profile.
	Tags []string `json:"tags,omitempty"`
	// Payloads in the profile.
	PayloadTypes []string `json:"payloadTypes,omitempty"`
}

// NewGetNetworkSmProfiles200ResponseInner instantiates a new GetNetworkSmProfiles200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSmProfiles200ResponseInner() *GetNetworkSmProfiles200ResponseInner {
	this := GetNetworkSmProfiles200ResponseInner{}
	return &this
}

// NewGetNetworkSmProfiles200ResponseInnerWithDefaults instantiates a new GetNetworkSmProfiles200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSmProfiles200ResponseInnerWithDefaults() *GetNetworkSmProfiles200ResponseInner {
	this := GetNetworkSmProfiles200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkSmProfiles200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmProfiles200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkSmProfiles200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkSmProfiles200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkSmProfiles200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmProfiles200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkSmProfiles200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkSmProfiles200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetNetworkSmProfiles200ResponseInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmProfiles200ResponseInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetNetworkSmProfiles200ResponseInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetNetworkSmProfiles200ResponseInner) SetDescription(v string) {
	o.Description = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *GetNetworkSmProfiles200ResponseInner) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmProfiles200ResponseInner) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *GetNetworkSmProfiles200ResponseInner) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *GetNetworkSmProfiles200ResponseInner) SetScope(v string) {
	o.Scope = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GetNetworkSmProfiles200ResponseInner) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmProfiles200ResponseInner) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GetNetworkSmProfiles200ResponseInner) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GetNetworkSmProfiles200ResponseInner) SetTags(v []string) {
	o.Tags = v
}

// GetPayloadTypes returns the PayloadTypes field value if set, zero value otherwise.
func (o *GetNetworkSmProfiles200ResponseInner) GetPayloadTypes() []string {
	if o == nil || IsNil(o.PayloadTypes) {
		var ret []string
		return ret
	}
	return o.PayloadTypes
}

// GetPayloadTypesOk returns a tuple with the PayloadTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSmProfiles200ResponseInner) GetPayloadTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.PayloadTypes) {
		return nil, false
	}
	return o.PayloadTypes, true
}

// HasPayloadTypes returns a boolean if a field has been set.
func (o *GetNetworkSmProfiles200ResponseInner) HasPayloadTypes() bool {
	if o != nil && !IsNil(o.PayloadTypes) {
		return true
	}

	return false
}

// SetPayloadTypes gets a reference to the given []string and assigns it to the PayloadTypes field.
func (o *GetNetworkSmProfiles200ResponseInner) SetPayloadTypes(v []string) {
	o.PayloadTypes = v
}

func (o GetNetworkSmProfiles200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSmProfiles200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.PayloadTypes) {
		toSerialize["payloadTypes"] = o.PayloadTypes
	}
	return toSerialize, nil
}

type NullableGetNetworkSmProfiles200ResponseInner struct {
	value *GetNetworkSmProfiles200ResponseInner
	isSet bool
}

func (v NullableGetNetworkSmProfiles200ResponseInner) Get() *GetNetworkSmProfiles200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkSmProfiles200ResponseInner) Set(val *GetNetworkSmProfiles200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSmProfiles200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSmProfiles200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSmProfiles200ResponseInner(val *GetNetworkSmProfiles200ResponseInner) *NullableGetNetworkSmProfiles200ResponseInner {
	return &NullableGetNetworkSmProfiles200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkSmProfiles200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSmProfiles200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


