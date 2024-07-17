# UpdateNetworkSwitchStackRoutingInterface200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceId** | Pointer to **string** | The id | [optional] 
**Name** | Pointer to **string** | The name | [optional] 
**Subnet** | Pointer to **string** | IPv4 subnet | [optional] 
**InterfaceIp** | Pointer to **string** | IPv4 address | [optional] 
**MulticastRouting** | Pointer to **string** | Multicast routing status | [optional] 
**VlanId** | Pointer to **int32** | VLAN id | [optional] 
**OspfSettings** | Pointer to [**GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings**](GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings.md) |  | [optional] 
**OspfV3** | Pointer to [**GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3**](GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3.md) |  | [optional] 
**Ipv6** | Pointer to [**GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6**](GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6.md) |  | [optional] 
**UplinkV4** | Pointer to **bool** | Whether this is the switch&#39;s IPv4 uplink | [optional] 
**UplinkV6** | Pointer to **bool** | Whether this is the switch&#39;s IPv6 uplink | [optional] 

## Methods

### NewUpdateNetworkSwitchStackRoutingInterface200Response

`func NewUpdateNetworkSwitchStackRoutingInterface200Response() *UpdateNetworkSwitchStackRoutingInterface200Response`

NewUpdateNetworkSwitchStackRoutingInterface200Response instantiates a new UpdateNetworkSwitchStackRoutingInterface200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkSwitchStackRoutingInterface200ResponseWithDefaults

`func NewUpdateNetworkSwitchStackRoutingInterface200ResponseWithDefaults() *UpdateNetworkSwitchStackRoutingInterface200Response`

NewUpdateNetworkSwitchStackRoutingInterface200ResponseWithDefaults instantiates a new UpdateNetworkSwitchStackRoutingInterface200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceId

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetName

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetInterfaceIp

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) GetInterfaceIp() string`

GetInterfaceIp returns the InterfaceIp field if non-nil, zero value otherwise.

### GetInterfaceIpOk

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) GetInterfaceIpOk() (*string, bool)`

GetInterfaceIpOk returns a tuple with the InterfaceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIp

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) SetInterfaceIp(v string)`

SetInterfaceIp sets InterfaceIp field to given value.

### HasInterfaceIp

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) HasInterfaceIp() bool`

HasInterfaceIp returns a boolean if a field has been set.

### GetMulticastRouting

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) GetMulticastRouting() string`

GetMulticastRouting returns the MulticastRouting field if non-nil, zero value otherwise.

### GetMulticastRoutingOk

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) GetMulticastRoutingOk() (*string, bool)`

GetMulticastRoutingOk returns a tuple with the MulticastRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastRouting

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) SetMulticastRouting(v string)`

SetMulticastRouting sets MulticastRouting field to given value.

### HasMulticastRouting

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) HasMulticastRouting() bool`

HasMulticastRouting returns a boolean if a field has been set.

### GetVlanId

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetOspfSettings

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) GetOspfSettings() GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings`

GetOspfSettings returns the OspfSettings field if non-nil, zero value otherwise.

### GetOspfSettingsOk

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) GetOspfSettingsOk() (*GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings, bool)`

GetOspfSettingsOk returns a tuple with the OspfSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfSettings

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) SetOspfSettings(v GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfSettings)`

SetOspfSettings sets OspfSettings field to given value.

### HasOspfSettings

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) HasOspfSettings() bool`

HasOspfSettings returns a boolean if a field has been set.

### GetOspfV3

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) GetOspfV3() GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3`

GetOspfV3 returns the OspfV3 field if non-nil, zero value otherwise.

### GetOspfV3Ok

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) GetOspfV3Ok() (*GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3, bool)`

GetOspfV3Ok returns a tuple with the OspfV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfV3

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) SetOspfV3(v GetDeviceSwitchRoutingInterfaces200ResponseInnerOspfV3)`

SetOspfV3 sets OspfV3 field to given value.

### HasOspfV3

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) HasOspfV3() bool`

HasOspfV3 returns a boolean if a field has been set.

### GetIpv6

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) GetIpv6() GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) GetIpv6Ok() (*GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) SetIpv6(v GetDeviceSwitchRoutingInterfaces200ResponseInnerIpv6)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetUplinkV4

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) GetUplinkV4() bool`

GetUplinkV4 returns the UplinkV4 field if non-nil, zero value otherwise.

### GetUplinkV4Ok

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) GetUplinkV4Ok() (*bool, bool)`

GetUplinkV4Ok returns a tuple with the UplinkV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkV4

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) SetUplinkV4(v bool)`

SetUplinkV4 sets UplinkV4 field to given value.

### HasUplinkV4

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) HasUplinkV4() bool`

HasUplinkV4 returns a boolean if a field has been set.

### GetUplinkV6

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) GetUplinkV6() bool`

GetUplinkV6 returns the UplinkV6 field if non-nil, zero value otherwise.

### GetUplinkV6Ok

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) GetUplinkV6Ok() (*bool, bool)`

GetUplinkV6Ok returns a tuple with the UplinkV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkV6

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) SetUplinkV6(v bool)`

SetUplinkV6 sets UplinkV6 field to given value.

### HasUplinkV6

`func (o *UpdateNetworkSwitchStackRoutingInterface200Response) HasUplinkV6() bool`

HasUplinkV6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


