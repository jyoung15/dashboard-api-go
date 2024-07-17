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

// checks if the CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices{}

// CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices The devices involved in the swap
type CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices struct {
	Old CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevicesOld `json:"old"`
	New CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevicesNew `json:"new"`
}

type _CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices

// NewCreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices instantiates a new CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices(old CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevicesOld, new CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevicesNew) *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices {
	this := CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices{}
	this.Old = old
	this.New = new
	return &this
}

// NewCreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevicesWithDefaults instantiates a new CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevicesWithDefaults() *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices {
	this := CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices{}
	return &this
}

// GetOld returns the Old field value
func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices) GetOld() CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevicesOld {
	if o == nil {
		var ret CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevicesOld
		return ret
	}

	return o.Old
}

// GetOldOk returns a tuple with the Old field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices) GetOldOk() (*CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevicesOld, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Old, true
}

// SetOld sets field value
func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices) SetOld(v CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevicesOld) {
	o.Old = v
}

// GetNew returns the New field value
func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices) GetNew() CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevicesNew {
	if o == nil {
		var ret CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevicesNew
		return ret
	}

	return o.New
}

// GetNewOk returns a tuple with the New field value
// and a boolean to check if the value has been set.
func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices) GetNewOk() (*CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevicesNew, bool) {
	if o == nil {
		return nil, false
	}
	return &o.New, true
}

// SetNew sets field value
func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices) SetNew(v CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevicesNew) {
	o.New = v
}

func (o CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["old"] = o.Old
	toSerialize["new"] = o.New
	return toSerialize, nil
}

func (o *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"old",
		"new",
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

	varCreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices := _CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices)

	if err != nil {
		return err
	}

	*o = CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices(varCreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices)

	return err
}

type NullableCreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices struct {
	value *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices
	isSet bool
}

func (v NullableCreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices) Get() *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices {
	return v.value
}

func (v *NullableCreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices) Set(val *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices(val *CreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices) *NullableCreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices {
	return &NullableCreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices{value: val, isSet: true}
}

func (v NullableCreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrganizationInventoryDevicesSwapsBulk207ResponseSwapsInnerDevices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


