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

// checks if the AssignOrganizationLicensesSeats200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignOrganizationLicensesSeats200Response{}

// AssignOrganizationLicensesSeats200Response struct for AssignOrganizationLicensesSeats200Response
type AssignOrganizationLicensesSeats200Response struct {
	// Resulting licenses from the move
	ResultingLicenses []GetOrganizationLicenses200ResponseInner `json:"resultingLicenses,omitempty"`
}

// NewAssignOrganizationLicensesSeats200Response instantiates a new AssignOrganizationLicensesSeats200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignOrganizationLicensesSeats200Response() *AssignOrganizationLicensesSeats200Response {
	this := AssignOrganizationLicensesSeats200Response{}
	return &this
}

// NewAssignOrganizationLicensesSeats200ResponseWithDefaults instantiates a new AssignOrganizationLicensesSeats200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignOrganizationLicensesSeats200ResponseWithDefaults() *AssignOrganizationLicensesSeats200Response {
	this := AssignOrganizationLicensesSeats200Response{}
	return &this
}

// GetResultingLicenses returns the ResultingLicenses field value if set, zero value otherwise.
func (o *AssignOrganizationLicensesSeats200Response) GetResultingLicenses() []GetOrganizationLicenses200ResponseInner {
	if o == nil || IsNil(o.ResultingLicenses) {
		var ret []GetOrganizationLicenses200ResponseInner
		return ret
	}
	return o.ResultingLicenses
}

// GetResultingLicensesOk returns a tuple with the ResultingLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignOrganizationLicensesSeats200Response) GetResultingLicensesOk() ([]GetOrganizationLicenses200ResponseInner, bool) {
	if o == nil || IsNil(o.ResultingLicenses) {
		return nil, false
	}
	return o.ResultingLicenses, true
}

// HasResultingLicenses returns a boolean if a field has been set.
func (o *AssignOrganizationLicensesSeats200Response) HasResultingLicenses() bool {
	if o != nil && !IsNil(o.ResultingLicenses) {
		return true
	}

	return false
}

// SetResultingLicenses gets a reference to the given []GetOrganizationLicenses200ResponseInner and assigns it to the ResultingLicenses field.
func (o *AssignOrganizationLicensesSeats200Response) SetResultingLicenses(v []GetOrganizationLicenses200ResponseInner) {
	o.ResultingLicenses = v
}

func (o AssignOrganizationLicensesSeats200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignOrganizationLicensesSeats200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResultingLicenses) {
		toSerialize["resultingLicenses"] = o.ResultingLicenses
	}
	return toSerialize, nil
}

type NullableAssignOrganizationLicensesSeats200Response struct {
	value *AssignOrganizationLicensesSeats200Response
	isSet bool
}

func (v NullableAssignOrganizationLicensesSeats200Response) Get() *AssignOrganizationLicensesSeats200Response {
	return v.value
}

func (v *NullableAssignOrganizationLicensesSeats200Response) Set(val *AssignOrganizationLicensesSeats200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignOrganizationLicensesSeats200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignOrganizationLicensesSeats200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignOrganizationLicensesSeats200Response(val *AssignOrganizationLicensesSeats200Response) *NullableAssignOrganizationLicensesSeats200Response {
	return &NullableAssignOrganizationLicensesSeats200Response{value: val, isSet: true}
}

func (v NullableAssignOrganizationLicensesSeats200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignOrganizationLicensesSeats200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


