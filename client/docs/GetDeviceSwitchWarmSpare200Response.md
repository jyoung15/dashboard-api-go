# GetDeviceSwitchWarmSpare200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable or disable warm spare for a switch | [optional] 
**PrimarySerial** | Pointer to **string** | Serial number of the primary switch | [optional] 
**SpareSerial** | Pointer to **string** | Serial number of the warm spare switch | [optional] 

## Methods

### NewGetDeviceSwitchWarmSpare200Response

`func NewGetDeviceSwitchWarmSpare200Response() *GetDeviceSwitchWarmSpare200Response`

NewGetDeviceSwitchWarmSpare200Response instantiates a new GetDeviceSwitchWarmSpare200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceSwitchWarmSpare200ResponseWithDefaults

`func NewGetDeviceSwitchWarmSpare200ResponseWithDefaults() *GetDeviceSwitchWarmSpare200Response`

NewGetDeviceSwitchWarmSpare200ResponseWithDefaults instantiates a new GetDeviceSwitchWarmSpare200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GetDeviceSwitchWarmSpare200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetDeviceSwitchWarmSpare200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetDeviceSwitchWarmSpare200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetDeviceSwitchWarmSpare200Response) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPrimarySerial

`func (o *GetDeviceSwitchWarmSpare200Response) GetPrimarySerial() string`

GetPrimarySerial returns the PrimarySerial field if non-nil, zero value otherwise.

### GetPrimarySerialOk

`func (o *GetDeviceSwitchWarmSpare200Response) GetPrimarySerialOk() (*string, bool)`

GetPrimarySerialOk returns a tuple with the PrimarySerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySerial

`func (o *GetDeviceSwitchWarmSpare200Response) SetPrimarySerial(v string)`

SetPrimarySerial sets PrimarySerial field to given value.

### HasPrimarySerial

`func (o *GetDeviceSwitchWarmSpare200Response) HasPrimarySerial() bool`

HasPrimarySerial returns a boolean if a field has been set.

### GetSpareSerial

`func (o *GetDeviceSwitchWarmSpare200Response) GetSpareSerial() string`

GetSpareSerial returns the SpareSerial field if non-nil, zero value otherwise.

### GetSpareSerialOk

`func (o *GetDeviceSwitchWarmSpare200Response) GetSpareSerialOk() (*string, bool)`

GetSpareSerialOk returns a tuple with the SpareSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpareSerial

`func (o *GetDeviceSwitchWarmSpare200Response) SetSpareSerial(v string)`

SetSpareSerial sets SpareSerial field to given value.

### HasSpareSerial

`func (o *GetDeviceSwitchWarmSpare200Response) HasSpareSerial() bool`

HasSpareSerial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


