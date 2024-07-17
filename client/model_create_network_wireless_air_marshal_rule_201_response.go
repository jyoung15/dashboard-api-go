/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the CreateNetworkWirelessAirMarshalRule201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkWirelessAirMarshalRule201Response{}

// CreateNetworkWirelessAirMarshalRule201Response struct for CreateNetworkWirelessAirMarshalRule201Response
type CreateNetworkWirelessAirMarshalRule201Response struct {
	Network *CreateNetworkWirelessAirMarshalRule201ResponseNetwork `json:"network,omitempty"`
	// Indicates whether or not clients are allowed to       connect to rogue SSIDs by default. (blocked by default)
	RuleId *string `json:"ruleId,omitempty"`
	// Indicates whether or not clients are allowed to       connect to rogue SSIDs by default. (blocked by default)
	Type *string `json:"type,omitempty"`
	// Updated at timestamp
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// Created at timestamp
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Match *CreateNetworkWirelessAirMarshalRule201ResponseMatch `json:"match,omitempty"`
}

// NewCreateNetworkWirelessAirMarshalRule201Response instantiates a new CreateNetworkWirelessAirMarshalRule201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkWirelessAirMarshalRule201Response() *CreateNetworkWirelessAirMarshalRule201Response {
	this := CreateNetworkWirelessAirMarshalRule201Response{}
	return &this
}

// NewCreateNetworkWirelessAirMarshalRule201ResponseWithDefaults instantiates a new CreateNetworkWirelessAirMarshalRule201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkWirelessAirMarshalRule201ResponseWithDefaults() *CreateNetworkWirelessAirMarshalRule201Response {
	this := CreateNetworkWirelessAirMarshalRule201Response{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *CreateNetworkWirelessAirMarshalRule201Response) GetNetwork() CreateNetworkWirelessAirMarshalRule201ResponseNetwork {
	if o == nil || IsNil(o.Network) {
		var ret CreateNetworkWirelessAirMarshalRule201ResponseNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessAirMarshalRule201Response) GetNetworkOk() (*CreateNetworkWirelessAirMarshalRule201ResponseNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *CreateNetworkWirelessAirMarshalRule201Response) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given CreateNetworkWirelessAirMarshalRule201ResponseNetwork and assigns it to the Network field.
func (o *CreateNetworkWirelessAirMarshalRule201Response) SetNetwork(v CreateNetworkWirelessAirMarshalRule201ResponseNetwork) {
	o.Network = &v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *CreateNetworkWirelessAirMarshalRule201Response) GetRuleId() string {
	if o == nil || IsNil(o.RuleId) {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessAirMarshalRule201Response) GetRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RuleId) {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *CreateNetworkWirelessAirMarshalRule201Response) HasRuleId() bool {
	if o != nil && !IsNil(o.RuleId) {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *CreateNetworkWirelessAirMarshalRule201Response) SetRuleId(v string) {
	o.RuleId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateNetworkWirelessAirMarshalRule201Response) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessAirMarshalRule201Response) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateNetworkWirelessAirMarshalRule201Response) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateNetworkWirelessAirMarshalRule201Response) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CreateNetworkWirelessAirMarshalRule201Response) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessAirMarshalRule201Response) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CreateNetworkWirelessAirMarshalRule201Response) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CreateNetworkWirelessAirMarshalRule201Response) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CreateNetworkWirelessAirMarshalRule201Response) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessAirMarshalRule201Response) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CreateNetworkWirelessAirMarshalRule201Response) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CreateNetworkWirelessAirMarshalRule201Response) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *CreateNetworkWirelessAirMarshalRule201Response) GetMatch() CreateNetworkWirelessAirMarshalRule201ResponseMatch {
	if o == nil || IsNil(o.Match) {
		var ret CreateNetworkWirelessAirMarshalRule201ResponseMatch
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessAirMarshalRule201Response) GetMatchOk() (*CreateNetworkWirelessAirMarshalRule201ResponseMatch, bool) {
	if o == nil || IsNil(o.Match) {
		return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *CreateNetworkWirelessAirMarshalRule201Response) HasMatch() bool {
	if o != nil && !IsNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given CreateNetworkWirelessAirMarshalRule201ResponseMatch and assigns it to the Match field.
func (o *CreateNetworkWirelessAirMarshalRule201Response) SetMatch(v CreateNetworkWirelessAirMarshalRule201ResponseMatch) {
	o.Match = &v
}

func (o CreateNetworkWirelessAirMarshalRule201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkWirelessAirMarshalRule201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.RuleId) {
		toSerialize["ruleId"] = o.RuleId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Match) {
		toSerialize["match"] = o.Match
	}
	return toSerialize, nil
}

type NullableCreateNetworkWirelessAirMarshalRule201Response struct {
	value *CreateNetworkWirelessAirMarshalRule201Response
	isSet bool
}

func (v NullableCreateNetworkWirelessAirMarshalRule201Response) Get() *CreateNetworkWirelessAirMarshalRule201Response {
	return v.value
}

func (v *NullableCreateNetworkWirelessAirMarshalRule201Response) Set(val *CreateNetworkWirelessAirMarshalRule201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkWirelessAirMarshalRule201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkWirelessAirMarshalRule201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkWirelessAirMarshalRule201Response(val *CreateNetworkWirelessAirMarshalRule201Response) *NullableCreateNetworkWirelessAirMarshalRule201Response {
	return &NullableCreateNetworkWirelessAirMarshalRule201Response{value: val, isSet: true}
}

func (v NullableCreateNetworkWirelessAirMarshalRule201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkWirelessAirMarshalRule201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


