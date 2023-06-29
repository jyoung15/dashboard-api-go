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

// checks if the GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner{}

// GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner struct for GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner
type GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner struct {
	// The network ID of the hub.
	HubId *string `json:"hubId,omitempty"`
	// Indicates whether default route traffic should be sent to this hub.
	UseDefaultRoute *bool `json:"useDefaultRoute,omitempty"`
}

// NewGetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner instantiates a new GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner() *GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner {
	this := GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner{}
	return &this
}

// NewGetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInnerWithDefaults instantiates a new GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInnerWithDefaults() *GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner {
	this := GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner{}
	return &this
}

// GetHubId returns the HubId field value if set, zero value otherwise.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner) GetHubId() string {
	if o == nil || IsNil(o.HubId) {
		var ret string
		return ret
	}
	return *o.HubId
}

// GetHubIdOk returns a tuple with the HubId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner) GetHubIdOk() (*string, bool) {
	if o == nil || IsNil(o.HubId) {
		return nil, false
	}
	return o.HubId, true
}

// HasHubId returns a boolean if a field has been set.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner) HasHubId() bool {
	if o != nil && !IsNil(o.HubId) {
		return true
	}

	return false
}

// SetHubId gets a reference to the given string and assigns it to the HubId field.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner) SetHubId(v string) {
	o.HubId = &v
}

// GetUseDefaultRoute returns the UseDefaultRoute field value if set, zero value otherwise.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner) GetUseDefaultRoute() bool {
	if o == nil || IsNil(o.UseDefaultRoute) {
		var ret bool
		return ret
	}
	return *o.UseDefaultRoute
}

// GetUseDefaultRouteOk returns a tuple with the UseDefaultRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner) GetUseDefaultRouteOk() (*bool, bool) {
	if o == nil || IsNil(o.UseDefaultRoute) {
		return nil, false
	}
	return o.UseDefaultRoute, true
}

// HasUseDefaultRoute returns a boolean if a field has been set.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner) HasUseDefaultRoute() bool {
	if o != nil && !IsNil(o.UseDefaultRoute) {
		return true
	}

	return false
}

// SetUseDefaultRoute gets a reference to the given bool and assigns it to the UseDefaultRoute field.
func (o *GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner) SetUseDefaultRoute(v bool) {
	o.UseDefaultRoute = &v
}

func (o GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HubId) {
		toSerialize["hubId"] = o.HubId
	}
	if !IsNil(o.UseDefaultRoute) {
		toSerialize["useDefaultRoute"] = o.UseDefaultRoute
	}
	return toSerialize, nil
}

type NullableGetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner struct {
	value *GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner
	isSet bool
}

func (v NullableGetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner) Get() *GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner {
	return v.value
}

func (v *NullableGetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner) Set(val *GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner(val *GetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner) *NullableGetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner {
	return &NullableGetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner{value: val, isSet: true}
}

func (v NullableGetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkApplianceVpnSiteToSiteVpn200ResponseHubsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


