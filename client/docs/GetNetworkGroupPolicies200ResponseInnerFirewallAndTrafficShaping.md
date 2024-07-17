# GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to **string** | How firewall and traffic shaping rules are enforced. Can be &#39;network default&#39;, &#39;ignore&#39; or &#39;custom&#39;. | [optional] 
**TrafficShapingRules** | Pointer to [**[]GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner**](GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner.md) |     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules.  | [optional] 
**L3FirewallRules** | Pointer to [**[]GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL3FirewallRulesInner**](GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL3FirewallRulesInner.md) | An ordered array of the L3 firewall rules | [optional] 
**L7FirewallRules** | Pointer to [**[]GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner**](GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner.md) | An ordered array of L7 firewall rules | [optional] 

## Methods

### NewGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping

`func NewGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping() *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping`

NewGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping instantiates a new GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingWithDefaults

`func NewGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingWithDefaults() *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping`

NewGetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingWithDefaults instantiates a new GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) GetSettings() string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) GetSettingsOk() (*string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) SetSettings(v string)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetTrafficShapingRules

`func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) GetTrafficShapingRules() []GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner`

GetTrafficShapingRules returns the TrafficShapingRules field if non-nil, zero value otherwise.

### GetTrafficShapingRulesOk

`func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) GetTrafficShapingRulesOk() (*[]GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner, bool)`

GetTrafficShapingRulesOk returns a tuple with the TrafficShapingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficShapingRules

`func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) SetTrafficShapingRules(v []GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingTrafficShapingRulesInner)`

SetTrafficShapingRules sets TrafficShapingRules field to given value.

### HasTrafficShapingRules

`func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) HasTrafficShapingRules() bool`

HasTrafficShapingRules returns a boolean if a field has been set.

### GetL3FirewallRules

`func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) GetL3FirewallRules() []GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL3FirewallRulesInner`

GetL3FirewallRules returns the L3FirewallRules field if non-nil, zero value otherwise.

### GetL3FirewallRulesOk

`func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) GetL3FirewallRulesOk() (*[]GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL3FirewallRulesInner, bool)`

GetL3FirewallRulesOk returns a tuple with the L3FirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3FirewallRules

`func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) SetL3FirewallRules(v []GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL3FirewallRulesInner)`

SetL3FirewallRules sets L3FirewallRules field to given value.

### HasL3FirewallRules

`func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) HasL3FirewallRules() bool`

HasL3FirewallRules returns a boolean if a field has been set.

### GetL7FirewallRules

`func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) GetL7FirewallRules() []GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner`

GetL7FirewallRules returns the L7FirewallRules field if non-nil, zero value otherwise.

### GetL7FirewallRulesOk

`func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) GetL7FirewallRulesOk() (*[]GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner, bool)`

GetL7FirewallRulesOk returns a tuple with the L7FirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7FirewallRules

`func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) SetL7FirewallRules(v []GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShapingL7FirewallRulesInner)`

SetL7FirewallRules sets L7FirewallRules field to given value.

### HasL7FirewallRules

`func (o *GetNetworkGroupPolicies200ResponseInnerFirewallAndTrafficShaping) HasL7FirewallRules() bool`

HasL7FirewallRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


