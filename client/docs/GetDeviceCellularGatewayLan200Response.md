# GetDeviceCellularGatewayLan200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceName** | Pointer to **string** | Name of the MG. | [optional] 
**DeviceLanIp** | Pointer to **string** | Lan IP of the MG | [optional] 
**DeviceSubnet** | Pointer to **string** | Subnet configuration of the MG. | [optional] 
**FixedIpAssignments** | Pointer to [**[]GetDeviceCellularGatewayLan200ResponseFixedIpAssignmentsInner**](GetDeviceCellularGatewayLan200ResponseFixedIpAssignmentsInner.md) | list of all fixed IP assignments for a single MG | [optional] 
**ReservedIpRanges** | Pointer to [**[]GetDeviceCellularGatewayLan200ResponseReservedIpRangesInner**](GetDeviceCellularGatewayLan200ResponseReservedIpRangesInner.md) | list of all reserved IP ranges for a single MG | [optional] 

## Methods

### NewGetDeviceCellularGatewayLan200Response

`func NewGetDeviceCellularGatewayLan200Response() *GetDeviceCellularGatewayLan200Response`

NewGetDeviceCellularGatewayLan200Response instantiates a new GetDeviceCellularGatewayLan200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceCellularGatewayLan200ResponseWithDefaults

`func NewGetDeviceCellularGatewayLan200ResponseWithDefaults() *GetDeviceCellularGatewayLan200Response`

NewGetDeviceCellularGatewayLan200ResponseWithDefaults instantiates a new GetDeviceCellularGatewayLan200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceName

`func (o *GetDeviceCellularGatewayLan200Response) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *GetDeviceCellularGatewayLan200Response) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *GetDeviceCellularGatewayLan200Response) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *GetDeviceCellularGatewayLan200Response) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceLanIp

`func (o *GetDeviceCellularGatewayLan200Response) GetDeviceLanIp() string`

GetDeviceLanIp returns the DeviceLanIp field if non-nil, zero value otherwise.

### GetDeviceLanIpOk

`func (o *GetDeviceCellularGatewayLan200Response) GetDeviceLanIpOk() (*string, bool)`

GetDeviceLanIpOk returns a tuple with the DeviceLanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceLanIp

`func (o *GetDeviceCellularGatewayLan200Response) SetDeviceLanIp(v string)`

SetDeviceLanIp sets DeviceLanIp field to given value.

### HasDeviceLanIp

`func (o *GetDeviceCellularGatewayLan200Response) HasDeviceLanIp() bool`

HasDeviceLanIp returns a boolean if a field has been set.

### GetDeviceSubnet

`func (o *GetDeviceCellularGatewayLan200Response) GetDeviceSubnet() string`

GetDeviceSubnet returns the DeviceSubnet field if non-nil, zero value otherwise.

### GetDeviceSubnetOk

`func (o *GetDeviceCellularGatewayLan200Response) GetDeviceSubnetOk() (*string, bool)`

GetDeviceSubnetOk returns a tuple with the DeviceSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSubnet

`func (o *GetDeviceCellularGatewayLan200Response) SetDeviceSubnet(v string)`

SetDeviceSubnet sets DeviceSubnet field to given value.

### HasDeviceSubnet

`func (o *GetDeviceCellularGatewayLan200Response) HasDeviceSubnet() bool`

HasDeviceSubnet returns a boolean if a field has been set.

### GetFixedIpAssignments

`func (o *GetDeviceCellularGatewayLan200Response) GetFixedIpAssignments() []GetDeviceCellularGatewayLan200ResponseFixedIpAssignmentsInner`

GetFixedIpAssignments returns the FixedIpAssignments field if non-nil, zero value otherwise.

### GetFixedIpAssignmentsOk

`func (o *GetDeviceCellularGatewayLan200Response) GetFixedIpAssignmentsOk() (*[]GetDeviceCellularGatewayLan200ResponseFixedIpAssignmentsInner, bool)`

GetFixedIpAssignmentsOk returns a tuple with the FixedIpAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIpAssignments

`func (o *GetDeviceCellularGatewayLan200Response) SetFixedIpAssignments(v []GetDeviceCellularGatewayLan200ResponseFixedIpAssignmentsInner)`

SetFixedIpAssignments sets FixedIpAssignments field to given value.

### HasFixedIpAssignments

`func (o *GetDeviceCellularGatewayLan200Response) HasFixedIpAssignments() bool`

HasFixedIpAssignments returns a boolean if a field has been set.

### GetReservedIpRanges

`func (o *GetDeviceCellularGatewayLan200Response) GetReservedIpRanges() []GetDeviceCellularGatewayLan200ResponseReservedIpRangesInner`

GetReservedIpRanges returns the ReservedIpRanges field if non-nil, zero value otherwise.

### GetReservedIpRangesOk

`func (o *GetDeviceCellularGatewayLan200Response) GetReservedIpRangesOk() (*[]GetDeviceCellularGatewayLan200ResponseReservedIpRangesInner, bool)`

GetReservedIpRangesOk returns a tuple with the ReservedIpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedIpRanges

`func (o *GetDeviceCellularGatewayLan200Response) SetReservedIpRanges(v []GetDeviceCellularGatewayLan200ResponseReservedIpRangesInner)`

SetReservedIpRanges sets ReservedIpRanges field to given value.

### HasReservedIpRanges

`func (o *GetDeviceCellularGatewayLan200Response) HasReservedIpRanges() bool`

HasReservedIpRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


