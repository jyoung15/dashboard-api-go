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

// checks if the GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts{}

// GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts Numeric breakdown of network, organizations, entitlement counts
type GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts struct {
	Seats *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCountsSeats `json:"seats,omitempty"`
	// Number of networks bound to this subscription
	Networks *int32 `json:"networks,omitempty"`
	// Number of organizations bound to this subscription
	Organizations *int32 `json:"organizations,omitempty"`
}

// NewGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts instantiates a new GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts() *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts {
	this := GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts{}
	return &this
}

// NewGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCountsWithDefaults instantiates a new GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCountsWithDefaults() *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts {
	this := GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts{}
	return &this
}

// GetSeats returns the Seats field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts) GetSeats() GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCountsSeats {
	if o == nil || IsNil(o.Seats) {
		var ret GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCountsSeats
		return ret
	}
	return *o.Seats
}

// GetSeatsOk returns a tuple with the Seats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts) GetSeatsOk() (*GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCountsSeats, bool) {
	if o == nil || IsNil(o.Seats) {
		return nil, false
	}
	return o.Seats, true
}

// HasSeats returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts) HasSeats() bool {
	if o != nil && !IsNil(o.Seats) {
		return true
	}

	return false
}

// SetSeats gets a reference to the given GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCountsSeats and assigns it to the Seats field.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts) SetSeats(v GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCountsSeats) {
	o.Seats = &v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts) GetNetworks() int32 {
	if o == nil || IsNil(o.Networks) {
		var ret int32
		return ret
	}
	return *o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts) GetNetworksOk() (*int32, bool) {
	if o == nil || IsNil(o.Networks) {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts) HasNetworks() bool {
	if o != nil && !IsNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given int32 and assigns it to the Networks field.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts) SetNetworks(v int32) {
	o.Networks = &v
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts) GetOrganizations() int32 {
	if o == nil || IsNil(o.Organizations) {
		var ret int32
		return ret
	}
	return *o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts) GetOrganizationsOk() (*int32, bool) {
	if o == nil || IsNil(o.Organizations) {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts) HasOrganizations() bool {
	if o != nil && !IsNil(o.Organizations) {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given int32 and assigns it to the Organizations field.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts) SetOrganizations(v int32) {
	o.Organizations = &v
}

func (o GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Seats) {
		toSerialize["seats"] = o.Seats
	}
	if !IsNil(o.Networks) {
		toSerialize["networks"] = o.Networks
	}
	if !IsNil(o.Organizations) {
		toSerialize["organizations"] = o.Organizations
	}
	return toSerialize, nil
}

type NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts struct {
	value *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts
	isSet bool
}

func (v NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts) Get() *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts {
	return v.value
}

func (v *NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts) Set(val *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts(val *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts) *NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts {
	return &NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts{value: val, isSet: true}
}

func (v NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


