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

// checks if the GetNetworkSwitchRoutingOspf200ResponseAreasInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchRoutingOspf200ResponseAreasInner{}

// GetNetworkSwitchRoutingOspf200ResponseAreasInner struct for GetNetworkSwitchRoutingOspf200ResponseAreasInner
type GetNetworkSwitchRoutingOspf200ResponseAreasInner struct {
	// OSPF area ID
	AreaId *string `json:"areaId,omitempty"`
	// Name of the OSPF area
	AreaName *string `json:"areaName,omitempty"`
	// Area types in OSPF. Must be one of: [\"normal\", \"stub\", \"nssa\"]
	AreaType *string `json:"areaType,omitempty"`
}

// NewGetNetworkSwitchRoutingOspf200ResponseAreasInner instantiates a new GetNetworkSwitchRoutingOspf200ResponseAreasInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchRoutingOspf200ResponseAreasInner() *GetNetworkSwitchRoutingOspf200ResponseAreasInner {
	this := GetNetworkSwitchRoutingOspf200ResponseAreasInner{}
	return &this
}

// NewGetNetworkSwitchRoutingOspf200ResponseAreasInnerWithDefaults instantiates a new GetNetworkSwitchRoutingOspf200ResponseAreasInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchRoutingOspf200ResponseAreasInnerWithDefaults() *GetNetworkSwitchRoutingOspf200ResponseAreasInner {
	this := GetNetworkSwitchRoutingOspf200ResponseAreasInner{}
	return &this
}

// GetAreaId returns the AreaId field value if set, zero value otherwise.
func (o *GetNetworkSwitchRoutingOspf200ResponseAreasInner) GetAreaId() string {
	if o == nil || IsNil(o.AreaId) {
		var ret string
		return ret
	}
	return *o.AreaId
}

// GetAreaIdOk returns a tuple with the AreaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchRoutingOspf200ResponseAreasInner) GetAreaIdOk() (*string, bool) {
	if o == nil || IsNil(o.AreaId) {
		return nil, false
	}
	return o.AreaId, true
}

// HasAreaId returns a boolean if a field has been set.
func (o *GetNetworkSwitchRoutingOspf200ResponseAreasInner) HasAreaId() bool {
	if o != nil && !IsNil(o.AreaId) {
		return true
	}

	return false
}

// SetAreaId gets a reference to the given string and assigns it to the AreaId field.
func (o *GetNetworkSwitchRoutingOspf200ResponseAreasInner) SetAreaId(v string) {
	o.AreaId = &v
}

// GetAreaName returns the AreaName field value if set, zero value otherwise.
func (o *GetNetworkSwitchRoutingOspf200ResponseAreasInner) GetAreaName() string {
	if o == nil || IsNil(o.AreaName) {
		var ret string
		return ret
	}
	return *o.AreaName
}

// GetAreaNameOk returns a tuple with the AreaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchRoutingOspf200ResponseAreasInner) GetAreaNameOk() (*string, bool) {
	if o == nil || IsNil(o.AreaName) {
		return nil, false
	}
	return o.AreaName, true
}

// HasAreaName returns a boolean if a field has been set.
func (o *GetNetworkSwitchRoutingOspf200ResponseAreasInner) HasAreaName() bool {
	if o != nil && !IsNil(o.AreaName) {
		return true
	}

	return false
}

// SetAreaName gets a reference to the given string and assigns it to the AreaName field.
func (o *GetNetworkSwitchRoutingOspf200ResponseAreasInner) SetAreaName(v string) {
	o.AreaName = &v
}

// GetAreaType returns the AreaType field value if set, zero value otherwise.
func (o *GetNetworkSwitchRoutingOspf200ResponseAreasInner) GetAreaType() string {
	if o == nil || IsNil(o.AreaType) {
		var ret string
		return ret
	}
	return *o.AreaType
}

// GetAreaTypeOk returns a tuple with the AreaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchRoutingOspf200ResponseAreasInner) GetAreaTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AreaType) {
		return nil, false
	}
	return o.AreaType, true
}

// HasAreaType returns a boolean if a field has been set.
func (o *GetNetworkSwitchRoutingOspf200ResponseAreasInner) HasAreaType() bool {
	if o != nil && !IsNil(o.AreaType) {
		return true
	}

	return false
}

// SetAreaType gets a reference to the given string and assigns it to the AreaType field.
func (o *GetNetworkSwitchRoutingOspf200ResponseAreasInner) SetAreaType(v string) {
	o.AreaType = &v
}

func (o GetNetworkSwitchRoutingOspf200ResponseAreasInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchRoutingOspf200ResponseAreasInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AreaId) {
		toSerialize["areaId"] = o.AreaId
	}
	if !IsNil(o.AreaName) {
		toSerialize["areaName"] = o.AreaName
	}
	if !IsNil(o.AreaType) {
		toSerialize["areaType"] = o.AreaType
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchRoutingOspf200ResponseAreasInner struct {
	value *GetNetworkSwitchRoutingOspf200ResponseAreasInner
	isSet bool
}

func (v NullableGetNetworkSwitchRoutingOspf200ResponseAreasInner) Get() *GetNetworkSwitchRoutingOspf200ResponseAreasInner {
	return v.value
}

func (v *NullableGetNetworkSwitchRoutingOspf200ResponseAreasInner) Set(val *GetNetworkSwitchRoutingOspf200ResponseAreasInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchRoutingOspf200ResponseAreasInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchRoutingOspf200ResponseAreasInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchRoutingOspf200ResponseAreasInner(val *GetNetworkSwitchRoutingOspf200ResponseAreasInner) *NullableGetNetworkSwitchRoutingOspf200ResponseAreasInner {
	return &NullableGetNetworkSwitchRoutingOspf200ResponseAreasInner{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchRoutingOspf200ResponseAreasInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchRoutingOspf200ResponseAreasInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


