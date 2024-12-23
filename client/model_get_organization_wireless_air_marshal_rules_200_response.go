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

// checks if the GetOrganizationWirelessAirMarshalRules200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWirelessAirMarshalRules200Response{}

// GetOrganizationWirelessAirMarshalRules200Response struct for GetOrganizationWirelessAirMarshalRules200Response
type GetOrganizationWirelessAirMarshalRules200Response struct {
	// List of rules
	Items []GetOrganizationWirelessAirMarshalRules200ResponseItemsInner `json:"items,omitempty"`
	Meta *GetOrganizationWirelessAirMarshalRules200ResponseMeta `json:"meta,omitempty"`
}

// NewGetOrganizationWirelessAirMarshalRules200Response instantiates a new GetOrganizationWirelessAirMarshalRules200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWirelessAirMarshalRules200Response() *GetOrganizationWirelessAirMarshalRules200Response {
	this := GetOrganizationWirelessAirMarshalRules200Response{}
	return &this
}

// NewGetOrganizationWirelessAirMarshalRules200ResponseWithDefaults instantiates a new GetOrganizationWirelessAirMarshalRules200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWirelessAirMarshalRules200ResponseWithDefaults() *GetOrganizationWirelessAirMarshalRules200Response {
	this := GetOrganizationWirelessAirMarshalRules200Response{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *GetOrganizationWirelessAirMarshalRules200Response) GetItems() []GetOrganizationWirelessAirMarshalRules200ResponseItemsInner {
	if o == nil || IsNil(o.Items) {
		var ret []GetOrganizationWirelessAirMarshalRules200ResponseItemsInner
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessAirMarshalRules200Response) GetItemsOk() ([]GetOrganizationWirelessAirMarshalRules200ResponseItemsInner, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GetOrganizationWirelessAirMarshalRules200Response) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []GetOrganizationWirelessAirMarshalRules200ResponseItemsInner and assigns it to the Items field.
func (o *GetOrganizationWirelessAirMarshalRules200Response) SetItems(v []GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetOrganizationWirelessAirMarshalRules200Response) GetMeta() GetOrganizationWirelessAirMarshalRules200ResponseMeta {
	if o == nil || IsNil(o.Meta) {
		var ret GetOrganizationWirelessAirMarshalRules200ResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessAirMarshalRules200Response) GetMetaOk() (*GetOrganizationWirelessAirMarshalRules200ResponseMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetOrganizationWirelessAirMarshalRules200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given GetOrganizationWirelessAirMarshalRules200ResponseMeta and assigns it to the Meta field.
func (o *GetOrganizationWirelessAirMarshalRules200Response) SetMeta(v GetOrganizationWirelessAirMarshalRules200ResponseMeta) {
	o.Meta = &v
}

func (o GetOrganizationWirelessAirMarshalRules200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWirelessAirMarshalRules200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetOrganizationWirelessAirMarshalRules200Response struct {
	value *GetOrganizationWirelessAirMarshalRules200Response
	isSet bool
}

func (v NullableGetOrganizationWirelessAirMarshalRules200Response) Get() *GetOrganizationWirelessAirMarshalRules200Response {
	return v.value
}

func (v *NullableGetOrganizationWirelessAirMarshalRules200Response) Set(val *GetOrganizationWirelessAirMarshalRules200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWirelessAirMarshalRules200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWirelessAirMarshalRules200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWirelessAirMarshalRules200Response(val *GetOrganizationWirelessAirMarshalRules200Response) *NullableGetOrganizationWirelessAirMarshalRules200Response {
	return &NullableGetOrganizationWirelessAirMarshalRules200Response{value: val, isSet: true}
}

func (v NullableGetOrganizationWirelessAirMarshalRules200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWirelessAirMarshalRules200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


