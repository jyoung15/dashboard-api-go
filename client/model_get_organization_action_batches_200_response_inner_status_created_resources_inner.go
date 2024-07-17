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

// checks if the GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner{}

// GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner struct for GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner
type GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner struct {
	// ID of the created resource
	Id *string `json:"id,omitempty"`
	// URI, not including base, of the created resource
	Uri *string `json:"uri,omitempty"`
}

// NewGetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner instantiates a new GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner() *GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner {
	this := GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner{}
	return &this
}

// NewGetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInnerWithDefaults instantiates a new GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInnerWithDefaults() *GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner {
	this := GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner) SetId(v string) {
	o.Id = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner) SetUri(v string) {
	o.Uri = &v
}

func (o GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	return toSerialize, nil
}

type NullableGetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner struct {
	value *GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner
	isSet bool
}

func (v NullableGetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner) Get() *GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner {
	return v.value
}

func (v *NullableGetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner) Set(val *GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner(val *GetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner) *NullableGetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner {
	return &NullableGetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner{value: val, isSet: true}
}

func (v NullableGetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationActionBatches200ResponseInnerStatusCreatedResourcesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


