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

// checks if the GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup{}

// GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup The staged upgrade group
type GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup struct {
	// Id of the Staged Upgrade Group
	Id *string `json:"id,omitempty"`
	// Name of the Staged Upgrade Group
	Name *string `json:"name,omitempty"`
	// Description of the Staged Upgrade Group
	Description *string `json:"description,omitempty"`
}

// NewGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup instantiates a new GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup() *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup {
	this := GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup{}
	return &this
}

// NewGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroupWithDefaults instantiates a new GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroupWithDefaults() *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup {
	this := GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup) SetDescription(v string) {
	o.Description = &v
}

func (o GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup struct {
	value *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup
	isSet bool
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup) Get() *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup {
	return v.value
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup) Set(val *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup(val *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup) *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup {
	return &NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup{value: val, isSet: true}
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


