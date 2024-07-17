# GetDevice200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the device | [optional] 
**Lat** | Pointer to **float32** | Latitude of the device | [optional] 
**Lng** | Pointer to **float32** | Longitude of the device | [optional] 
**Address** | Pointer to **string** | Physical address of the device | [optional] 
**Notes** | Pointer to **string** | Notes for the device, limited to 255 characters | [optional] 
**Tags** | Pointer to **[]string** | List of tags assigned to the device | [optional] 
**NetworkId** | Pointer to **string** | ID of the network the device belongs to | [optional] 
**Serial** | Pointer to **string** | Serial number of the device | [optional] 
**Model** | Pointer to **string** | Model of the device | [optional] 
**Mac** | Pointer to **string** | MAC address of the device | [optional] 
**LanIp** | Pointer to **string** | LAN IP address of the device | [optional] 
**Firmware** | Pointer to **string** | Firmware version of the device | [optional] 
**FloorPlanId** | Pointer to **string** | The floor plan to associate to this device. null disassociates the device from the floorplan. | [optional] 
**Details** | Pointer to [**[]GetDevice200ResponseDetailsInner**](GetDevice200ResponseDetailsInner.md) | Additional device information | [optional] 
**BeaconIdParams** | Pointer to [**GetDevice200ResponseBeaconIdParams**](GetDevice200ResponseBeaconIdParams.md) |  | [optional] 

## Methods

### NewGetDevice200Response

`func NewGetDevice200Response() *GetDevice200Response`

NewGetDevice200Response instantiates a new GetDevice200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDevice200ResponseWithDefaults

`func NewGetDevice200ResponseWithDefaults() *GetDevice200Response`

NewGetDevice200ResponseWithDefaults instantiates a new GetDevice200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetDevice200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDevice200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDevice200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDevice200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLat

`func (o *GetDevice200Response) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *GetDevice200Response) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *GetDevice200Response) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *GetDevice200Response) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLng

`func (o *GetDevice200Response) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *GetDevice200Response) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *GetDevice200Response) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *GetDevice200Response) HasLng() bool`

HasLng returns a boolean if a field has been set.

### GetAddress

`func (o *GetDevice200Response) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetDevice200Response) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetDevice200Response) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetDevice200Response) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetNotes

`func (o *GetDevice200Response) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *GetDevice200Response) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *GetDevice200Response) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *GetDevice200Response) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetTags

`func (o *GetDevice200Response) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetDevice200Response) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetDevice200Response) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetDevice200Response) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNetworkId

`func (o *GetDevice200Response) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *GetDevice200Response) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *GetDevice200Response) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *GetDevice200Response) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetSerial

`func (o *GetDevice200Response) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetDevice200Response) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetDevice200Response) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetDevice200Response) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetModel

`func (o *GetDevice200Response) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GetDevice200Response) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GetDevice200Response) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GetDevice200Response) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetMac

`func (o *GetDevice200Response) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetDevice200Response) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetDevice200Response) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetDevice200Response) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetLanIp

`func (o *GetDevice200Response) GetLanIp() string`

GetLanIp returns the LanIp field if non-nil, zero value otherwise.

### GetLanIpOk

`func (o *GetDevice200Response) GetLanIpOk() (*string, bool)`

GetLanIpOk returns a tuple with the LanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanIp

`func (o *GetDevice200Response) SetLanIp(v string)`

SetLanIp sets LanIp field to given value.

### HasLanIp

`func (o *GetDevice200Response) HasLanIp() bool`

HasLanIp returns a boolean if a field has been set.

### GetFirmware

`func (o *GetDevice200Response) GetFirmware() string`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *GetDevice200Response) GetFirmwareOk() (*string, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *GetDevice200Response) SetFirmware(v string)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *GetDevice200Response) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetFloorPlanId

`func (o *GetDevice200Response) GetFloorPlanId() string`

GetFloorPlanId returns the FloorPlanId field if non-nil, zero value otherwise.

### GetFloorPlanIdOk

`func (o *GetDevice200Response) GetFloorPlanIdOk() (*string, bool)`

GetFloorPlanIdOk returns a tuple with the FloorPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorPlanId

`func (o *GetDevice200Response) SetFloorPlanId(v string)`

SetFloorPlanId sets FloorPlanId field to given value.

### HasFloorPlanId

`func (o *GetDevice200Response) HasFloorPlanId() bool`

HasFloorPlanId returns a boolean if a field has been set.

### GetDetails

`func (o *GetDevice200Response) GetDetails() []GetDevice200ResponseDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetDevice200Response) GetDetailsOk() (*[]GetDevice200ResponseDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetDevice200Response) SetDetails(v []GetDevice200ResponseDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetDevice200Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetBeaconIdParams

`func (o *GetDevice200Response) GetBeaconIdParams() GetDevice200ResponseBeaconIdParams`

GetBeaconIdParams returns the BeaconIdParams field if non-nil, zero value otherwise.

### GetBeaconIdParamsOk

`func (o *GetDevice200Response) GetBeaconIdParamsOk() (*GetDevice200ResponseBeaconIdParams, bool)`

GetBeaconIdParamsOk returns a tuple with the BeaconIdParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconIdParams

`func (o *GetDevice200Response) SetBeaconIdParams(v GetDevice200ResponseBeaconIdParams)`

SetBeaconIdParams sets BeaconIdParams field to given value.

### HasBeaconIdParams

`func (o *GetDevice200Response) HasBeaconIdParams() bool`

HasBeaconIdParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


