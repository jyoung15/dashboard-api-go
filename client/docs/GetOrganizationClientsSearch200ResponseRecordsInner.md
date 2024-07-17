# GetOrganizationClientsSearch200ResponseRecordsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**GetOrganizationClientsSearch200ResponseRecordsInnerNetwork**](GetOrganizationClientsSearch200ResponseRecordsInnerNetwork.md) |  | [optional] 
**Ip** | Pointer to **string** | The IP address of the client | [optional] 
**Ip6** | Pointer to **string** | The IPv6 address of the client | [optional] 
**Description** | Pointer to **string** | Short description of the client | [optional] 
**FirstSeen** | Pointer to **int32** | Timestamp client was first seen in the network | [optional] 
**LastSeen** | Pointer to **int32** | Timestamp client was last seen in the network | [optional] 
**Os** | Pointer to **string** | The operating system of the client | [optional] 
**User** | Pointer to **string** | The username of the user of the client | [optional] 
**Vlan** | Pointer to **string** | The name of the VLAN that the client is connected to | [optional] 
**Ssid** | Pointer to **string** | The name of the SSID that the client is connected to | [optional] 
**Switchport** | Pointer to **string** | The switch port the client is connected to | [optional] 
**WirelessCapabilities** | Pointer to **string** | Wireless capabilities of the client | [optional] 
**SmInstalled** | Pointer to **bool** | Status of SM for the client | [optional] 
**RecentDeviceMac** | Pointer to **string** | The MAC address of the node that the device was last connected to | [optional] 
**ClientVpnConnections** | Pointer to [**[]GetOrganizationClientsSearch200ResponseRecordsInnerClientVpnConnectionsInner**](GetOrganizationClientsSearch200ResponseRecordsInnerClientVpnConnectionsInner.md) | VPN connections associated with the client | [optional] 
**Lldp** | Pointer to **[][]string** | The link layer discover protocol settings for the client | [optional] 
**Cdp** | Pointer to **[][]string** | The Cisco discover protocol settings for the client | [optional] 
**Status** | Pointer to **string** | The connection status of the client | [optional] 

## Methods

### NewGetOrganizationClientsSearch200ResponseRecordsInner

`func NewGetOrganizationClientsSearch200ResponseRecordsInner() *GetOrganizationClientsSearch200ResponseRecordsInner`

NewGetOrganizationClientsSearch200ResponseRecordsInner instantiates a new GetOrganizationClientsSearch200ResponseRecordsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationClientsSearch200ResponseRecordsInnerWithDefaults

`func NewGetOrganizationClientsSearch200ResponseRecordsInnerWithDefaults() *GetOrganizationClientsSearch200ResponseRecordsInner`

NewGetOrganizationClientsSearch200ResponseRecordsInnerWithDefaults instantiates a new GetOrganizationClientsSearch200ResponseRecordsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetNetwork() GetOrganizationClientsSearch200ResponseRecordsInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetNetworkOk() (*GetOrganizationClientsSearch200ResponseRecordsInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) SetNetwork(v GetOrganizationClientsSearch200ResponseRecordsInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetIp

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIp6

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetIp6() string`

GetIp6 returns the Ip6 field if non-nil, zero value otherwise.

### GetIp6Ok

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetIp6Ok() (*string, bool)`

GetIp6Ok returns a tuple with the Ip6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp6

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) SetIp6(v string)`

SetIp6 sets Ip6 field to given value.

### HasIp6

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) HasIp6() bool`

HasIp6 returns a boolean if a field has been set.

### GetDescription

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFirstSeen

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetFirstSeen() int32`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetFirstSeenOk() (*int32, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) SetFirstSeen(v int32)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetLastSeen

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetLastSeen() int32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetLastSeenOk() (*int32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) SetLastSeen(v int32)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetOs

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetUser

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVlan

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetVlan() string`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetVlanOk() (*string, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) SetVlan(v string)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetSsid

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetSwitchport

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetSwitchport() string`

GetSwitchport returns the Switchport field if non-nil, zero value otherwise.

### GetSwitchportOk

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetSwitchportOk() (*string, bool)`

GetSwitchportOk returns a tuple with the Switchport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchport

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) SetSwitchport(v string)`

SetSwitchport sets Switchport field to given value.

### HasSwitchport

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) HasSwitchport() bool`

HasSwitchport returns a boolean if a field has been set.

### GetWirelessCapabilities

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetWirelessCapabilities() string`

GetWirelessCapabilities returns the WirelessCapabilities field if non-nil, zero value otherwise.

### GetWirelessCapabilitiesOk

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetWirelessCapabilitiesOk() (*string, bool)`

GetWirelessCapabilitiesOk returns a tuple with the WirelessCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessCapabilities

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) SetWirelessCapabilities(v string)`

SetWirelessCapabilities sets WirelessCapabilities field to given value.

### HasWirelessCapabilities

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) HasWirelessCapabilities() bool`

HasWirelessCapabilities returns a boolean if a field has been set.

### GetSmInstalled

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetSmInstalled() bool`

GetSmInstalled returns the SmInstalled field if non-nil, zero value otherwise.

### GetSmInstalledOk

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetSmInstalledOk() (*bool, bool)`

GetSmInstalledOk returns a tuple with the SmInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmInstalled

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) SetSmInstalled(v bool)`

SetSmInstalled sets SmInstalled field to given value.

### HasSmInstalled

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) HasSmInstalled() bool`

HasSmInstalled returns a boolean if a field has been set.

### GetRecentDeviceMac

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetRecentDeviceMac() string`

GetRecentDeviceMac returns the RecentDeviceMac field if non-nil, zero value otherwise.

### GetRecentDeviceMacOk

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetRecentDeviceMacOk() (*string, bool)`

GetRecentDeviceMacOk returns a tuple with the RecentDeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentDeviceMac

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) SetRecentDeviceMac(v string)`

SetRecentDeviceMac sets RecentDeviceMac field to given value.

### HasRecentDeviceMac

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) HasRecentDeviceMac() bool`

HasRecentDeviceMac returns a boolean if a field has been set.

### GetClientVpnConnections

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetClientVpnConnections() []GetOrganizationClientsSearch200ResponseRecordsInnerClientVpnConnectionsInner`

GetClientVpnConnections returns the ClientVpnConnections field if non-nil, zero value otherwise.

### GetClientVpnConnectionsOk

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetClientVpnConnectionsOk() (*[]GetOrganizationClientsSearch200ResponseRecordsInnerClientVpnConnectionsInner, bool)`

GetClientVpnConnectionsOk returns a tuple with the ClientVpnConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVpnConnections

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) SetClientVpnConnections(v []GetOrganizationClientsSearch200ResponseRecordsInnerClientVpnConnectionsInner)`

SetClientVpnConnections sets ClientVpnConnections field to given value.

### HasClientVpnConnections

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) HasClientVpnConnections() bool`

HasClientVpnConnections returns a boolean if a field has been set.

### GetLldp

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetLldp() [][]string`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetLldpOk() (*[][]string, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) SetLldp(v [][]string)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) HasLldp() bool`

HasLldp returns a boolean if a field has been set.

### GetCdp

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetCdp() [][]string`

GetCdp returns the Cdp field if non-nil, zero value otherwise.

### GetCdpOk

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetCdpOk() (*[][]string, bool)`

GetCdpOk returns a tuple with the Cdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdp

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) SetCdp(v [][]string)`

SetCdp sets Cdp field to given value.

### HasCdp

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) HasCdp() bool`

HasCdp returns a boolean if a field has been set.

### GetStatus

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetOrganizationClientsSearch200ResponseRecordsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


