# GetNetworkApplianceFirewallFirewalledServices200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** | Appliance service name | [optional] 
**Access** | Pointer to **string** | A string indicating the rule for which IPs are allowed to use the specified service | [optional] 
**AllowedIps** | Pointer to **[]string** | An array of allowed IPs that can access the service | [optional] 

## Methods

### NewGetNetworkApplianceFirewallFirewalledServices200ResponseInner

`func NewGetNetworkApplianceFirewallFirewalledServices200ResponseInner() *GetNetworkApplianceFirewallFirewalledServices200ResponseInner`

NewGetNetworkApplianceFirewallFirewalledServices200ResponseInner instantiates a new GetNetworkApplianceFirewallFirewalledServices200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkApplianceFirewallFirewalledServices200ResponseInnerWithDefaults

`func NewGetNetworkApplianceFirewallFirewalledServices200ResponseInnerWithDefaults() *GetNetworkApplianceFirewallFirewalledServices200ResponseInner`

NewGetNetworkApplianceFirewallFirewalledServices200ResponseInnerWithDefaults instantiates a new GetNetworkApplianceFirewallFirewalledServices200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *GetNetworkApplianceFirewallFirewalledServices200ResponseInner) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GetNetworkApplianceFirewallFirewalledServices200ResponseInner) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GetNetworkApplianceFirewallFirewalledServices200ResponseInner) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *GetNetworkApplianceFirewallFirewalledServices200ResponseInner) HasService() bool`

HasService returns a boolean if a field has been set.

### GetAccess

`func (o *GetNetworkApplianceFirewallFirewalledServices200ResponseInner) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *GetNetworkApplianceFirewallFirewalledServices200ResponseInner) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *GetNetworkApplianceFirewallFirewalledServices200ResponseInner) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *GetNetworkApplianceFirewallFirewalledServices200ResponseInner) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetAllowedIps

`func (o *GetNetworkApplianceFirewallFirewalledServices200ResponseInner) GetAllowedIps() []string`

GetAllowedIps returns the AllowedIps field if non-nil, zero value otherwise.

### GetAllowedIpsOk

`func (o *GetNetworkApplianceFirewallFirewalledServices200ResponseInner) GetAllowedIpsOk() (*[]string, bool)`

GetAllowedIpsOk returns a tuple with the AllowedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIps

`func (o *GetNetworkApplianceFirewallFirewalledServices200ResponseInner) SetAllowedIps(v []string)`

SetAllowedIps sets AllowedIps field to given value.

### HasAllowedIps

`func (o *GetNetworkApplianceFirewallFirewalledServices200ResponseInner) HasAllowedIps() bool`

HasAllowedIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


