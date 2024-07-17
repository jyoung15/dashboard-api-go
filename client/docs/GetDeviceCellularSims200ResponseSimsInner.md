# GetDeviceCellularSims200ResponseSimsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slot** | Pointer to **string** | SIM slot being configured. Must be &#39;sim1&#39; on single-sim devices. Use &#39;sim3&#39; for eSIM. | [optional] 
**IsPrimary** | Pointer to **bool** | If true, this SIM is activated on platform bootup. It must be true on single-SIM devices and is a required field for dual-SIM MGs unless it is being configured using &#39;simOrdering&#39;. | [optional] [default to false]
**Apns** | Pointer to [**[]GetDeviceCellularSims200ResponseSimsInnerApnsInner**](GetDeviceCellularSims200ResponseSimsInnerApnsInner.md) | APN configurations. If empty, the default APN will be used. | [optional] [default to []]

## Methods

### NewGetDeviceCellularSims200ResponseSimsInner

`func NewGetDeviceCellularSims200ResponseSimsInner() *GetDeviceCellularSims200ResponseSimsInner`

NewGetDeviceCellularSims200ResponseSimsInner instantiates a new GetDeviceCellularSims200ResponseSimsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceCellularSims200ResponseSimsInnerWithDefaults

`func NewGetDeviceCellularSims200ResponseSimsInnerWithDefaults() *GetDeviceCellularSims200ResponseSimsInner`

NewGetDeviceCellularSims200ResponseSimsInnerWithDefaults instantiates a new GetDeviceCellularSims200ResponseSimsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlot

`func (o *GetDeviceCellularSims200ResponseSimsInner) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *GetDeviceCellularSims200ResponseSimsInner) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *GetDeviceCellularSims200ResponseSimsInner) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *GetDeviceCellularSims200ResponseSimsInner) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetIsPrimary

`func (o *GetDeviceCellularSims200ResponseSimsInner) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *GetDeviceCellularSims200ResponseSimsInner) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *GetDeviceCellularSims200ResponseSimsInner) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *GetDeviceCellularSims200ResponseSimsInner) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetApns

`func (o *GetDeviceCellularSims200ResponseSimsInner) GetApns() []GetDeviceCellularSims200ResponseSimsInnerApnsInner`

GetApns returns the Apns field if non-nil, zero value otherwise.

### GetApnsOk

`func (o *GetDeviceCellularSims200ResponseSimsInner) GetApnsOk() (*[]GetDeviceCellularSims200ResponseSimsInnerApnsInner, bool)`

GetApnsOk returns a tuple with the Apns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApns

`func (o *GetDeviceCellularSims200ResponseSimsInner) SetApns(v []GetDeviceCellularSims200ResponseSimsInnerApnsInner)`

SetApns sets Apns field to given value.

### HasApns

`func (o *GetDeviceCellularSims200ResponseSimsInner) HasApns() bool`

HasApns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


