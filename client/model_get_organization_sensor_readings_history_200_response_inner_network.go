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

// checks if the GetOrganizationSensorReadingsHistory200ResponseInnerNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSensorReadingsHistory200ResponseInnerNetwork{}

// GetOrganizationSensorReadingsHistory200ResponseInnerNetwork Network to which the sensor belongs.
type GetOrganizationSensorReadingsHistory200ResponseInnerNetwork struct {
	// ID of the network.
	Id *string `json:"id,omitempty"`
	// Name of the network.
	Name *string `json:"name,omitempty"`
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerNetwork instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSensorReadingsHistory200ResponseInnerNetwork() *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerNetwork{}
	return &this
}

// NewGetOrganizationSensorReadingsHistory200ResponseInnerNetworkWithDefaults instantiates a new GetOrganizationSensorReadingsHistory200ResponseInnerNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSensorReadingsHistory200ResponseInnerNetworkWithDefaults() *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork {
	this := GetOrganizationSensorReadingsHistory200ResponseInnerNetwork{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) SetName(v string) {
	o.Name = &v
}

func (o GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGetOrganizationSensorReadingsHistory200ResponseInnerNetwork struct {
	value *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork
	isSet bool
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerNetwork) Get() *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork {
	return v.value
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerNetwork) Set(val *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSensorReadingsHistory200ResponseInnerNetwork(val *GetOrganizationSensorReadingsHistory200ResponseInnerNetwork) *NullableGetOrganizationSensorReadingsHistory200ResponseInnerNetwork {
	return &NullableGetOrganizationSensorReadingsHistory200ResponseInnerNetwork{value: val, isSet: true}
}

func (v NullableGetOrganizationSensorReadingsHistory200ResponseInnerNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSensorReadingsHistory200ResponseInnerNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


