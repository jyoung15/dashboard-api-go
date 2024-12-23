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

// checks if the GetNetworkGroupPolicies200ResponseInnerSchedulingMonday type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkGroupPolicies200ResponseInnerSchedulingMonday{}

// GetNetworkGroupPolicies200ResponseInnerSchedulingMonday The schedule object for Monday.
type GetNetworkGroupPolicies200ResponseInnerSchedulingMonday struct {
	// Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	Active *bool `json:"active,omitempty"`
	// The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	From *string `json:"from,omitempty"`
	// The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
	To *string `json:"to,omitempty"`
}

// NewGetNetworkGroupPolicies200ResponseInnerSchedulingMonday instantiates a new GetNetworkGroupPolicies200ResponseInnerSchedulingMonday object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkGroupPolicies200ResponseInnerSchedulingMonday() *GetNetworkGroupPolicies200ResponseInnerSchedulingMonday {
	this := GetNetworkGroupPolicies200ResponseInnerSchedulingMonday{}
	return &this
}

// NewGetNetworkGroupPolicies200ResponseInnerSchedulingMondayWithDefaults instantiates a new GetNetworkGroupPolicies200ResponseInnerSchedulingMonday object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkGroupPolicies200ResponseInnerSchedulingMondayWithDefaults() *GetNetworkGroupPolicies200ResponseInnerSchedulingMonday {
	this := GetNetworkGroupPolicies200ResponseInnerSchedulingMonday{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *GetNetworkGroupPolicies200ResponseInnerSchedulingMonday) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerSchedulingMonday) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerSchedulingMonday) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *GetNetworkGroupPolicies200ResponseInnerSchedulingMonday) SetActive(v bool) {
	o.Active = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *GetNetworkGroupPolicies200ResponseInnerSchedulingMonday) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerSchedulingMonday) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerSchedulingMonday) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *GetNetworkGroupPolicies200ResponseInnerSchedulingMonday) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *GetNetworkGroupPolicies200ResponseInnerSchedulingMonday) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerSchedulingMonday) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerSchedulingMonday) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *GetNetworkGroupPolicies200ResponseInnerSchedulingMonday) SetTo(v string) {
	o.To = &v
}

func (o GetNetworkGroupPolicies200ResponseInnerSchedulingMonday) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkGroupPolicies200ResponseInnerSchedulingMonday) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	return toSerialize, nil
}

type NullableGetNetworkGroupPolicies200ResponseInnerSchedulingMonday struct {
	value *GetNetworkGroupPolicies200ResponseInnerSchedulingMonday
	isSet bool
}

func (v NullableGetNetworkGroupPolicies200ResponseInnerSchedulingMonday) Get() *GetNetworkGroupPolicies200ResponseInnerSchedulingMonday {
	return v.value
}

func (v *NullableGetNetworkGroupPolicies200ResponseInnerSchedulingMonday) Set(val *GetNetworkGroupPolicies200ResponseInnerSchedulingMonday) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkGroupPolicies200ResponseInnerSchedulingMonday) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkGroupPolicies200ResponseInnerSchedulingMonday) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkGroupPolicies200ResponseInnerSchedulingMonday(val *GetNetworkGroupPolicies200ResponseInnerSchedulingMonday) *NullableGetNetworkGroupPolicies200ResponseInnerSchedulingMonday {
	return &NullableGetNetworkGroupPolicies200ResponseInnerSchedulingMonday{value: val, isSet: true}
}

func (v NullableGetNetworkGroupPolicies200ResponseInnerSchedulingMonday) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkGroupPolicies200ResponseInnerSchedulingMonday) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


