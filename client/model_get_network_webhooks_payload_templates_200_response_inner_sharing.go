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

// checks if the GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing{}

// GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing Information on which entities have access to the template
type GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing struct {
	ByNetwork *GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork `json:"byNetwork,omitempty"`
}

// NewGetNetworkWebhooksPayloadTemplates200ResponseInnerSharing instantiates a new GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWebhooksPayloadTemplates200ResponseInnerSharing() *GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing {
	this := GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing{}
	return &this
}

// NewGetNetworkWebhooksPayloadTemplates200ResponseInnerSharingWithDefaults instantiates a new GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWebhooksPayloadTemplates200ResponseInnerSharingWithDefaults() *GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing {
	this := GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing{}
	return &this
}

// GetByNetwork returns the ByNetwork field value if set, zero value otherwise.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing) GetByNetwork() GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork {
	if o == nil || IsNil(o.ByNetwork) {
		var ret GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork
		return ret
	}
	return *o.ByNetwork
}

// GetByNetworkOk returns a tuple with the ByNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing) GetByNetworkOk() (*GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork, bool) {
	if o == nil || IsNil(o.ByNetwork) {
		return nil, false
	}
	return o.ByNetwork, true
}

// HasByNetwork returns a boolean if a field has been set.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing) HasByNetwork() bool {
	if o != nil && !IsNil(o.ByNetwork) {
		return true
	}

	return false
}

// SetByNetwork gets a reference to the given GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork and assigns it to the ByNetwork field.
func (o *GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing) SetByNetwork(v GetNetworkWebhooksPayloadTemplates200ResponseInnerSharingByNetwork) {
	o.ByNetwork = &v
}

func (o GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ByNetwork) {
		toSerialize["byNetwork"] = o.ByNetwork
	}
	return toSerialize, nil
}

type NullableGetNetworkWebhooksPayloadTemplates200ResponseInnerSharing struct {
	value *GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing
	isSet bool
}

func (v NullableGetNetworkWebhooksPayloadTemplates200ResponseInnerSharing) Get() *GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing {
	return v.value
}

func (v *NullableGetNetworkWebhooksPayloadTemplates200ResponseInnerSharing) Set(val *GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWebhooksPayloadTemplates200ResponseInnerSharing) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWebhooksPayloadTemplates200ResponseInnerSharing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWebhooksPayloadTemplates200ResponseInnerSharing(val *GetNetworkWebhooksPayloadTemplates200ResponseInnerSharing) *NullableGetNetworkWebhooksPayloadTemplates200ResponseInnerSharing {
	return &NullableGetNetworkWebhooksPayloadTemplates200ResponseInnerSharing{value: val, isSet: true}
}

func (v NullableGetNetworkWebhooksPayloadTemplates200ResponseInnerSharing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWebhooksPayloadTemplates200ResponseInnerSharing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


