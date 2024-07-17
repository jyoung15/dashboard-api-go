# ClaimNetworkDevices200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serials** | Pointer to **[]string** | The serials of the devices | [optional] 
**Errors** | Pointer to [**[]ClaimNetworkDevices200ResponseErrorsInner**](ClaimNetworkDevices200ResponseErrorsInner.md) | Errors for devices that were not added | [optional] 

## Methods

### NewClaimNetworkDevices200Response

`func NewClaimNetworkDevices200Response() *ClaimNetworkDevices200Response`

NewClaimNetworkDevices200Response instantiates a new ClaimNetworkDevices200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClaimNetworkDevices200ResponseWithDefaults

`func NewClaimNetworkDevices200ResponseWithDefaults() *ClaimNetworkDevices200Response`

NewClaimNetworkDevices200ResponseWithDefaults instantiates a new ClaimNetworkDevices200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerials

`func (o *ClaimNetworkDevices200Response) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *ClaimNetworkDevices200Response) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *ClaimNetworkDevices200Response) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *ClaimNetworkDevices200Response) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetErrors

`func (o *ClaimNetworkDevices200Response) GetErrors() []ClaimNetworkDevices200ResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ClaimNetworkDevices200Response) GetErrorsOk() (*[]ClaimNetworkDevices200ResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ClaimNetworkDevices200Response) SetErrors(v []ClaimNetworkDevices200ResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ClaimNetworkDevices200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


