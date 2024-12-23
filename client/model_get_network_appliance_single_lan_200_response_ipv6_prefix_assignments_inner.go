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

// checks if the GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner{}

// GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner struct for GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner
type GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner struct {
	// Auto assign a /64 prefix from the origin to the single LAN
	Autonomous *bool `json:"autonomous,omitempty"`
	// Manual configuration of a /64 prefix on the single LAN
	StaticPrefix *string `json:"staticPrefix,omitempty"`
	// Manual configuration of the IPv6 Appliance IP
	StaticApplianceIp6 *string `json:"staticApplianceIp6,omitempty"`
	Origin *CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin `json:"origin,omitempty"`
}

// NewGetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner instantiates a new GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner() *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner {
	this := GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner{}
	return &this
}

// NewGetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInnerWithDefaults instantiates a new GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInnerWithDefaults() *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner {
	this := GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner{}
	return &this
}

// GetAutonomous returns the Autonomous field value if set, zero value otherwise.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) GetAutonomous() bool {
	if o == nil || IsNil(o.Autonomous) {
		var ret bool
		return ret
	}
	return *o.Autonomous
}

// GetAutonomousOk returns a tuple with the Autonomous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) GetAutonomousOk() (*bool, bool) {
	if o == nil || IsNil(o.Autonomous) {
		return nil, false
	}
	return o.Autonomous, true
}

// HasAutonomous returns a boolean if a field has been set.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) HasAutonomous() bool {
	if o != nil && !IsNil(o.Autonomous) {
		return true
	}

	return false
}

// SetAutonomous gets a reference to the given bool and assigns it to the Autonomous field.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) SetAutonomous(v bool) {
	o.Autonomous = &v
}

// GetStaticPrefix returns the StaticPrefix field value if set, zero value otherwise.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) GetStaticPrefix() string {
	if o == nil || IsNil(o.StaticPrefix) {
		var ret string
		return ret
	}
	return *o.StaticPrefix
}

// GetStaticPrefixOk returns a tuple with the StaticPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) GetStaticPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.StaticPrefix) {
		return nil, false
	}
	return o.StaticPrefix, true
}

// HasStaticPrefix returns a boolean if a field has been set.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) HasStaticPrefix() bool {
	if o != nil && !IsNil(o.StaticPrefix) {
		return true
	}

	return false
}

// SetStaticPrefix gets a reference to the given string and assigns it to the StaticPrefix field.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) SetStaticPrefix(v string) {
	o.StaticPrefix = &v
}

// GetStaticApplianceIp6 returns the StaticApplianceIp6 field value if set, zero value otherwise.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) GetStaticApplianceIp6() string {
	if o == nil || IsNil(o.StaticApplianceIp6) {
		var ret string
		return ret
	}
	return *o.StaticApplianceIp6
}

// GetStaticApplianceIp6Ok returns a tuple with the StaticApplianceIp6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) GetStaticApplianceIp6Ok() (*string, bool) {
	if o == nil || IsNil(o.StaticApplianceIp6) {
		return nil, false
	}
	return o.StaticApplianceIp6, true
}

// HasStaticApplianceIp6 returns a boolean if a field has been set.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) HasStaticApplianceIp6() bool {
	if o != nil && !IsNil(o.StaticApplianceIp6) {
		return true
	}

	return false
}

// SetStaticApplianceIp6 gets a reference to the given string and assigns it to the StaticApplianceIp6 field.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) SetStaticApplianceIp6(v string) {
	o.StaticApplianceIp6 = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) GetOrigin() CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin {
	if o == nil || IsNil(o.Origin) {
		var ret CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) GetOriginOk() (*CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin and assigns it to the Origin field.
func (o *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) SetOrigin(v CreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin) {
	o.Origin = &v
}

func (o GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Autonomous) {
		toSerialize["autonomous"] = o.Autonomous
	}
	if !IsNil(o.StaticPrefix) {
		toSerialize["staticPrefix"] = o.StaticPrefix
	}
	if !IsNil(o.StaticApplianceIp6) {
		toSerialize["staticApplianceIp6"] = o.StaticApplianceIp6
	}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner struct {
	value *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner
	isSet bool
}

func (v NullableGetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) Get() *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner {
	return v.value
}

func (v *NullableGetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) Set(val *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner(val *GetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) *NullableGetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner {
	return &NullableGetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceSingleLan200ResponseIpv6PrefixAssignmentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


