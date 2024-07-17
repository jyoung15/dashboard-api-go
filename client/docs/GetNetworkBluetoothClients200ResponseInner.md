# GetNetworkBluetoothClients200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the client | [optional] 
**Mac** | Pointer to **string** | MAC address of the client | [optional] 
**NetworkId** | Pointer to **string** | Network ID | [optional] 
**Name** | Pointer to **string** | Name of the client | [optional] 
**DeviceName** | Pointer to **string** | Bluetooth device name | [optional] 
**Manufacturer** | Pointer to **string** | Name of the manufacturer | [optional] 
**LastSeen** | Pointer to **int32** | Epoch timestamp of the device&#39;s last appearance | [optional] 
**SeenByDeviceMac** | Pointer to **string** | Seen by device MAC | [optional] 
**InSightAlert** | Pointer to **bool** | Device in sight alert | [optional] 
**OutOfSightAlert** | Pointer to **bool** | Device out of sight alert | [optional] 
**Tags** | Pointer to **[]string** | A list of tags applied to the device | [optional] 

## Methods

### NewGetNetworkBluetoothClients200ResponseInner

`func NewGetNetworkBluetoothClients200ResponseInner() *GetNetworkBluetoothClients200ResponseInner`

NewGetNetworkBluetoothClients200ResponseInner instantiates a new GetNetworkBluetoothClients200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkBluetoothClients200ResponseInnerWithDefaults

`func NewGetNetworkBluetoothClients200ResponseInnerWithDefaults() *GetNetworkBluetoothClients200ResponseInner`

NewGetNetworkBluetoothClients200ResponseInnerWithDefaults instantiates a new GetNetworkBluetoothClients200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkBluetoothClients200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkBluetoothClients200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkBluetoothClients200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkBluetoothClients200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMac

`func (o *GetNetworkBluetoothClients200ResponseInner) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetNetworkBluetoothClients200ResponseInner) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetNetworkBluetoothClients200ResponseInner) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetNetworkBluetoothClients200ResponseInner) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetNetworkId

`func (o *GetNetworkBluetoothClients200ResponseInner) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *GetNetworkBluetoothClients200ResponseInner) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *GetNetworkBluetoothClients200ResponseInner) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *GetNetworkBluetoothClients200ResponseInner) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkBluetoothClients200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkBluetoothClients200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkBluetoothClients200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkBluetoothClients200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDeviceName

`func (o *GetNetworkBluetoothClients200ResponseInner) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *GetNetworkBluetoothClients200ResponseInner) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *GetNetworkBluetoothClients200ResponseInner) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *GetNetworkBluetoothClients200ResponseInner) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetManufacturer

`func (o *GetNetworkBluetoothClients200ResponseInner) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *GetNetworkBluetoothClients200ResponseInner) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *GetNetworkBluetoothClients200ResponseInner) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *GetNetworkBluetoothClients200ResponseInner) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetLastSeen

`func (o *GetNetworkBluetoothClients200ResponseInner) GetLastSeen() int32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *GetNetworkBluetoothClients200ResponseInner) GetLastSeenOk() (*int32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *GetNetworkBluetoothClients200ResponseInner) SetLastSeen(v int32)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *GetNetworkBluetoothClients200ResponseInner) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetSeenByDeviceMac

`func (o *GetNetworkBluetoothClients200ResponseInner) GetSeenByDeviceMac() string`

GetSeenByDeviceMac returns the SeenByDeviceMac field if non-nil, zero value otherwise.

### GetSeenByDeviceMacOk

`func (o *GetNetworkBluetoothClients200ResponseInner) GetSeenByDeviceMacOk() (*string, bool)`

GetSeenByDeviceMacOk returns a tuple with the SeenByDeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeenByDeviceMac

`func (o *GetNetworkBluetoothClients200ResponseInner) SetSeenByDeviceMac(v string)`

SetSeenByDeviceMac sets SeenByDeviceMac field to given value.

### HasSeenByDeviceMac

`func (o *GetNetworkBluetoothClients200ResponseInner) HasSeenByDeviceMac() bool`

HasSeenByDeviceMac returns a boolean if a field has been set.

### GetInSightAlert

`func (o *GetNetworkBluetoothClients200ResponseInner) GetInSightAlert() bool`

GetInSightAlert returns the InSightAlert field if non-nil, zero value otherwise.

### GetInSightAlertOk

`func (o *GetNetworkBluetoothClients200ResponseInner) GetInSightAlertOk() (*bool, bool)`

GetInSightAlertOk returns a tuple with the InSightAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInSightAlert

`func (o *GetNetworkBluetoothClients200ResponseInner) SetInSightAlert(v bool)`

SetInSightAlert sets InSightAlert field to given value.

### HasInSightAlert

`func (o *GetNetworkBluetoothClients200ResponseInner) HasInSightAlert() bool`

HasInSightAlert returns a boolean if a field has been set.

### GetOutOfSightAlert

`func (o *GetNetworkBluetoothClients200ResponseInner) GetOutOfSightAlert() bool`

GetOutOfSightAlert returns the OutOfSightAlert field if non-nil, zero value otherwise.

### GetOutOfSightAlertOk

`func (o *GetNetworkBluetoothClients200ResponseInner) GetOutOfSightAlertOk() (*bool, bool)`

GetOutOfSightAlertOk returns a tuple with the OutOfSightAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfSightAlert

`func (o *GetNetworkBluetoothClients200ResponseInner) SetOutOfSightAlert(v bool)`

SetOutOfSightAlert sets OutOfSightAlert field to given value.

### HasOutOfSightAlert

`func (o *GetNetworkBluetoothClients200ResponseInner) HasOutOfSightAlert() bool`

HasOutOfSightAlert returns a boolean if a field has been set.

### GetTags

`func (o *GetNetworkBluetoothClients200ResponseInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetNetworkBluetoothClients200ResponseInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetNetworkBluetoothClients200ResponseInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetNetworkBluetoothClients200ResponseInner) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


