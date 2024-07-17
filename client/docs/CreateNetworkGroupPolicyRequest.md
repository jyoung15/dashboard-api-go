# CreateNetworkGroupPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name for your group policy. Required. | 
**Scheduling** | Pointer to [**GetNetworkGroupPolicies200ResponseInnerScheduling**](GetNetworkGroupPolicies200ResponseInnerScheduling.md) |  | [optional] 
**Bandwidth** | Pointer to [**GetNetworkGroupPolicies200ResponseInnerBandwidth**](GetNetworkGroupPolicies200ResponseInnerBandwidth.md) |  | [optional] 
**FirewallAndTrafficShaping** | Pointer to [**GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping**](GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping.md) |  | [optional] 
**ContentFiltering** | Pointer to [**GetNetworkGroupPolicies200ResponseInnerContentFiltering**](GetNetworkGroupPolicies200ResponseInnerContentFiltering.md) |  | [optional] 
**SplashAuthSettings** | Pointer to **string** | Whether clients bound to your policy will bypass splash authorization or behave according to the network&#39;s rules. Can be one of &#39;network default&#39; or &#39;bypass&#39;. Only available if your network has a wireless configuration. | [optional] 
**VlanTagging** | Pointer to [**GetNetworkGroupPolicies200ResponseInnerVlanTagging**](GetNetworkGroupPolicies200ResponseInnerVlanTagging.md) |  | [optional] 
**BonjourForwarding** | Pointer to [**GetNetworkGroupPolicies200ResponseInnerBonjourForwarding**](GetNetworkGroupPolicies200ResponseInnerBonjourForwarding.md) |  | [optional] 

## Methods

### NewCreateNetworkGroupPolicyRequest

`func NewCreateNetworkGroupPolicyRequest(name string, ) *CreateNetworkGroupPolicyRequest`

NewCreateNetworkGroupPolicyRequest instantiates a new CreateNetworkGroupPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkGroupPolicyRequestWithDefaults

`func NewCreateNetworkGroupPolicyRequestWithDefaults() *CreateNetworkGroupPolicyRequest`

NewCreateNetworkGroupPolicyRequestWithDefaults instantiates a new CreateNetworkGroupPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNetworkGroupPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNetworkGroupPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNetworkGroupPolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetScheduling

`func (o *CreateNetworkGroupPolicyRequest) GetScheduling() GetNetworkGroupPolicies200ResponseInnerScheduling`

GetScheduling returns the Scheduling field if non-nil, zero value otherwise.

### GetSchedulingOk

`func (o *CreateNetworkGroupPolicyRequest) GetSchedulingOk() (*GetNetworkGroupPolicies200ResponseInnerScheduling, bool)`

GetSchedulingOk returns a tuple with the Scheduling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduling

`func (o *CreateNetworkGroupPolicyRequest) SetScheduling(v GetNetworkGroupPolicies200ResponseInnerScheduling)`

SetScheduling sets Scheduling field to given value.

### HasScheduling

`func (o *CreateNetworkGroupPolicyRequest) HasScheduling() bool`

HasScheduling returns a boolean if a field has been set.

### GetBandwidth

`func (o *CreateNetworkGroupPolicyRequest) GetBandwidth() GetNetworkGroupPolicies200ResponseInnerBandwidth`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *CreateNetworkGroupPolicyRequest) GetBandwidthOk() (*GetNetworkGroupPolicies200ResponseInnerBandwidth, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *CreateNetworkGroupPolicyRequest) SetBandwidth(v GetNetworkGroupPolicies200ResponseInnerBandwidth)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *CreateNetworkGroupPolicyRequest) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetFirewallAndTrafficShaping

`func (o *CreateNetworkGroupPolicyRequest) GetFirewallAndTrafficShaping() GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping`

GetFirewallAndTrafficShaping returns the FirewallAndTrafficShaping field if non-nil, zero value otherwise.

### GetFirewallAndTrafficShapingOk

`func (o *CreateNetworkGroupPolicyRequest) GetFirewallAndTrafficShapingOk() (*GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping, bool)`

GetFirewallAndTrafficShapingOk returns a tuple with the FirewallAndTrafficShaping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallAndTrafficShaping

`func (o *CreateNetworkGroupPolicyRequest) SetFirewallAndTrafficShaping(v GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping)`

SetFirewallAndTrafficShaping sets FirewallAndTrafficShaping field to given value.

### HasFirewallAndTrafficShaping

`func (o *CreateNetworkGroupPolicyRequest) HasFirewallAndTrafficShaping() bool`

HasFirewallAndTrafficShaping returns a boolean if a field has been set.

### GetContentFiltering

`func (o *CreateNetworkGroupPolicyRequest) GetContentFiltering() GetNetworkGroupPolicies200ResponseInnerContentFiltering`

GetContentFiltering returns the ContentFiltering field if non-nil, zero value otherwise.

### GetContentFilteringOk

`func (o *CreateNetworkGroupPolicyRequest) GetContentFilteringOk() (*GetNetworkGroupPolicies200ResponseInnerContentFiltering, bool)`

GetContentFilteringOk returns a tuple with the ContentFiltering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFiltering

`func (o *CreateNetworkGroupPolicyRequest) SetContentFiltering(v GetNetworkGroupPolicies200ResponseInnerContentFiltering)`

SetContentFiltering sets ContentFiltering field to given value.

### HasContentFiltering

`func (o *CreateNetworkGroupPolicyRequest) HasContentFiltering() bool`

HasContentFiltering returns a boolean if a field has been set.

### GetSplashAuthSettings

`func (o *CreateNetworkGroupPolicyRequest) GetSplashAuthSettings() string`

GetSplashAuthSettings returns the SplashAuthSettings field if non-nil, zero value otherwise.

### GetSplashAuthSettingsOk

`func (o *CreateNetworkGroupPolicyRequest) GetSplashAuthSettingsOk() (*string, bool)`

GetSplashAuthSettingsOk returns a tuple with the SplashAuthSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashAuthSettings

`func (o *CreateNetworkGroupPolicyRequest) SetSplashAuthSettings(v string)`

SetSplashAuthSettings sets SplashAuthSettings field to given value.

### HasSplashAuthSettings

`func (o *CreateNetworkGroupPolicyRequest) HasSplashAuthSettings() bool`

HasSplashAuthSettings returns a boolean if a field has been set.

### GetVlanTagging

`func (o *CreateNetworkGroupPolicyRequest) GetVlanTagging() GetNetworkGroupPolicies200ResponseInnerVlanTagging`

GetVlanTagging returns the VlanTagging field if non-nil, zero value otherwise.

### GetVlanTaggingOk

`func (o *CreateNetworkGroupPolicyRequest) GetVlanTaggingOk() (*GetNetworkGroupPolicies200ResponseInnerVlanTagging, bool)`

GetVlanTaggingOk returns a tuple with the VlanTagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTagging

`func (o *CreateNetworkGroupPolicyRequest) SetVlanTagging(v GetNetworkGroupPolicies200ResponseInnerVlanTagging)`

SetVlanTagging sets VlanTagging field to given value.

### HasVlanTagging

`func (o *CreateNetworkGroupPolicyRequest) HasVlanTagging() bool`

HasVlanTagging returns a boolean if a field has been set.

### GetBonjourForwarding

`func (o *CreateNetworkGroupPolicyRequest) GetBonjourForwarding() GetNetworkGroupPolicies200ResponseInnerBonjourForwarding`

GetBonjourForwarding returns the BonjourForwarding field if non-nil, zero value otherwise.

### GetBonjourForwardingOk

`func (o *CreateNetworkGroupPolicyRequest) GetBonjourForwardingOk() (*GetNetworkGroupPolicies200ResponseInnerBonjourForwarding, bool)`

GetBonjourForwardingOk returns a tuple with the BonjourForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonjourForwarding

`func (o *CreateNetworkGroupPolicyRequest) SetBonjourForwarding(v GetNetworkGroupPolicies200ResponseInnerBonjourForwarding)`

SetBonjourForwarding sets BonjourForwarding field to given value.

### HasBonjourForwarding

`func (o *CreateNetworkGroupPolicyRequest) HasBonjourForwarding() bool`

HasBonjourForwarding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


