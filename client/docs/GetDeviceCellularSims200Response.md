# GetDeviceCellularSims200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sims** | Pointer to [**[]GetDeviceCellularSims200ResponseSimsInner**](GetDeviceCellularSims200ResponseSimsInner.md) | List of SIMs. If a SIM was previously configured and not specified in this request, it will remain unchanged. | [optional] 
**SimOrdering** | Pointer to **[]string** | Specifies the ordering of all SIMs for an MG: primary, secondary, and not-in-use (when applicable). It&#39;s required for devices with 3 or more SIMs and can be used in place of &#39;isPrimary&#39; for dual-SIM devices. To indicate eSIM, use &#39;sim3&#39;. Sim failover will occur only between primary and secondary sim slots. | [optional] 
**SimFailover** | Pointer to [**GetDeviceCellularSims200ResponseSimFailover**](GetDeviceCellularSims200ResponseSimFailover.md) |  | [optional] 

## Methods

### NewGetDeviceCellularSims200Response

`func NewGetDeviceCellularSims200Response() *GetDeviceCellularSims200Response`

NewGetDeviceCellularSims200Response instantiates a new GetDeviceCellularSims200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceCellularSims200ResponseWithDefaults

`func NewGetDeviceCellularSims200ResponseWithDefaults() *GetDeviceCellularSims200Response`

NewGetDeviceCellularSims200ResponseWithDefaults instantiates a new GetDeviceCellularSims200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSims

`func (o *GetDeviceCellularSims200Response) GetSims() []GetDeviceCellularSims200ResponseSimsInner`

GetSims returns the Sims field if non-nil, zero value otherwise.

### GetSimsOk

`func (o *GetDeviceCellularSims200Response) GetSimsOk() (*[]GetDeviceCellularSims200ResponseSimsInner, bool)`

GetSimsOk returns a tuple with the Sims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSims

`func (o *GetDeviceCellularSims200Response) SetSims(v []GetDeviceCellularSims200ResponseSimsInner)`

SetSims sets Sims field to given value.

### HasSims

`func (o *GetDeviceCellularSims200Response) HasSims() bool`

HasSims returns a boolean if a field has been set.

### GetSimOrdering

`func (o *GetDeviceCellularSims200Response) GetSimOrdering() []string`

GetSimOrdering returns the SimOrdering field if non-nil, zero value otherwise.

### GetSimOrderingOk

`func (o *GetDeviceCellularSims200Response) GetSimOrderingOk() (*[]string, bool)`

GetSimOrderingOk returns a tuple with the SimOrdering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimOrdering

`func (o *GetDeviceCellularSims200Response) SetSimOrdering(v []string)`

SetSimOrdering sets SimOrdering field to given value.

### HasSimOrdering

`func (o *GetDeviceCellularSims200Response) HasSimOrdering() bool`

HasSimOrdering returns a boolean if a field has been set.

### GetSimFailover

`func (o *GetDeviceCellularSims200Response) GetSimFailover() GetDeviceCellularSims200ResponseSimFailover`

GetSimFailover returns the SimFailover field if non-nil, zero value otherwise.

### GetSimFailoverOk

`func (o *GetDeviceCellularSims200Response) GetSimFailoverOk() (*GetDeviceCellularSims200ResponseSimFailover, bool)`

GetSimFailoverOk returns a tuple with the SimFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimFailover

`func (o *GetDeviceCellularSims200Response) SetSimFailover(v GetDeviceCellularSims200ResponseSimFailover)`

SetSimFailover sets SimFailover field to given value.

### HasSimFailover

`func (o *GetDeviceCellularSims200Response) HasSimFailover() bool`

HasSimFailover returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


