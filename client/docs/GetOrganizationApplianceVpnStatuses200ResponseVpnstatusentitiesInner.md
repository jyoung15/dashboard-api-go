# GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | Pointer to **string** | Network Id | [optional] 
**NetworkName** | Pointer to **string** | Network name | [optional] 
**DeviceSerial** | Pointer to **string** | Serial number of the device | [optional] 
**DeviceStatus** | Pointer to **string** | Device Status | [optional] 
**Uplinks** | Pointer to [**[]GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerUplinksInner**](GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerUplinksInner.md) | List of Uplink Information | [optional] 
**VpnMode** | Pointer to **string** | VPN Mode | [optional] 
**ExportedSubnets** | Pointer to [**[]GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner**](GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner.md) | List of Exported Subnets | [optional] 
**MerakiVpnPeers** | Pointer to [**[]GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerMerakiVpnPeersInner**](GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerMerakiVpnPeersInner.md) | Meraki VPN Peers | [optional] 
**ThirdPartyVpnPeers** | Pointer to [**[]GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerThirdPartyVpnPeersInner**](GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerThirdPartyVpnPeersInner.md) | Third Party VPN Peers | [optional] 

## Methods

### NewGetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner

`func NewGetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner() *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner`

NewGetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner instantiates a new GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerWithDefaults

`func NewGetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerWithDefaults() *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner`

NewGetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerWithDefaults instantiates a new GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkName

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetDeviceSerial

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) GetDeviceSerial() string`

GetDeviceSerial returns the DeviceSerial field if non-nil, zero value otherwise.

### GetDeviceSerialOk

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) GetDeviceSerialOk() (*string, bool)`

GetDeviceSerialOk returns a tuple with the DeviceSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSerial

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) SetDeviceSerial(v string)`

SetDeviceSerial sets DeviceSerial field to given value.

### HasDeviceSerial

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) HasDeviceSerial() bool`

HasDeviceSerial returns a boolean if a field has been set.

### GetDeviceStatus

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) GetDeviceStatus() string`

GetDeviceStatus returns the DeviceStatus field if non-nil, zero value otherwise.

### GetDeviceStatusOk

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) GetDeviceStatusOk() (*string, bool)`

GetDeviceStatusOk returns a tuple with the DeviceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatus

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) SetDeviceStatus(v string)`

SetDeviceStatus sets DeviceStatus field to given value.

### HasDeviceStatus

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) HasDeviceStatus() bool`

HasDeviceStatus returns a boolean if a field has been set.

### GetUplinks

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) GetUplinks() []GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerUplinksInner`

GetUplinks returns the Uplinks field if non-nil, zero value otherwise.

### GetUplinksOk

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) GetUplinksOk() (*[]GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerUplinksInner, bool)`

GetUplinksOk returns a tuple with the Uplinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinks

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) SetUplinks(v []GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerUplinksInner)`

SetUplinks sets Uplinks field to given value.

### HasUplinks

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) HasUplinks() bool`

HasUplinks returns a boolean if a field has been set.

### GetVpnMode

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) GetVpnMode() string`

GetVpnMode returns the VpnMode field if non-nil, zero value otherwise.

### GetVpnModeOk

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) GetVpnModeOk() (*string, bool)`

GetVpnModeOk returns a tuple with the VpnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnMode

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) SetVpnMode(v string)`

SetVpnMode sets VpnMode field to given value.

### HasVpnMode

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) HasVpnMode() bool`

HasVpnMode returns a boolean if a field has been set.

### GetExportedSubnets

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) GetExportedSubnets() []GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner`

GetExportedSubnets returns the ExportedSubnets field if non-nil, zero value otherwise.

### GetExportedSubnetsOk

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) GetExportedSubnetsOk() (*[]GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner, bool)`

GetExportedSubnetsOk returns a tuple with the ExportedSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedSubnets

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) SetExportedSubnets(v []GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerExportedSubnetsInner)`

SetExportedSubnets sets ExportedSubnets field to given value.

### HasExportedSubnets

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) HasExportedSubnets() bool`

HasExportedSubnets returns a boolean if a field has been set.

### GetMerakiVpnPeers

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) GetMerakiVpnPeers() []GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerMerakiVpnPeersInner`

GetMerakiVpnPeers returns the MerakiVpnPeers field if non-nil, zero value otherwise.

### GetMerakiVpnPeersOk

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) GetMerakiVpnPeersOk() (*[]GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerMerakiVpnPeersInner, bool)`

GetMerakiVpnPeersOk returns a tuple with the MerakiVpnPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerakiVpnPeers

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) SetMerakiVpnPeers(v []GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerMerakiVpnPeersInner)`

SetMerakiVpnPeers sets MerakiVpnPeers field to given value.

### HasMerakiVpnPeers

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) HasMerakiVpnPeers() bool`

HasMerakiVpnPeers returns a boolean if a field has been set.

### GetThirdPartyVpnPeers

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) GetThirdPartyVpnPeers() []GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerThirdPartyVpnPeersInner`

GetThirdPartyVpnPeers returns the ThirdPartyVpnPeers field if non-nil, zero value otherwise.

### GetThirdPartyVpnPeersOk

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) GetThirdPartyVpnPeersOk() (*[]GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerThirdPartyVpnPeersInner, bool)`

GetThirdPartyVpnPeersOk returns a tuple with the ThirdPartyVpnPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyVpnPeers

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) SetThirdPartyVpnPeers(v []GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInnerThirdPartyVpnPeersInner)`

SetThirdPartyVpnPeers sets ThirdPartyVpnPeers field to given value.

### HasThirdPartyVpnPeers

`func (o *GetOrganizationApplianceVpnStatuses200ResponseVpnstatusentitiesInner) HasThirdPartyVpnPeers() bool`

HasThirdPartyVpnPeers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


