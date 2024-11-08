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

// checks if the CreateNetworkFirmwareUpgradesStagedGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkFirmwareUpgradesStagedGroupRequest{}

// CreateNetworkFirmwareUpgradesStagedGroupRequest struct for CreateNetworkFirmwareUpgradesStagedGroupRequest
type CreateNetworkFirmwareUpgradesStagedGroupRequest struct {
	// Name of the Staged Upgrade Group. Length must be 1 to 255 characters
	Name string `json:"name"`
	// Description of the Staged Upgrade Group. Length must be 1 to 255 characters
	Description *string `json:"description,omitempty"`
	// Boolean indicating the default Group. Any device that does not have a group explicitly assigned will upgrade with this group
	IsDefault bool `json:"isDefault"`
	AssignedDevices *CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices `json:"assignedDevices,omitempty"`
}

type _CreateNetworkFirmwareUpgradesStagedGroupRequest CreateNetworkFirmwareUpgradesStagedGroupRequest

// NewCreateNetworkFirmwareUpgradesStagedGroupRequest instantiates a new CreateNetworkFirmwareUpgradesStagedGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkFirmwareUpgradesStagedGroupRequest(name string, isDefault bool) *CreateNetworkFirmwareUpgradesStagedGroupRequest {
	this := CreateNetworkFirmwareUpgradesStagedGroupRequest{}
	this.Name = name
	this.IsDefault = isDefault
	return &this
}

// NewCreateNetworkFirmwareUpgradesStagedGroupRequestWithDefaults instantiates a new CreateNetworkFirmwareUpgradesStagedGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkFirmwareUpgradesStagedGroupRequestWithDefaults() *CreateNetworkFirmwareUpgradesStagedGroupRequest {
	this := CreateNetworkFirmwareUpgradesStagedGroupRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) SetDescription(v string) {
	o.Description = &v
}

// GetIsDefault returns the IsDefault field value
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) SetIsDefault(v bool) {
	o.IsDefault = v
}

// GetAssignedDevices returns the AssignedDevices field value if set, zero value otherwise.
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) GetAssignedDevices() CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices {
	if o == nil || IsNil(o.AssignedDevices) {
		var ret CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices
		return ret
	}
	return *o.AssignedDevices
}

// GetAssignedDevicesOk returns a tuple with the AssignedDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) GetAssignedDevicesOk() (*CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices, bool) {
	if o == nil || IsNil(o.AssignedDevices) {
		return nil, false
	}
	return o.AssignedDevices, true
}

// HasAssignedDevices returns a boolean if a field has been set.
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) HasAssignedDevices() bool {
	if o != nil && !IsNil(o.AssignedDevices) {
		return true
	}

	return false
}

// SetAssignedDevices gets a reference to the given CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices and assigns it to the AssignedDevices field.
func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) SetAssignedDevices(v CreateNetworkFirmwareUpgradesStagedGroupRequestAssignedDevices) {
	o.AssignedDevices = &v
}

func (o CreateNetworkFirmwareUpgradesStagedGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkFirmwareUpgradesStagedGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["isDefault"] = o.IsDefault
	if !IsNil(o.AssignedDevices) {
		toSerialize["assignedDevices"] = o.AssignedDevices
	}
	return toSerialize, nil
}

func (o *CreateNetworkFirmwareUpgradesStagedGroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"isDefault",
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

	varCreateNetworkFirmwareUpgradesStagedGroupRequest := _CreateNetworkFirmwareUpgradesStagedGroupRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateNetworkFirmwareUpgradesStagedGroupRequest)

	if err != nil {
		return err
	}

	*o = CreateNetworkFirmwareUpgradesStagedGroupRequest(varCreateNetworkFirmwareUpgradesStagedGroupRequest)

	return err
}

type NullableCreateNetworkFirmwareUpgradesStagedGroupRequest struct {
	value *CreateNetworkFirmwareUpgradesStagedGroupRequest
	isSet bool
}

func (v NullableCreateNetworkFirmwareUpgradesStagedGroupRequest) Get() *CreateNetworkFirmwareUpgradesStagedGroupRequest {
	return v.value
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedGroupRequest) Set(val *CreateNetworkFirmwareUpgradesStagedGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkFirmwareUpgradesStagedGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkFirmwareUpgradesStagedGroupRequest(val *CreateNetworkFirmwareUpgradesStagedGroupRequest) *NullableCreateNetworkFirmwareUpgradesStagedGroupRequest {
	return &NullableCreateNetworkFirmwareUpgradesStagedGroupRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkFirmwareUpgradesStagedGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


