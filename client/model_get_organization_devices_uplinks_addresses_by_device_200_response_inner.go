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

// checks if the GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner{}

// GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner struct for GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner
type GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner struct {
	// The device MAC address.
	Mac *string `json:"mac,omitempty"`
	// The device name.
	Name *string `json:"name,omitempty"`
	Network *GetOrganizationDevicesAvailabilities200ResponseInnerNetwork `json:"network,omitempty"`
	// Device product type.
	ProductType *string `json:"productType,omitempty"`
	// The device serial number.
	Serial *string `json:"serial,omitempty"`
	// List of custom tags for the device.
	Tags []string `json:"tags,omitempty"`
	// List of device uplink addresses information.
	Uplinks []GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner `json:"uplinks,omitempty"`
}

// NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInner instantiates a new GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInner() *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner {
	this := GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner{}
	return &this
}

// NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerWithDefaults instantiates a new GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerWithDefaults() *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner {
	this := GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner{}
	return &this
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) SetMac(v string) {
	o.Mac = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetNetwork() GetOrganizationDevicesAvailabilities200ResponseInnerNetwork {
	if o == nil || IsNil(o.Network) {
		var ret GetOrganizationDevicesAvailabilities200ResponseInnerNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetNetworkOk() (*GetOrganizationDevicesAvailabilities200ResponseInnerNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given GetOrganizationDevicesAvailabilities200ResponseInnerNetwork and assigns it to the Network field.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) SetNetwork(v GetOrganizationDevicesAvailabilities200ResponseInnerNetwork) {
	o.Network = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetProductType() string {
	if o == nil || IsNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetProductTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductType) {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) HasProductType() bool {
	if o != nil && !IsNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) SetProductType(v string) {
	o.ProductType = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) SetSerial(v string) {
	o.Serial = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) SetTags(v []string) {
	o.Tags = v
}

// GetUplinks returns the Uplinks field value if set, zero value otherwise.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetUplinks() []GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner {
	if o == nil || IsNil(o.Uplinks) {
		var ret []GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner
		return ret
	}
	return o.Uplinks
}

// GetUplinksOk returns a tuple with the Uplinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) GetUplinksOk() ([]GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner, bool) {
	if o == nil || IsNil(o.Uplinks) {
		return nil, false
	}
	return o.Uplinks, true
}

// HasUplinks returns a boolean if a field has been set.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) HasUplinks() bool {
	if o != nil && !IsNil(o.Uplinks) {
		return true
	}

	return false
}

// SetUplinks gets a reference to the given []GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner and assigns it to the Uplinks field.
func (o *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) SetUplinks(v []GetOrganizationDevicesUplinksAddressesByDevice200ResponseInnerUplinksInner) {
	o.Uplinks = v
}

func (o GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Uplinks) {
		toSerialize["uplinks"] = o.Uplinks
	}
	return toSerialize, nil
}

type NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInner struct {
	value *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) Get() *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) Set(val *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInner(val *GetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) *NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInner {
	return &NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationDevicesUplinksAddressesByDevice200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


