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

// checks if the GetOrganizationWirelessAirMarshalSettingsByNetwork200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationWirelessAirMarshalSettingsByNetwork200Response{}

// GetOrganizationWirelessAirMarshalSettingsByNetwork200Response struct for GetOrganizationWirelessAirMarshalSettingsByNetwork200Response
type GetOrganizationWirelessAirMarshalSettingsByNetwork200Response struct {
	// List of settings
	Items []UpdateNetworkWirelessAirMarshalSettings200Response `json:"items,omitempty"`
	Meta *GetOrganizationWirelessAirMarshalSettingsByNetwork200ResponseMeta `json:"meta,omitempty"`
}

// NewGetOrganizationWirelessAirMarshalSettingsByNetwork200Response instantiates a new GetOrganizationWirelessAirMarshalSettingsByNetwork200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationWirelessAirMarshalSettingsByNetwork200Response() *GetOrganizationWirelessAirMarshalSettingsByNetwork200Response {
	this := GetOrganizationWirelessAirMarshalSettingsByNetwork200Response{}
	return &this
}

// NewGetOrganizationWirelessAirMarshalSettingsByNetwork200ResponseWithDefaults instantiates a new GetOrganizationWirelessAirMarshalSettingsByNetwork200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationWirelessAirMarshalSettingsByNetwork200ResponseWithDefaults() *GetOrganizationWirelessAirMarshalSettingsByNetwork200Response {
	this := GetOrganizationWirelessAirMarshalSettingsByNetwork200Response{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *GetOrganizationWirelessAirMarshalSettingsByNetwork200Response) GetItems() []UpdateNetworkWirelessAirMarshalSettings200Response {
	if o == nil || IsNil(o.Items) {
		var ret []UpdateNetworkWirelessAirMarshalSettings200Response
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessAirMarshalSettingsByNetwork200Response) GetItemsOk() ([]UpdateNetworkWirelessAirMarshalSettings200Response, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GetOrganizationWirelessAirMarshalSettingsByNetwork200Response) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []UpdateNetworkWirelessAirMarshalSettings200Response and assigns it to the Items field.
func (o *GetOrganizationWirelessAirMarshalSettingsByNetwork200Response) SetItems(v []UpdateNetworkWirelessAirMarshalSettings200Response) {
	o.Items = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetOrganizationWirelessAirMarshalSettingsByNetwork200Response) GetMeta() GetOrganizationWirelessAirMarshalSettingsByNetwork200ResponseMeta {
	if o == nil || IsNil(o.Meta) {
		var ret GetOrganizationWirelessAirMarshalSettingsByNetwork200ResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationWirelessAirMarshalSettingsByNetwork200Response) GetMetaOk() (*GetOrganizationWirelessAirMarshalSettingsByNetwork200ResponseMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetOrganizationWirelessAirMarshalSettingsByNetwork200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given GetOrganizationWirelessAirMarshalSettingsByNetwork200ResponseMeta and assigns it to the Meta field.
func (o *GetOrganizationWirelessAirMarshalSettingsByNetwork200Response) SetMeta(v GetOrganizationWirelessAirMarshalSettingsByNetwork200ResponseMeta) {
	o.Meta = &v
}

func (o GetOrganizationWirelessAirMarshalSettingsByNetwork200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationWirelessAirMarshalSettingsByNetwork200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetOrganizationWirelessAirMarshalSettingsByNetwork200Response struct {
	value *GetOrganizationWirelessAirMarshalSettingsByNetwork200Response
	isSet bool
}

func (v NullableGetOrganizationWirelessAirMarshalSettingsByNetwork200Response) Get() *GetOrganizationWirelessAirMarshalSettingsByNetwork200Response {
	return v.value
}

func (v *NullableGetOrganizationWirelessAirMarshalSettingsByNetwork200Response) Set(val *GetOrganizationWirelessAirMarshalSettingsByNetwork200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationWirelessAirMarshalSettingsByNetwork200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationWirelessAirMarshalSettingsByNetwork200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationWirelessAirMarshalSettingsByNetwork200Response(val *GetOrganizationWirelessAirMarshalSettingsByNetwork200Response) *NullableGetOrganizationWirelessAirMarshalSettingsByNetwork200Response {
	return &NullableGetOrganizationWirelessAirMarshalSettingsByNetwork200Response{value: val, isSet: true}
}

func (v NullableGetOrganizationWirelessAirMarshalSettingsByNetwork200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationWirelessAirMarshalSettingsByNetwork200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


