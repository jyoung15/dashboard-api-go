# GetNetworkApplianceVpnBgp200ResponseNeighborsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | The IPv4 address of the neighbor | [optional] 
**Ipv6** | Pointer to [**GetNetworkApplianceVpnBgp200ResponseNeighborsInnerIpv6**](GetNetworkApplianceVpnBgp200ResponseNeighborsInnerIpv6.md) |  | [optional] 
**RemoteAsNumber** | Pointer to **int32** | Remote AS number of the neighbor | [optional] 
**ReceiveLimit** | Pointer to **int32** | The maximum number of routes that the appliance can receive from the neighbor | [optional] 
**AllowTransit** | Pointer to **bool** | Whether the appliance will advertise routes learned from other Autonomous Systems | [optional] 
**EbgpHoldTimer** | Pointer to **int32** | The eBGP hold time in seconds for the neighbor | [optional] 
**EbgpMultihop** | Pointer to **int32** | The number of hops the appliance must traverse to establish a peering relationship with the neighbor | [optional] 
**SourceInterface** | Pointer to **string** | The output interface the appliance uses to establish a peering relationship with the neighbor | [optional] 
**NextHopIp** | Pointer to **string** | The IPv4 address of the neighbor that will establish a TCP session with the appliance | [optional] 
**TtlSecurity** | Pointer to [**GetNetworkApplianceVpnBgp200ResponseNeighborsInnerTtlSecurity**](GetNetworkApplianceVpnBgp200ResponseNeighborsInnerTtlSecurity.md) |  | [optional] 
**Authentication** | Pointer to [**GetNetworkApplianceVpnBgp200ResponseNeighborsInnerAuthentication**](GetNetworkApplianceVpnBgp200ResponseNeighborsInnerAuthentication.md) |  | [optional] 

## Methods

### NewGetNetworkApplianceVpnBgp200ResponseNeighborsInner

`func NewGetNetworkApplianceVpnBgp200ResponseNeighborsInner() *GetNetworkApplianceVpnBgp200ResponseNeighborsInner`

NewGetNetworkApplianceVpnBgp200ResponseNeighborsInner instantiates a new GetNetworkApplianceVpnBgp200ResponseNeighborsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkApplianceVpnBgp200ResponseNeighborsInnerWithDefaults

`func NewGetNetworkApplianceVpnBgp200ResponseNeighborsInnerWithDefaults() *GetNetworkApplianceVpnBgp200ResponseNeighborsInner`

NewGetNetworkApplianceVpnBgp200ResponseNeighborsInnerWithDefaults instantiates a new GetNetworkApplianceVpnBgp200ResponseNeighborsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) GetIpv6() GetNetworkApplianceVpnBgp200ResponseNeighborsInnerIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) GetIpv6Ok() (*GetNetworkApplianceVpnBgp200ResponseNeighborsInnerIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) SetIpv6(v GetNetworkApplianceVpnBgp200ResponseNeighborsInnerIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetRemoteAsNumber

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) GetRemoteAsNumber() int32`

GetRemoteAsNumber returns the RemoteAsNumber field if non-nil, zero value otherwise.

### GetRemoteAsNumberOk

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) GetRemoteAsNumberOk() (*int32, bool)`

GetRemoteAsNumberOk returns a tuple with the RemoteAsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAsNumber

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) SetRemoteAsNumber(v int32)`

SetRemoteAsNumber sets RemoteAsNumber field to given value.

### HasRemoteAsNumber

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) HasRemoteAsNumber() bool`

HasRemoteAsNumber returns a boolean if a field has been set.

### GetReceiveLimit

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) GetReceiveLimit() int32`

GetReceiveLimit returns the ReceiveLimit field if non-nil, zero value otherwise.

### GetReceiveLimitOk

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) GetReceiveLimitOk() (*int32, bool)`

GetReceiveLimitOk returns a tuple with the ReceiveLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveLimit

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) SetReceiveLimit(v int32)`

SetReceiveLimit sets ReceiveLimit field to given value.

### HasReceiveLimit

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) HasReceiveLimit() bool`

HasReceiveLimit returns a boolean if a field has been set.

### GetAllowTransit

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) GetAllowTransit() bool`

GetAllowTransit returns the AllowTransit field if non-nil, zero value otherwise.

### GetAllowTransitOk

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) GetAllowTransitOk() (*bool, bool)`

GetAllowTransitOk returns a tuple with the AllowTransit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTransit

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) SetAllowTransit(v bool)`

SetAllowTransit sets AllowTransit field to given value.

### HasAllowTransit

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) HasAllowTransit() bool`

HasAllowTransit returns a boolean if a field has been set.

### GetEbgpHoldTimer

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) GetEbgpHoldTimer() int32`

GetEbgpHoldTimer returns the EbgpHoldTimer field if non-nil, zero value otherwise.

### GetEbgpHoldTimerOk

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) GetEbgpHoldTimerOk() (*int32, bool)`

GetEbgpHoldTimerOk returns a tuple with the EbgpHoldTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpHoldTimer

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) SetEbgpHoldTimer(v int32)`

SetEbgpHoldTimer sets EbgpHoldTimer field to given value.

### HasEbgpHoldTimer

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) HasEbgpHoldTimer() bool`

HasEbgpHoldTimer returns a boolean if a field has been set.

### GetEbgpMultihop

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) GetEbgpMultihop() int32`

GetEbgpMultihop returns the EbgpMultihop field if non-nil, zero value otherwise.

### GetEbgpMultihopOk

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) GetEbgpMultihopOk() (*int32, bool)`

GetEbgpMultihopOk returns a tuple with the EbgpMultihop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbgpMultihop

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) SetEbgpMultihop(v int32)`

SetEbgpMultihop sets EbgpMultihop field to given value.

### HasEbgpMultihop

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) HasEbgpMultihop() bool`

HasEbgpMultihop returns a boolean if a field has been set.

### GetSourceInterface

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) GetSourceInterface() string`

GetSourceInterface returns the SourceInterface field if non-nil, zero value otherwise.

### GetSourceInterfaceOk

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) GetSourceInterfaceOk() (*string, bool)`

GetSourceInterfaceOk returns a tuple with the SourceInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInterface

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) SetSourceInterface(v string)`

SetSourceInterface sets SourceInterface field to given value.

### HasSourceInterface

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) HasSourceInterface() bool`

HasSourceInterface returns a boolean if a field has been set.

### GetNextHopIp

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) GetNextHopIp() string`

GetNextHopIp returns the NextHopIp field if non-nil, zero value otherwise.

### GetNextHopIpOk

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) GetNextHopIpOk() (*string, bool)`

GetNextHopIpOk returns a tuple with the NextHopIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHopIp

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) SetNextHopIp(v string)`

SetNextHopIp sets NextHopIp field to given value.

### HasNextHopIp

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) HasNextHopIp() bool`

HasNextHopIp returns a boolean if a field has been set.

### GetTtlSecurity

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) GetTtlSecurity() GetNetworkApplianceVpnBgp200ResponseNeighborsInnerTtlSecurity`

GetTtlSecurity returns the TtlSecurity field if non-nil, zero value otherwise.

### GetTtlSecurityOk

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) GetTtlSecurityOk() (*GetNetworkApplianceVpnBgp200ResponseNeighborsInnerTtlSecurity, bool)`

GetTtlSecurityOk returns a tuple with the TtlSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlSecurity

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) SetTtlSecurity(v GetNetworkApplianceVpnBgp200ResponseNeighborsInnerTtlSecurity)`

SetTtlSecurity sets TtlSecurity field to given value.

### HasTtlSecurity

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) HasTtlSecurity() bool`

HasTtlSecurity returns a boolean if a field has been set.

### GetAuthentication

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) GetAuthentication() GetNetworkApplianceVpnBgp200ResponseNeighborsInnerAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) GetAuthenticationOk() (*GetNetworkApplianceVpnBgp200ResponseNeighborsInnerAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) SetAuthentication(v GetNetworkApplianceVpnBgp200ResponseNeighborsInnerAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *GetNetworkApplianceVpnBgp200ResponseNeighborsInner) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


