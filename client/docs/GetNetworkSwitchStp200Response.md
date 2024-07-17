# GetNetworkSwitchStp200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RstpEnabled** | Pointer to **bool** | The spanning tree protocol status in network | [optional] 
**StpBridgePriority** | Pointer to [**[]GetNetworkSwitchStp200ResponseStpBridgePriorityInner**](GetNetworkSwitchStp200ResponseStpBridgePriorityInner.md) | STP bridge priority for switches/stacks or switch templates. An empty array will clear the STP bridge priority settings. | [optional] 

## Methods

### NewGetNetworkSwitchStp200Response

`func NewGetNetworkSwitchStp200Response() *GetNetworkSwitchStp200Response`

NewGetNetworkSwitchStp200Response instantiates a new GetNetworkSwitchStp200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSwitchStp200ResponseWithDefaults

`func NewGetNetworkSwitchStp200ResponseWithDefaults() *GetNetworkSwitchStp200Response`

NewGetNetworkSwitchStp200ResponseWithDefaults instantiates a new GetNetworkSwitchStp200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRstpEnabled

`func (o *GetNetworkSwitchStp200Response) GetRstpEnabled() bool`

GetRstpEnabled returns the RstpEnabled field if non-nil, zero value otherwise.

### GetRstpEnabledOk

`func (o *GetNetworkSwitchStp200Response) GetRstpEnabledOk() (*bool, bool)`

GetRstpEnabledOk returns a tuple with the RstpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRstpEnabled

`func (o *GetNetworkSwitchStp200Response) SetRstpEnabled(v bool)`

SetRstpEnabled sets RstpEnabled field to given value.

### HasRstpEnabled

`func (o *GetNetworkSwitchStp200Response) HasRstpEnabled() bool`

HasRstpEnabled returns a boolean if a field has been set.

### GetStpBridgePriority

`func (o *GetNetworkSwitchStp200Response) GetStpBridgePriority() []GetNetworkSwitchStp200ResponseStpBridgePriorityInner`

GetStpBridgePriority returns the StpBridgePriority field if non-nil, zero value otherwise.

### GetStpBridgePriorityOk

`func (o *GetNetworkSwitchStp200Response) GetStpBridgePriorityOk() (*[]GetNetworkSwitchStp200ResponseStpBridgePriorityInner, bool)`

GetStpBridgePriorityOk returns a tuple with the StpBridgePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpBridgePriority

`func (o *GetNetworkSwitchStp200Response) SetStpBridgePriority(v []GetNetworkSwitchStp200ResponseStpBridgePriorityInner)`

SetStpBridgePriority sets StpBridgePriority field to given value.

### HasStpBridgePriority

`func (o *GetNetworkSwitchStp200Response) HasStpBridgePriority() bool`

HasStpBridgePriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


