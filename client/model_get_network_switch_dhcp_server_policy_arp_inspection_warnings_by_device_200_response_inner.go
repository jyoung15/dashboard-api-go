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

// checks if the GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner{}

// GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner struct for GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner
type GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner struct {
	// Switch serial.
	Serial *string `json:"serial,omitempty"`
	// Switch name.
	Name *string `json:"name,omitempty"`
	// Url link to switch.
	Url *string `json:"url,omitempty"`
	// Whether this switch supports Dynamic ARP Inspection.
	SupportsInspection *bool `json:"supportsInspection,omitempty"`
	// Whether this switch has a trusted DAI port. Always false if supportsInspection is false.
	HasTrustedPort *bool `json:"hasTrustedPort,omitempty"`
}

// NewGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner instantiates a new GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner() *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner {
	this := GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner{}
	return &this
}

// NewGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInnerWithDefaults instantiates a new GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInnerWithDefaults() *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner {
	this := GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) SetSerial(v string) {
	o.Serial = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) SetUrl(v string) {
	o.Url = &v
}

// GetSupportsInspection returns the SupportsInspection field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) GetSupportsInspection() bool {
	if o == nil || IsNil(o.SupportsInspection) {
		var ret bool
		return ret
	}
	return *o.SupportsInspection
}

// GetSupportsInspectionOk returns a tuple with the SupportsInspection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) GetSupportsInspectionOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsInspection) {
		return nil, false
	}
	return o.SupportsInspection, true
}

// HasSupportsInspection returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) HasSupportsInspection() bool {
	if o != nil && !IsNil(o.SupportsInspection) {
		return true
	}

	return false
}

// SetSupportsInspection gets a reference to the given bool and assigns it to the SupportsInspection field.
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) SetSupportsInspection(v bool) {
	o.SupportsInspection = &v
}

// GetHasTrustedPort returns the HasTrustedPort field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) GetHasTrustedPort() bool {
	if o == nil || IsNil(o.HasTrustedPort) {
		var ret bool
		return ret
	}
	return *o.HasTrustedPort
}

// GetHasTrustedPortOk returns a tuple with the HasTrustedPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) GetHasTrustedPortOk() (*bool, bool) {
	if o == nil || IsNil(o.HasTrustedPort) {
		return nil, false
	}
	return o.HasTrustedPort, true
}

// HasHasTrustedPort returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) HasHasTrustedPort() bool {
	if o != nil && !IsNil(o.HasTrustedPort) {
		return true
	}

	return false
}

// SetHasTrustedPort gets a reference to the given bool and assigns it to the HasTrustedPort field.
func (o *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) SetHasTrustedPort(v bool) {
	o.HasTrustedPort = &v
}

func (o GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.SupportsInspection) {
		toSerialize["supportsInspection"] = o.SupportsInspection
	}
	if !IsNil(o.HasTrustedPort) {
		toSerialize["hasTrustedPort"] = o.HasTrustedPort
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner struct {
	value *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner
	isSet bool
}

func (v NullableGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) Get() *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner {
	return v.value
}

func (v *NullableGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) Set(val *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner(val *GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) *NullableGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner {
	return &NullableGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


