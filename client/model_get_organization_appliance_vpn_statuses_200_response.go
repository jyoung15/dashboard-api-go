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

// checks if the GetOrganizationApplianceVpnStatuses200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationApplianceVpnStatuses200Response{}

// GetOrganizationApplianceVpnStatuses200Response struct for GetOrganizationApplianceVpnStatuses200Response
type GetOrganizationApplianceVpnStatuses200Response struct {
	// The list of VPN Status for networks
	Vpnstatusentities []GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner `json:"vpnstatusentities,omitempty"`
}

// NewGetOrganizationApplianceVpnStatuses200Response instantiates a new GetOrganizationApplianceVpnStatuses200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationApplianceVpnStatuses200Response() *GetOrganizationApplianceVpnStatuses200Response {
	this := GetOrganizationApplianceVpnStatuses200Response{}
	return &this
}

// NewGetOrganizationApplianceVpnStatuses200ResponseWithDefaults instantiates a new GetOrganizationApplianceVpnStatuses200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationApplianceVpnStatuses200ResponseWithDefaults() *GetOrganizationApplianceVpnStatuses200Response {
	this := GetOrganizationApplianceVpnStatuses200Response{}
	return &this
}

// GetVpnstatusentities returns the Vpnstatusentities field value if set, zero value otherwise.
func (o *GetOrganizationApplianceVpnStatuses200Response) GetVpnstatusentities() []GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner {
	if o == nil || IsNil(o.Vpnstatusentities) {
		var ret []GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner
		return ret
	}
	return o.Vpnstatusentities
}

// GetVpnstatusentitiesOk returns a tuple with the Vpnstatusentities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceVpnStatuses200Response) GetVpnstatusentitiesOk() ([]GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner, bool) {
	if o == nil || IsNil(o.Vpnstatusentities) {
		return nil, false
	}
	return o.Vpnstatusentities, true
}

// HasVpnstatusentities returns a boolean if a field has been set.
func (o *GetOrganizationApplianceVpnStatuses200Response) HasVpnstatusentities() bool {
	if o != nil && !IsNil(o.Vpnstatusentities) {
		return true
	}

	return false
}

// SetVpnstatusentities gets a reference to the given []GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner and assigns it to the Vpnstatusentities field.
func (o *GetOrganizationApplianceVpnStatuses200Response) SetVpnstatusentities(v []GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) {
	o.Vpnstatusentities = v
}

func (o GetOrganizationApplianceVpnStatuses200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationApplianceVpnStatuses200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Vpnstatusentities) {
		toSerialize["vpnstatusentities"] = o.Vpnstatusentities
	}
	return toSerialize, nil
}

type NullableGetOrganizationApplianceVpnStatuses200Response struct {
	value *GetOrganizationApplianceVpnStatuses200Response
	isSet bool
}

func (v NullableGetOrganizationApplianceVpnStatuses200Response) Get() *GetOrganizationApplianceVpnStatuses200Response {
	return v.value
}

func (v *NullableGetOrganizationApplianceVpnStatuses200Response) Set(val *GetOrganizationApplianceVpnStatuses200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationApplianceVpnStatuses200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationApplianceVpnStatuses200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationApplianceVpnStatuses200Response(val *GetOrganizationApplianceVpnStatuses200Response) *NullableGetOrganizationApplianceVpnStatuses200Response {
	return &NullableGetOrganizationApplianceVpnStatuses200Response{value: val, isSet: true}
}

func (v NullableGetOrganizationApplianceVpnStatuses200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationApplianceVpnStatuses200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


