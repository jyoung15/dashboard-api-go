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

// checks if the BulkUpdateOrganizationDevicesDetailsRequestDetailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkUpdateOrganizationDevicesDetailsRequestDetailsInner{}

// BulkUpdateOrganizationDevicesDetailsRequestDetailsInner struct for BulkUpdateOrganizationDevicesDetailsRequestDetailsInner
type BulkUpdateOrganizationDevicesDetailsRequestDetailsInner struct {
	// Name of device detail
	Name string `json:"name"`
	// Value of device detail
	Value *string `json:"value,omitempty"`
}

type _BulkUpdateOrganizationDevicesDetailsRequestDetailsInner BulkUpdateOrganizationDevicesDetailsRequestDetailsInner

// NewBulkUpdateOrganizationDevicesDetailsRequestDetailsInner instantiates a new BulkUpdateOrganizationDevicesDetailsRequestDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkUpdateOrganizationDevicesDetailsRequestDetailsInner(name string) *BulkUpdateOrganizationDevicesDetailsRequestDetailsInner {
	this := BulkUpdateOrganizationDevicesDetailsRequestDetailsInner{}
	this.Name = name
	return &this
}

// NewBulkUpdateOrganizationDevicesDetailsRequestDetailsInnerWithDefaults instantiates a new BulkUpdateOrganizationDevicesDetailsRequestDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkUpdateOrganizationDevicesDetailsRequestDetailsInnerWithDefaults() *BulkUpdateOrganizationDevicesDetailsRequestDetailsInner {
	this := BulkUpdateOrganizationDevicesDetailsRequestDetailsInner{}
	return &this
}

// GetName returns the Name field value
func (o *BulkUpdateOrganizationDevicesDetailsRequestDetailsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BulkUpdateOrganizationDevicesDetailsRequestDetailsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BulkUpdateOrganizationDevicesDetailsRequestDetailsInner) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BulkUpdateOrganizationDevicesDetailsRequestDetailsInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUpdateOrganizationDevicesDetailsRequestDetailsInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BulkUpdateOrganizationDevicesDetailsRequestDetailsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *BulkUpdateOrganizationDevicesDetailsRequestDetailsInner) SetValue(v string) {
	o.Value = &v
}

func (o BulkUpdateOrganizationDevicesDetailsRequestDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkUpdateOrganizationDevicesDetailsRequestDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

func (o *BulkUpdateOrganizationDevicesDetailsRequestDetailsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varBulkUpdateOrganizationDevicesDetailsRequestDetailsInner := _BulkUpdateOrganizationDevicesDetailsRequestDetailsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBulkUpdateOrganizationDevicesDetailsRequestDetailsInner)

	if err != nil {
		return err
	}

	*o = BulkUpdateOrganizationDevicesDetailsRequestDetailsInner(varBulkUpdateOrganizationDevicesDetailsRequestDetailsInner)

	return err
}

type NullableBulkUpdateOrganizationDevicesDetailsRequestDetailsInner struct {
	value *BulkUpdateOrganizationDevicesDetailsRequestDetailsInner
	isSet bool
}

func (v NullableBulkUpdateOrganizationDevicesDetailsRequestDetailsInner) Get() *BulkUpdateOrganizationDevicesDetailsRequestDetailsInner {
	return v.value
}

func (v *NullableBulkUpdateOrganizationDevicesDetailsRequestDetailsInner) Set(val *BulkUpdateOrganizationDevicesDetailsRequestDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkUpdateOrganizationDevicesDetailsRequestDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkUpdateOrganizationDevicesDetailsRequestDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkUpdateOrganizationDevicesDetailsRequestDetailsInner(val *BulkUpdateOrganizationDevicesDetailsRequestDetailsInner) *NullableBulkUpdateOrganizationDevicesDetailsRequestDetailsInner {
	return &NullableBulkUpdateOrganizationDevicesDetailsRequestDetailsInner{value: val, isSet: true}
}

func (v NullableBulkUpdateOrganizationDevicesDetailsRequestDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkUpdateOrganizationDevicesDetailsRequestDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


