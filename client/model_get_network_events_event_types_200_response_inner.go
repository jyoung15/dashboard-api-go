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

// GetNetworkEventsEventTypes200ResponseInner struct for GetNetworkEventsEventTypes200ResponseInner
type GetNetworkEventsEventTypes200ResponseInner struct {
	// Event category
	Category *string `json:"category,omitempty"`
	// Event type
	Type *string `json:"type,omitempty"`
	// Description of the event
	Description *string `json:"description,omitempty"`
}

// NewGetNetworkEventsEventTypes200ResponseInner instantiates a new GetNetworkEventsEventTypes200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkEventsEventTypes200ResponseInner() *GetNetworkEventsEventTypes200ResponseInner {
	this := GetNetworkEventsEventTypes200ResponseInner{}
	return &this
}

// NewGetNetworkEventsEventTypes200ResponseInnerWithDefaults instantiates a new GetNetworkEventsEventTypes200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkEventsEventTypes200ResponseInnerWithDefaults() *GetNetworkEventsEventTypes200ResponseInner {
	this := GetNetworkEventsEventTypes200ResponseInner{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *GetNetworkEventsEventTypes200ResponseInner) GetCategory() string {
	if o == nil || isNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkEventsEventTypes200ResponseInner) GetCategoryOk() (*string, bool) {
	if o == nil || isNil(o.Category) {
    return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *GetNetworkEventsEventTypes200ResponseInner) HasCategory() bool {
	if o != nil && !isNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *GetNetworkEventsEventTypes200ResponseInner) SetCategory(v string) {
	o.Category = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetNetworkEventsEventTypes200ResponseInner) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkEventsEventTypes200ResponseInner) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetNetworkEventsEventTypes200ResponseInner) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetNetworkEventsEventTypes200ResponseInner) SetType(v string) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetNetworkEventsEventTypes200ResponseInner) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkEventsEventTypes200ResponseInner) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetNetworkEventsEventTypes200ResponseInner) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetNetworkEventsEventTypes200ResponseInner) SetDescription(v string) {
	o.Description = &v
}

func (o GetNetworkEventsEventTypes200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkEventsEventTypes200ResponseInner struct {
	value *GetNetworkEventsEventTypes200ResponseInner
	isSet bool
}

func (v NullableGetNetworkEventsEventTypes200ResponseInner) Get() *GetNetworkEventsEventTypes200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkEventsEventTypes200ResponseInner) Set(val *GetNetworkEventsEventTypes200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkEventsEventTypes200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkEventsEventTypes200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkEventsEventTypes200ResponseInner(val *GetNetworkEventsEventTypes200ResponseInner) *NullableGetNetworkEventsEventTypes200ResponseInner {
	return &NullableGetNetworkEventsEventTypes200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkEventsEventTypes200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkEventsEventTypes200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


