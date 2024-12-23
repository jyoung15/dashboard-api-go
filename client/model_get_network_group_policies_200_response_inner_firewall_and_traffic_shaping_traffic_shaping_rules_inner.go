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

// checks if the GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner{}

// GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner struct for GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner
type GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner struct {
	//     A list of objects describing the definitions of your traffic shaping rule. At least one definition is required. 
	Definitions []UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner `json:"definitions"`
	PerClientBandwidthLimits *UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits `json:"perClientBandwidthLimits,omitempty"`
	//     The DSCP tag applied by your rule. null means 'Do not change DSCP tag'.     For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint. 
	DscpTagValue *int32 `json:"dscpTagValue,omitempty"`
	//     The PCP tag applied by your rule. Can be 0 (lowest priority) through 7 (highest priority).     null means 'Do not set PCP tag'. 
	PcpTagValue *int32 `json:"pcpTagValue,omitempty"`
	//     A string, indicating the priority level for packets bound to your rule.     Can be 'low', 'normal' or 'high'. 
	Priority *string `json:"priority,omitempty"`
}

type _GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner

// NewGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner instantiates a new GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner(definitions []UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner {
	this := GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner{}
	this.Definitions = definitions
	return &this
}

// NewGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInnerWithDefaults instantiates a new GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInnerWithDefaults() *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner {
	this := GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner{}
	return &this
}

// GetDefinitions returns the Definitions field value
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) GetDefinitions() []UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner {
	if o == nil {
		var ret []UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner
		return ret
	}

	return o.Definitions
}

// GetDefinitionsOk returns a tuple with the Definitions field value
// and a boolean to check if the value has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) GetDefinitionsOk() ([]UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Definitions, true
}

// SetDefinitions sets field value
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) SetDefinitions(v []UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerDefinitionsInner) {
	o.Definitions = v
}

// GetPerClientBandwidthLimits returns the PerClientBandwidthLimits field value if set, zero value otherwise.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) GetPerClientBandwidthLimits() UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits {
	if o == nil || IsNil(o.PerClientBandwidthLimits) {
		var ret UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits
		return ret
	}
	return *o.PerClientBandwidthLimits
}

// GetPerClientBandwidthLimitsOk returns a tuple with the PerClientBandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) GetPerClientBandwidthLimitsOk() (*UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits, bool) {
	if o == nil || IsNil(o.PerClientBandwidthLimits) {
		return nil, false
	}
	return o.PerClientBandwidthLimits, true
}

// HasPerClientBandwidthLimits returns a boolean if a field has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) HasPerClientBandwidthLimits() bool {
	if o != nil && !IsNil(o.PerClientBandwidthLimits) {
		return true
	}

	return false
}

// SetPerClientBandwidthLimits gets a reference to the given UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits and assigns it to the PerClientBandwidthLimits field.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) SetPerClientBandwidthLimits(v UpdateNetworkApplianceTrafficShapingRulesRequestRulesInnerPerClientBandwidthLimits) {
	o.PerClientBandwidthLimits = &v
}

// GetDscpTagValue returns the DscpTagValue field value if set, zero value otherwise.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) GetDscpTagValue() int32 {
	if o == nil || IsNil(o.DscpTagValue) {
		var ret int32
		return ret
	}
	return *o.DscpTagValue
}

// GetDscpTagValueOk returns a tuple with the DscpTagValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) GetDscpTagValueOk() (*int32, bool) {
	if o == nil || IsNil(o.DscpTagValue) {
		return nil, false
	}
	return o.DscpTagValue, true
}

// HasDscpTagValue returns a boolean if a field has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) HasDscpTagValue() bool {
	if o != nil && !IsNil(o.DscpTagValue) {
		return true
	}

	return false
}

// SetDscpTagValue gets a reference to the given int32 and assigns it to the DscpTagValue field.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) SetDscpTagValue(v int32) {
	o.DscpTagValue = &v
}

// GetPcpTagValue returns the PcpTagValue field value if set, zero value otherwise.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) GetPcpTagValue() int32 {
	if o == nil || IsNil(o.PcpTagValue) {
		var ret int32
		return ret
	}
	return *o.PcpTagValue
}

// GetPcpTagValueOk returns a tuple with the PcpTagValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) GetPcpTagValueOk() (*int32, bool) {
	if o == nil || IsNil(o.PcpTagValue) {
		return nil, false
	}
	return o.PcpTagValue, true
}

// HasPcpTagValue returns a boolean if a field has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) HasPcpTagValue() bool {
	if o != nil && !IsNil(o.PcpTagValue) {
		return true
	}

	return false
}

// SetPcpTagValue gets a reference to the given int32 and assigns it to the PcpTagValue field.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) SetPcpTagValue(v int32) {
	o.PcpTagValue = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) SetPriority(v string) {
	o.Priority = &v
}

func (o GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["definitions"] = o.Definitions
	if !IsNil(o.PerClientBandwidthLimits) {
		toSerialize["perClientBandwidthLimits"] = o.PerClientBandwidthLimits
	}
	if !IsNil(o.DscpTagValue) {
		toSerialize["dscpTagValue"] = o.DscpTagValue
	}
	if !IsNil(o.PcpTagValue) {
		toSerialize["pcpTagValue"] = o.PcpTagValue
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	return toSerialize, nil
}

func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"definitions",
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

	varGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner := _GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner)

	if err != nil {
		return err
	}

	*o = GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner(varGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner)

	return err
}

type NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner struct {
	value *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner
	isSet bool
}

func (v NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) Get() *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner {
	return v.value
}

func (v *NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) Set(val *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner(val *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) *NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner {
	return &NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner{value: val, isSet: true}
}

func (v NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


