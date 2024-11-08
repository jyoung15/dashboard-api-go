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

// checks if the CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner{}

// CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner struct for CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner
type CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner struct {
	// Model of the AP
	Model *string `json:"model,omitempty"`
	// Band to use for each flex radio. For example, ['6'] will set the AP's first flex radio to 6 GHz
	Bands []string `json:"bands,omitempty"`
}

// NewCreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner instantiates a new CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner() *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner {
	this := CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner{}
	return &this
}

// NewCreateNetworkWirelessRfProfileRequestFlexRadiosByModelInnerWithDefaults instantiates a new CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkWirelessRfProfileRequestFlexRadiosByModelInnerWithDefaults() *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner {
	this := CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner{}
	return &this
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) SetModel(v string) {
	o.Model = &v
}

// GetBands returns the Bands field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) GetBands() []string {
	if o == nil || IsNil(o.Bands) {
		var ret []string
		return ret
	}
	return o.Bands
}

// GetBandsOk returns a tuple with the Bands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) GetBandsOk() ([]string, bool) {
	if o == nil || IsNil(o.Bands) {
		return nil, false
	}
	return o.Bands, true
}

// HasBands returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) HasBands() bool {
	if o != nil && !IsNil(o.Bands) {
		return true
	}

	return false
}

// SetBands gets a reference to the given []string and assigns it to the Bands field.
func (o *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) SetBands(v []string) {
	o.Bands = v
}

func (o CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.Bands) {
		toSerialize["bands"] = o.Bands
	}
	return toSerialize, nil
}

type NullableCreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner struct {
	value *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner
	isSet bool
}

func (v NullableCreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) Get() *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner {
	return v.value
}

func (v *NullableCreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) Set(val *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner(val *CreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) *NullableCreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner {
	return &NullableCreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner{value: val, isSet: true}
}

func (v NullableCreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkWirelessRfProfileRequestFlexRadiosByModelInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


