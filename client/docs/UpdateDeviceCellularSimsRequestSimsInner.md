# UpdateDeviceCellularSimsRequestSimsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slot** | Pointer to **string** | SIM slot being configured. Must be &#39;sim1&#39; on single-sim devices. Use &#39;sim3&#39; for eSIM. | [optional] 
**IsPrimary** | Pointer to **bool** | If true, this SIM is activated on platform bootup. It must be true on single-SIM devices and is a required field for dual-SIM MGs unless it is being configured using &#39;simOrdering&#39;. | [optional] [default to false]
**Apns** | Pointer to [**[]GetDeviceCellularSims200ResponseSimsInnerApnsInner**](GetDeviceCellularSims200ResponseSimsInnerApnsInner.md) | APN configurations. If empty, the default APN will be used. | [optional] [default to []]
**SimOrder** | Pointer to **int32** | Priority of SIM slot being configured. Use a value between 1 and total number of SIMs available. The value must be unique for each SIM. | [optional] 

## Methods

### NewUpdateDeviceCellularSimsRequestSimsInner

`func NewUpdateDeviceCellularSimsRequestSimsInner() *UpdateDeviceCellularSimsRequestSimsInner`

NewUpdateDeviceCellularSimsRequestSimsInner instantiates a new UpdateDeviceCellularSimsRequestSimsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceCellularSimsRequestSimsInnerWithDefaults

`func NewUpdateDeviceCellularSimsRequestSimsInnerWithDefaults() *UpdateDeviceCellularSimsRequestSimsInner`

NewUpdateDeviceCellularSimsRequestSimsInnerWithDefaults instantiates a new UpdateDeviceCellularSimsRequestSimsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlot

`func (o *UpdateDeviceCellularSimsRequestSimsInner) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *UpdateDeviceCellularSimsRequestSimsInner) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *UpdateDeviceCellularSimsRequestSimsInner) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *UpdateDeviceCellularSimsRequestSimsInner) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetIsPrimary

`func (o *UpdateDeviceCellularSimsRequestSimsInner) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *UpdateDeviceCellularSimsRequestSimsInner) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *UpdateDeviceCellularSimsRequestSimsInner) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *UpdateDeviceCellularSimsRequestSimsInner) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetApns

`func (o *UpdateDeviceCellularSimsRequestSimsInner) GetApns() []GetDeviceCellularSims200ResponseSimsInnerApnsInner`

GetApns returns the Apns field if non-nil, zero value otherwise.

### GetApnsOk

`func (o *UpdateDeviceCellularSimsRequestSimsInner) GetApnsOk() (*[]GetDeviceCellularSims200ResponseSimsInnerApnsInner, bool)`

GetApnsOk returns a tuple with the Apns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApns

`func (o *UpdateDeviceCellularSimsRequestSimsInner) SetApns(v []GetDeviceCellularSims200ResponseSimsInnerApnsInner)`

SetApns sets Apns field to given value.

### HasApns

`func (o *UpdateDeviceCellularSimsRequestSimsInner) HasApns() bool`

HasApns returns a boolean if a field has been set.

### GetSimOrder

`func (o *UpdateDeviceCellularSimsRequestSimsInner) GetSimOrder() int32`

GetSimOrder returns the SimOrder field if non-nil, zero value otherwise.

### GetSimOrderOk

`func (o *UpdateDeviceCellularSimsRequestSimsInner) GetSimOrderOk() (*int32, bool)`

GetSimOrderOk returns a tuple with the SimOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimOrder

`func (o *UpdateDeviceCellularSimsRequestSimsInner) SetSimOrder(v int32)`

SetSimOrder sets SimOrder field to given value.

### HasSimOrder

`func (o *UpdateDeviceCellularSimsRequestSimsInner) HasSimOrder() bool`

HasSimOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


