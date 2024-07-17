# GetDeviceWirelessElectronicShelfLabel200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApEslId** | Pointer to **int32** | An identifier for the device used by the ESL system | [optional] 
**Serial** | Pointer to **string** | The serial number of the device | [optional] 
**Channel** | Pointer to **string** | Desired ESL channel for the device, or &#39;Auto&#39; (case insensitive) to use the recommended channel | [optional] 
**Enabled** | Pointer to **bool** | Turn ESL features on and off for this device | [optional] 
**NetworkId** | Pointer to **string** | The identifier for the device&#39;s network | [optional] 
**Hostname** | Pointer to **string** | Hostname of the ESL management service | [optional] 
**Provider** | Pointer to **string** | The service providing ESL functionality | [optional] 

## Methods

### NewGetDeviceWirelessElectronicShelfLabel200Response

`func NewGetDeviceWirelessElectronicShelfLabel200Response() *GetDeviceWirelessElectronicShelfLabel200Response`

NewGetDeviceWirelessElectronicShelfLabel200Response instantiates a new GetDeviceWirelessElectronicShelfLabel200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceWirelessElectronicShelfLabel200ResponseWithDefaults

`func NewGetDeviceWirelessElectronicShelfLabel200ResponseWithDefaults() *GetDeviceWirelessElectronicShelfLabel200Response`

NewGetDeviceWirelessElectronicShelfLabel200ResponseWithDefaults instantiates a new GetDeviceWirelessElectronicShelfLabel200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApEslId

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetApEslId() int32`

GetApEslId returns the ApEslId field if non-nil, zero value otherwise.

### GetApEslIdOk

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetApEslIdOk() (*int32, bool)`

GetApEslIdOk returns a tuple with the ApEslId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApEslId

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) SetApEslId(v int32)`

SetApEslId sets ApEslId field to given value.

### HasApEslId

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) HasApEslId() bool`

HasApEslId returns a boolean if a field has been set.

### GetSerial

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetChannel

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetEnabled

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNetworkId

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetHostname

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetProvider

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *GetDeviceWirelessElectronicShelfLabel200Response) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


