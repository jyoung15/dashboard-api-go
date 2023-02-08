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

// GetNetworkApplianceSettings200Response struct for GetNetworkApplianceSettings200Response
type GetNetworkApplianceSettings200Response struct {
	// Client tracking method of a network
	ClientTrackingMethod *string `json:"clientTrackingMethod,omitempty"`
	// Deployment mode of a network
	DeploymentMode *string `json:"deploymentMode,omitempty"`
	DynamicDns *GetNetworkApplianceSettings200ResponseDynamicDns `json:"dynamicDns,omitempty"`
}

// NewGetNetworkApplianceSettings200Response instantiates a new GetNetworkApplianceSettings200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceSettings200Response() *GetNetworkApplianceSettings200Response {
	this := GetNetworkApplianceSettings200Response{}
	return &this
}

// NewGetNetworkApplianceSettings200ResponseWithDefaults instantiates a new GetNetworkApplianceSettings200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceSettings200ResponseWithDefaults() *GetNetworkApplianceSettings200Response {
	this := GetNetworkApplianceSettings200Response{}
	return &this
}

// GetClientTrackingMethod returns the ClientTrackingMethod field value if set, zero value otherwise.
func (o *GetNetworkApplianceSettings200Response) GetClientTrackingMethod() string {
	if o == nil || isNil(o.ClientTrackingMethod) {
		var ret string
		return ret
	}
	return *o.ClientTrackingMethod
}

// GetClientTrackingMethodOk returns a tuple with the ClientTrackingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSettings200Response) GetClientTrackingMethodOk() (*string, bool) {
	if o == nil || isNil(o.ClientTrackingMethod) {
    return nil, false
	}
	return o.ClientTrackingMethod, true
}

// HasClientTrackingMethod returns a boolean if a field has been set.
func (o *GetNetworkApplianceSettings200Response) HasClientTrackingMethod() bool {
	if o != nil && !isNil(o.ClientTrackingMethod) {
		return true
	}

	return false
}

// SetClientTrackingMethod gets a reference to the given string and assigns it to the ClientTrackingMethod field.
func (o *GetNetworkApplianceSettings200Response) SetClientTrackingMethod(v string) {
	o.ClientTrackingMethod = &v
}

// GetDeploymentMode returns the DeploymentMode field value if set, zero value otherwise.
func (o *GetNetworkApplianceSettings200Response) GetDeploymentMode() string {
	if o == nil || isNil(o.DeploymentMode) {
		var ret string
		return ret
	}
	return *o.DeploymentMode
}

// GetDeploymentModeOk returns a tuple with the DeploymentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSettings200Response) GetDeploymentModeOk() (*string, bool) {
	if o == nil || isNil(o.DeploymentMode) {
    return nil, false
	}
	return o.DeploymentMode, true
}

// HasDeploymentMode returns a boolean if a field has been set.
func (o *GetNetworkApplianceSettings200Response) HasDeploymentMode() bool {
	if o != nil && !isNil(o.DeploymentMode) {
		return true
	}

	return false
}

// SetDeploymentMode gets a reference to the given string and assigns it to the DeploymentMode field.
func (o *GetNetworkApplianceSettings200Response) SetDeploymentMode(v string) {
	o.DeploymentMode = &v
}

// GetDynamicDns returns the DynamicDns field value if set, zero value otherwise.
func (o *GetNetworkApplianceSettings200Response) GetDynamicDns() GetNetworkApplianceSettings200ResponseDynamicDns {
	if o == nil || isNil(o.DynamicDns) {
		var ret GetNetworkApplianceSettings200ResponseDynamicDns
		return ret
	}
	return *o.DynamicDns
}

// GetDynamicDnsOk returns a tuple with the DynamicDns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceSettings200Response) GetDynamicDnsOk() (*GetNetworkApplianceSettings200ResponseDynamicDns, bool) {
	if o == nil || isNil(o.DynamicDns) {
    return nil, false
	}
	return o.DynamicDns, true
}

// HasDynamicDns returns a boolean if a field has been set.
func (o *GetNetworkApplianceSettings200Response) HasDynamicDns() bool {
	if o != nil && !isNil(o.DynamicDns) {
		return true
	}

	return false
}

// SetDynamicDns gets a reference to the given GetNetworkApplianceSettings200ResponseDynamicDns and assigns it to the DynamicDns field.
func (o *GetNetworkApplianceSettings200Response) SetDynamicDns(v GetNetworkApplianceSettings200ResponseDynamicDns) {
	o.DynamicDns = &v
}

func (o GetNetworkApplianceSettings200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ClientTrackingMethod) {
		toSerialize["clientTrackingMethod"] = o.ClientTrackingMethod
	}
	if !isNil(o.DeploymentMode) {
		toSerialize["deploymentMode"] = o.DeploymentMode
	}
	if !isNil(o.DynamicDns) {
		toSerialize["dynamicDns"] = o.DynamicDns
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkApplianceSettings200Response struct {
	value *GetNetworkApplianceSettings200Response
	isSet bool
}

func (v NullableGetNetworkApplianceSettings200Response) Get() *GetNetworkApplianceSettings200Response {
	return v.value
}

func (v *NullableGetNetworkApplianceSettings200Response) Set(val *GetNetworkApplianceSettings200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceSettings200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceSettings200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceSettings200Response(val *GetNetworkApplianceSettings200Response) *NullableGetNetworkApplianceSettings200Response {
	return &NullableGetNetworkApplianceSettings200Response{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceSettings200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceSettings200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


