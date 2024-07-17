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

// checks if the GetNetworkSplashLoginAttempts200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSplashLoginAttempts200ResponseInner{}

// GetNetworkSplashLoginAttempts200ResponseInner struct for GetNetworkSplashLoginAttempts200ResponseInner
type GetNetworkSplashLoginAttempts200ResponseInner struct {
	// User name
	Name *string `json:"name,omitempty"`
	// User login identifier
	Login *string `json:"login,omitempty"`
	// SSID name
	Ssid *string `json:"ssid,omitempty"`
	// Login timestamp
	LoginAt *time.Time `json:"loginAt,omitempty"`
	// Gateway device mac address
	GatewayDeviceMac *string `json:"gatewayDeviceMac,omitempty"`
	// Client mac address
	ClientMac *string `json:"clientMac,omitempty"`
	// Client ID
	ClientId *string `json:"clientId,omitempty"`
	// Authorization status
	Authorization *string `json:"authorization,omitempty"`
}

// NewGetNetworkSplashLoginAttempts200ResponseInner instantiates a new GetNetworkSplashLoginAttempts200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSplashLoginAttempts200ResponseInner() *GetNetworkSplashLoginAttempts200ResponseInner {
	this := GetNetworkSplashLoginAttempts200ResponseInner{}
	return &this
}

// NewGetNetworkSplashLoginAttempts200ResponseInnerWithDefaults instantiates a new GetNetworkSplashLoginAttempts200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSplashLoginAttempts200ResponseInnerWithDefaults() *GetNetworkSplashLoginAttempts200ResponseInner {
	this := GetNetworkSplashLoginAttempts200ResponseInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetLogin() string {
	if o == nil || IsNil(o.Login) {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetLoginOk() (*string, bool) {
	if o == nil || IsNil(o.Login) {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) HasLogin() bool {
	if o != nil && !IsNil(o.Login) {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) SetLogin(v string) {
	o.Login = &v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetSsid() string {
	if o == nil || IsNil(o.Ssid) {
		var ret string
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetSsidOk() (*string, bool) {
	if o == nil || IsNil(o.Ssid) {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) HasSsid() bool {
	if o != nil && !IsNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given string and assigns it to the Ssid field.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) SetSsid(v string) {
	o.Ssid = &v
}

// GetLoginAt returns the LoginAt field value if set, zero value otherwise.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetLoginAt() time.Time {
	if o == nil || IsNil(o.LoginAt) {
		var ret time.Time
		return ret
	}
	return *o.LoginAt
}

// GetLoginAtOk returns a tuple with the LoginAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetLoginAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LoginAt) {
		return nil, false
	}
	return o.LoginAt, true
}

// HasLoginAt returns a boolean if a field has been set.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) HasLoginAt() bool {
	if o != nil && !IsNil(o.LoginAt) {
		return true
	}

	return false
}

// SetLoginAt gets a reference to the given time.Time and assigns it to the LoginAt field.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) SetLoginAt(v time.Time) {
	o.LoginAt = &v
}

// GetGatewayDeviceMac returns the GatewayDeviceMac field value if set, zero value otherwise.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetGatewayDeviceMac() string {
	if o == nil || IsNil(o.GatewayDeviceMac) {
		var ret string
		return ret
	}
	return *o.GatewayDeviceMac
}

// GetGatewayDeviceMacOk returns a tuple with the GatewayDeviceMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetGatewayDeviceMacOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayDeviceMac) {
		return nil, false
	}
	return o.GatewayDeviceMac, true
}

// HasGatewayDeviceMac returns a boolean if a field has been set.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) HasGatewayDeviceMac() bool {
	if o != nil && !IsNil(o.GatewayDeviceMac) {
		return true
	}

	return false
}

// SetGatewayDeviceMac gets a reference to the given string and assigns it to the GatewayDeviceMac field.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) SetGatewayDeviceMac(v string) {
	o.GatewayDeviceMac = &v
}

// GetClientMac returns the ClientMac field value if set, zero value otherwise.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetClientMac() string {
	if o == nil || IsNil(o.ClientMac) {
		var ret string
		return ret
	}
	return *o.ClientMac
}

// GetClientMacOk returns a tuple with the ClientMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetClientMacOk() (*string, bool) {
	if o == nil || IsNil(o.ClientMac) {
		return nil, false
	}
	return o.ClientMac, true
}

// HasClientMac returns a boolean if a field has been set.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) HasClientMac() bool {
	if o != nil && !IsNil(o.ClientMac) {
		return true
	}

	return false
}

// SetClientMac gets a reference to the given string and assigns it to the ClientMac field.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) SetClientMac(v string) {
	o.ClientMac = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) SetClientId(v string) {
	o.ClientId = &v
}

// GetAuthorization returns the Authorization field value if set, zero value otherwise.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetAuthorization() string {
	if o == nil || IsNil(o.Authorization) {
		var ret string
		return ret
	}
	return *o.Authorization
}

// GetAuthorizationOk returns a tuple with the Authorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetAuthorizationOk() (*string, bool) {
	if o == nil || IsNil(o.Authorization) {
		return nil, false
	}
	return o.Authorization, true
}

// HasAuthorization returns a boolean if a field has been set.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) HasAuthorization() bool {
	if o != nil && !IsNil(o.Authorization) {
		return true
	}

	return false
}

// SetAuthorization gets a reference to the given string and assigns it to the Authorization field.
func (o *GetNetworkSplashLoginAttempts200ResponseInner) SetAuthorization(v string) {
	o.Authorization = &v
}

func (o GetNetworkSplashLoginAttempts200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSplashLoginAttempts200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Login) {
		toSerialize["login"] = o.Login
	}
	if !IsNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	if !IsNil(o.LoginAt) {
		toSerialize["loginAt"] = o.LoginAt
	}
	if !IsNil(o.GatewayDeviceMac) {
		toSerialize["gatewayDeviceMac"] = o.GatewayDeviceMac
	}
	if !IsNil(o.ClientMac) {
		toSerialize["clientMac"] = o.ClientMac
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.Authorization) {
		toSerialize["authorization"] = o.Authorization
	}
	return toSerialize, nil
}

type NullableGetNetworkSplashLoginAttempts200ResponseInner struct {
	value *GetNetworkSplashLoginAttempts200ResponseInner
	isSet bool
}

func (v NullableGetNetworkSplashLoginAttempts200ResponseInner) Get() *GetNetworkSplashLoginAttempts200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkSplashLoginAttempts200ResponseInner) Set(val *GetNetworkSplashLoginAttempts200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSplashLoginAttempts200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSplashLoginAttempts200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSplashLoginAttempts200ResponseInner(val *GetNetworkSplashLoginAttempts200ResponseInner) *NullableGetNetworkSplashLoginAttempts200ResponseInner {
	return &NullableGetNetworkSplashLoginAttempts200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkSplashLoginAttempts200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSplashLoginAttempts200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


