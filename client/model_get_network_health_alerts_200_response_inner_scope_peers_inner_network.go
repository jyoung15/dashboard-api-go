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

// checks if the GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork{}

// GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork Network of the peer
type GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork struct {
	// Name of the network
	Name *string `json:"name,omitempty"`
	// Id of the network
	Id *string `json:"id,omitempty"`
}

// NewGetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork instantiates a new GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork() *GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork {
	this := GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork{}
	return &this
}

// NewGetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetworkWithDefaults instantiates a new GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetworkWithDefaults() *GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork {
	this := GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork) SetId(v string) {
	o.Id = &v
}

func (o GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableGetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork struct {
	value *GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork
	isSet bool
}

func (v NullableGetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork) Get() *GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork {
	return v.value
}

func (v *NullableGetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork) Set(val *GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork(val *GetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork) *NullableGetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork {
	return &NullableGetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork{value: val, isSet: true}
}

func (v NullableGetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkHealthAlerts200ResponseInnerScopePeersInnerNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


