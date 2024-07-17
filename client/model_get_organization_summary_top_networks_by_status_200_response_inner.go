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

// checks if the GetOrganizationSummaryTopNetworksByStatus200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationSummaryTopNetworksByStatus200ResponseInner{}

// GetOrganizationSummaryTopNetworksByStatus200ResponseInner struct for GetOrganizationSummaryTopNetworksByStatus200ResponseInner
type GetOrganizationSummaryTopNetworksByStatus200ResponseInner struct {
	// Network identifier
	NetworkId *string `json:"networkId,omitempty"`
	// Network name
	Name *string `json:"name,omitempty"`
	// Network clients list URL
	Url *string `json:"url,omitempty"`
	// Network tags
	Tags []string `json:"tags,omitempty"`
	Clients *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerClients `json:"clients,omitempty"`
	Statuses *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses `json:"statuses,omitempty"`
	Devices *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices `json:"devices,omitempty"`
	// Product types in network
	ProductTypes []string `json:"productTypes,omitempty"`
}

// NewGetOrganizationSummaryTopNetworksByStatus200ResponseInner instantiates a new GetOrganizationSummaryTopNetworksByStatus200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationSummaryTopNetworksByStatus200ResponseInner() *GetOrganizationSummaryTopNetworksByStatus200ResponseInner {
	this := GetOrganizationSummaryTopNetworksByStatus200ResponseInner{}
	return &this
}

// NewGetOrganizationSummaryTopNetworksByStatus200ResponseInnerWithDefaults instantiates a new GetOrganizationSummaryTopNetworksByStatus200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationSummaryTopNetworksByStatus200ResponseInnerWithDefaults() *GetOrganizationSummaryTopNetworksByStatus200ResponseInner {
	this := GetOrganizationSummaryTopNetworksByStatus200ResponseInner{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetNetworkId() string {
	if o == nil || IsNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) SetUrl(v string) {
	o.Url = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) SetTags(v []string) {
	o.Tags = v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetClients() GetOrganizationSummaryTopNetworksByStatus200ResponseInnerClients {
	if o == nil || IsNil(o.Clients) {
		var ret GetOrganizationSummaryTopNetworksByStatus200ResponseInnerClients
		return ret
	}
	return *o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetClientsOk() (*GetOrganizationSummaryTopNetworksByStatus200ResponseInnerClients, bool) {
	if o == nil || IsNil(o.Clients) {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) HasClients() bool {
	if o != nil && !IsNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given GetOrganizationSummaryTopNetworksByStatus200ResponseInnerClients and assigns it to the Clients field.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) SetClients(v GetOrganizationSummaryTopNetworksByStatus200ResponseInnerClients) {
	o.Clients = &v
}

// GetStatuses returns the Statuses field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetStatuses() GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses {
	if o == nil || IsNil(o.Statuses) {
		var ret GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses
		return ret
	}
	return *o.Statuses
}

// GetStatusesOk returns a tuple with the Statuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetStatusesOk() (*GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses, bool) {
	if o == nil || IsNil(o.Statuses) {
		return nil, false
	}
	return o.Statuses, true
}

// HasStatuses returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) HasStatuses() bool {
	if o != nil && !IsNil(o.Statuses) {
		return true
	}

	return false
}

// SetStatuses gets a reference to the given GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses and assigns it to the Statuses field.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) SetStatuses(v GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses) {
	o.Statuses = &v
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetDevices() GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices {
	if o == nil || IsNil(o.Devices) {
		var ret GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices
		return ret
	}
	return *o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetDevicesOk() (*GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices, bool) {
	if o == nil || IsNil(o.Devices) {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) HasDevices() bool {
	if o != nil && !IsNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices and assigns it to the Devices field.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) SetDevices(v GetOrganizationSummaryTopNetworksByStatus200ResponseInnerDevices) {
	o.Devices = &v
}

// GetProductTypes returns the ProductTypes field value if set, zero value otherwise.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetProductTypes() []string {
	if o == nil || IsNil(o.ProductTypes) {
		var ret []string
		return ret
	}
	return o.ProductTypes
}

// GetProductTypesOk returns a tuple with the ProductTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) GetProductTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ProductTypes) {
		return nil, false
	}
	return o.ProductTypes, true
}

// HasProductTypes returns a boolean if a field has been set.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) HasProductTypes() bool {
	if o != nil && !IsNil(o.ProductTypes) {
		return true
	}

	return false
}

// SetProductTypes gets a reference to the given []string and assigns it to the ProductTypes field.
func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) SetProductTypes(v []string) {
	o.ProductTypes = v
}

func (o GetOrganizationSummaryTopNetworksByStatus200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationSummaryTopNetworksByStatus200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Clients) {
		toSerialize["clients"] = o.Clients
	}
	if !IsNil(o.Statuses) {
		toSerialize["statuses"] = o.Statuses
	}
	if !IsNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if !IsNil(o.ProductTypes) {
		toSerialize["productTypes"] = o.ProductTypes
	}
	return toSerialize, nil
}

type NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInner struct {
	value *GetOrganizationSummaryTopNetworksByStatus200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInner) Get() *GetOrganizationSummaryTopNetworksByStatus200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInner) Set(val *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationSummaryTopNetworksByStatus200ResponseInner(val *GetOrganizationSummaryTopNetworksByStatus200ResponseInner) *NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInner {
	return &NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationSummaryTopNetworksByStatus200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

