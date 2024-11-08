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

// checks if the UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest{}

// UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest struct for UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest
type UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest struct {
	// The list of VPN peers
	Peers []UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner `json:"peers"`
}

type _UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest

// NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest instantiates a new UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest(peers []UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest {
	this := UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest{}
	this.Peers = peers
	return &this
}

// NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestWithDefaults instantiates a new UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestWithDefaults() *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest {
	this := UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest{}
	return &this
}

// GetPeers returns the Peers field value
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) GetPeers() []UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner {
	if o == nil {
		var ret []UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner
		return ret
	}

	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) GetPeersOk() ([]UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Peers, true
}

// SetPeers sets field value
func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) SetPeers(v []UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner) {
	o.Peers = v
}

func (o UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["peers"] = o.Peers
	return toSerialize, nil
}

func (o *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"peers",
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

	varUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest := _UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest)

	if err != nil {
		return err
	}

	*o = UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest(varUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest)

	return err
}

type NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest struct {
	value *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest
	isSet bool
}

func (v NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) Get() *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest {
	return v.value
}

func (v *NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) Set(val *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest(val *UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) *NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest {
	return &NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


