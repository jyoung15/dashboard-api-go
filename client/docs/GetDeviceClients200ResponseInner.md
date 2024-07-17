# GetDeviceClients200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the client | [optional] 
**Mac** | Pointer to **string** | The MAC address of the client | [optional] 
**Description** | Pointer to **string** | Short description of the client | [optional] 
**MdnsName** | Pointer to **string** | The client&#39;s MDNS name | [optional] 
**DhcpHostname** | Pointer to **string** | The client&#39;s DHCP hostname | [optional] 
**User** | Pointer to **string** | The client user&#39;s name | [optional] 
**Ip** | Pointer to **string** | The IP address of the client | [optional] 
**Vlan** | Pointer to **string** | The client-assigned name of the VLAN the client is connected to | [optional] 
**NamedVlan** | Pointer to **string** | The owner-assigned name of the VLAN the client is connected to | [optional] 
**Switchport** | Pointer to **string** | The name of the switchport with clients on it, if the device is a switch | [optional] 
**AdaptivePolicyGroup** | Pointer to **string** | A description of the adaptive policy group | [optional] 
**Usage** | Pointer to [**GetDeviceClients200ResponseInnerUsage**](GetDeviceClients200ResponseInnerUsage.md) |  | [optional] 

## Methods

### NewGetDeviceClients200ResponseInner

`func NewGetDeviceClients200ResponseInner() *GetDeviceClients200ResponseInner`

NewGetDeviceClients200ResponseInner instantiates a new GetDeviceClients200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceClients200ResponseInnerWithDefaults

`func NewGetDeviceClients200ResponseInnerWithDefaults() *GetDeviceClients200ResponseInner`

NewGetDeviceClients200ResponseInnerWithDefaults instantiates a new GetDeviceClients200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetDeviceClients200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDeviceClients200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDeviceClients200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetDeviceClients200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMac

`func (o *GetDeviceClients200ResponseInner) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetDeviceClients200ResponseInner) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetDeviceClients200ResponseInner) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetDeviceClients200ResponseInner) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetDescription

`func (o *GetDeviceClients200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetDeviceClients200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetDeviceClients200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetDeviceClients200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMdnsName

`func (o *GetDeviceClients200ResponseInner) GetMdnsName() string`

GetMdnsName returns the MdnsName field if non-nil, zero value otherwise.

### GetMdnsNameOk

`func (o *GetDeviceClients200ResponseInner) GetMdnsNameOk() (*string, bool)`

GetMdnsNameOk returns a tuple with the MdnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdnsName

`func (o *GetDeviceClients200ResponseInner) SetMdnsName(v string)`

SetMdnsName sets MdnsName field to given value.

### HasMdnsName

`func (o *GetDeviceClients200ResponseInner) HasMdnsName() bool`

HasMdnsName returns a boolean if a field has been set.

### GetDhcpHostname

`func (o *GetDeviceClients200ResponseInner) GetDhcpHostname() string`

GetDhcpHostname returns the DhcpHostname field if non-nil, zero value otherwise.

### GetDhcpHostnameOk

`func (o *GetDeviceClients200ResponseInner) GetDhcpHostnameOk() (*string, bool)`

GetDhcpHostnameOk returns a tuple with the DhcpHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpHostname

`func (o *GetDeviceClients200ResponseInner) SetDhcpHostname(v string)`

SetDhcpHostname sets DhcpHostname field to given value.

### HasDhcpHostname

`func (o *GetDeviceClients200ResponseInner) HasDhcpHostname() bool`

HasDhcpHostname returns a boolean if a field has been set.

### GetUser

`func (o *GetDeviceClients200ResponseInner) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetDeviceClients200ResponseInner) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetDeviceClients200ResponseInner) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *GetDeviceClients200ResponseInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetIp

`func (o *GetDeviceClients200ResponseInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetDeviceClients200ResponseInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetDeviceClients200ResponseInner) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GetDeviceClients200ResponseInner) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetVlan

`func (o *GetDeviceClients200ResponseInner) GetVlan() string`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *GetDeviceClients200ResponseInner) GetVlanOk() (*string, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *GetDeviceClients200ResponseInner) SetVlan(v string)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *GetDeviceClients200ResponseInner) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetNamedVlan

`func (o *GetDeviceClients200ResponseInner) GetNamedVlan() string`

GetNamedVlan returns the NamedVlan field if non-nil, zero value otherwise.

### GetNamedVlanOk

`func (o *GetDeviceClients200ResponseInner) GetNamedVlanOk() (*string, bool)`

GetNamedVlanOk returns a tuple with the NamedVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedVlan

`func (o *GetDeviceClients200ResponseInner) SetNamedVlan(v string)`

SetNamedVlan sets NamedVlan field to given value.

### HasNamedVlan

`func (o *GetDeviceClients200ResponseInner) HasNamedVlan() bool`

HasNamedVlan returns a boolean if a field has been set.

### GetSwitchport

`func (o *GetDeviceClients200ResponseInner) GetSwitchport() string`

GetSwitchport returns the Switchport field if non-nil, zero value otherwise.

### GetSwitchportOk

`func (o *GetDeviceClients200ResponseInner) GetSwitchportOk() (*string, bool)`

GetSwitchportOk returns a tuple with the Switchport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchport

`func (o *GetDeviceClients200ResponseInner) SetSwitchport(v string)`

SetSwitchport sets Switchport field to given value.

### HasSwitchport

`func (o *GetDeviceClients200ResponseInner) HasSwitchport() bool`

HasSwitchport returns a boolean if a field has been set.

### GetAdaptivePolicyGroup

`func (o *GetDeviceClients200ResponseInner) GetAdaptivePolicyGroup() string`

GetAdaptivePolicyGroup returns the AdaptivePolicyGroup field if non-nil, zero value otherwise.

### GetAdaptivePolicyGroupOk

`func (o *GetDeviceClients200ResponseInner) GetAdaptivePolicyGroupOk() (*string, bool)`

GetAdaptivePolicyGroupOk returns a tuple with the AdaptivePolicyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdaptivePolicyGroup

`func (o *GetDeviceClients200ResponseInner) SetAdaptivePolicyGroup(v string)`

SetAdaptivePolicyGroup sets AdaptivePolicyGroup field to given value.

### HasAdaptivePolicyGroup

`func (o *GetDeviceClients200ResponseInner) HasAdaptivePolicyGroup() bool`

HasAdaptivePolicyGroup returns a boolean if a field has been set.

### GetUsage

`func (o *GetDeviceClients200ResponseInner) GetUsage() GetDeviceClients200ResponseInnerUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *GetDeviceClients200ResponseInner) GetUsageOk() (*GetDeviceClients200ResponseInnerUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *GetDeviceClients200ResponseInner) SetUsage(v GetDeviceClients200ResponseInnerUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *GetDeviceClients200ResponseInner) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


