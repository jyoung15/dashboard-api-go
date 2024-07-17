# VmxNetworkDevicesClaim200Response

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
**Imei** | Pointer to **string** | IMEI of the device, if applicable | [optional] 
**Mac** | Pointer to **string** | MAC address of the device | [optional] 
**LanIp** | Pointer to **string** | LAN IP address of the device | [optional] 
**Firmware** | Pointer to **string** | Firmware version of the device | [optional] 
**ProductType** | Pointer to **string** | Product type of the device | [optional] 
**Details** | Pointer to [**[]GetDevice200ResponseDetailsInner**](GetDevice200ResponseDetailsInner.md) | Additional device information | [optional] 

## Methods

### NewVmxNetworkDevicesClaim200Response

`func NewVmxNetworkDevicesClaim200Response() *VmxNetworkDevicesClaim200Response`

NewVmxNetworkDevicesClaim200Response instantiates a new VmxNetworkDevicesClaim200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmxNetworkDevicesClaim200ResponseWithDefaults

`func NewVmxNetworkDevicesClaim200ResponseWithDefaults() *VmxNetworkDevicesClaim200Response`

NewVmxNetworkDevicesClaim200ResponseWithDefaults instantiates a new VmxNetworkDevicesClaim200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VmxNetworkDevicesClaim200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmxNetworkDevicesClaim200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmxNetworkDevicesClaim200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VmxNetworkDevicesClaim200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLat

`func (o *VmxNetworkDevicesClaim200Response) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *VmxNetworkDevicesClaim200Response) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *VmxNetworkDevicesClaim200Response) SetLat(v float32)`

SetLat sets Lat field to given value.

### HasLat

`func (o *VmxNetworkDevicesClaim200Response) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLng

`func (o *VmxNetworkDevicesClaim200Response) GetLng() float32`

GetLng returns the Lng field if non-nil, zero value otherwise.

### GetLngOk

`func (o *VmxNetworkDevicesClaim200Response) GetLngOk() (*float32, bool)`

GetLngOk returns a tuple with the Lng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLng

`func (o *VmxNetworkDevicesClaim200Response) SetLng(v float32)`

SetLng sets Lng field to given value.

### HasLng

`func (o *VmxNetworkDevicesClaim200Response) HasLng() bool`

HasLng returns a boolean if a field has been set.

### GetAddress

`func (o *VmxNetworkDevicesClaim200Response) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VmxNetworkDevicesClaim200Response) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VmxNetworkDevicesClaim200Response) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VmxNetworkDevicesClaim200Response) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetNotes

`func (o *VmxNetworkDevicesClaim200Response) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *VmxNetworkDevicesClaim200Response) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *VmxNetworkDevicesClaim200Response) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *VmxNetworkDevicesClaim200Response) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetTags

`func (o *VmxNetworkDevicesClaim200Response) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VmxNetworkDevicesClaim200Response) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VmxNetworkDevicesClaim200Response) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VmxNetworkDevicesClaim200Response) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNetworkId

`func (o *VmxNetworkDevicesClaim200Response) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *VmxNetworkDevicesClaim200Response) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *VmxNetworkDevicesClaim200Response) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *VmxNetworkDevicesClaim200Response) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetSerial

`func (o *VmxNetworkDevicesClaim200Response) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *VmxNetworkDevicesClaim200Response) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *VmxNetworkDevicesClaim200Response) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *VmxNetworkDevicesClaim200Response) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetModel

`func (o *VmxNetworkDevicesClaim200Response) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *VmxNetworkDevicesClaim200Response) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *VmxNetworkDevicesClaim200Response) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *VmxNetworkDevicesClaim200Response) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetImei

`func (o *VmxNetworkDevicesClaim200Response) GetImei() string`

GetImei returns the Imei field if non-nil, zero value otherwise.

### GetImeiOk

`func (o *VmxNetworkDevicesClaim200Response) GetImeiOk() (*string, bool)`

GetImeiOk returns a tuple with the Imei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImei

`func (o *VmxNetworkDevicesClaim200Response) SetImei(v string)`

SetImei sets Imei field to given value.

### HasImei

`func (o *VmxNetworkDevicesClaim200Response) HasImei() bool`

HasImei returns a boolean if a field has been set.

### GetMac

`func (o *VmxNetworkDevicesClaim200Response) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *VmxNetworkDevicesClaim200Response) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *VmxNetworkDevicesClaim200Response) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *VmxNetworkDevicesClaim200Response) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetLanIp

`func (o *VmxNetworkDevicesClaim200Response) GetLanIp() string`

GetLanIp returns the LanIp field if non-nil, zero value otherwise.

### GetLanIpOk

`func (o *VmxNetworkDevicesClaim200Response) GetLanIpOk() (*string, bool)`

GetLanIpOk returns a tuple with the LanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanIp

`func (o *VmxNetworkDevicesClaim200Response) SetLanIp(v string)`

SetLanIp sets LanIp field to given value.

### HasLanIp

`func (o *VmxNetworkDevicesClaim200Response) HasLanIp() bool`

HasLanIp returns a boolean if a field has been set.

### GetFirmware

`func (o *VmxNetworkDevicesClaim200Response) GetFirmware() string`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *VmxNetworkDevicesClaim200Response) GetFirmwareOk() (*string, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *VmxNetworkDevicesClaim200Response) SetFirmware(v string)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *VmxNetworkDevicesClaim200Response) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetProductType

`func (o *VmxNetworkDevicesClaim200Response) GetProductType() string`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *VmxNetworkDevicesClaim200Response) GetProductTypeOk() (*string, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *VmxNetworkDevicesClaim200Response) SetProductType(v string)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *VmxNetworkDevicesClaim200Response) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetDetails

`func (o *VmxNetworkDevicesClaim200Response) GetDetails() []GetDevice200ResponseDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *VmxNetworkDevicesClaim200Response) GetDetailsOk() (*[]GetDevice200ResponseDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *VmxNetworkDevicesClaim200Response) SetDetails(v []GetDevice200ResponseDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *VmxNetworkDevicesClaim200Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


