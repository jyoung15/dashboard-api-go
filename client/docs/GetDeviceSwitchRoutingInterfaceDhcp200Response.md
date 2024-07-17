# GetDeviceSwitchRoutingInterfaceDhcp200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpMode** | Pointer to **string** | The DHCP mode options for the switch stack interface (&#39;dhcpDisabled&#39;, &#39;dhcpRelay&#39; or &#39;dhcpServer&#39;) | [optional] 
**DhcpRelayServerIps** | Pointer to **[]string** | The DHCP relay server IPs to which DHCP packets would get relayed for the switch stack interface | [optional] 
**DhcpLeaseTime** | Pointer to **string** | The DHCP lease time config for the dhcp server running on the switch stack interface (&#39;30 minutes&#39;, &#39;1 hour&#39;, &#39;4 hours&#39;, &#39;12 hours&#39;, &#39;1 day&#39; or &#39;1 week&#39;) | [optional] 
**DnsNameserversOption** | Pointer to **string** | The DHCP name server option for the dhcp server running on the switch stack interface (&#39;googlePublicDns&#39;, &#39;openDns&#39; or &#39;custom&#39;) | [optional] 
**DnsCustomNameservers** | Pointer to **[]string** | The DHCP name server IPs when DHCP name server option is &#39;custom&#39; | [optional] 
**BootOptionsEnabled** | Pointer to **bool** | Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch stack interface | [optional] 
**BootNextServer** | Pointer to **string** | The PXE boot server IP for the DHCP server running on the switch stack interface | [optional] 
**BootFileName** | Pointer to **string** | The PXE boot server file name for the DHCP server running on the switch stack interface | [optional] 
**DhcpOptions** | Pointer to [**[]GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner**](GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner.md) | Array of DHCP options consisting of code, type and value for the DHCP server running on the switch stack interface | [optional] 
**ReservedIpRanges** | Pointer to [**[]GetDeviceSwitchRoutingInterfaceDhcp200ResponseReservedIpRangesInner**](GetDeviceSwitchRoutingInterfaceDhcp200ResponseReservedIpRangesInner.md) | Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface | [optional] 
**FixedIpAssignments** | Pointer to [**[]GetDeviceSwitchRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner**](GetDeviceSwitchRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner.md) | Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface | [optional] 

## Methods

### NewGetDeviceSwitchRoutingInterfaceDhcp200Response

`func NewGetDeviceSwitchRoutingInterfaceDhcp200Response() *GetDeviceSwitchRoutingInterfaceDhcp200Response`

NewGetDeviceSwitchRoutingInterfaceDhcp200Response instantiates a new GetDeviceSwitchRoutingInterfaceDhcp200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceSwitchRoutingInterfaceDhcp200ResponseWithDefaults

`func NewGetDeviceSwitchRoutingInterfaceDhcp200ResponseWithDefaults() *GetDeviceSwitchRoutingInterfaceDhcp200Response`

NewGetDeviceSwitchRoutingInterfaceDhcp200ResponseWithDefaults instantiates a new GetDeviceSwitchRoutingInterfaceDhcp200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpMode

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) GetDhcpMode() string`

GetDhcpMode returns the DhcpMode field if non-nil, zero value otherwise.

### GetDhcpModeOk

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) GetDhcpModeOk() (*string, bool)`

GetDhcpModeOk returns a tuple with the DhcpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpMode

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) SetDhcpMode(v string)`

SetDhcpMode sets DhcpMode field to given value.

### HasDhcpMode

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) HasDhcpMode() bool`

HasDhcpMode returns a boolean if a field has been set.

### GetDhcpRelayServerIps

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) GetDhcpRelayServerIps() []string`

GetDhcpRelayServerIps returns the DhcpRelayServerIps field if non-nil, zero value otherwise.

### GetDhcpRelayServerIpsOk

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) GetDhcpRelayServerIpsOk() (*[]string, bool)`

GetDhcpRelayServerIpsOk returns a tuple with the DhcpRelayServerIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRelayServerIps

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) SetDhcpRelayServerIps(v []string)`

SetDhcpRelayServerIps sets DhcpRelayServerIps field to given value.

### HasDhcpRelayServerIps

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) HasDhcpRelayServerIps() bool`

HasDhcpRelayServerIps returns a boolean if a field has been set.

### GetDhcpLeaseTime

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) GetDhcpLeaseTime() string`

GetDhcpLeaseTime returns the DhcpLeaseTime field if non-nil, zero value otherwise.

### GetDhcpLeaseTimeOk

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) GetDhcpLeaseTimeOk() (*string, bool)`

GetDhcpLeaseTimeOk returns a tuple with the DhcpLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpLeaseTime

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) SetDhcpLeaseTime(v string)`

SetDhcpLeaseTime sets DhcpLeaseTime field to given value.

### HasDhcpLeaseTime

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) HasDhcpLeaseTime() bool`

HasDhcpLeaseTime returns a boolean if a field has been set.

### GetDnsNameserversOption

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) GetDnsNameserversOption() string`

GetDnsNameserversOption returns the DnsNameserversOption field if non-nil, zero value otherwise.

### GetDnsNameserversOptionOk

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) GetDnsNameserversOptionOk() (*string, bool)`

GetDnsNameserversOptionOk returns a tuple with the DnsNameserversOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNameserversOption

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) SetDnsNameserversOption(v string)`

SetDnsNameserversOption sets DnsNameserversOption field to given value.

### HasDnsNameserversOption

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) HasDnsNameserversOption() bool`

HasDnsNameserversOption returns a boolean if a field has been set.

### GetDnsCustomNameservers

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) GetDnsCustomNameservers() []string`

GetDnsCustomNameservers returns the DnsCustomNameservers field if non-nil, zero value otherwise.

### GetDnsCustomNameserversOk

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) GetDnsCustomNameserversOk() (*[]string, bool)`

GetDnsCustomNameserversOk returns a tuple with the DnsCustomNameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCustomNameservers

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) SetDnsCustomNameservers(v []string)`

SetDnsCustomNameservers sets DnsCustomNameservers field to given value.

### HasDnsCustomNameservers

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) HasDnsCustomNameservers() bool`

HasDnsCustomNameservers returns a boolean if a field has been set.

### GetBootOptionsEnabled

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) GetBootOptionsEnabled() bool`

GetBootOptionsEnabled returns the BootOptionsEnabled field if non-nil, zero value otherwise.

### GetBootOptionsEnabledOk

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) GetBootOptionsEnabledOk() (*bool, bool)`

GetBootOptionsEnabledOk returns a tuple with the BootOptionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOptionsEnabled

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) SetBootOptionsEnabled(v bool)`

SetBootOptionsEnabled sets BootOptionsEnabled field to given value.

### HasBootOptionsEnabled

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) HasBootOptionsEnabled() bool`

HasBootOptionsEnabled returns a boolean if a field has been set.

### GetBootNextServer

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) GetBootNextServer() string`

GetBootNextServer returns the BootNextServer field if non-nil, zero value otherwise.

### GetBootNextServerOk

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) GetBootNextServerOk() (*string, bool)`

GetBootNextServerOk returns a tuple with the BootNextServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootNextServer

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) SetBootNextServer(v string)`

SetBootNextServer sets BootNextServer field to given value.

### HasBootNextServer

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) HasBootNextServer() bool`

HasBootNextServer returns a boolean if a field has been set.

### GetBootFileName

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) GetBootFileName() string`

GetBootFileName returns the BootFileName field if non-nil, zero value otherwise.

### GetBootFileNameOk

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) GetBootFileNameOk() (*string, bool)`

GetBootFileNameOk returns a tuple with the BootFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootFileName

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) SetBootFileName(v string)`

SetBootFileName sets BootFileName field to given value.

### HasBootFileName

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) HasBootFileName() bool`

HasBootFileName returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) GetDhcpOptions() []GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) GetDhcpOptionsOk() (*[]GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) SetDhcpOptions(v []GetDeviceSwitchRoutingInterfaceDhcp200ResponseDhcpOptionsInner)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetReservedIpRanges

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) GetReservedIpRanges() []GetDeviceSwitchRoutingInterfaceDhcp200ResponseReservedIpRangesInner`

GetReservedIpRanges returns the ReservedIpRanges field if non-nil, zero value otherwise.

### GetReservedIpRangesOk

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) GetReservedIpRangesOk() (*[]GetDeviceSwitchRoutingInterfaceDhcp200ResponseReservedIpRangesInner, bool)`

GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIpRanges

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) SetReservedIpRanges(v []GetDeviceSwitchRoutingInterfaceDhcp200ResponseReservedIpRangesInner)`

SetReservedIpRanges sets ReservedIpRanges field to given value.

### HasReservedIpRanges

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) HasReservedIpRanges() bool`

HasReservedIpRanges returns a boolean if a field has been set.

### GetFixedIpAssignments

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) GetFixedIpAssignments() []GetDeviceSwitchRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner`

GetFixedIpAssignments returns the FixedIpAssignments field if non-nil, zero value otherwise.

### GetFixedIpAssignmentsOk

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) GetFixedIpAssignmentsOk() (*[]GetDeviceSwitchRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner, bool)`

GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAssignments

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) SetFixedIpAssignments(v []GetDeviceSwitchRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner)`

SetFixedIpAssignments sets FixedIpAssignments field to given value.

### HasFixedIpAssignments

`func (o *GetDeviceSwitchRoutingInterfaceDhcp200Response) HasFixedIpAssignments() bool`

HasFixedIpAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


