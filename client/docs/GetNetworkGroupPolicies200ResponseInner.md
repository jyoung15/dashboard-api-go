# GetNetworkGroupPolicies200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupPolicyId** | Pointer to **string** | The ID of the group policy | [optional] 
**Scheduling** | Pointer to [**GetNetworkGroupPolicies200ResponseInnerScheduling**](GetNetworkGroupPolicies200ResponseInnerScheduling.md) |  | [optional] 
**Bandwidth** | Pointer to [**GetNetworkGroupPolicies200ResponseInnerBandwidth**](GetNetworkGroupPolicies200ResponseInnerBandwidth.md) |  | [optional] 
**FirewallAndTrafficShaping** | Pointer to [**GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping**](GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping.md) |  | [optional] 
**ContentFiltering** | Pointer to [**GetNetworkGroupPolicies200ResponseInnerContentFiltering**](GetNetworkGroupPolicies200ResponseInnerContentFiltering.md) |  | [optional] 
**SplashAuthSettings** | Pointer to **string** | Whether clients bound to your policy will bypass splash authorization or behave according to the network&#39;s rules. Can be one of &#39;network default&#39; or &#39;bypass&#39;. Only available if your network has a wireless configuration. | [optional] 
**VlanTagging** | Pointer to [**GetNetworkGroupPolicies200ResponseInnerVlanTagging**](GetNetworkGroupPolicies200ResponseInnerVlanTagging.md) |  | [optional] 
**BonjourForwarding** | Pointer to [**GetNetworkGroupPolicies200ResponseInnerBonjourForwarding**](GetNetworkGroupPolicies200ResponseInnerBonjourForwarding.md) |  | [optional] 

## Methods

### NewGetNetworkGroupPolicies200ResponseInner

`func NewGetNetworkGroupPolicies200ResponseInner() *GetNetworkGroupPolicies200ResponseInner`

NewGetNetworkGroupPolicies200ResponseInner instantiates a new GetNetworkGroupPolicies200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkGroupPolicies200ResponseInnerWithDefaults

`func NewGetNetworkGroupPolicies200ResponseInnerWithDefaults() *GetNetworkGroupPolicies200ResponseInner`

NewGetNetworkGroupPolicies200ResponseInnerWithDefaults instantiates a new GetNetworkGroupPolicies200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupPolicyId

`func (o *GetNetworkGroupPolicies200ResponseInner) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *GetNetworkGroupPolicies200ResponseInner) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *GetNetworkGroupPolicies200ResponseInner) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *GetNetworkGroupPolicies200ResponseInner) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.

### GetScheduling

`func (o *GetNetworkGroupPolicies200ResponseInner) GetScheduling() GetNetworkGroupPolicies200ResponseInnerScheduling`

GetScheduling returns the Scheduling field if non-nil, zero value otherwise.

### GetSchedulingOk

`func (o *GetNetworkGroupPolicies200ResponseInner) GetSchedulingOk() (*GetNetworkGroupPolicies200ResponseInnerScheduling, bool)`

GetSchedulingOk returns a tuple with the Scheduling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduling

`func (o *GetNetworkGroupPolicies200ResponseInner) SetScheduling(v GetNetworkGroupPolicies200ResponseInnerScheduling)`

SetScheduling sets Scheduling field to given value.

### HasScheduling

`func (o *GetNetworkGroupPolicies200ResponseInner) HasScheduling() bool`

HasScheduling returns a boolean if a field has been set.

### GetBandwidth

`func (o *GetNetworkGroupPolicies200ResponseInner) GetBandwidth() GetNetworkGroupPolicies200ResponseInnerBandwidth`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *GetNetworkGroupPolicies200ResponseInner) GetBandwidthOk() (*GetNetworkGroupPolicies200ResponseInnerBandwidth, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *GetNetworkGroupPolicies200ResponseInner) SetBandwidth(v GetNetworkGroupPolicies200ResponseInnerBandwidth)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *GetNetworkGroupPolicies200ResponseInner) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetFirewallAndTrafficShaping

`func (o *GetNetworkGroupPolicies200ResponseInner) GetFirewallAndTrafficShaping() GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping`

GetFirewallAndTrafficShaping returns the FirewallAndTrafficShaping field if non-nil, zero value otherwise.

### GetFirewallAndTrafficShapingOk

`func (o *GetNetworkGroupPolicies200ResponseInner) GetFirewallAndTrafficShapingOk() (*GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping, bool)`

GetFirewallAndTrafficShapingOk returns a tuple with the FirewallAndTrafficShaping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallAndTrafficShaping

`func (o *GetNetworkGroupPolicies200ResponseInner) SetFirewallAndTrafficShaping(v GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping)`

SetFirewallAndTrafficShaping sets FirewallAndTrafficShaping field to given value.

### HasFirewallAndTrafficShaping

`func (o *GetNetworkGroupPolicies200ResponseInner) HasFirewallAndTrafficShaping() bool`

HasFirewallAndTrafficShaping returns a boolean if a field has been set.

### GetContentFiltering

`func (o *GetNetworkGroupPolicies200ResponseInner) GetContentFiltering() GetNetworkGroupPolicies200ResponseInnerContentFiltering`

GetContentFiltering returns the ContentFiltering field if non-nil, zero value otherwise.

### GetContentFilteringOk

`func (o *GetNetworkGroupPolicies200ResponseInner) GetContentFilteringOk() (*GetNetworkGroupPolicies200ResponseInnerContentFiltering, bool)`

GetContentFilteringOk returns a tuple with the ContentFiltering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentFiltering

`func (o *GetNetworkGroupPolicies200ResponseInner) SetContentFiltering(v GetNetworkGroupPolicies200ResponseInnerContentFiltering)`

SetContentFiltering sets ContentFiltering field to given value.

### HasContentFiltering

`func (o *GetNetworkGroupPolicies200ResponseInner) HasContentFiltering() bool`

HasContentFiltering returns a boolean if a field has been set.

### GetSplashAuthSettings

`func (o *GetNetworkGroupPolicies200ResponseInner) GetSplashAuthSettings() string`

GetSplashAuthSettings returns the SplashAuthSettings field if non-nil, zero value otherwise.

### GetSplashAuthSettingsOk

`func (o *GetNetworkGroupPolicies200ResponseInner) GetSplashAuthSettingsOk() (*string, bool)`

GetSplashAuthSettingsOk returns a tuple with the SplashAuthSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashAuthSettings

`func (o *GetNetworkGroupPolicies200ResponseInner) SetSplashAuthSettings(v string)`

SetSplashAuthSettings sets SplashAuthSettings field to given value.

### HasSplashAuthSettings

`func (o *GetNetworkGroupPolicies200ResponseInner) HasSplashAuthSettings() bool`

HasSplashAuthSettings returns a boolean if a field has been set.

### GetVlanTagging

`func (o *GetNetworkGroupPolicies200ResponseInner) GetVlanTagging() GetNetworkGroupPolicies200ResponseInnerVlanTagging`

GetVlanTagging returns the VlanTagging field if non-nil, zero value otherwise.

### GetVlanTaggingOk

`func (o *GetNetworkGroupPolicies200ResponseInner) GetVlanTaggingOk() (*GetNetworkGroupPolicies200ResponseInnerVlanTagging, bool)`

GetVlanTaggingOk returns a tuple with the VlanTagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTagging

`func (o *GetNetworkGroupPolicies200ResponseInner) SetVlanTagging(v GetNetworkGroupPolicies200ResponseInnerVlanTagging)`

SetVlanTagging sets VlanTagging field to given value.

### HasVlanTagging

`func (o *GetNetworkGroupPolicies200ResponseInner) HasVlanTagging() bool`

HasVlanTagging returns a boolean if a field has been set.

### GetBonjourForwarding

`func (o *GetNetworkGroupPolicies200ResponseInner) GetBonjourForwarding() GetNetworkGroupPolicies200ResponseInnerBonjourForwarding`

GetBonjourForwarding returns the BonjourForwarding field if non-nil, zero value otherwise.

### GetBonjourForwardingOk

`func (o *GetNetworkGroupPolicies200ResponseInner) GetBonjourForwardingOk() (*GetNetworkGroupPolicies200ResponseInnerBonjourForwarding, bool)`

GetBonjourForwardingOk returns a tuple with the BonjourForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBonjourForwarding

`func (o *GetNetworkGroupPolicies200ResponseInner) SetBonjourForwarding(v GetNetworkGroupPolicies200ResponseInnerBonjourForwarding)`

SetBonjourForwarding sets BonjourForwarding field to given value.

### HasBonjourForwarding

`func (o *GetNetworkGroupPolicies200ResponseInner) HasBonjourForwarding() bool`

HasBonjourForwarding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


