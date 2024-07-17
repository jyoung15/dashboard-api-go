# GetNetworkWirelessMeshStatuses200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | The serial number for the device. | [optional] 
**MeshRoute** | Pointer to **[]string** | List of device serials that make up the mesh. | [optional] 
**LatestMeshPerformance** | Pointer to [**GetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformance**](GetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformance.md) |  | [optional] 

## Methods

### NewGetNetworkWirelessMeshStatuses200ResponseInner

`func NewGetNetworkWirelessMeshStatuses200ResponseInner() *GetNetworkWirelessMeshStatuses200ResponseInner`

NewGetNetworkWirelessMeshStatuses200ResponseInner instantiates a new GetNetworkWirelessMeshStatuses200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWirelessMeshStatuses200ResponseInnerWithDefaults

`func NewGetNetworkWirelessMeshStatuses200ResponseInnerWithDefaults() *GetNetworkWirelessMeshStatuses200ResponseInner`

NewGetNetworkWirelessMeshStatuses200ResponseInnerWithDefaults instantiates a new GetNetworkWirelessMeshStatuses200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *GetNetworkWirelessMeshStatuses200ResponseInner) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetNetworkWirelessMeshStatuses200ResponseInner) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetNetworkWirelessMeshStatuses200ResponseInner) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetNetworkWirelessMeshStatuses200ResponseInner) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetMeshRoute

`func (o *GetNetworkWirelessMeshStatuses200ResponseInner) GetMeshRoute() []string`

GetMeshRoute returns the MeshRoute field if non-nil, zero value otherwise.

### GetMeshRouteOk

`func (o *GetNetworkWirelessMeshStatuses200ResponseInner) GetMeshRouteOk() (*[]string, bool)`

GetMeshRouteOk returns a tuple with the MeshRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshRoute

`func (o *GetNetworkWirelessMeshStatuses200ResponseInner) SetMeshRoute(v []string)`

SetMeshRoute sets MeshRoute field to given value.

### HasMeshRoute

`func (o *GetNetworkWirelessMeshStatuses200ResponseInner) HasMeshRoute() bool`

HasMeshRoute returns a boolean if a field has been set.

### GetLatestMeshPerformance

`func (o *GetNetworkWirelessMeshStatuses200ResponseInner) GetLatestMeshPerformance() GetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformance`

GetLatestMeshPerformance returns the LatestMeshPerformance field if non-nil, zero value otherwise.

### GetLatestMeshPerformanceOk

`func (o *GetNetworkWirelessMeshStatuses200ResponseInner) GetLatestMeshPerformanceOk() (*GetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformance, bool)`

GetLatestMeshPerformanceOk returns a tuple with the LatestMeshPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestMeshPerformance

`func (o *GetNetworkWirelessMeshStatuses200ResponseInner) SetLatestMeshPerformance(v GetNetworkWirelessMeshStatuses200ResponseInnerLatestMeshPerformance)`

SetLatestMeshPerformance sets LatestMeshPerformance field to given value.

### HasLatestMeshPerformance

`func (o *GetNetworkWirelessMeshStatuses200ResponseInner) HasLatestMeshPerformance() bool`

HasLatestMeshPerformance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


