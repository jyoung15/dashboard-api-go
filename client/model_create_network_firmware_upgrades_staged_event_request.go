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

// checks if the CreateNetworkFirmwareUpgradesStagedEventRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkFirmwareUpgradesStagedEventRequest{}

// CreateNetworkFirmwareUpgradesStagedEventRequest struct for CreateNetworkFirmwareUpgradesStagedEventRequest
type CreateNetworkFirmwareUpgradesStagedEventRequest struct {
	Products *CreateNetworkFirmwareUpgradesStagedEventRequestProducts `json:"products,omitempty"`
	// All firmware upgrade stages in the network with their start time.
	Stages []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner `json:"stages"`
}

type _CreateNetworkFirmwareUpgradesStagedEventRequest CreateNetworkFirmwareUpgradesStagedEventRequest

// NewCreateNetworkFirmwareUpgradesStagedEventRequest instantiates a new CreateNetworkFirmwareUpgradesStagedEventRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkFirmwareUpgradesStagedEventRequest(stages []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) *CreateNetworkFirmwareUpgradesStagedEventRequest {
	this := CreateNetworkFirmwareUpgradesStagedEventRequest{}
	this.Stages = stages
	return &this
}

// NewCreateNetworkFirmwareUpgradesStagedEventRequestWithDefaults instantiates a new CreateNetworkFirmwareUpgradesStagedEventRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkFirmwareUpgradesStagedEventRequestWithDefaults() *CreateNetworkFirmwareUpgradesStagedEventRequest {
	this := CreateNetworkFirmwareUpgradesStagedEventRequest{}
	return &this
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *CreateNetworkFirmwareUpgradesStagedEventRequest) GetProducts() CreateNetworkFirmwareUpgradesStagedEventRequestProducts {
	if o == nil || IsNil(o.Products) {
		var ret CreateNetworkFirmwareUpgradesStagedEventRequestProducts
		return ret
	}
	return *o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesStagedEventRequest) GetProductsOk() (*CreateNetworkFirmwareUpgradesStagedEventRequestProducts, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *CreateNetworkFirmwareUpgradesStagedEventRequest) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given CreateNetworkFirmwareUpgradesStagedEventRequestProducts and assigns it to the Products field.
func (o *CreateNetworkFirmwareUpgradesStagedEventRequest) SetProducts(v CreateNetworkFirmwareUpgradesStagedEventRequestProducts) {
	o.Products = &v
}

// GetStages returns the Stages field value
func (o *CreateNetworkFirmwareUpgradesStagedEventRequest) GetStages() []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner {
	if o == nil {
		var ret []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner
		return ret
	}

	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkFirmwareUpgradesStagedEventRequest) GetStagesOk() ([]UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stages, true
}

// SetStages sets field value
func (o *CreateNetworkFirmwareUpgradesStagedEventRequest) SetStages(v []UpdateNetworkFirmwareUpgradesStagedEventsRequestStagesInner) {
	o.Stages = v
}

func (o CreateNetworkFirmwareUpgradesStagedEventRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkFirmwareUpgradesStagedEventRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	toSerialize["stages"] = o.Stages
	return toSerialize, nil
}

func (o *CreateNetworkFirmwareUpgradesStagedEventRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"stages",
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

	varCreateNetworkFirmwareUpgradesStagedEventRequest := _CreateNetworkFirmwareUpgradesStagedEventRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateNetworkFirmwareUpgradesStagedEventRequest)

	if err != nil {
		return err
	}

	*o = CreateNetworkFirmwareUpgradesStagedEventRequest(varCreateNetworkFirmwareUpgradesStagedEventRequest)

	return err
}

type NullableCreateNetworkFirmwareUpgradesStagedEventRequest struct {
	value *CreateNetworkFirmwareUpgradesStagedEventRequest
	isSet bool
}

func (v NullableCreateNetworkFirmwareUpgradesStagedEventRequest) Get() *CreateNetworkFirmwareUpgradesStagedEventRequest {
	return v.value
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedEventRequest) Set(val *CreateNetworkFirmwareUpgradesStagedEventRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkFirmwareUpgradesStagedEventRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedEventRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkFirmwareUpgradesStagedEventRequest(val *CreateNetworkFirmwareUpgradesStagedEventRequest) *NullableCreateNetworkFirmwareUpgradesStagedEventRequest {
	return &NullableCreateNetworkFirmwareUpgradesStagedEventRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkFirmwareUpgradesStagedEventRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkFirmwareUpgradesStagedEventRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


