# GetOrganizationClientsSearch200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The ID of the client | [optional] 
**Mac** | Pointer to **string** | The MAC address of the client | [optional] 
**Manufacturer** | Pointer to **string** | Manufacturer of the client | [optional] 
**Records** | Pointer to [**[]GetOrganizationClientsSearch200ResponseRecordsInner**](GetOrganizationClientsSearch200ResponseRecordsInner.md) | The clients that appear on any networks within an organization | [optional] 

## Methods

### NewGetOrganizationClientsSearch200Response

`func NewGetOrganizationClientsSearch200Response() *GetOrganizationClientsSearch200Response`

NewGetOrganizationClientsSearch200Response instantiates a new GetOrganizationClientsSearch200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationClientsSearch200ResponseWithDefaults

`func NewGetOrganizationClientsSearch200ResponseWithDefaults() *GetOrganizationClientsSearch200Response`

NewGetOrganizationClientsSearch200ResponseWithDefaults instantiates a new GetOrganizationClientsSearch200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *GetOrganizationClientsSearch200Response) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GetOrganizationClientsSearch200Response) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GetOrganizationClientsSearch200Response) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GetOrganizationClientsSearch200Response) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetMac

`func (o *GetOrganizationClientsSearch200Response) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetOrganizationClientsSearch200Response) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetOrganizationClientsSearch200Response) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetOrganizationClientsSearch200Response) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetManufacturer

`func (o *GetOrganizationClientsSearch200Response) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *GetOrganizationClientsSearch200Response) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *GetOrganizationClientsSearch200Response) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *GetOrganizationClientsSearch200Response) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetRecords

`func (o *GetOrganizationClientsSearch200Response) GetRecords() []GetOrganizationClientsSearch200ResponseRecordsInner`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *GetOrganizationClientsSearch200Response) GetRecordsOk() (*[]GetOrganizationClientsSearch200ResponseRecordsInner, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *GetOrganizationClientsSearch200Response) SetRecords(v []GetOrganizationClientsSearch200ResponseRecordsInner)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *GetOrganizationClientsSearch200Response) HasRecords() bool`

HasRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


