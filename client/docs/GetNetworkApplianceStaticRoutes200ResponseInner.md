# GetNetworkApplianceStaticRoutes200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Route ID | [optional] 
**IpVersion** | Pointer to **int32** | IP protocol version | [optional] 
**NetworkId** | Pointer to **string** | Network ID | [optional] 
**Enabled** | Pointer to **bool** | Whether the route is enabled or not | [optional] 
**Name** | Pointer to **string** | Name of the route | [optional] 
**Subnet** | Pointer to **string** | Subnet of the route | [optional] 
**GatewayIp** | Pointer to **string** | Gateway IP address (next hop) | [optional] 
**FixedIpAssignments** | Pointer to [**map[string]GetNetworkApplianceStaticRoutes200ResponseInnerFixedIpAssignmentsValue**](GetNetworkApplianceStaticRoutes200ResponseInnerFixedIpAssignmentsValue.md) | Fixed DHCP IP assignments on the route | [optional] 
**ReservedIpRanges** | Pointer to [**[]GetNetworkApplianceStaticRoutes200ResponseInnerReservedIpRangesInner**](GetNetworkApplianceStaticRoutes200ResponseInnerReservedIpRangesInner.md) | DHCP reserved IP ranges | [optional] 
**GatewayVlanId** | Pointer to **int32** | Gateway VLAN ID | [optional] 

## Methods

### NewGetNetworkApplianceStaticRoutes200ResponseInner

`func NewGetNetworkApplianceStaticRoutes200ResponseInner() *GetNetworkApplianceStaticRoutes200ResponseInner`

NewGetNetworkApplianceStaticRoutes200ResponseInner instantiates a new GetNetworkApplianceStaticRoutes200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkApplianceStaticRoutes200ResponseInnerWithDefaults

`func NewGetNetworkApplianceStaticRoutes200ResponseInnerWithDefaults() *GetNetworkApplianceStaticRoutes200ResponseInner`

NewGetNetworkApplianceStaticRoutes200ResponseInnerWithDefaults instantiates a new GetNetworkApplianceStaticRoutes200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpVersion

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) GetIpVersion() int32`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) GetIpVersionOk() (*int32, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) SetIpVersion(v int32)`

SetIpVersion sets IpVersion field to given value.

### HasIpVersion

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) HasIpVersion() bool`

HasIpVersion returns a boolean if a field has been set.

### GetNetworkId

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetEnabled

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubnet

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetGatewayIp

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) GetGatewayIp() string`

GetGatewayIp returns the GatewayIp field if non-nil, zero value otherwise.

### GetGatewayIpOk

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) GetGatewayIpOk() (*string, bool)`

GetGatewayIpOk returns a tuple with the GatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIp

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) SetGatewayIp(v string)`

SetGatewayIp sets GatewayIp field to given value.

### HasGatewayIp

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) HasGatewayIp() bool`

HasGatewayIp returns a boolean if a field has been set.

### GetFixedIpAssignments

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) GetFixedIpAssignments() map[string]GetNetworkApplianceStaticRoutes200ResponseInnerFixedIpAssignmentsValue`

GetFixedIpAssignments returns the FixedIpAssignments field if non-nil, zero value otherwise.

### GetFixedIpAssignmentsOk

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) GetFixedIpAssignmentsOk() (*map[string]GetNetworkApplianceStaticRoutes200ResponseInnerFixedIpAssignmentsValue, bool)`

GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAssignments

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) SetFixedIpAssignments(v map[string]GetNetworkApplianceStaticRoutes200ResponseInnerFixedIpAssignmentsValue)`

SetFixedIpAssignments sets FixedIpAssignments field to given value.

### HasFixedIpAssignments

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) HasFixedIpAssignments() bool`

HasFixedIpAssignments returns a boolean if a field has been set.

### GetReservedIpRanges

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) GetReservedIpRanges() []GetNetworkApplianceStaticRoutes200ResponseInnerReservedIpRangesInner`

GetReservedIpRanges returns the ReservedIpRanges field if non-nil, zero value otherwise.

### GetReservedIpRangesOk

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) GetReservedIpRangesOk() (*[]GetNetworkApplianceStaticRoutes200ResponseInnerReservedIpRangesInner, bool)`

GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIpRanges

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) SetReservedIpRanges(v []GetNetworkApplianceStaticRoutes200ResponseInnerReservedIpRangesInner)`

SetReservedIpRanges sets ReservedIpRanges field to given value.

### HasReservedIpRanges

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) HasReservedIpRanges() bool`

HasReservedIpRanges returns a boolean if a field has been set.

### GetGatewayVlanId

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) GetGatewayVlanId() int32`

GetGatewayVlanId returns the GatewayVlanId field if non-nil, zero value otherwise.

### GetGatewayVlanIdOk

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) GetGatewayVlanIdOk() (*int32, bool)`

GetGatewayVlanIdOk returns a tuple with the GatewayVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayVlanId

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) SetGatewayVlanId(v int32)`

SetGatewayVlanId sets GatewayVlanId field to given value.

### HasGatewayVlanId

`func (o *GetNetworkApplianceStaticRoutes200ResponseInner) HasGatewayVlanId() bool`

HasGatewayVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


