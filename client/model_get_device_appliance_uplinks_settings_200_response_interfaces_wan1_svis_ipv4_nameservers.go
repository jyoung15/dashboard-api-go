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

// checks if the GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers{}

// GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers The nameserver settings for this SVI.
type GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers struct {
	// Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.
	Addresses []string `json:"addresses,omitempty"`
}

// NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers instantiates a new GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers() *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers {
	this := GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers{}
	return &this
}

// NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4NameserversWithDefaults instantiates a new GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4NameserversWithDefaults() *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers {
	this := GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers) GetAddresses() []string {
	if o == nil || IsNil(o.Addresses) {
		var ret []string
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers) GetAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []string and assigns it to the Addresses field.
func (o *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers) SetAddresses(v []string) {
	o.Addresses = v
}

func (o GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	return toSerialize, nil
}

type NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers struct {
	value *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers
	isSet bool
}

func (v NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers) Get() *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers {
	return v.value
}

func (v *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers) Set(val *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers(val *GetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers) *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers {
	return &NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers{value: val, isSet: true}
}

func (v NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceApplianceUplinksSettings200ResponseInterfacesWan1SvisIpv4Nameservers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


