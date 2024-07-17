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

// checks if the GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf{}

// GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf struct for GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf
type GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf struct {
	// comparative operator used to filter for all values less than the given value
	Lt *time.Time `json:"lt,omitempty"`
	// comparative operator used to filter for all values greater than the given value
	Gt *time.Time `json:"gt,omitempty"`
	// comparative operator used to filter for all values less than or equal to the given value
	Lte *time.Time `json:"lte,omitempty"`
	// comparative operator used to filter for all values greater than or equal to the given value
	Gte *time.Time `json:"gte,omitempty"`
	// comparative operator used to filter for all values not equal to the given value
	Neq *time.Time `json:"neq,omitempty"`
}

// NewGetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf instantiates a new GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf() *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf {
	this := GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf{}
	return &this
}

// NewGetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOfWithDefaults instantiates a new GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOfWithDefaults() *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf {
	this := GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf{}
	return &this
}

// GetLt returns the Lt field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) GetLt() time.Time {
	if o == nil || IsNil(o.Lt) {
		var ret time.Time
		return ret
	}
	return *o.Lt
}

// GetLtOk returns a tuple with the Lt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) GetLtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Lt) {
		return nil, false
	}
	return o.Lt, true
}

// HasLt returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) HasLt() bool {
	if o != nil && !IsNil(o.Lt) {
		return true
	}

	return false
}

// SetLt gets a reference to the given time.Time and assigns it to the Lt field.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) SetLt(v time.Time) {
	o.Lt = &v
}

// GetGt returns the Gt field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) GetGt() time.Time {
	if o == nil || IsNil(o.Gt) {
		var ret time.Time
		return ret
	}
	return *o.Gt
}

// GetGtOk returns a tuple with the Gt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) GetGtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Gt) {
		return nil, false
	}
	return o.Gt, true
}

// HasGt returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) HasGt() bool {
	if o != nil && !IsNil(o.Gt) {
		return true
	}

	return false
}

// SetGt gets a reference to the given time.Time and assigns it to the Gt field.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) SetGt(v time.Time) {
	o.Gt = &v
}

// GetLte returns the Lte field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) GetLte() time.Time {
	if o == nil || IsNil(o.Lte) {
		var ret time.Time
		return ret
	}
	return *o.Lte
}

// GetLteOk returns a tuple with the Lte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) GetLteOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Lte) {
		return nil, false
	}
	return o.Lte, true
}

// HasLte returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) HasLte() bool {
	if o != nil && !IsNil(o.Lte) {
		return true
	}

	return false
}

// SetLte gets a reference to the given time.Time and assigns it to the Lte field.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) SetLte(v time.Time) {
	o.Lte = &v
}

// GetGte returns the Gte field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) GetGte() time.Time {
	if o == nil || IsNil(o.Gte) {
		var ret time.Time
		return ret
	}
	return *o.Gte
}

// GetGteOk returns a tuple with the Gte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) GetGteOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Gte) {
		return nil, false
	}
	return o.Gte, true
}

// HasGte returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) HasGte() bool {
	if o != nil && !IsNil(o.Gte) {
		return true
	}

	return false
}

// SetGte gets a reference to the given time.Time and assigns it to the Gte field.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) SetGte(v time.Time) {
	o.Gte = &v
}

// GetNeq returns the Neq field value if set, zero value otherwise.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) GetNeq() time.Time {
	if o == nil || IsNil(o.Neq) {
		var ret time.Time
		return ret
	}
	return *o.Neq
}

// GetNeqOk returns a tuple with the Neq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) GetNeqOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Neq) {
		return nil, false
	}
	return o.Neq, true
}

// HasNeq returns a boolean if a field has been set.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) HasNeq() bool {
	if o != nil && !IsNil(o.Neq) {
		return true
	}

	return false
}

// SetNeq gets a reference to the given time.Time and assigns it to the Neq field.
func (o *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) SetNeq(v time.Time) {
	o.Neq = &v
}

func (o GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lt) {
		toSerialize["lt"] = o.Lt
	}
	if !IsNil(o.Gt) {
		toSerialize["gt"] = o.Gt
	}
	if !IsNil(o.Lte) {
		toSerialize["lte"] = o.Lte
	}
	if !IsNil(o.Gte) {
		toSerialize["gte"] = o.Gte
	}
	if !IsNil(o.Neq) {
		toSerialize["neq"] = o.Neq
	}
	return toSerialize, nil
}

type NullableGetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf struct {
	value *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf
	isSet bool
}

func (v NullableGetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) Get() *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf {
	return v.value
}

func (v *NullableGetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) Set(val *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf(val *GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) *NullableGetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf {
	return &NullableGetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf{value: val, isSet: true}
}

func (v NullableGetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


