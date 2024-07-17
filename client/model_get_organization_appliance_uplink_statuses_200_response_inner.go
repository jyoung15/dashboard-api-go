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

// checks if the GetOrganizationApplianceUplinkStatuses200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationApplianceUplinkStatuses200ResponseInner{}

// GetOrganizationApplianceUplinkStatuses200ResponseInner struct for GetOrganizationApplianceUplinkStatuses200ResponseInner
type GetOrganizationApplianceUplinkStatuses200ResponseInner struct {
	// Network identifier
	NetworkId *string `json:"networkId,omitempty"`
	// The uplink serial
	Serial *string `json:"serial,omitempty"`
	// The uplink model
	Model *string `json:"model,omitempty"`
	// Last reported time for the device
	LastReportedAt *time.Time `json:"lastReportedAt,omitempty"`
	HighAvailability *GetOrganizationApplianceUplinkStatuses200ResponseInnerHighAvailability `json:"highAvailability,omitempty"`
	// Uplinks
	Uplinks []GetOrganizationApplianceUplinkStatuses200ResponseInnerUplinksInner `json:"uplinks,omitempty"`
}

// NewGetOrganizationApplianceUplinkStatuses200ResponseInner instantiates a new GetOrganizationApplianceUplinkStatuses200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationApplianceUplinkStatuses200ResponseInner() *GetOrganizationApplianceUplinkStatuses200ResponseInner {
	this := GetOrganizationApplianceUplinkStatuses200ResponseInner{}
	return &this
}

// NewGetOrganizationApplianceUplinkStatuses200ResponseInnerWithDefaults instantiates a new GetOrganizationApplianceUplinkStatuses200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationApplianceUplinkStatuses200ResponseInnerWithDefaults() *GetOrganizationApplianceUplinkStatuses200ResponseInner {
	this := GetOrganizationApplianceUplinkStatuses200ResponseInner{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetNetworkId() string {
	if o == nil || IsNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) SetSerial(v string) {
	o.Serial = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) SetModel(v string) {
	o.Model = &v
}

// GetLastReportedAt returns the LastReportedAt field value if set, zero value otherwise.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetLastReportedAt() time.Time {
	if o == nil || IsNil(o.LastReportedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastReportedAt
}

// GetLastReportedAtOk returns a tuple with the LastReportedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetLastReportedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastReportedAt) {
		return nil, false
	}
	return o.LastReportedAt, true
}

// HasLastReportedAt returns a boolean if a field has been set.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) HasLastReportedAt() bool {
	if o != nil && !IsNil(o.LastReportedAt) {
		return true
	}

	return false
}

// SetLastReportedAt gets a reference to the given time.Time and assigns it to the LastReportedAt field.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) SetLastReportedAt(v time.Time) {
	o.LastReportedAt = &v
}

// GetHighAvailability returns the HighAvailability field value if set, zero value otherwise.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetHighAvailability() GetOrganizationApplianceUplinkStatuses200ResponseInnerHighAvailability {
	if o == nil || IsNil(o.HighAvailability) {
		var ret GetOrganizationApplianceUplinkStatuses200ResponseInnerHighAvailability
		return ret
	}
	return *o.HighAvailability
}

// GetHighAvailabilityOk returns a tuple with the HighAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetHighAvailabilityOk() (*GetOrganizationApplianceUplinkStatuses200ResponseInnerHighAvailability, bool) {
	if o == nil || IsNil(o.HighAvailability) {
		return nil, false
	}
	return o.HighAvailability, true
}

// HasHighAvailability returns a boolean if a field has been set.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) HasHighAvailability() bool {
	if o != nil && !IsNil(o.HighAvailability) {
		return true
	}

	return false
}

// SetHighAvailability gets a reference to the given GetOrganizationApplianceUplinkStatuses200ResponseInnerHighAvailability and assigns it to the HighAvailability field.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) SetHighAvailability(v GetOrganizationApplianceUplinkStatuses200ResponseInnerHighAvailability) {
	o.HighAvailability = &v
}

// GetUplinks returns the Uplinks field value if set, zero value otherwise.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetUplinks() []GetOrganizationApplianceUplinkStatuses200ResponseInnerUplinksInner {
	if o == nil || IsNil(o.Uplinks) {
		var ret []GetOrganizationApplianceUplinkStatuses200ResponseInnerUplinksInner
		return ret
	}
	return o.Uplinks
}

// GetUplinksOk returns a tuple with the Uplinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) GetUplinksOk() ([]GetOrganizationApplianceUplinkStatuses200ResponseInnerUplinksInner, bool) {
	if o == nil || IsNil(o.Uplinks) {
		return nil, false
	}
	return o.Uplinks, true
}

// HasUplinks returns a boolean if a field has been set.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) HasUplinks() bool {
	if o != nil && !IsNil(o.Uplinks) {
		return true
	}

	return false
}

// SetUplinks gets a reference to the given []GetOrganizationApplianceUplinkStatuses200ResponseInnerUplinksInner and assigns it to the Uplinks field.
func (o *GetOrganizationApplianceUplinkStatuses200ResponseInner) SetUplinks(v []GetOrganizationApplianceUplinkStatuses200ResponseInnerUplinksInner) {
	o.Uplinks = v
}

func (o GetOrganizationApplianceUplinkStatuses200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationApplianceUplinkStatuses200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.LastReportedAt) {
		toSerialize["lastReportedAt"] = o.LastReportedAt
	}
	if !IsNil(o.HighAvailability) {
		toSerialize["highAvailability"] = o.HighAvailability
	}
	if !IsNil(o.Uplinks) {
		toSerialize["uplinks"] = o.Uplinks
	}
	return toSerialize, nil
}

type NullableGetOrganizationApplianceUplinkStatuses200ResponseInner struct {
	value *GetOrganizationApplianceUplinkStatuses200ResponseInner
	isSet bool
}

func (v NullableGetOrganizationApplianceUplinkStatuses200ResponseInner) Get() *GetOrganizationApplianceUplinkStatuses200ResponseInner {
	return v.value
}

func (v *NullableGetOrganizationApplianceUplinkStatuses200ResponseInner) Set(val *GetOrganizationApplianceUplinkStatuses200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationApplianceUplinkStatuses200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationApplianceUplinkStatuses200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationApplianceUplinkStatuses200ResponseInner(val *GetOrganizationApplianceUplinkStatuses200ResponseInner) *NullableGetOrganizationApplianceUplinkStatuses200ResponseInner {
	return &NullableGetOrganizationApplianceUplinkStatuses200ResponseInner{value: val, isSet: true}
}

func (v NullableGetOrganizationApplianceUplinkStatuses200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationApplianceUplinkStatuses200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

