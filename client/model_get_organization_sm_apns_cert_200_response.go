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

// checks if the GetOrganizationSmApnsCert200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSmApnsCert200Response{}

// GetOrganizationSmApnsCert200Response struct for GetOrganizationSmApnsCert200Response
type GetOrganizationSmApnsCert200Response struct {
	// Organization APNS Certificate used by devices to communication with Apple
	Certificate *string `json:"certificate,omitempty"`
}

// NewGetOrganizationSmApnsCert200Response instantiates a new GetOrganizationSmApnsCert200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSmApnsCert200Response() *GetOrganizationSmApnsCert200Response {
	this := GetOrganizationSmApnsCert200Response{}
	return &this
}

// NewGetOrganizationSmApnsCert200ResponseWithDefaults instantiates a new GetOrganizationSmApnsCert200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSmApnsCert200ResponseWithDefaults() *GetOrganizationSmApnsCert200Response {
	this := GetOrganizationSmApnsCert200Response{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *GetOrganizationSmApnsCert200Response) GetCertificate() string {
	if o == nil || IsNil(o.Certificate) {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSmApnsCert200Response) GetCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *GetOrganizationSmApnsCert200Response) HasCertificate() bool {
	if o != nil && !IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *GetOrganizationSmApnsCert200Response) SetCertificate(v string) {
	o.Certificate = &v
}

func (o GetOrganizationSmApnsCert200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSmApnsCert200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	return toSerialize, nil
}

type NullableGetOrganizationSmApnsCert200Response struct {
	value *GetOrganizationSmApnsCert200Response
	isSet bool
}

func (v NullableGetOrganizationSmApnsCert200Response) Get() *GetOrganizationSmApnsCert200Response {
	return v.value
}

func (v *NullableGetOrganizationSmApnsCert200Response) Set(val *GetOrganizationSmApnsCert200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSmApnsCert200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSmApnsCert200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSmApnsCert200Response(val *GetOrganizationSmApnsCert200Response) *NullableGetOrganizationSmApnsCert200Response {
	return &NullableGetOrganizationSmApnsCert200Response{value: val, isSet: true}
}

func (v NullableGetOrganizationSmApnsCert200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSmApnsCert200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


