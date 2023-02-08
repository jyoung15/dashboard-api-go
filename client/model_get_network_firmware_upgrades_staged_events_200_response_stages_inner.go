/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 February, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: v1.30.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner struct for GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner
type GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner struct {
	Group *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup `json:"group,omitempty"`
	Milestones *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones `json:"milestones,omitempty"`
	// Current upgrade status of the group
	Status *string `json:"status,omitempty"`
}

// NewGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner instantiates a new GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner() *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner {
	this := GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner{}
	return &this
}

// NewGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerWithDefaults instantiates a new GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerWithDefaults() *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner {
	this := GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner) GetGroup() GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup {
	if o == nil || isNil(o.Group) {
		var ret GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner) GetGroupOk() (*GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup, bool) {
	if o == nil || isNil(o.Group) {
    return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner) HasGroup() bool {
	if o != nil && !isNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup and assigns it to the Group field.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner) SetGroup(v GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerGroup) {
	o.Group = &v
}

// GetMilestones returns the Milestones field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner) GetMilestones() GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones {
	if o == nil || isNil(o.Milestones) {
		var ret GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones
		return ret
	}
	return *o.Milestones
}

// GetMilestonesOk returns a tuple with the Milestones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner) GetMilestonesOk() (*GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones, bool) {
	if o == nil || isNil(o.Milestones) {
    return nil, false
	}
	return o.Milestones, true
}

// HasMilestones returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner) HasMilestones() bool {
	if o != nil && !isNil(o.Milestones) {
		return true
	}

	return false
}

// SetMilestones gets a reference to the given GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones and assigns it to the Milestones field.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner) SetMilestones(v GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInnerMilestones) {
	o.Milestones = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner) SetStatus(v string) {
	o.Status = &v
}

func (o GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !isNil(o.Milestones) {
		toSerialize["milestones"] = o.Milestones
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner struct {
	value *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner
	isSet bool
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner) Get() *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner {
	return v.value
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner) Set(val *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner(val *GetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner) *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner {
	return &NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner{value: val, isSet: true}
}

func (v NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFirmwareUpgradesStagedEvents200ResponseStagesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


