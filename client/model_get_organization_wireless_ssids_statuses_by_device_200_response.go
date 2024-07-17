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

// checks if the GetOrganizationWirelessSsidsStatusesByDevice200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWirelessSsidsStatusesByDevice200Response{}

// GetOrganizationWirelessSsidsStatusesByDevice200Response struct for GetOrganizationWirelessSsidsStatusesByDevice200Response
type GetOrganizationWirelessSsidsStatusesByDevice200Response struct {
	// The top-level propery containing all status data.
	Items []GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner `json:"items,omitempty"`
	Meta *GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta `json:"meta,omitempty"`
}

// NewGetOrganizationWirelessSsidsStatusesByDevice200Response instantiates a new GetOrganizationWirelessSsidsStatusesByDevice200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWirelessSsidsStatusesByDevice200Response() *GetOrganizationWirelessSsidsStatusesByDevice200Response {
	this := GetOrganizationWirelessSsidsStatusesByDevice200Response{}
	return &this
}

// NewGetOrganizationWirelessSsidsStatusesByDevice200ResponseWithDefaults instantiates a new GetOrganizationWirelessSsidsStatusesByDevice200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWirelessSsidsStatusesByDevice200ResponseWithDefaults() *GetOrganizationWirelessSsidsStatusesByDevice200Response {
	this := GetOrganizationWirelessSsidsStatusesByDevice200Response{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *GetOrganizationWirelessSsidsStatusesByDevice200Response) GetItems() []GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner {
	if o == nil || IsNil(o.Items) {
		var ret []GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessSsidsStatusesByDevice200Response) GetItemsOk() ([]GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GetOrganizationWirelessSsidsStatusesByDevice200Response) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner and assigns it to the Items field.
func (o *GetOrganizationWirelessSsidsStatusesByDevice200Response) SetItems(v []GetOrganizationWirelessSsidsStatusesByDevice200ResponseItemsInner) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetOrganizationWirelessSsidsStatusesByDevice200Response) GetMeta() GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta {
	if o == nil || IsNil(o.Meta) {
		var ret GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessSsidsStatusesByDevice200Response) GetMetaOk() (*GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetOrganizationWirelessSsidsStatusesByDevice200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta and assigns it to the Meta field.
func (o *GetOrganizationWirelessSsidsStatusesByDevice200Response) SetMeta(v GetOrganizationWirelessSsidsStatusesByDevice200ResponseMeta) {
	o.Meta = &v
}

func (o GetOrganizationWirelessSsidsStatusesByDevice200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWirelessSsidsStatusesByDevice200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetOrganizationWirelessSsidsStatusesByDevice200Response struct {
	value *GetOrganizationWirelessSsidsStatusesByDevice200Response
	isSet bool
}

func (v NullableGetOrganizationWirelessSsidsStatusesByDevice200Response) Get() *GetOrganizationWirelessSsidsStatusesByDevice200Response {
	return v.value
}

func (v *NullableGetOrganizationWirelessSsidsStatusesByDevice200Response) Set(val *GetOrganizationWirelessSsidsStatusesByDevice200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWirelessSsidsStatusesByDevice200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWirelessSsidsStatusesByDevice200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWirelessSsidsStatusesByDevice200Response(val *GetOrganizationWirelessSsidsStatusesByDevice200Response) *NullableGetOrganizationWirelessSsidsStatusesByDevice200Response {
	return &NullableGetOrganizationWirelessSsidsStatusesByDevice200Response{value: val, isSet: true}
}

func (v NullableGetOrganizationWirelessSsidsStatusesByDevice200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWirelessSsidsStatusesByDevice200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

