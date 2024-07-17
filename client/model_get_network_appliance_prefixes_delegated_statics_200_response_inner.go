/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner{}

// GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner struct for GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner
type GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner struct {
	// Static delegated prefix id.
	StaticDelegatedPrefixId *string `json:"staticDelegatedPrefixId,omitempty"`
	// IPv6 prefix/prefix length.
	Prefix *string `json:"prefix,omitempty"`
	Origin *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin `json:"origin,omitempty"`
	// Identifying description for the prefix.
	Description *string `json:"description,omitempty"`
	// Prefix creation time.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Prefix Updated time.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewGetNetworkAppliancePrefixesDelegatedStatics200ResponseInner instantiates a new GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkAppliancePrefixesDelegatedStatics200ResponseInner() *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner {
	this := GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner{}
	return &this
}

// NewGetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerWithDefaults instantiates a new GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerWithDefaults() *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner {
	this := GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner{}
	return &this
}

// GetStaticDelegatedPrefixId returns the StaticDelegatedPrefixId field value if set, zero value otherwise.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetStaticDelegatedPrefixId() string {
	if o == nil || IsNil(o.StaticDelegatedPrefixId) {
		var ret string
		return ret
	}
	return *o.StaticDelegatedPrefixId
}

// GetStaticDelegatedPrefixIdOk returns a tuple with the StaticDelegatedPrefixId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetStaticDelegatedPrefixIdOk() (*string, bool) {
	if o == nil || IsNil(o.StaticDelegatedPrefixId) {
		return nil, false
	}
	return o.StaticDelegatedPrefixId, true
}

// HasStaticDelegatedPrefixId returns a boolean if a field has been set.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) HasStaticDelegatedPrefixId() bool {
	if o != nil && !IsNil(o.StaticDelegatedPrefixId) {
		return true
	}

	return false
}

// SetStaticDelegatedPrefixId gets a reference to the given string and assigns it to the StaticDelegatedPrefixId field.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) SetStaticDelegatedPrefixId(v string) {
	o.StaticDelegatedPrefixId = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) SetPrefix(v string) {
	o.Prefix = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetOrigin() GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin {
	if o == nil || IsNil(o.Origin) {
		var ret GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetOriginOk() (*GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin and assigns it to the Origin field.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) SetOrigin(v GetNetworkAppliancePrefixesDelegatedStatics200ResponseInnerOrigin) {
	o.Origin = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StaticDelegatedPrefixId) {
		toSerialize["staticDelegatedPrefixId"] = o.StaticDelegatedPrefixId
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableGetNetworkAppliancePrefixesDelegatedStatics200ResponseInner struct {
	value *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner
	isSet bool
}

func (v NullableGetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) Get() *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) Set(val *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkAppliancePrefixesDelegatedStatics200ResponseInner(val *GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) *NullableGetNetworkAppliancePrefixesDelegatedStatics200ResponseInner {
	return &NullableGetNetworkAppliancePrefixesDelegatedStatics200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkAppliancePrefixesDelegatedStatics200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


