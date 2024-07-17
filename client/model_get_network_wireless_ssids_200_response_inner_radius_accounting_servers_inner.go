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

// checks if the GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner{}

// GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner struct for GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner
type GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner struct {
	// IP address (or FQDN) to which the APs will send RADIUS accounting messages
	Host *string `json:"host,omitempty"`
	// Port on the RADIUS server that is listening for accounting messages
	Port *int32 `json:"port,omitempty"`
	// The ID of the Openroaming Certificate attached to radius server
	OpenRoamingCertificateId *int32 `json:"openRoamingCertificateId,omitempty"`
	// Certificate used for authorization for the RADSEC Server
	CaCertificate *string `json:"caCertificate,omitempty"`
}

// NewGetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner instantiates a new GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner() *GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner {
	this := GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner{}
	return &this
}

// NewGetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInnerWithDefaults instantiates a new GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInnerWithDefaults() *GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner {
	this := GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) SetPort(v int32) {
	o.Port = &v
}

// GetOpenRoamingCertificateId returns the OpenRoamingCertificateId field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) GetOpenRoamingCertificateId() int32 {
	if o == nil || IsNil(o.OpenRoamingCertificateId) {
		var ret int32
		return ret
	}
	return *o.OpenRoamingCertificateId
}

// GetOpenRoamingCertificateIdOk returns a tuple with the OpenRoamingCertificateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) GetOpenRoamingCertificateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OpenRoamingCertificateId) {
		return nil, false
	}
	return o.OpenRoamingCertificateId, true
}

// HasOpenRoamingCertificateId returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) HasOpenRoamingCertificateId() bool {
	if o != nil && !IsNil(o.OpenRoamingCertificateId) {
		return true
	}

	return false
}

// SetOpenRoamingCertificateId gets a reference to the given int32 and assigns it to the OpenRoamingCertificateId field.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) SetOpenRoamingCertificateId(v int32) {
	o.OpenRoamingCertificateId = &v
}

// GetCaCertificate returns the CaCertificate field value if set, zero value otherwise.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) GetCaCertificate() string {
	if o == nil || IsNil(o.CaCertificate) {
		var ret string
		return ret
	}
	return *o.CaCertificate
}

// GetCaCertificateOk returns a tuple with the CaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) GetCaCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.CaCertificate) {
		return nil, false
	}
	return o.CaCertificate, true
}

// HasCaCertificate returns a boolean if a field has been set.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) HasCaCertificate() bool {
	if o != nil && !IsNil(o.CaCertificate) {
		return true
	}

	return false
}

// SetCaCertificate gets a reference to the given string and assigns it to the CaCertificate field.
func (o *GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) SetCaCertificate(v string) {
	o.CaCertificate = &v
}

func (o GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.OpenRoamingCertificateId) {
		toSerialize["openRoamingCertificateId"] = o.OpenRoamingCertificateId
	}
	if !IsNil(o.CaCertificate) {
		toSerialize["caCertificate"] = o.CaCertificate
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner struct {
	value *GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner
	isSet bool
}

func (v NullableGetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) Get() *GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner {
	return v.value
}

func (v *NullableGetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) Set(val *GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner(val *GetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) *NullableGetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner {
	return &NullableGetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessSsids200ResponseInnerRadiusAccountingServersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


