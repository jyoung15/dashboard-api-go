# GetNetworkApplianceVpnBgp200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether BGP is enabled on the appliance | [optional] 
**AsNumber** | Pointer to **int32** | The number of the Autonomous System to which the appliance belongs | [optional] 
**IbgpHoldTimer** | Pointer to **int32** | The iBGP hold time in seconds | [optional] 
**Neighbors** | Pointer to [**[]GetNetworkApplianceVpnBgp200ResponseNeighborsInner**](GetNetworkApplianceVpnBgp200ResponseNeighborsInner.md) | List of eBGP neighbor configurations | [optional] 

## Methods

### NewGetNetworkApplianceVpnBgp200Response

`func NewGetNetworkApplianceVpnBgp200Response() *GetNetworkApplianceVpnBgp200Response`

NewGetNetworkApplianceVpnBgp200Response instantiates a new GetNetworkApplianceVpnBgp200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkApplianceVpnBgp200ResponseWithDefaults

`func NewGetNetworkApplianceVpnBgp200ResponseWithDefaults() *GetNetworkApplianceVpnBgp200Response`

NewGetNetworkApplianceVpnBgp200ResponseWithDefaults instantiates a new GetNetworkApplianceVpnBgp200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GetNetworkApplianceVpnBgp200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkApplianceVpnBgp200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkApplianceVpnBgp200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkApplianceVpnBgp200Response) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAsNumber

`func (o *GetNetworkApplianceVpnBgp200Response) GetAsNumber() int32`

GetAsNumber returns the AsNumber field if non-nil, zero value otherwise.

### GetAsNumberOk

`func (o *GetNetworkApplianceVpnBgp200Response) GetAsNumberOk() (*int32, bool)`

GetAsNumberOk returns a tuple with the AsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsNumber

`func (o *GetNetworkApplianceVpnBgp200Response) SetAsNumber(v int32)`

SetAsNumber sets AsNumber field to given value.

### HasAsNumber

`func (o *GetNetworkApplianceVpnBgp200Response) HasAsNumber() bool`

HasAsNumber returns a boolean if a field has been set.

### GetIbgpHoldTimer

`func (o *GetNetworkApplianceVpnBgp200Response) GetIbgpHoldTimer() int32`

GetIbgpHoldTimer returns the IbgpHoldTimer field if non-nil, zero value otherwise.

### GetIbgpHoldTimerOk

`func (o *GetNetworkApplianceVpnBgp200Response) GetIbgpHoldTimerOk() (*int32, bool)`

GetIbgpHoldTimerOk returns a tuple with the IbgpHoldTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbgpHoldTimer

`func (o *GetNetworkApplianceVpnBgp200Response) SetIbgpHoldTimer(v int32)`

SetIbgpHoldTimer sets IbgpHoldTimer field to given value.

### HasIbgpHoldTimer

`func (o *GetNetworkApplianceVpnBgp200Response) HasIbgpHoldTimer() bool`

HasIbgpHoldTimer returns a boolean if a field has been set.

### GetNeighbors

`func (o *GetNetworkApplianceVpnBgp200Response) GetNeighbors() []GetNetworkApplianceVpnBgp200ResponseNeighborsInner`

GetNeighbors returns the Neighbors field if non-nil, zero value otherwise.

### GetNeighborsOk

`func (o *GetNetworkApplianceVpnBgp200Response) GetNeighborsOk() (*[]GetNetworkApplianceVpnBgp200ResponseNeighborsInner, bool)`

GetNeighborsOk returns a tuple with the Neighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighbors

`func (o *GetNetworkApplianceVpnBgp200Response) SetNeighbors(v []GetNetworkApplianceVpnBgp200ResponseNeighborsInner)`

SetNeighbors sets Neighbors field to given value.

### HasNeighbors

`func (o *GetNetworkApplianceVpnBgp200Response) HasNeighbors() bool`

HasNeighbors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


