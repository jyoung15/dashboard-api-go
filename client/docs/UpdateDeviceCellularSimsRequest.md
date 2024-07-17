# UpdateDeviceCellularSimsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sims** | Pointer to [**[]UpdateDeviceCellularSimsRequestSimsInner**](UpdateDeviceCellularSimsRequestSimsInner.md) | List of SIMs. If a SIM was previously configured and not specified in this request, it will remain unchanged. | [optional] 
**SimOrdering** | Pointer to **[]string** | Specifies the ordering of all SIMs for an MG: primary, secondary, and not-in-use (when applicable). It&#39;s required for devices with 3 or more SIMs and can be used in place of &#39;isPrimary&#39; for dual-SIM devices. To indicate eSIM, use &#39;sim3&#39;. Sim failover will occur only between primary and secondary sim slots. | [optional] 
**SimFailover** | Pointer to [**UpdateDeviceCellularSimsRequestSimFailover**](UpdateDeviceCellularSimsRequestSimFailover.md) |  | [optional] 

## Methods

### NewUpdateDeviceCellularSimsRequest

`func NewUpdateDeviceCellularSimsRequest() *UpdateDeviceCellularSimsRequest`

NewUpdateDeviceCellularSimsRequest instantiates a new UpdateDeviceCellularSimsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceCellularSimsRequestWithDefaults

`func NewUpdateDeviceCellularSimsRequestWithDefaults() *UpdateDeviceCellularSimsRequest`

NewUpdateDeviceCellularSimsRequestWithDefaults instantiates a new UpdateDeviceCellularSimsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSims

`func (o *UpdateDeviceCellularSimsRequest) GetSims() []UpdateDeviceCellularSimsRequestSimsInner`

GetSims returns the Sims field if non-nil, zero value otherwise.

### GetSimsOk

`func (o *UpdateDeviceCellularSimsRequest) GetSimsOk() (*[]UpdateDeviceCellularSimsRequestSimsInner, bool)`

GetSimsOk returns a tuple with the Sims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSims

`func (o *UpdateDeviceCellularSimsRequest) SetSims(v []UpdateDeviceCellularSimsRequestSimsInner)`

SetSims sets Sims field to given value.

### HasSims

`func (o *UpdateDeviceCellularSimsRequest) HasSims() bool`

HasSims returns a boolean if a field has been set.

### GetSimOrdering

`func (o *UpdateDeviceCellularSimsRequest) GetSimOrdering() []string`

GetSimOrdering returns the SimOrdering field if non-nil, zero value otherwise.

### GetSimOrderingOk

`func (o *UpdateDeviceCellularSimsRequest) GetSimOrderingOk() (*[]string, bool)`

GetSimOrderingOk returns a tuple with the SimOrdering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimOrdering

`func (o *UpdateDeviceCellularSimsRequest) SetSimOrdering(v []string)`

SetSimOrdering sets SimOrdering field to given value.

### HasSimOrdering

`func (o *UpdateDeviceCellularSimsRequest) HasSimOrdering() bool`

HasSimOrdering returns a boolean if a field has been set.

### GetSimFailover

`func (o *UpdateDeviceCellularSimsRequest) GetSimFailover() UpdateDeviceCellularSimsRequestSimFailover`

GetSimFailover returns the SimFailover field if non-nil, zero value otherwise.

### GetSimFailoverOk

`func (o *UpdateDeviceCellularSimsRequest) GetSimFailoverOk() (*UpdateDeviceCellularSimsRequestSimFailover, bool)`

GetSimFailoverOk returns a tuple with the SimFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimFailover

`func (o *UpdateDeviceCellularSimsRequest) SetSimFailover(v UpdateDeviceCellularSimsRequestSimFailover)`

SetSimFailover sets SimFailover field to given value.

### HasSimFailover

`func (o *UpdateDeviceCellularSimsRequest) HasSimFailover() bool`

HasSimFailover returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


