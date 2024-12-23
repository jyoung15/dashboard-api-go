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

// checks if the GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner{}

// GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner struct for GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner
type GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner struct {
	// The id of the application
	Id *string `json:"id,omitempty"`
	// The name of the application
	Name *string `json:"name,omitempty"`
}

// NewGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner instantiates a new GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner() *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner {
	this := GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner{}
	return &this
}

// NewGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInnerWithDefaults instantiates a new GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInnerWithDefaults() *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner {
	this := GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner) SetName(v string) {
	o.Name = &v
}

func (o GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner struct {
	value *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner
	isSet bool
}

func (v NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner) Get() *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner {
	return v.value
}

func (v *NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner) Set(val *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner(val *GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner) *NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner {
	return &NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200ResponseApplicationCategoriesInnerApplicationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


