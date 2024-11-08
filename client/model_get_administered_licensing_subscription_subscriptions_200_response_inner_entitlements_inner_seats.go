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

// checks if the GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats{}

// GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats Seat distribution
type GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats struct {
	// Number of seats in use
	Assigned *int32 `json:"assigned,omitempty"`
	// Number of seats available for use
	Available *int32 `json:"available,omitempty"`
	// Total number of seats provided by this subscription for this sku
	Limit *int32 `json:"limit,omitempty"`
}

// NewGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats instantiates a new GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats() *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats {
	this := GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats{}
	return &this
}

// NewGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeatsWithDefaults instantiates a new GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeatsWithDefaults() *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats {
	this := GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats{}
	return &this
}

// GetAssigned returns the Assigned field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats) GetAssigned() int32 {
	if o == nil || IsNil(o.Assigned) {
		var ret int32
		return ret
	}
	return *o.Assigned
}

// GetAssignedOk returns a tuple with the Assigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats) GetAssignedOk() (*int32, bool) {
	if o == nil || IsNil(o.Assigned) {
		return nil, false
	}
	return o.Assigned, true
}

// HasAssigned returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats) HasAssigned() bool {
	if o != nil && !IsNil(o.Assigned) {
		return true
	}

	return false
}

// SetAssigned gets a reference to the given int32 and assigns it to the Assigned field.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats) SetAssigned(v int32) {
	o.Assigned = &v
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats) GetAvailable() int32 {
	if o == nil || IsNil(o.Available) {
		var ret int32
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats) GetAvailableOk() (*int32, bool) {
	if o == nil || IsNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats) HasAvailable() bool {
	if o != nil && !IsNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given int32 and assigns it to the Available field.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats) SetAvailable(v int32) {
	o.Available = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats) SetLimit(v int32) {
	o.Limit = &v
}

func (o GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assigned) {
		toSerialize["assigned"] = o.Assigned
	}
	if !IsNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	return toSerialize, nil
}

type NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats struct {
	value *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats
	isSet bool
}

func (v NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats) Get() *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats {
	return v.value
}

func (v *NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats) Set(val *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats(val *GetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats) *NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats {
	return &NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats{value: val, isSet: true}
}

func (v NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAdministeredLicensingSubscriptionSubscriptions200ResponseInnerEntitlementsInnerSeats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


