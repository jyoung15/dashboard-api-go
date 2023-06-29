/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner{}

// CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner struct for CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner
type CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner struct {
	// Profile identifier.
	Profile string `json:"profile"`
	// Port identifier of switch port. For modules, the identifier is \"SlotNumber_ModuleType_PortNumber\" (Ex: \"1_8X10G_1\"), otherwise it is just the port number (Ex: \"8\").
	PortId string `json:"portId"`
}

// NewCreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner instantiates a new CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner(profile string, portId string) *CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner {
	this := CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner{}
	this.Profile = profile
	this.PortId = portId
	return &this
}

// NewCreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInnerWithDefaults instantiates a new CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInnerWithDefaults() *CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner {
	this := CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner{}
	return &this
}

// GetProfile returns the Profile field value
func (o *CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner) GetProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner) GetProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner) SetProfile(v string) {
	o.Profile = v
}

// GetPortId returns the PortId field value
func (o *CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner) GetPortId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner) GetPortIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PortId, true
}

// SetPortId sets field value
func (o *CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner) SetPortId(v string) {
	o.PortId = v
}

func (o CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["profile"] = o.Profile
	toSerialize["portId"] = o.PortId
	return toSerialize, nil
}

type NullableCreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner struct {
	value *CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner
	isSet bool
}

func (v NullableCreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner) Get() *CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner {
	return v.value
}

func (v *NullableCreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner) Set(val *CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner(val *CreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner) *NullableCreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner {
	return &NullableCreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner{value: val, isSet: true}
}

func (v NullableCreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkSwitchLinkAggregationRequestSwitchProfilePortsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


