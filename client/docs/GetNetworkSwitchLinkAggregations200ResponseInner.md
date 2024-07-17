# GetNetworkSwitchLinkAggregations200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID for the link aggregation. | [optional] 
**SwitchPorts** | Pointer to [**[]GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner**](GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner.md) | The ID for the link aggregation. | [optional] 

## Methods

### NewGetNetworkSwitchLinkAggregations200ResponseInner

`func NewGetNetworkSwitchLinkAggregations200ResponseInner() *GetNetworkSwitchLinkAggregations200ResponseInner`

NewGetNetworkSwitchLinkAggregations200ResponseInner instantiates a new GetNetworkSwitchLinkAggregations200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSwitchLinkAggregations200ResponseInnerWithDefaults

`func NewGetNetworkSwitchLinkAggregations200ResponseInnerWithDefaults() *GetNetworkSwitchLinkAggregations200ResponseInner`

NewGetNetworkSwitchLinkAggregations200ResponseInnerWithDefaults instantiates a new GetNetworkSwitchLinkAggregations200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkSwitchLinkAggregations200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkSwitchLinkAggregations200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkSwitchLinkAggregations200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkSwitchLinkAggregations200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSwitchPorts

`func (o *GetNetworkSwitchLinkAggregations200ResponseInner) GetSwitchPorts() []GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner`

GetSwitchPorts returns the SwitchPorts field if non-nil, zero value otherwise.

### GetSwitchPortsOk

`func (o *GetNetworkSwitchLinkAggregations200ResponseInner) GetSwitchPortsOk() (*[]GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner, bool)`

GetSwitchPortsOk returns a tuple with the SwitchPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPorts

`func (o *GetNetworkSwitchLinkAggregations200ResponseInner) SetSwitchPorts(v []GetNetworkSwitchLinkAggregations200ResponseInnerSwitchPortsInner)`

SetSwitchPorts sets SwitchPorts field to given value.

### HasSwitchPorts

`func (o *GetNetworkSwitchLinkAggregations200ResponseInner) HasSwitchPorts() bool`

HasSwitchPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


