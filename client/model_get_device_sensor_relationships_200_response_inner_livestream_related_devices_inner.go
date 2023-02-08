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

// GetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner struct for GetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner
type GetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner struct {
	// The serial of the related device
	Serial *string `json:"serial,omitempty"`
	// The product type of the related device
	ProductType *string `json:"productType,omitempty"`
}

// NewGetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner instantiates a new GetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner() *GetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner {
	this := GetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner{}
	return &this
}

// NewGetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInnerWithDefaults instantiates a new GetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInnerWithDefaults() *GetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner {
	this := GetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner) SetSerial(v string) {
	o.Serial = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *GetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner) GetProductType() string {
	if o == nil || isNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner) GetProductTypeOk() (*string, bool) {
	if o == nil || isNil(o.ProductType) {
    return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *GetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner) HasProductType() bool {
	if o != nil && !isNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *GetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner) SetProductType(v string) {
	o.ProductType = &v
}

func (o GetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	return json.Marshal(toSerialize)
}

type NullableGetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner struct {
	value *GetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner
	isSet bool
}

func (v NullableGetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner) Get() *GetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner {
	return v.value
}

func (v *NullableGetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner) Set(val *GetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner(val *GetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner) *NullableGetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner {
	return &NullableGetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner{value: val, isSet: true}
}

func (v NullableGetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceSensorRelationships200ResponseInnerLivestreamRelatedDevicesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


