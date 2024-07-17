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

// checks if the GetOrganizationSnmp200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSnmp200Response{}

// GetOrganizationSnmp200Response struct for GetOrganizationSnmp200Response
type GetOrganizationSnmp200Response struct {
	// Boolean indicating whether SNMP version 2c is enabled for the organization.
	V2cEnabled *bool `json:"v2cEnabled,omitempty"`
	// The community string for SNMP version 2c, if enabled.
	V2CommunityString *string `json:"v2CommunityString,omitempty"`
	// Boolean indicating whether SNMP version 3 is enabled for the organization.
	V3Enabled *bool `json:"v3Enabled,omitempty"`
	// The user for SNMP version 3, if enabled.
	V3User *string `json:"v3User,omitempty"`
	// The SNMP version 3 authentication mode. Can be either 'MD5' or 'SHA'.
	V3AuthMode *string `json:"v3AuthMode,omitempty"`
	// The SNMP version 3 privacy mode. Can be either 'DES' or 'AES128'.
	V3PrivMode *string `json:"v3PrivMode,omitempty"`
	// The list of IPv4 addresses that are allowed to access the SNMP server.
	PeerIps []string `json:"peerIps,omitempty"`
	// The hostname of the SNMP server.
	Hostname *string `json:"hostname,omitempty"`
	// The port of the SNMP server.
	Port *int32 `json:"port,omitempty"`
}

// NewGetOrganizationSnmp200Response instantiates a new GetOrganizationSnmp200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSnmp200Response() *GetOrganizationSnmp200Response {
	this := GetOrganizationSnmp200Response{}
	return &this
}

// NewGetOrganizationSnmp200ResponseWithDefaults instantiates a new GetOrganizationSnmp200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSnmp200ResponseWithDefaults() *GetOrganizationSnmp200Response {
	this := GetOrganizationSnmp200Response{}
	return &this
}

// GetV2cEnabled returns the V2cEnabled field value if set, zero value otherwise.
func (o *GetOrganizationSnmp200Response) GetV2cEnabled() bool {
	if o == nil || IsNil(o.V2cEnabled) {
		var ret bool
		return ret
	}
	return *o.V2cEnabled
}

// GetV2cEnabledOk returns a tuple with the V2cEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSnmp200Response) GetV2cEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.V2cEnabled) {
		return nil, false
	}
	return o.V2cEnabled, true
}

// HasV2cEnabled returns a boolean if a field has been set.
func (o *GetOrganizationSnmp200Response) HasV2cEnabled() bool {
	if o != nil && !IsNil(o.V2cEnabled) {
		return true
	}

	return false
}

// SetV2cEnabled gets a reference to the given bool and assigns it to the V2cEnabled field.
func (o *GetOrganizationSnmp200Response) SetV2cEnabled(v bool) {
	o.V2cEnabled = &v
}

// GetV2CommunityString returns the V2CommunityString field value if set, zero value otherwise.
func (o *GetOrganizationSnmp200Response) GetV2CommunityString() string {
	if o == nil || IsNil(o.V2CommunityString) {
		var ret string
		return ret
	}
	return *o.V2CommunityString
}

// GetV2CommunityStringOk returns a tuple with the V2CommunityString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSnmp200Response) GetV2CommunityStringOk() (*string, bool) {
	if o == nil || IsNil(o.V2CommunityString) {
		return nil, false
	}
	return o.V2CommunityString, true
}

// HasV2CommunityString returns a boolean if a field has been set.
func (o *GetOrganizationSnmp200Response) HasV2CommunityString() bool {
	if o != nil && !IsNil(o.V2CommunityString) {
		return true
	}

	return false
}

// SetV2CommunityString gets a reference to the given string and assigns it to the V2CommunityString field.
func (o *GetOrganizationSnmp200Response) SetV2CommunityString(v string) {
	o.V2CommunityString = &v
}

// GetV3Enabled returns the V3Enabled field value if set, zero value otherwise.
func (o *GetOrganizationSnmp200Response) GetV3Enabled() bool {
	if o == nil || IsNil(o.V3Enabled) {
		var ret bool
		return ret
	}
	return *o.V3Enabled
}

// GetV3EnabledOk returns a tuple with the V3Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSnmp200Response) GetV3EnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.V3Enabled) {
		return nil, false
	}
	return o.V3Enabled, true
}

// HasV3Enabled returns a boolean if a field has been set.
func (o *GetOrganizationSnmp200Response) HasV3Enabled() bool {
	if o != nil && !IsNil(o.V3Enabled) {
		return true
	}

	return false
}

// SetV3Enabled gets a reference to the given bool and assigns it to the V3Enabled field.
func (o *GetOrganizationSnmp200Response) SetV3Enabled(v bool) {
	o.V3Enabled = &v
}

// GetV3User returns the V3User field value if set, zero value otherwise.
func (o *GetOrganizationSnmp200Response) GetV3User() string {
	if o == nil || IsNil(o.V3User) {
		var ret string
		return ret
	}
	return *o.V3User
}

// GetV3UserOk returns a tuple with the V3User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSnmp200Response) GetV3UserOk() (*string, bool) {
	if o == nil || IsNil(o.V3User) {
		return nil, false
	}
	return o.V3User, true
}

// HasV3User returns a boolean if a field has been set.
func (o *GetOrganizationSnmp200Response) HasV3User() bool {
	if o != nil && !IsNil(o.V3User) {
		return true
	}

	return false
}

// SetV3User gets a reference to the given string and assigns it to the V3User field.
func (o *GetOrganizationSnmp200Response) SetV3User(v string) {
	o.V3User = &v
}

// GetV3AuthMode returns the V3AuthMode field value if set, zero value otherwise.
func (o *GetOrganizationSnmp200Response) GetV3AuthMode() string {
	if o == nil || IsNil(o.V3AuthMode) {
		var ret string
		return ret
	}
	return *o.V3AuthMode
}

// GetV3AuthModeOk returns a tuple with the V3AuthMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSnmp200Response) GetV3AuthModeOk() (*string, bool) {
	if o == nil || IsNil(o.V3AuthMode) {
		return nil, false
	}
	return o.V3AuthMode, true
}

// HasV3AuthMode returns a boolean if a field has been set.
func (o *GetOrganizationSnmp200Response) HasV3AuthMode() bool {
	if o != nil && !IsNil(o.V3AuthMode) {
		return true
	}

	return false
}

// SetV3AuthMode gets a reference to the given string and assigns it to the V3AuthMode field.
func (o *GetOrganizationSnmp200Response) SetV3AuthMode(v string) {
	o.V3AuthMode = &v
}

// GetV3PrivMode returns the V3PrivMode field value if set, zero value otherwise.
func (o *GetOrganizationSnmp200Response) GetV3PrivMode() string {
	if o == nil || IsNil(o.V3PrivMode) {
		var ret string
		return ret
	}
	return *o.V3PrivMode
}

// GetV3PrivModeOk returns a tuple with the V3PrivMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSnmp200Response) GetV3PrivModeOk() (*string, bool) {
	if o == nil || IsNil(o.V3PrivMode) {
		return nil, false
	}
	return o.V3PrivMode, true
}

// HasV3PrivMode returns a boolean if a field has been set.
func (o *GetOrganizationSnmp200Response) HasV3PrivMode() bool {
	if o != nil && !IsNil(o.V3PrivMode) {
		return true
	}

	return false
}

// SetV3PrivMode gets a reference to the given string and assigns it to the V3PrivMode field.
func (o *GetOrganizationSnmp200Response) SetV3PrivMode(v string) {
	o.V3PrivMode = &v
}

// GetPeerIps returns the PeerIps field value if set, zero value otherwise.
func (o *GetOrganizationSnmp200Response) GetPeerIps() []string {
	if o == nil || IsNil(o.PeerIps) {
		var ret []string
		return ret
	}
	return o.PeerIps
}

// GetPeerIpsOk returns a tuple with the PeerIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSnmp200Response) GetPeerIpsOk() ([]string, bool) {
	if o == nil || IsNil(o.PeerIps) {
		return nil, false
	}
	return o.PeerIps, true
}

// HasPeerIps returns a boolean if a field has been set.
func (o *GetOrganizationSnmp200Response) HasPeerIps() bool {
	if o != nil && !IsNil(o.PeerIps) {
		return true
	}

	return false
}

// SetPeerIps gets a reference to the given []string and assigns it to the PeerIps field.
func (o *GetOrganizationSnmp200Response) SetPeerIps(v []string) {
	o.PeerIps = v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *GetOrganizationSnmp200Response) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSnmp200Response) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *GetOrganizationSnmp200Response) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *GetOrganizationSnmp200Response) SetHostname(v string) {
	o.Hostname = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *GetOrganizationSnmp200Response) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSnmp200Response) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *GetOrganizationSnmp200Response) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *GetOrganizationSnmp200Response) SetPort(v int32) {
	o.Port = &v
}

func (o GetOrganizationSnmp200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSnmp200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.V2cEnabled) {
		toSerialize["v2cEnabled"] = o.V2cEnabled
	}
	if !IsNil(o.V2CommunityString) {
		toSerialize["v2CommunityString"] = o.V2CommunityString
	}
	if !IsNil(o.V3Enabled) {
		toSerialize["v3Enabled"] = o.V3Enabled
	}
	if !IsNil(o.V3User) {
		toSerialize["v3User"] = o.V3User
	}
	if !IsNil(o.V3AuthMode) {
		toSerialize["v3AuthMode"] = o.V3AuthMode
	}
	if !IsNil(o.V3PrivMode) {
		toSerialize["v3PrivMode"] = o.V3PrivMode
	}
	if !IsNil(o.PeerIps) {
		toSerialize["peerIps"] = o.PeerIps
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return toSerialize, nil
}

type NullableGetOrganizationSnmp200Response struct {
	value *GetOrganizationSnmp200Response
	isSet bool
}

func (v NullableGetOrganizationSnmp200Response) Get() *GetOrganizationSnmp200Response {
	return v.value
}

func (v *NullableGetOrganizationSnmp200Response) Set(val *GetOrganizationSnmp200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSnmp200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSnmp200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSnmp200Response(val *GetOrganizationSnmp200Response) *NullableGetOrganizationSnmp200Response {
	return &NullableGetOrganizationSnmp200Response{value: val, isSet: true}
}

func (v NullableGetOrganizationSnmp200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSnmp200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

