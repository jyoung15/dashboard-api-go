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

// checks if the GetOrganizationAlertsProfiles200ResponseInnerRecipients type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationAlertsProfiles200ResponseInnerRecipients{}

// GetOrganizationAlertsProfiles200ResponseInnerRecipients List of recipients that will recieve the alert.
type GetOrganizationAlertsProfiles200ResponseInnerRecipients struct {
	// A list of emails that will receive information about the alert
	Emails []string `json:"emails,omitempty"`
	// A list base64 encoded urls of webhook endpoints that will receive information about the alert
	HttpServerIds []string `json:"httpServerIds,omitempty"`
}

// NewGetOrganizationAlertsProfiles200ResponseInnerRecipients instantiates a new GetOrganizationAlertsProfiles200ResponseInnerRecipients object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationAlertsProfiles200ResponseInnerRecipients() *GetOrganizationAlertsProfiles200ResponseInnerRecipients {
	this := GetOrganizationAlertsProfiles200ResponseInnerRecipients{}
	return &this
}

// NewGetOrganizationAlertsProfiles200ResponseInnerRecipientsWithDefaults instantiates a new GetOrganizationAlertsProfiles200ResponseInnerRecipients object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationAlertsProfiles200ResponseInnerRecipientsWithDefaults() *GetOrganizationAlertsProfiles200ResponseInnerRecipients {
	this := GetOrganizationAlertsProfiles200ResponseInnerRecipients{}
	return &this
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *GetOrganizationAlertsProfiles200ResponseInnerRecipients) GetEmails() []string {
	if o == nil || IsNil(o.Emails) {
		var ret []string
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAlertsProfiles200ResponseInnerRecipients) GetEmailsOk() ([]string, bool) {
	if o == nil || IsNil(o.Emails) {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *GetOrganizationAlertsProfiles200ResponseInnerRecipients) HasEmails() bool {
	if o != nil && !IsNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []string and assigns it to the Emails field.
func (o *GetOrganizationAlertsProfiles200ResponseInnerRecipients) SetEmails(v []string) {
	o.Emails = v
}

// GetHttpServerIds returns the HttpServerIds field value if set, zero value otherwise.
func (o *GetOrganizationAlertsProfiles200ResponseInnerRecipients) GetHttpServerIds() []string {
	if o == nil || IsNil(o.HttpServerIds) {
		var ret []string
		return ret
	}
	return o.HttpServerIds
}

// GetHttpServerIdsOk returns a tuple with the HttpServerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAlertsProfiles200ResponseInnerRecipients) GetHttpServerIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.HttpServerIds) {
		return nil, false
	}
	return o.HttpServerIds, true
}

// HasHttpServerIds returns a boolean if a field has been set.
func (o *GetOrganizationAlertsProfiles200ResponseInnerRecipients) HasHttpServerIds() bool {
	if o != nil && !IsNil(o.HttpServerIds) {
		return true
	}

	return false
}

// SetHttpServerIds gets a reference to the given []string and assigns it to the HttpServerIds field.
func (o *GetOrganizationAlertsProfiles200ResponseInnerRecipients) SetHttpServerIds(v []string) {
	o.HttpServerIds = v
}

func (o GetOrganizationAlertsProfiles200ResponseInnerRecipients) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationAlertsProfiles200ResponseInnerRecipients) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Emails) {
		toSerialize["emails"] = o.Emails
	}
	if !IsNil(o.HttpServerIds) {
		toSerialize["httpServerIds"] = o.HttpServerIds
	}
	return toSerialize, nil
}

type NullableGetOrganizationAlertsProfiles200ResponseInnerRecipients struct {
	value *GetOrganizationAlertsProfiles200ResponseInnerRecipients
	isSet bool
}

func (v NullableGetOrganizationAlertsProfiles200ResponseInnerRecipients) Get() *GetOrganizationAlertsProfiles200ResponseInnerRecipients {
	return v.value
}

func (v *NullableGetOrganizationAlertsProfiles200ResponseInnerRecipients) Set(val *GetOrganizationAlertsProfiles200ResponseInnerRecipients) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationAlertsProfiles200ResponseInnerRecipients) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationAlertsProfiles200ResponseInnerRecipients) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationAlertsProfiles200ResponseInnerRecipients(val *GetOrganizationAlertsProfiles200ResponseInnerRecipients) *NullableGetOrganizationAlertsProfiles200ResponseInnerRecipients {
	return &NullableGetOrganizationAlertsProfiles200ResponseInnerRecipients{value: val, isSet: true}
}

func (v NullableGetOrganizationAlertsProfiles200ResponseInnerRecipients) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationAlertsProfiles200ResponseInnerRecipients) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


