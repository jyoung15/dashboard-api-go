/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones{}

// UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones The Staged Upgrade Milestones for the specific stage
type UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones struct {
	// The start time of the staged upgrade stage. (In ISO-8601 format, in the time zone of the network.)
	ScheduledFor string `json:"scheduledFor"`
}

type _UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones

// NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones instantiates a new UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones(scheduledFor string) *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones {
	this := UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones{}
	this.ScheduledFor = scheduledFor
	return &this
}

// NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestonesWithDefaults instantiates a new UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestonesWithDefaults() *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones {
	this := UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones{}
	return &this
}

// GetScheduledFor returns the ScheduledFor field value
func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) GetScheduledFor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScheduledFor
}

// GetScheduledForOk returns a tuple with the ScheduledFor field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) GetScheduledForOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledFor, true
}

// SetScheduledFor sets field value
func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) SetScheduledFor(v string) {
	o.ScheduledFor = v
}

func (o UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["scheduledFor"] = o.ScheduledFor
	return toSerialize, nil
}

func (o *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"scheduledFor",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones := _UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones)

	if err != nil {
		return err
	}

	*o = UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones(varUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones)

	return err
}

type NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones struct {
	value *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones
	isSet bool
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) Get() *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones {
	return v.value
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) Set(val *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones(val *UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) *NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones {
	return &NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones{value: val, isSet: true}
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInnerMilestones) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


