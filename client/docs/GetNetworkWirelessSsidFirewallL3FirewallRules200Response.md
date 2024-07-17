# GetNetworkWirelessSsidFirewallL3FirewallRules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]GetNetworkWirelessSsidFirewallL3FirewallRules200ResponseRulesInner**](GetNetworkWirelessSsidFirewallL3FirewallRules200ResponseRulesInner.md) | An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule). | [optional] 

## Methods

### NewGetNetworkWirelessSsidFirewallL3FirewallRules200Response

`func NewGetNetworkWirelessSsidFirewallL3FirewallRules200Response() *GetNetworkWirelessSsidFirewallL3FirewallRules200Response`

NewGetNetworkWirelessSsidFirewallL3FirewallRules200Response instantiates a new GetNetworkWirelessSsidFirewallL3FirewallRules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWirelessSsidFirewallL3FirewallRules200ResponseWithDefaults

`func NewGetNetworkWirelessSsidFirewallL3FirewallRules200ResponseWithDefaults() *GetNetworkWirelessSsidFirewallL3FirewallRules200Response`

NewGetNetworkWirelessSsidFirewallL3FirewallRules200ResponseWithDefaults instantiates a new GetNetworkWirelessSsidFirewallL3FirewallRules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *GetNetworkWirelessSsidFirewallL3FirewallRules200Response) GetRules() []GetNetworkWirelessSsidFirewallL3FirewallRules200ResponseRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GetNetworkWirelessSsidFirewallL3FirewallRules200Response) GetRulesOk() (*[]GetNetworkWirelessSsidFirewallL3FirewallRules200ResponseRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GetNetworkWirelessSsidFirewallL3FirewallRules200Response) SetRules(v []GetNetworkWirelessSsidFirewallL3FirewallRules200ResponseRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *GetNetworkWirelessSsidFirewallL3FirewallRules200Response) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


