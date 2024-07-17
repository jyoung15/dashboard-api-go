# GetDeviceSwitchPortsStatusesPackets200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortId** | Pointer to **string** | The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module. | [optional] 
**Packets** | Pointer to [**[]GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner**](GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner.md) | The packet counts on the switch. | [optional] 

## Methods

### NewGetDeviceSwitchPortsStatusesPackets200ResponseInner

`func NewGetDeviceSwitchPortsStatusesPackets200ResponseInner() *GetDeviceSwitchPortsStatusesPackets200ResponseInner`

NewGetDeviceSwitchPortsStatusesPackets200ResponseInner instantiates a new GetDeviceSwitchPortsStatusesPackets200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceSwitchPortsStatusesPackets200ResponseInnerWithDefaults

`func NewGetDeviceSwitchPortsStatusesPackets200ResponseInnerWithDefaults() *GetDeviceSwitchPortsStatusesPackets200ResponseInner`

NewGetDeviceSwitchPortsStatusesPackets200ResponseInnerWithDefaults instantiates a new GetDeviceSwitchPortsStatusesPackets200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInner) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInner) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInner) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInner) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetPackets

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInner) GetPackets() []GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner`

GetPackets returns the Packets field if non-nil, zero value otherwise.

### GetPacketsOk

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInner) GetPacketsOk() (*[]GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner, bool)`

GetPacketsOk returns a tuple with the Packets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackets

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInner) SetPackets(v []GetDeviceSwitchPortsStatusesPackets200ResponseInnerPacketsInner)`

SetPackets sets Packets field to given value.

### HasPackets

`func (o *GetDeviceSwitchPortsStatusesPackets200ResponseInner) HasPackets() bool`

HasPackets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


