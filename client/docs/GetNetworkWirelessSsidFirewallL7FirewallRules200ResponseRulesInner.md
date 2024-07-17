# GetNetworkWirelessSsidFirewallL7FirewallRules200ResponseRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | Pointer to **string** | &#39;Deny&#39; traffic specified by this rule | [optional] 
**Type** | Pointer to **string** | Type of the L7 firewall rule. One of: &#39;application&#39;, &#39;applicationCategory&#39;, &#39;host&#39;, &#39;port&#39;, &#39;ipRange&#39; | [optional] 
**Value** | Pointer to **string** | The value of what needs to get blocked. Format of the value varies depending on type of the firewall rule selected. | [optional] 

## Methods

### NewGetNetworkWirelessSsidFirewallL7FirewallRules200ResponseRulesInner

`func NewGetNetworkWirelessSsidFirewallL7FirewallRules200ResponseRulesInner() *GetNetworkWirelessSsidFirewallL7FirewallRules200ResponseRulesInner`

NewGetNetworkWirelessSsidFirewallL7FirewallRules200ResponseRulesInner instantiates a new GetNetworkWirelessSsidFirewallL7FirewallRules200ResponseRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWirelessSsidFirewallL7FirewallRules200ResponseRulesInnerWithDefaults

`func NewGetNetworkWirelessSsidFirewallL7FirewallRules200ResponseRulesInnerWithDefaults() *GetNetworkWirelessSsidFirewallL7FirewallRules200ResponseRulesInner`

NewGetNetworkWirelessSsidFirewallL7FirewallRules200ResponseRulesInnerWithDefaults instantiates a new GetNetworkWirelessSsidFirewallL7FirewallRules200ResponseRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *GetNetworkWirelessSsidFirewallL7FirewallRules200ResponseRulesInner) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GetNetworkWirelessSsidFirewallL7FirewallRules200ResponseRulesInner) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GetNetworkWirelessSsidFirewallL7FirewallRules200ResponseRulesInner) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *GetNetworkWirelessSsidFirewallL7FirewallRules200ResponseRulesInner) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetType

`func (o *GetNetworkWirelessSsidFirewallL7FirewallRules200ResponseRulesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNetworkWirelessSsidFirewallL7FirewallRules200ResponseRulesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNetworkWirelessSsidFirewallL7FirewallRules200ResponseRulesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetNetworkWirelessSsidFirewallL7FirewallRules200ResponseRulesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *GetNetworkWirelessSsidFirewallL7FirewallRules200ResponseRulesInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetNetworkWirelessSsidFirewallL7FirewallRules200ResponseRulesInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetNetworkWirelessSsidFirewallL7FirewallRules200ResponseRulesInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetNetworkWirelessSsidFirewallL7FirewallRules200ResponseRulesInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


