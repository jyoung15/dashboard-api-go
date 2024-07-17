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

// checks if the GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner{}

// GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner struct for GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner
type GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner struct {
	// URL of affected device
	Url *string `json:"url,omitempty"`
	// Name of affected device
	Name *string `json:"name,omitempty"`
	// Order of affected device in array
	Order *int32 `json:"order,omitempty"`
	// Type of affected device
	ProductType *string `json:"productType,omitempty"`
	// Serial of affected device
	Serial *string `json:"serial,omitempty"`
	// MAC address of affected device
	Mac *string `json:"mac,omitempty"`
	// IMEI of affected device
	Imei *string `json:"imei,omitempty"`
	// The device tags
	Tags []string `json:"tags,omitempty"`
	Lldp *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp `json:"lldp,omitempty"`
}

// NewGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner instantiates a new GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner() *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner {
	this := GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner{}
	return &this
}

// NewGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerWithDefaults instantiates a new GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerWithDefaults() *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner {
	this := GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) SetUrl(v string) {
	o.Url = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) SetName(v string) {
	o.Name = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) SetOrder(v int32) {
	o.Order = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) GetProductType() string {
	if o == nil || IsNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) GetProductTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProductType) {
		return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) HasProductType() bool {
	if o != nil && !IsNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) SetProductType(v string) {
	o.ProductType = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) SetSerial(v string) {
	o.Serial = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) SetMac(v string) {
	o.Mac = &v
}

// GetImei returns the Imei field value if set, zero value otherwise.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) GetImei() string {
	if o == nil || IsNil(o.Imei) {
		var ret string
		return ret
	}
	return *o.Imei
}

// GetImeiOk returns a tuple with the Imei field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) GetImeiOk() (*string, bool) {
	if o == nil || IsNil(o.Imei) {
		return nil, false
	}
	return o.Imei, true
}

// HasImei returns a boolean if a field has been set.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) HasImei() bool {
	if o != nil && !IsNil(o.Imei) {
		return true
	}

	return false
}

// SetImei gets a reference to the given string and assigns it to the Imei field.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) SetImei(v string) {
	o.Imei = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) SetTags(v []string) {
	o.Tags = v
}

// GetLldp returns the Lldp field value if set, zero value otherwise.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) GetLldp() GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp {
	if o == nil || IsNil(o.Lldp) {
		var ret GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp
		return ret
	}
	return *o.Lldp
}

// GetLldpOk returns a tuple with the Lldp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) GetLldpOk() (*GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp, bool) {
	if o == nil || IsNil(o.Lldp) {
		return nil, false
	}
	return o.Lldp, true
}

// HasLldp returns a boolean if a field has been set.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) HasLldp() bool {
	if o != nil && !IsNil(o.Lldp) {
		return true
	}

	return false
}

// SetLldp gets a reference to the given GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp and assigns it to the Lldp field.
func (o *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) SetLldp(v GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInnerLldp) {
	o.Lldp = &v
}

func (o GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.Imei) {
		toSerialize["imei"] = o.Imei
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Lldp) {
		toSerialize["lldp"] = o.Lldp
	}
	return toSerialize, nil
}

type NullableGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner struct {
	value *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner
	isSet bool
}

func (v NullableGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) Get() *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner {
	return v.value
}

func (v *NullableGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) Set(val *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner(val *GetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) *NullableGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner {
	return &NullableGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner{value: val, isSet: true}
}

func (v NullableGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationAssuranceAlerts200ResponseInnerScopeDevicesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

