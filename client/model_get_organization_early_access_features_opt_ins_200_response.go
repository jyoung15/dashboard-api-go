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

// checks if the GetOrganizationEarlyAccessFeaturesOptIns200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationEarlyAccessFeaturesOptIns200Response{}

// GetOrganizationEarlyAccessFeaturesOptIns200Response struct for GetOrganizationEarlyAccessFeaturesOptIns200Response
type GetOrganizationEarlyAccessFeaturesOptIns200Response struct {
	// ID of Early Access Feature
	Id *string `json:"id,omitempty"`
	// Name of Early Access Feature
	ShortName *string `json:"shortName,omitempty"`
	// Networks assigned to the Early Access Feature
	LimitScopeToNetworks []GetOrganizationEarlyAccessFeaturesOptIns200ResponseLimitScopeToNetworksInner `json:"limitScopeToNetworks,omitempty"`
	OptOutEligibility *GetOrganizationEarlyAccessFeaturesOptIns200ResponseOptOutEligibility `json:"optOutEligibility,omitempty"`
	// Time when Early Access Feature was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

// NewGetOrganizationEarlyAccessFeaturesOptIns200Response instantiates a new GetOrganizationEarlyAccessFeaturesOptIns200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationEarlyAccessFeaturesOptIns200Response() *GetOrganizationEarlyAccessFeaturesOptIns200Response {
	this := GetOrganizationEarlyAccessFeaturesOptIns200Response{}
	return &this
}

// NewGetOrganizationEarlyAccessFeaturesOptIns200ResponseWithDefaults instantiates a new GetOrganizationEarlyAccessFeaturesOptIns200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationEarlyAccessFeaturesOptIns200ResponseWithDefaults() *GetOrganizationEarlyAccessFeaturesOptIns200Response {
	this := GetOrganizationEarlyAccessFeaturesOptIns200Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) SetId(v string) {
	o.Id = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) GetShortName() string {
	if o == nil || IsNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) GetShortNameOk() (*string, bool) {
	if o == nil || IsNil(o.ShortName) {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) HasShortName() bool {
	if o != nil && !IsNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) SetShortName(v string) {
	o.ShortName = &v
}

// GetLimitScopeToNetworks returns the LimitScopeToNetworks field value if set, zero value otherwise.
func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) GetLimitScopeToNetworks() []GetOrganizationEarlyAccessFeaturesOptIns200ResponseLimitScopeToNetworksInner {
	if o == nil || IsNil(o.LimitScopeToNetworks) {
		var ret []GetOrganizationEarlyAccessFeaturesOptIns200ResponseLimitScopeToNetworksInner
		return ret
	}
	return o.LimitScopeToNetworks
}

// GetLimitScopeToNetworksOk returns a tuple with the LimitScopeToNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) GetLimitScopeToNetworksOk() ([]GetOrganizationEarlyAccessFeaturesOptIns200ResponseLimitScopeToNetworksInner, bool) {
	if o == nil || IsNil(o.LimitScopeToNetworks) {
		return nil, false
	}
	return o.LimitScopeToNetworks, true
}

// HasLimitScopeToNetworks returns a boolean if a field has been set.
func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) HasLimitScopeToNetworks() bool {
	if o != nil && !IsNil(o.LimitScopeToNetworks) {
		return true
	}

	return false
}

// SetLimitScopeToNetworks gets a reference to the given []GetOrganizationEarlyAccessFeaturesOptIns200ResponseLimitScopeToNetworksInner and assigns it to the LimitScopeToNetworks field.
func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) SetLimitScopeToNetworks(v []GetOrganizationEarlyAccessFeaturesOptIns200ResponseLimitScopeToNetworksInner) {
	o.LimitScopeToNetworks = v
}

// GetOptOutEligibility returns the OptOutEligibility field value if set, zero value otherwise.
func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) GetOptOutEligibility() GetOrganizationEarlyAccessFeaturesOptIns200ResponseOptOutEligibility {
	if o == nil || IsNil(o.OptOutEligibility) {
		var ret GetOrganizationEarlyAccessFeaturesOptIns200ResponseOptOutEligibility
		return ret
	}
	return *o.OptOutEligibility
}

// GetOptOutEligibilityOk returns a tuple with the OptOutEligibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) GetOptOutEligibilityOk() (*GetOrganizationEarlyAccessFeaturesOptIns200ResponseOptOutEligibility, bool) {
	if o == nil || IsNil(o.OptOutEligibility) {
		return nil, false
	}
	return o.OptOutEligibility, true
}

// HasOptOutEligibility returns a boolean if a field has been set.
func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) HasOptOutEligibility() bool {
	if o != nil && !IsNil(o.OptOutEligibility) {
		return true
	}

	return false
}

// SetOptOutEligibility gets a reference to the given GetOrganizationEarlyAccessFeaturesOptIns200ResponseOptOutEligibility and assigns it to the OptOutEligibility field.
func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) SetOptOutEligibility(v GetOrganizationEarlyAccessFeaturesOptIns200ResponseOptOutEligibility) {
	o.OptOutEligibility = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o GetOrganizationEarlyAccessFeaturesOptIns200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationEarlyAccessFeaturesOptIns200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ShortName) {
		toSerialize["shortName"] = o.ShortName
	}
	if !IsNil(o.LimitScopeToNetworks) {
		toSerialize["limitScopeToNetworks"] = o.LimitScopeToNetworks
	}
	if !IsNil(o.OptOutEligibility) {
		toSerialize["optOutEligibility"] = o.OptOutEligibility
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableGetOrganizationEarlyAccessFeaturesOptIns200Response struct {
	value *GetOrganizationEarlyAccessFeaturesOptIns200Response
	isSet bool
}

func (v NullableGetOrganizationEarlyAccessFeaturesOptIns200Response) Get() *GetOrganizationEarlyAccessFeaturesOptIns200Response {
	return v.value
}

func (v *NullableGetOrganizationEarlyAccessFeaturesOptIns200Response) Set(val *GetOrganizationEarlyAccessFeaturesOptIns200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationEarlyAccessFeaturesOptIns200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationEarlyAccessFeaturesOptIns200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationEarlyAccessFeaturesOptIns200Response(val *GetOrganizationEarlyAccessFeaturesOptIns200Response) *NullableGetOrganizationEarlyAccessFeaturesOptIns200Response {
	return &NullableGetOrganizationEarlyAccessFeaturesOptIns200Response{value: val, isSet: true}
}

func (v NullableGetOrganizationEarlyAccessFeaturesOptIns200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationEarlyAccessFeaturesOptIns200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


