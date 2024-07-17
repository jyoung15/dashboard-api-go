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

// checks if the GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations{}

// GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations A hash of destinations for this specific alert
type GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations struct {
	// A list of emails that will receive information about the alert
	Emails []string `json:"emails,omitempty"`
	// A list of phone numbers that will receive text messages about the alert. Only available for sensors status alerts.
	SmsNumbers []string `json:"smsNumbers,omitempty"`
	// If true, then all network admins will receive emails for this alert
	AllAdmins *bool `json:"allAdmins,omitempty"`
	// If true, then an SNMP trap will be sent for this alert if there is an SNMP trap server configured for this network
	Snmp *bool `json:"snmp,omitempty"`
	// A list of HTTP server IDs to send a Webhook to for this alert
	HttpServerIds []string `json:"httpServerIds,omitempty"`
}

// NewGetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations instantiates a new GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations() *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations {
	this := GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations{}
	return &this
}

// NewGetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinationsWithDefaults instantiates a new GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinationsWithDefaults() *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations {
	this := GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations{}
	return &this
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) GetEmails() []string {
	if o == nil || IsNil(o.Emails) {
		var ret []string
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) GetEmailsOk() ([]string, bool) {
	if o == nil || IsNil(o.Emails) {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) HasEmails() bool {
	if o != nil && !IsNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []string and assigns it to the Emails field.
func (o *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) SetEmails(v []string) {
	o.Emails = v
}

// GetSmsNumbers returns the SmsNumbers field value if set, zero value otherwise.
func (o *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) GetSmsNumbers() []string {
	if o == nil || IsNil(o.SmsNumbers) {
		var ret []string
		return ret
	}
	return o.SmsNumbers
}

// GetSmsNumbersOk returns a tuple with the SmsNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) GetSmsNumbersOk() ([]string, bool) {
	if o == nil || IsNil(o.SmsNumbers) {
		return nil, false
	}
	return o.SmsNumbers, true
}

// HasSmsNumbers returns a boolean if a field has been set.
func (o *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) HasSmsNumbers() bool {
	if o != nil && !IsNil(o.SmsNumbers) {
		return true
	}

	return false
}

// SetSmsNumbers gets a reference to the given []string and assigns it to the SmsNumbers field.
func (o *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) SetSmsNumbers(v []string) {
	o.SmsNumbers = v
}

// GetAllAdmins returns the AllAdmins field value if set, zero value otherwise.
func (o *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) GetAllAdmins() bool {
	if o == nil || IsNil(o.AllAdmins) {
		var ret bool
		return ret
	}
	return *o.AllAdmins
}

// GetAllAdminsOk returns a tuple with the AllAdmins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) GetAllAdminsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllAdmins) {
		return nil, false
	}
	return o.AllAdmins, true
}

// HasAllAdmins returns a boolean if a field has been set.
func (o *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) HasAllAdmins() bool {
	if o != nil && !IsNil(o.AllAdmins) {
		return true
	}

	return false
}

// SetAllAdmins gets a reference to the given bool and assigns it to the AllAdmins field.
func (o *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) SetAllAdmins(v bool) {
	o.AllAdmins = &v
}

// GetSnmp returns the Snmp field value if set, zero value otherwise.
func (o *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) GetSnmp() bool {
	if o == nil || IsNil(o.Snmp) {
		var ret bool
		return ret
	}
	return *o.Snmp
}

// GetSnmpOk returns a tuple with the Snmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) GetSnmpOk() (*bool, bool) {
	if o == nil || IsNil(o.Snmp) {
		return nil, false
	}
	return o.Snmp, true
}

// HasSnmp returns a boolean if a field has been set.
func (o *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) HasSnmp() bool {
	if o != nil && !IsNil(o.Snmp) {
		return true
	}

	return false
}

// SetSnmp gets a reference to the given bool and assigns it to the Snmp field.
func (o *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) SetSnmp(v bool) {
	o.Snmp = &v
}

// GetHttpServerIds returns the HttpServerIds field value if set, zero value otherwise.
func (o *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) GetHttpServerIds() []string {
	if o == nil || IsNil(o.HttpServerIds) {
		var ret []string
		return ret
	}
	return o.HttpServerIds
}

// GetHttpServerIdsOk returns a tuple with the HttpServerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) GetHttpServerIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.HttpServerIds) {
		return nil, false
	}
	return o.HttpServerIds, true
}

// HasHttpServerIds returns a boolean if a field has been set.
func (o *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) HasHttpServerIds() bool {
	if o != nil && !IsNil(o.HttpServerIds) {
		return true
	}

	return false
}

// SetHttpServerIds gets a reference to the given []string and assigns it to the HttpServerIds field.
func (o *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) SetHttpServerIds(v []string) {
	o.HttpServerIds = v
}

func (o GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Emails) {
		toSerialize["emails"] = o.Emails
	}
	if !IsNil(o.SmsNumbers) {
		toSerialize["smsNumbers"] = o.SmsNumbers
	}
	if !IsNil(o.AllAdmins) {
		toSerialize["allAdmins"] = o.AllAdmins
	}
	if !IsNil(o.Snmp) {
		toSerialize["snmp"] = o.Snmp
	}
	if !IsNil(o.HttpServerIds) {
		toSerialize["httpServerIds"] = o.HttpServerIds
	}
	return toSerialize, nil
}

type NullableGetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations struct {
	value *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations
	isSet bool
}

func (v NullableGetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) Get() *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations {
	return v.value
}

func (v *NullableGetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) Set(val *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations(val *GetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) *NullableGetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations {
	return &NullableGetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations{value: val, isSet: true}
}

func (v NullableGetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkAlertsSettings200ResponseAlertsInnerAlertDestinations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


