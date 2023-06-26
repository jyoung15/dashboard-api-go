/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// RenewOrganizationLicensesSeatsRequest struct for RenewOrganizationLicensesSeatsRequest
type RenewOrganizationLicensesSeatsRequest struct {
	// The ID of the SM license to renew. This license must already be assigned to an SM network
	LicenseIdToRenew string `json:"licenseIdToRenew"`
	// The SM license to use to renew the seats on 'licenseIdToRenew'. This license must have at least as many seats available as there are seats on 'licenseIdToRenew'
	UnusedLicenseId string `json:"unusedLicenseId"`
}

// NewRenewOrganizationLicensesSeatsRequest instantiates a new RenewOrganizationLicensesSeatsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenewOrganizationLicensesSeatsRequest(licenseIdToRenew string, unusedLicenseId string) *RenewOrganizationLicensesSeatsRequest {
	this := RenewOrganizationLicensesSeatsRequest{}
	this.LicenseIdToRenew = licenseIdToRenew
	this.UnusedLicenseId = unusedLicenseId
	return &this
}

// NewRenewOrganizationLicensesSeatsRequestWithDefaults instantiates a new RenewOrganizationLicensesSeatsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenewOrganizationLicensesSeatsRequestWithDefaults() *RenewOrganizationLicensesSeatsRequest {
	this := RenewOrganizationLicensesSeatsRequest{}
	return &this
}

// GetLicenseIdToRenew returns the LicenseIdToRenew field value
func (o *RenewOrganizationLicensesSeatsRequest) GetLicenseIdToRenew() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LicenseIdToRenew
}

// GetLicenseIdToRenewOk returns a tuple with the LicenseIdToRenew field value
// and a boolean to check if the value has been set.
func (o *RenewOrganizationLicensesSeatsRequest) GetLicenseIdToRenewOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LicenseIdToRenew, true
}

// SetLicenseIdToRenew sets field value
func (o *RenewOrganizationLicensesSeatsRequest) SetLicenseIdToRenew(v string) {
	o.LicenseIdToRenew = v
}

// GetUnusedLicenseId returns the UnusedLicenseId field value
func (o *RenewOrganizationLicensesSeatsRequest) GetUnusedLicenseId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnusedLicenseId
}

// GetUnusedLicenseIdOk returns a tuple with the UnusedLicenseId field value
// and a boolean to check if the value has been set.
func (o *RenewOrganizationLicensesSeatsRequest) GetUnusedLicenseIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.UnusedLicenseId, true
}

// SetUnusedLicenseId sets field value
func (o *RenewOrganizationLicensesSeatsRequest) SetUnusedLicenseId(v string) {
	o.UnusedLicenseId = v
}

func (o RenewOrganizationLicensesSeatsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["licenseIdToRenew"] = o.LicenseIdToRenew
	}
	if true {
		toSerialize["unusedLicenseId"] = o.UnusedLicenseId
	}
	return json.Marshal(toSerialize)
}

type NullableRenewOrganizationLicensesSeatsRequest struct {
	value *RenewOrganizationLicensesSeatsRequest
	isSet bool
}

func (v NullableRenewOrganizationLicensesSeatsRequest) Get() *RenewOrganizationLicensesSeatsRequest {
	return v.value
}

func (v *NullableRenewOrganizationLicensesSeatsRequest) Set(val *RenewOrganizationLicensesSeatsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRenewOrganizationLicensesSeatsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRenewOrganizationLicensesSeatsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenewOrganizationLicensesSeatsRequest(val *RenewOrganizationLicensesSeatsRequest) *NullableRenewOrganizationLicensesSeatsRequest {
	return &NullableRenewOrganizationLicensesSeatsRequest{value: val, isSet: true}
}

func (v NullableRenewOrganizationLicensesSeatsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenewOrganizationLicensesSeatsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


