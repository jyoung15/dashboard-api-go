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

// checks if the GetOrganizationSummaryTopClientsByUsage200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSummaryTopClientsByUsage200ResponseInner{}

// GetOrganizationSummaryTopClientsByUsage200ResponseInner struct for GetOrganizationSummaryTopClientsByUsage200ResponseInner
type GetOrganizationSummaryTopClientsByUsage200ResponseInner struct {
	// Name of client
	Name *string `json:"name,omitempty"`
	// MAC address of client
	Mac *string `json:"mac,omitempty"`
	// ID of client
	Id *string `json:"id,omitempty"`
	Network *GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork `json:"network,omitempty"`
	Usage *GetOrganizationSummaryTopClientsByUsage200ResponseInnerUsage `json:"usage,omitempty"`
}

// NewGetOrganizationSummaryTopClientsByUsage200ResponseInner instantiates a new GetOrganizationSummaryTopClientsByUsage200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSummaryTopClientsByUsage200ResponseInner() *GetOrganizationSummaryTopClientsByUsage200ResponseInner {
	this := GetOrganizationSummaryTopClientsByUsage200ResponseInner{}
	return &this
}

// NewGetOrganizationSummaryTopClientsByUsage200ResponseInnerWithDefaults instantiates a new GetOrganizationSummaryTopClientsByUsage200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSummaryTopClientsByUsage200ResponseInnerWithDefaults() *GetOrganizationSummaryTopClientsByUsage200ResponseInner {
	this := GetOrganizationSummaryTopClientsByUsage200ResponseInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInner) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInner) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInner) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInner) SetMac(v string) {
	o.Mac = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInner) GetNetwork() GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork {
	if o == nil || IsNil(o.Network) {
		var ret GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInner) GetNetworkOk() (*GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInner) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork and assigns it to the Network field.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInner) SetNetwork(v GetOrganizationSummaryTopClientsByUsage200ResponseInnerNetwork) {
	o.Network = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInner) GetUsage() GetOrganizationSummaryTopClientsByUsage200ResponseInnerUsage {
	if o == nil || IsNil(o.Usage) {
		var ret GetOrganizationSummaryTopClientsByUsage200ResponseInnerUsage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInner) GetUsageOk() (*GetOrganizationSummaryTopClientsByUsage200ResponseInnerUsage, bool) {
	if o == nil || IsNil(o.Usage) {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInner) HasUsage() bool {
	if o != nil && !IsNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given GetOrganizationSummaryTopClientsByUsage200ResponseInnerUsage and assigns it to the Usage field.
func (o *GetOrganizationSummaryTopClientsByUsage200ResponseInner) SetUsage(v GetOrganizationSummaryTopClientsByUsage200ResponseInnerUsage) {
	o.Usage = &v
}

func (o GetOrganizationSummaryTopClientsByUsage200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSummaryTopClientsByUsage200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	return toSerialize, nil
}

type NullableGetOrganizationSummaryTopClientsByUsage200ResponseInner struct {
	value *GetOrganizationSummaryTopClientsByUsage200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationSummaryTopClientsByUsage200ResponseInner) Get() *GetOrganizationSummaryTopClientsByUsage200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationSummaryTopClientsByUsage200ResponseInner) Set(val *GetOrganizationSummaryTopClientsByUsage200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSummaryTopClientsByUsage200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSummaryTopClientsByUsage200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSummaryTopClientsByUsage200ResponseInner(val *GetOrganizationSummaryTopClientsByUsage200ResponseInner) *NullableGetOrganizationSummaryTopClientsByUsage200ResponseInner {
	return &NullableGetOrganizationSummaryTopClientsByUsage200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationSummaryTopClientsByUsage200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSummaryTopClientsByUsage200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


