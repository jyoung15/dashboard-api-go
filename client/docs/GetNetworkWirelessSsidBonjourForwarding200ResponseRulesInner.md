# GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Desctiption of the bonjour forwarding rule | [optional] 
**VlanId** | Pointer to **string** | The ID of the service VLAN. Required | [optional] 
**Services** | Pointer to **[]string** | A list of Bonjour services. At least one service must be specified. Available services are &#39;All Services&#39;, &#39;AirPlay&#39;, &#39;AFP&#39;, &#39;BitTorrent&#39;, &#39;FTP&#39;, &#39;iChat&#39;, &#39;iTunes&#39;, &#39;Printers&#39;, &#39;Samba&#39;, &#39;Scanners&#39; and &#39;SSH&#39; | [optional] 

## Methods

### NewGetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner

`func NewGetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner() *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner`

NewGetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner instantiates a new GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWirelessSsidBonjourForwarding200ResponseRulesInnerWithDefaults

`func NewGetNetworkWirelessSsidBonjourForwarding200ResponseRulesInnerWithDefaults() *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner`

NewGetNetworkWirelessSsidBonjourForwarding200ResponseRulesInnerWithDefaults instantiates a new GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVlanId

`func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) GetVlanId() string`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) GetVlanIdOk() (*string, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) SetVlanId(v string)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetServices

`func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *GetNetworkWirelessSsidBonjourForwarding200ResponseRulesInner) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


