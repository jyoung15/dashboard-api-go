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

// checks if the CreateNetworkApplianceVlanRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkApplianceVlanRequest{}

// CreateNetworkApplianceVlanRequest struct for CreateNetworkApplianceVlanRequest
type CreateNetworkApplianceVlanRequest struct {
	// The VLAN ID of the new VLAN (must be between 1 and 4094)
	Id string `json:"id"`
	// The name of the new VLAN
	Name string `json:"name"`
	// The subnet of the VLAN
	Subnet *string `json:"subnet,omitempty"`
	// The local IP of the appliance on the VLAN
	ApplianceIp *string `json:"applianceIp,omitempty"`
	// The id of the desired group policy to apply to the VLAN
	GroupPolicyId *string `json:"groupPolicyId,omitempty"`
	// Type of subnetting of the VLAN. Applicable only for template network.
	TemplateVlanType *string `json:"templateVlanType,omitempty"`
	// CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN.
	Cidr *string `json:"cidr,omitempty"`
	// Mask used for the subnet of all bound to the template networks. Applicable only for template network.
	Mask *int32 `json:"mask,omitempty"`
	Ipv6 *UpdateNetworkApplianceSingleLanRequestIpv6 `json:"ipv6,omitempty"`
	// The appliance's handling of DHCP requests on this VLAN. One of: 'Run a DHCP server', 'Relay DHCP to another server' or 'Do not respond to DHCP requests'
	DhcpHandling *string `json:"dhcpHandling,omitempty"`
	// The term of DHCP leases if the appliance is running a DHCP server on this VLAN. One of: '30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week'
	DhcpLeaseTime *string `json:"dhcpLeaseTime,omitempty"`
	MandatoryDhcp *GetNetworkApplianceVlans200ResponseInnerMandatoryDhcp `json:"mandatoryDhcp,omitempty"`
	// Use DHCP boot options specified in other properties
	DhcpBootOptionsEnabled *bool `json:"dhcpBootOptionsEnabled,omitempty"`
	// The list of DHCP options that will be included in DHCP responses. Each object in the list should have \"code\", \"type\", and \"value\" properties.
	DhcpOptions []GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner `json:"dhcpOptions,omitempty"`
}

type _CreateNetworkApplianceVlanRequest CreateNetworkApplianceVlanRequest

// NewCreateNetworkApplianceVlanRequest instantiates a new CreateNetworkApplianceVlanRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkApplianceVlanRequest(id string, name string) *CreateNetworkApplianceVlanRequest {
	this := CreateNetworkApplianceVlanRequest{}
	this.Id = id
	this.Name = name
	var templateVlanType string = "same"
	this.TemplateVlanType = &templateVlanType
	return &this
}

// NewCreateNetworkApplianceVlanRequestWithDefaults instantiates a new CreateNetworkApplianceVlanRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkApplianceVlanRequestWithDefaults() *CreateNetworkApplianceVlanRequest {
	this := CreateNetworkApplianceVlanRequest{}
	var templateVlanType string = "same"
	this.TemplateVlanType = &templateVlanType
	return &this
}

// GetId returns the Id field value
func (o *CreateNetworkApplianceVlanRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlanRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CreateNetworkApplianceVlanRequest) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CreateNetworkApplianceVlanRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlanRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNetworkApplianceVlanRequest) SetName(v string) {
	o.Name = v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlanRequest) GetSubnet() string {
	if o == nil || IsNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlanRequest) GetSubnetOk() (*string, bool) {
	if o == nil || IsNil(o.Subnet) {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlanRequest) HasSubnet() bool {
	if o != nil && !IsNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *CreateNetworkApplianceVlanRequest) SetSubnet(v string) {
	o.Subnet = &v
}

// GetApplianceIp returns the ApplianceIp field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlanRequest) GetApplianceIp() string {
	if o == nil || IsNil(o.ApplianceIp) {
		var ret string
		return ret
	}
	return *o.ApplianceIp
}

// GetApplianceIpOk returns a tuple with the ApplianceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlanRequest) GetApplianceIpOk() (*string, bool) {
	if o == nil || IsNil(o.ApplianceIp) {
		return nil, false
	}
	return o.ApplianceIp, true
}

// HasApplianceIp returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlanRequest) HasApplianceIp() bool {
	if o != nil && !IsNil(o.ApplianceIp) {
		return true
	}

	return false
}

// SetApplianceIp gets a reference to the given string and assigns it to the ApplianceIp field.
func (o *CreateNetworkApplianceVlanRequest) SetApplianceIp(v string) {
	o.ApplianceIp = &v
}

// GetGroupPolicyId returns the GroupPolicyId field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlanRequest) GetGroupPolicyId() string {
	if o == nil || IsNil(o.GroupPolicyId) {
		var ret string
		return ret
	}
	return *o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlanRequest) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupPolicyId) {
		return nil, false
	}
	return o.GroupPolicyId, true
}

// HasGroupPolicyId returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlanRequest) HasGroupPolicyId() bool {
	if o != nil && !IsNil(o.GroupPolicyId) {
		return true
	}

	return false
}

// SetGroupPolicyId gets a reference to the given string and assigns it to the GroupPolicyId field.
func (o *CreateNetworkApplianceVlanRequest) SetGroupPolicyId(v string) {
	o.GroupPolicyId = &v
}

// GetTemplateVlanType returns the TemplateVlanType field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlanRequest) GetTemplateVlanType() string {
	if o == nil || IsNil(o.TemplateVlanType) {
		var ret string
		return ret
	}
	return *o.TemplateVlanType
}

// GetTemplateVlanTypeOk returns a tuple with the TemplateVlanType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlanRequest) GetTemplateVlanTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateVlanType) {
		return nil, false
	}
	return o.TemplateVlanType, true
}

// HasTemplateVlanType returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlanRequest) HasTemplateVlanType() bool {
	if o != nil && !IsNil(o.TemplateVlanType) {
		return true
	}

	return false
}

// SetTemplateVlanType gets a reference to the given string and assigns it to the TemplateVlanType field.
func (o *CreateNetworkApplianceVlanRequest) SetTemplateVlanType(v string) {
	o.TemplateVlanType = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlanRequest) GetCidr() string {
	if o == nil || IsNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlanRequest) GetCidrOk() (*string, bool) {
	if o == nil || IsNil(o.Cidr) {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlanRequest) HasCidr() bool {
	if o != nil && !IsNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *CreateNetworkApplianceVlanRequest) SetCidr(v string) {
	o.Cidr = &v
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlanRequest) GetMask() int32 {
	if o == nil || IsNil(o.Mask) {
		var ret int32
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlanRequest) GetMaskOk() (*int32, bool) {
	if o == nil || IsNil(o.Mask) {
		return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlanRequest) HasMask() bool {
	if o != nil && !IsNil(o.Mask) {
		return true
	}

	return false
}

// SetMask gets a reference to the given int32 and assigns it to the Mask field.
func (o *CreateNetworkApplianceVlanRequest) SetMask(v int32) {
	o.Mask = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlanRequest) GetIpv6() UpdateNetworkApplianceSingleLanRequestIpv6 {
	if o == nil || IsNil(o.Ipv6) {
		var ret UpdateNetworkApplianceSingleLanRequestIpv6
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlanRequest) GetIpv6Ok() (*UpdateNetworkApplianceSingleLanRequestIpv6, bool) {
	if o == nil || IsNil(o.Ipv6) {
		return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlanRequest) HasIpv6() bool {
	if o != nil && !IsNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given UpdateNetworkApplianceSingleLanRequestIpv6 and assigns it to the Ipv6 field.
func (o *CreateNetworkApplianceVlanRequest) SetIpv6(v UpdateNetworkApplianceSingleLanRequestIpv6) {
	o.Ipv6 = &v
}

// GetDhcpHandling returns the DhcpHandling field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlanRequest) GetDhcpHandling() string {
	if o == nil || IsNil(o.DhcpHandling) {
		var ret string
		return ret
	}
	return *o.DhcpHandling
}

// GetDhcpHandlingOk returns a tuple with the DhcpHandling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlanRequest) GetDhcpHandlingOk() (*string, bool) {
	if o == nil || IsNil(o.DhcpHandling) {
		return nil, false
	}
	return o.DhcpHandling, true
}

// HasDhcpHandling returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlanRequest) HasDhcpHandling() bool {
	if o != nil && !IsNil(o.DhcpHandling) {
		return true
	}

	return false
}

// SetDhcpHandling gets a reference to the given string and assigns it to the DhcpHandling field.
func (o *CreateNetworkApplianceVlanRequest) SetDhcpHandling(v string) {
	o.DhcpHandling = &v
}

// GetDhcpLeaseTime returns the DhcpLeaseTime field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlanRequest) GetDhcpLeaseTime() string {
	if o == nil || IsNil(o.DhcpLeaseTime) {
		var ret string
		return ret
	}
	return *o.DhcpLeaseTime
}

// GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlanRequest) GetDhcpLeaseTimeOk() (*string, bool) {
	if o == nil || IsNil(o.DhcpLeaseTime) {
		return nil, false
	}
	return o.DhcpLeaseTime, true
}

// HasDhcpLeaseTime returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlanRequest) HasDhcpLeaseTime() bool {
	if o != nil && !IsNil(o.DhcpLeaseTime) {
		return true
	}

	return false
}

// SetDhcpLeaseTime gets a reference to the given string and assigns it to the DhcpLeaseTime field.
func (o *CreateNetworkApplianceVlanRequest) SetDhcpLeaseTime(v string) {
	o.DhcpLeaseTime = &v
}

// GetMandatoryDhcp returns the MandatoryDhcp field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlanRequest) GetMandatoryDhcp() GetNetworkApplianceVlans200ResponseInnerMandatoryDhcp {
	if o == nil || IsNil(o.MandatoryDhcp) {
		var ret GetNetworkApplianceVlans200ResponseInnerMandatoryDhcp
		return ret
	}
	return *o.MandatoryDhcp
}

// GetMandatoryDhcpOk returns a tuple with the MandatoryDhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlanRequest) GetMandatoryDhcpOk() (*GetNetworkApplianceVlans200ResponseInnerMandatoryDhcp, bool) {
	if o == nil || IsNil(o.MandatoryDhcp) {
		return nil, false
	}
	return o.MandatoryDhcp, true
}

// HasMandatoryDhcp returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlanRequest) HasMandatoryDhcp() bool {
	if o != nil && !IsNil(o.MandatoryDhcp) {
		return true
	}

	return false
}

// SetMandatoryDhcp gets a reference to the given GetNetworkApplianceVlans200ResponseInnerMandatoryDhcp and assigns it to the MandatoryDhcp field.
func (o *CreateNetworkApplianceVlanRequest) SetMandatoryDhcp(v GetNetworkApplianceVlans200ResponseInnerMandatoryDhcp) {
	o.MandatoryDhcp = &v
}

// GetDhcpBootOptionsEnabled returns the DhcpBootOptionsEnabled field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlanRequest) GetDhcpBootOptionsEnabled() bool {
	if o == nil || IsNil(o.DhcpBootOptionsEnabled) {
		var ret bool
		return ret
	}
	return *o.DhcpBootOptionsEnabled
}

// GetDhcpBootOptionsEnabledOk returns a tuple with the DhcpBootOptionsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlanRequest) GetDhcpBootOptionsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DhcpBootOptionsEnabled) {
		return nil, false
	}
	return o.DhcpBootOptionsEnabled, true
}

// HasDhcpBootOptionsEnabled returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlanRequest) HasDhcpBootOptionsEnabled() bool {
	if o != nil && !IsNil(o.DhcpBootOptionsEnabled) {
		return true
	}

	return false
}

// SetDhcpBootOptionsEnabled gets a reference to the given bool and assigns it to the DhcpBootOptionsEnabled field.
func (o *CreateNetworkApplianceVlanRequest) SetDhcpBootOptionsEnabled(v bool) {
	o.DhcpBootOptionsEnabled = &v
}

// GetDhcpOptions returns the DhcpOptions field value if set, zero value otherwise.
func (o *CreateNetworkApplianceVlanRequest) GetDhcpOptions() []GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner {
	if o == nil || IsNil(o.DhcpOptions) {
		var ret []GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner
		return ret
	}
	return o.DhcpOptions
}

// GetDhcpOptionsOk returns a tuple with the DhcpOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceVlanRequest) GetDhcpOptionsOk() ([]GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner, bool) {
	if o == nil || IsNil(o.DhcpOptions) {
		return nil, false
	}
	return o.DhcpOptions, true
}

// HasDhcpOptions returns a boolean if a field has been set.
func (o *CreateNetworkApplianceVlanRequest) HasDhcpOptions() bool {
	if o != nil && !IsNil(o.DhcpOptions) {
		return true
	}

	return false
}

// SetDhcpOptions gets a reference to the given []GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner and assigns it to the DhcpOptions field.
func (o *CreateNetworkApplianceVlanRequest) SetDhcpOptions(v []GetNetworkApplianceVlans200ResponseInnerDhcpOptionsInner) {
	o.DhcpOptions = v
}

func (o CreateNetworkApplianceVlanRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkApplianceVlanRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !IsNil(o.ApplianceIp) {
		toSerialize["applianceIp"] = o.ApplianceIp
	}
	if !IsNil(o.GroupPolicyId) {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	if !IsNil(o.TemplateVlanType) {
		toSerialize["templateVlanType"] = o.TemplateVlanType
	}
	if !IsNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !IsNil(o.Mask) {
		toSerialize["mask"] = o.Mask
	}
	if !IsNil(o.Ipv6) {
		toSerialize["ipv6"] = o.Ipv6
	}
	if !IsNil(o.DhcpHandling) {
		toSerialize["dhcpHandling"] = o.DhcpHandling
	}
	if !IsNil(o.DhcpLeaseTime) {
		toSerialize["dhcpLeaseTime"] = o.DhcpLeaseTime
	}
	if !IsNil(o.MandatoryDhcp) {
		toSerialize["mandatoryDhcp"] = o.MandatoryDhcp
	}
	if !IsNil(o.DhcpBootOptionsEnabled) {
		toSerialize["dhcpBootOptionsEnabled"] = o.DhcpBootOptionsEnabled
	}
	if !IsNil(o.DhcpOptions) {
		toSerialize["dhcpOptions"] = o.DhcpOptions
	}
	return toSerialize, nil
}

func (o *CreateNetworkApplianceVlanRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varCreateNetworkApplianceVlanRequest := _CreateNetworkApplianceVlanRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateNetworkApplianceVlanRequest)

	if err != nil {
		return err
	}

	*o = CreateNetworkApplianceVlanRequest(varCreateNetworkApplianceVlanRequest)

	return err
}

type NullableCreateNetworkApplianceVlanRequest struct {
	value *CreateNetworkApplianceVlanRequest
	isSet bool
}

func (v NullableCreateNetworkApplianceVlanRequest) Get() *CreateNetworkApplianceVlanRequest {
	return v.value
}

func (v *NullableCreateNetworkApplianceVlanRequest) Set(val *CreateNetworkApplianceVlanRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkApplianceVlanRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkApplianceVlanRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkApplianceVlanRequest(val *CreateNetworkApplianceVlanRequest) *NullableCreateNetworkApplianceVlanRequest {
	return &NullableCreateNetworkApplianceVlanRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkApplianceVlanRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkApplianceVlanRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


