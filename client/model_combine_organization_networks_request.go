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

// checks if the CombineOrganizationNetworksRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CombineOrganizationNetworksRequest{}

// CombineOrganizationNetworksRequest struct for CombineOrganizationNetworksRequest
type CombineOrganizationNetworksRequest struct {
	// The name of the combined network
	Name string `json:"name"`
	// A list of the network IDs that will be combined. If an ID of a combined network is included in this list, the other networks in the list will be grouped into that network
	NetworkIds []string `json:"networkIds"`
	// A unique identifier which can be used for device enrollment or easy access through the Meraki SM Registration page or the Self Service Portal. Please note that changing this field may cause existing bookmarks to break. All networks that are part of this combined network will have their enrollment string appended by '-network_type'. If left empty, all exisitng enrollment strings will be deleted.
	EnrollmentString *string `json:"enrollmentString,omitempty"`
}

type _CombineOrganizationNetworksRequest CombineOrganizationNetworksRequest

// NewCombineOrganizationNetworksRequest instantiates a new CombineOrganizationNetworksRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCombineOrganizationNetworksRequest(name string, networkIds []string) *CombineOrganizationNetworksRequest {
	this := CombineOrganizationNetworksRequest{}
	this.Name = name
	this.NetworkIds = networkIds
	return &this
}

// NewCombineOrganizationNetworksRequestWithDefaults instantiates a new CombineOrganizationNetworksRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCombineOrganizationNetworksRequestWithDefaults() *CombineOrganizationNetworksRequest {
	this := CombineOrganizationNetworksRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CombineOrganizationNetworksRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CombineOrganizationNetworksRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CombineOrganizationNetworksRequest) SetName(v string) {
	o.Name = v
}

// GetNetworkIds returns the NetworkIds field value
func (o *CombineOrganizationNetworksRequest) GetNetworkIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.NetworkIds
}

// GetNetworkIdsOk returns a tuple with the NetworkIds field value
// and a boolean to check if the value has been set.
func (o *CombineOrganizationNetworksRequest) GetNetworkIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkIds, true
}

// SetNetworkIds sets field value
func (o *CombineOrganizationNetworksRequest) SetNetworkIds(v []string) {
	o.NetworkIds = v
}

// GetEnrollmentString returns the EnrollmentString field value if set, zero value otherwise.
func (o *CombineOrganizationNetworksRequest) GetEnrollmentString() string {
	if o == nil || IsNil(o.EnrollmentString) {
		var ret string
		return ret
	}
	return *o.EnrollmentString
}

// GetEnrollmentStringOk returns a tuple with the EnrollmentString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CombineOrganizationNetworksRequest) GetEnrollmentStringOk() (*string, bool) {
	if o == nil || IsNil(o.EnrollmentString) {
		return nil, false
	}
	return o.EnrollmentString, true
}

// HasEnrollmentString returns a boolean if a field has been set.
func (o *CombineOrganizationNetworksRequest) HasEnrollmentString() bool {
	if o != nil && !IsNil(o.EnrollmentString) {
		return true
	}

	return false
}

// SetEnrollmentString gets a reference to the given string and assigns it to the EnrollmentString field.
func (o *CombineOrganizationNetworksRequest) SetEnrollmentString(v string) {
	o.EnrollmentString = &v
}

func (o CombineOrganizationNetworksRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CombineOrganizationNetworksRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["networkIds"] = o.NetworkIds
	if !IsNil(o.EnrollmentString) {
		toSerialize["enrollmentString"] = o.EnrollmentString
	}
	return toSerialize, nil
}

func (o *CombineOrganizationNetworksRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"networkIds",
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

	varCombineOrganizationNetworksRequest := _CombineOrganizationNetworksRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCombineOrganizationNetworksRequest)

	if err != nil {
		return err
	}

	*o = CombineOrganizationNetworksRequest(varCombineOrganizationNetworksRequest)

	return err
}

type NullableCombineOrganizationNetworksRequest struct {
	value *CombineOrganizationNetworksRequest
	isSet bool
}

func (v NullableCombineOrganizationNetworksRequest) Get() *CombineOrganizationNetworksRequest {
	return v.value
}

func (v *NullableCombineOrganizationNetworksRequest) Set(val *CombineOrganizationNetworksRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCombineOrganizationNetworksRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCombineOrganizationNetworksRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCombineOrganizationNetworksRequest(val *CombineOrganizationNetworksRequest) *NullableCombineOrganizationNetworksRequest {
	return &NullableCombineOrganizationNetworksRequest{value: val, isSet: true}
}

func (v NullableCombineOrganizationNetworksRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCombineOrganizationNetworksRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


