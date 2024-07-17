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

// checks if the GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner{}

// GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner struct for GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner
type GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner struct {
	// Subnet
	Subnet *string `json:"subnet,omitempty"`
	// Name of the subnet
	Name *string `json:"name,omitempty"`
}

// NewGetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner instantiates a new GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner() *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner {
	this := GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner{}
	return &this
}

// NewGetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInnerWithDefaults instantiates a new GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInnerWithDefaults() *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner {
	this := GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner{}
	return &this
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner) GetSubnet() string {
	if o == nil || IsNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner) GetSubnetOk() (*string, bool) {
	if o == nil || IsNil(o.Subnet) {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner) HasSubnet() bool {
	if o != nil && !IsNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner) SetSubnet(v string) {
	o.Subnet = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner) SetName(v string) {
	o.Name = &v
}

func (o GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner struct {
	value *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner
	isSet bool
}

func (v NullableGetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner) Get() *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner {
	return v.value
}

func (v *NullableGetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner) Set(val *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner(val *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner) *NullableGetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner {
	return &NullableGetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner{value: val, isSet: true}
}

func (v NullableGetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


