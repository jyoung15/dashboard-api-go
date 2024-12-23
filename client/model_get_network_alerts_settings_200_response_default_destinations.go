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

// checks if the GetNetworkAlertsSettings200ResponseDefaultDestinations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkAlertsSettings200ResponseDefaultDestinations{}

// GetNetworkAlertsSettings200ResponseDefaultDestinations The network-wide destinations for all alerts on the network.
type GetNetworkAlertsSettings200ResponseDefaultDestinations struct {
	// A list of emails that will receive the alert(s).
	Emails []string `json:"emails,omitempty"`
	// If true, then all network admins will receive emails.
	AllAdmins *bool `json:"allAdmins,omitempty"`
	// If true, then an SNMP trap will be sent if there is an SNMP trap server configured for this network.
	Snmp *bool `json:"snmp,omitempty"`
	// A list of HTTP server IDs to send a Webhook to
	HttpServerIds []string `json:"httpServerIds,omitempty"`
}

// NewGetNetworkAlertsSettings200ResponseDefaultDestinations instantiates a new GetNetworkAlertsSettings200ResponseDefaultDestinations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkAlertsSettings200ResponseDefaultDestinations() *GetNetworkAlertsSettings200ResponseDefaultDestinations {
	this := GetNetworkAlertsSettings200ResponseDefaultDestinations{}
	return &this
}

// NewGetNetworkAlertsSettings200ResponseDefaultDestinationsWithDefaults instantiates a new GetNetworkAlertsSettings200ResponseDefaultDestinations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkAlertsSettings200ResponseDefaultDestinationsWithDefaults() *GetNetworkAlertsSettings200ResponseDefaultDestinations {
	this := GetNetworkAlertsSettings200ResponseDefaultDestinations{}
	return &this
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) GetEmails() []string {
	if o == nil || IsNil(o.Emails) {
		var ret []string
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) GetEmailsOk() ([]string, bool) {
	if o == nil || IsNil(o.Emails) {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) HasEmails() bool {
	if o != nil && !IsNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []string and assigns it to the Emails field.
func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) SetEmails(v []string) {
	o.Emails = v
}

// GetAllAdmins returns the AllAdmins field value if set, zero value otherwise.
func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) GetAllAdmins() bool {
	if o == nil || IsNil(o.AllAdmins) {
		var ret bool
		return ret
	}
	return *o.AllAdmins
}

// GetAllAdminsOk returns a tuple with the AllAdmins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) GetAllAdminsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllAdmins) {
		return nil, false
	}
	return o.AllAdmins, true
}

// HasAllAdmins returns a boolean if a field has been set.
func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) HasAllAdmins() bool {
	if o != nil && !IsNil(o.AllAdmins) {
		return true
	}

	return false
}

// SetAllAdmins gets a reference to the given bool and assigns it to the AllAdmins field.
func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) SetAllAdmins(v bool) {
	o.AllAdmins = &v
}

// GetSnmp returns the Snmp field value if set, zero value otherwise.
func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) GetSnmp() bool {
	if o == nil || IsNil(o.Snmp) {
		var ret bool
		return ret
	}
	return *o.Snmp
}

// GetSnmpOk returns a tuple with the Snmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) GetSnmpOk() (*bool, bool) {
	if o == nil || IsNil(o.Snmp) {
		return nil, false
	}
	return o.Snmp, true
}

// HasSnmp returns a boolean if a field has been set.
func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) HasSnmp() bool {
	if o != nil && !IsNil(o.Snmp) {
		return true
	}

	return false
}

// SetSnmp gets a reference to the given bool and assigns it to the Snmp field.
func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) SetSnmp(v bool) {
	o.Snmp = &v
}

// GetHttpServerIds returns the HttpServerIds field value if set, zero value otherwise.
func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) GetHttpServerIds() []string {
	if o == nil || IsNil(o.HttpServerIds) {
		var ret []string
		return ret
	}
	return o.HttpServerIds
}

// GetHttpServerIdsOk returns a tuple with the HttpServerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) GetHttpServerIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.HttpServerIds) {
		return nil, false
	}
	return o.HttpServerIds, true
}

// HasHttpServerIds returns a boolean if a field has been set.
func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) HasHttpServerIds() bool {
	if o != nil && !IsNil(o.HttpServerIds) {
		return true
	}

	return false
}

// SetHttpServerIds gets a reference to the given []string and assigns it to the HttpServerIds field.
func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) SetHttpServerIds(v []string) {
	o.HttpServerIds = v
}

func (o GetNetworkAlertsSettings200ResponseDefaultDestinations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkAlertsSettings200ResponseDefaultDestinations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Emails) {
		toSerialize["emails"] = o.Emails
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

type NullableGetNetworkAlertsSettings200ResponseDefaultDestinations struct {
	value *GetNetworkAlertsSettings200ResponseDefaultDestinations
	isSet bool
}

func (v NullableGetNetworkAlertsSettings200ResponseDefaultDestinations) Get() *GetNetworkAlertsSettings200ResponseDefaultDestinations {
	return v.value
}

func (v *NullableGetNetworkAlertsSettings200ResponseDefaultDestinations) Set(val *GetNetworkAlertsSettings200ResponseDefaultDestinations) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkAlertsSettings200ResponseDefaultDestinations) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkAlertsSettings200ResponseDefaultDestinations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkAlertsSettings200ResponseDefaultDestinations(val *GetNetworkAlertsSettings200ResponseDefaultDestinations) *NullableGetNetworkAlertsSettings200ResponseDefaultDestinations {
	return &NullableGetNetworkAlertsSettings200ResponseDefaultDestinations{value: val, isSet: true}
}

func (v NullableGetNetworkAlertsSettings200ResponseDefaultDestinations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkAlertsSettings200ResponseDefaultDestinations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


