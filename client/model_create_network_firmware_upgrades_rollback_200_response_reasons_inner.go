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

// checks if the CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner{}

// CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner struct for CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner
type CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner struct {
	// Reason for the rollback
	Category *string `json:"category,omitempty"`
	// Additional comment about the rollback
	Comment *string `json:"comment,omitempty"`
}

// NewCreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner instantiates a new CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner() *CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner {
	this := CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner{}
	return &this
}

// NewCreateNetworkFirmwareUpgradesRollback200ResponseReasonsInnerWithDefaults instantiates a new CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkFirmwareUpgradesRollback200ResponseReasonsInnerWithDefaults() *CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner {
	this := CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner) SetCategory(v string) {
	o.Category = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner) SetComment(v string) {
	o.Comment = &v
}

func (o CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return toSerialize, nil
}

type NullableCreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner struct {
	value *CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner
	isSet bool
}

func (v NullableCreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner) Get() *CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner {
	return v.value
}

func (v *NullableCreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner) Set(val *CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner(val *CreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner) *NullableCreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner {
	return &NullableCreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner{value: val, isSet: true}
}

func (v NullableCreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkFirmwareUpgradesRollback200ResponseReasonsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


