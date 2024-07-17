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

// checks if the GetDeviceClients200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceClients200ResponseInner{}

// GetDeviceClients200ResponseInner struct for GetDeviceClients200ResponseInner
type GetDeviceClients200ResponseInner struct {
	// The ID of the client
	Id *string `json:"id,omitempty"`
	// The MAC address of the client
	Mac *string `json:"mac,omitempty"`
	// Short description of the client
	Description *string `json:"description,omitempty"`
	// The client's MDNS name
	MdnsName *string `json:"mdnsName,omitempty"`
	// The client's DHCP hostname
	DhcpHostname *string `json:"dhcpHostname,omitempty"`
	// The client user's name
	User *string `json:"user,omitempty"`
	// The IP address of the client
	Ip *string `json:"ip,omitempty"`
	// The client-assigned name of the VLAN the client is connected to
	Vlan *string `json:"vlan,omitempty"`
	// The owner-assigned name of the VLAN the client is connected to
	NamedVlan *string `json:"namedVlan,omitempty"`
	// The name of the switchport with clients on it, if the device is a switch
	Switchport *string `json:"switchport,omitempty"`
	// A description of the adaptive policy group
	AdaptivePolicyGroup *string `json:"adaptivePolicyGroup,omitempty"`
	Usage *GetDeviceClients200ResponseInnerUsage `json:"usage,omitempty"`
}

// NewGetDeviceClients200ResponseInner instantiates a new GetDeviceClients200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceClients200ResponseInner() *GetDeviceClients200ResponseInner {
	this := GetDeviceClients200ResponseInner{}
	return &this
}

// NewGetDeviceClients200ResponseInnerWithDefaults instantiates a new GetDeviceClients200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceClients200ResponseInnerWithDefaults() *GetDeviceClients200ResponseInner {
	this := GetDeviceClients200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetDeviceClients200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceClients200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetDeviceClients200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetDeviceClients200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *GetDeviceClients200ResponseInner) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceClients200ResponseInner) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *GetDeviceClients200ResponseInner) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *GetDeviceClients200ResponseInner) SetMac(v string) {
	o.Mac = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetDeviceClients200ResponseInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceClients200ResponseInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetDeviceClients200ResponseInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetDeviceClients200ResponseInner) SetDescription(v string) {
	o.Description = &v
}

// GetMdnsName returns the MdnsName field value if set, zero value otherwise.
func (o *GetDeviceClients200ResponseInner) GetMdnsName() string {
	if o == nil || IsNil(o.MdnsName) {
		var ret string
		return ret
	}
	return *o.MdnsName
}

// GetMdnsNameOk returns a tuple with the MdnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceClients200ResponseInner) GetMdnsNameOk() (*string, bool) {
	if o == nil || IsNil(o.MdnsName) {
		return nil, false
	}
	return o.MdnsName, true
}

// HasMdnsName returns a boolean if a field has been set.
func (o *GetDeviceClients200ResponseInner) HasMdnsName() bool {
	if o != nil && !IsNil(o.MdnsName) {
		return true
	}

	return false
}

// SetMdnsName gets a reference to the given string and assigns it to the MdnsName field.
func (o *GetDeviceClients200ResponseInner) SetMdnsName(v string) {
	o.MdnsName = &v
}

// GetDhcpHostname returns the DhcpHostname field value if set, zero value otherwise.
func (o *GetDeviceClients200ResponseInner) GetDhcpHostname() string {
	if o == nil || IsNil(o.DhcpHostname) {
		var ret string
		return ret
	}
	return *o.DhcpHostname
}

// GetDhcpHostnameOk returns a tuple with the DhcpHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceClients200ResponseInner) GetDhcpHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.DhcpHostname) {
		return nil, false
	}
	return o.DhcpHostname, true
}

// HasDhcpHostname returns a boolean if a field has been set.
func (o *GetDeviceClients200ResponseInner) HasDhcpHostname() bool {
	if o != nil && !IsNil(o.DhcpHostname) {
		return true
	}

	return false
}

// SetDhcpHostname gets a reference to the given string and assigns it to the DhcpHostname field.
func (o *GetDeviceClients200ResponseInner) SetDhcpHostname(v string) {
	o.DhcpHostname = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *GetDeviceClients200ResponseInner) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceClients200ResponseInner) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *GetDeviceClients200ResponseInner) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *GetDeviceClients200ResponseInner) SetUser(v string) {
	o.User = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *GetDeviceClients200ResponseInner) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceClients200ResponseInner) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *GetDeviceClients200ResponseInner) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *GetDeviceClients200ResponseInner) SetIp(v string) {
	o.Ip = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise.
func (o *GetDeviceClients200ResponseInner) GetVlan() string {
	if o == nil || IsNil(o.Vlan) {
		var ret string
		return ret
	}
	return *o.Vlan
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceClients200ResponseInner) GetVlanOk() (*string, bool) {
	if o == nil || IsNil(o.Vlan) {
		return nil, false
	}
	return o.Vlan, true
}

// HasVlan returns a boolean if a field has been set.
func (o *GetDeviceClients200ResponseInner) HasVlan() bool {
	if o != nil && !IsNil(o.Vlan) {
		return true
	}

	return false
}

// SetVlan gets a reference to the given string and assigns it to the Vlan field.
func (o *GetDeviceClients200ResponseInner) SetVlan(v string) {
	o.Vlan = &v
}

// GetNamedVlan returns the NamedVlan field value if set, zero value otherwise.
func (o *GetDeviceClients200ResponseInner) GetNamedVlan() string {
	if o == nil || IsNil(o.NamedVlan) {
		var ret string
		return ret
	}
	return *o.NamedVlan
}

// GetNamedVlanOk returns a tuple with the NamedVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceClients200ResponseInner) GetNamedVlanOk() (*string, bool) {
	if o == nil || IsNil(o.NamedVlan) {
		return nil, false
	}
	return o.NamedVlan, true
}

// HasNamedVlan returns a boolean if a field has been set.
func (o *GetDeviceClients200ResponseInner) HasNamedVlan() bool {
	if o != nil && !IsNil(o.NamedVlan) {
		return true
	}

	return false
}

// SetNamedVlan gets a reference to the given string and assigns it to the NamedVlan field.
func (o *GetDeviceClients200ResponseInner) SetNamedVlan(v string) {
	o.NamedVlan = &v
}

// GetSwitchport returns the Switchport field value if set, zero value otherwise.
func (o *GetDeviceClients200ResponseInner) GetSwitchport() string {
	if o == nil || IsNil(o.Switchport) {
		var ret string
		return ret
	}
	return *o.Switchport
}

// GetSwitchportOk returns a tuple with the Switchport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceClients200ResponseInner) GetSwitchportOk() (*string, bool) {
	if o == nil || IsNil(o.Switchport) {
		return nil, false
	}
	return o.Switchport, true
}

// HasSwitchport returns a boolean if a field has been set.
func (o *GetDeviceClients200ResponseInner) HasSwitchport() bool {
	if o != nil && !IsNil(o.Switchport) {
		return true
	}

	return false
}

// SetSwitchport gets a reference to the given string and assigns it to the Switchport field.
func (o *GetDeviceClients200ResponseInner) SetSwitchport(v string) {
	o.Switchport = &v
}

// GetAdaptivePolicyGroup returns the AdaptivePolicyGroup field value if set, zero value otherwise.
func (o *GetDeviceClients200ResponseInner) GetAdaptivePolicyGroup() string {
	if o == nil || IsNil(o.AdaptivePolicyGroup) {
		var ret string
		return ret
	}
	return *o.AdaptivePolicyGroup
}

// GetAdaptivePolicyGroupOk returns a tuple with the AdaptivePolicyGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceClients200ResponseInner) GetAdaptivePolicyGroupOk() (*string, bool) {
	if o == nil || IsNil(o.AdaptivePolicyGroup) {
		return nil, false
	}
	return o.AdaptivePolicyGroup, true
}

// HasAdaptivePolicyGroup returns a boolean if a field has been set.
func (o *GetDeviceClients200ResponseInner) HasAdaptivePolicyGroup() bool {
	if o != nil && !IsNil(o.AdaptivePolicyGroup) {
		return true
	}

	return false
}

// SetAdaptivePolicyGroup gets a reference to the given string and assigns it to the AdaptivePolicyGroup field.
func (o *GetDeviceClients200ResponseInner) SetAdaptivePolicyGroup(v string) {
	o.AdaptivePolicyGroup = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *GetDeviceClients200ResponseInner) GetUsage() GetDeviceClients200ResponseInnerUsage {
	if o == nil || IsNil(o.Usage) {
		var ret GetDeviceClients200ResponseInnerUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceClients200ResponseInner) GetUsageOk() (*GetDeviceClients200ResponseInnerUsage, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *GetDeviceClients200ResponseInner) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given GetDeviceClients200ResponseInnerUsage and assigns it to the Usage field.
func (o *GetDeviceClients200ResponseInner) SetUsage(v GetDeviceClients200ResponseInnerUsage) {
	o.Usage = &v
}

func (o GetDeviceClients200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceClients200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MdnsName) {
		toSerialize["mdnsName"] = o.MdnsName
	}
	if !IsNil(o.DhcpHostname) {
		toSerialize["dhcpHostname"] = o.DhcpHostname
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Vlan) {
		toSerialize["vlan"] = o.Vlan
	}
	if !IsNil(o.NamedVlan) {
		toSerialize["namedVlan"] = o.NamedVlan
	}
	if !IsNil(o.Switchport) {
		toSerialize["switchport"] = o.Switchport
	}
	if !IsNil(o.AdaptivePolicyGroup) {
		toSerialize["adaptivePolicyGroup"] = o.AdaptivePolicyGroup
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return toSerialize, nil
}

type NullableGetDeviceClients200ResponseInner struct {
	value *GetDeviceClients200ResponseInner
	isSet bool
}

func (v NullableGetDeviceClients200ResponseInner) Get() *GetDeviceClients200ResponseInner {
	return v.value
}

func (v *NullableGetDeviceClients200ResponseInner) Set(val *GetDeviceClients200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceClients200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceClients200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceClients200ResponseInner(val *GetDeviceClients200ResponseInner) *NullableGetDeviceClients200ResponseInner {
	return &NullableGetDeviceClients200ResponseInner{value: val, isSet: true}
}

func (v NullableGetDeviceClients200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceClients200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


