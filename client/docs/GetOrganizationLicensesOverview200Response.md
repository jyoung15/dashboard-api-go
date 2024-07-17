# GetOrganizationLicensesOverview200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | License status (Co-termination licensing only) | [optional] 
**ExpirationDate** | Pointer to **string** | License expiration date (Co-termination licensing only) | [optional] 
**LicensedDeviceCounts** | Pointer to **map[string]int32** | License counts (Co-termination licensing only) | [optional] 
**LicenseCount** | Pointer to **int32** | Total number of licenses (Per-device licensing only) | [optional] 
**States** | Pointer to [**GetOrganizationLicensesOverview200ResponseStates**](GetOrganizationLicensesOverview200ResponseStates.md) |  | [optional] 
**LicenseTypes** | Pointer to [**[]GetOrganizationLicensesOverview200ResponseLicenseTypesInner**](GetOrganizationLicensesOverview200ResponseLicenseTypesInner.md) | Data by license type (Per-device licensing only) | [optional] 
**SystemsManager** | Pointer to [**GetOrganizationLicensesOverview200ResponseSystemsManager**](GetOrganizationLicensesOverview200ResponseSystemsManager.md) |  | [optional] 

## Methods

### NewGetOrganizationLicensesOverview200Response

`func NewGetOrganizationLicensesOverview200Response() *GetOrganizationLicensesOverview200Response`

NewGetOrganizationLicensesOverview200Response instantiates a new GetOrganizationLicensesOverview200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationLicensesOverview200ResponseWithDefaults

`func NewGetOrganizationLicensesOverview200ResponseWithDefaults() *GetOrganizationLicensesOverview200Response`

NewGetOrganizationLicensesOverview200ResponseWithDefaults instantiates a new GetOrganizationLicensesOverview200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetOrganizationLicensesOverview200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOrganizationLicensesOverview200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOrganizationLicensesOverview200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetOrganizationLicensesOverview200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExpirationDate

`func (o *GetOrganizationLicensesOverview200Response) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *GetOrganizationLicensesOverview200Response) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *GetOrganizationLicensesOverview200Response) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *GetOrganizationLicensesOverview200Response) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetLicensedDeviceCounts

`func (o *GetOrganizationLicensesOverview200Response) GetLicensedDeviceCounts() map[string]int32`

GetLicensedDeviceCounts returns the LicensedDeviceCounts field if non-nil, zero value otherwise.

### GetLicensedDeviceCountsOk

`func (o *GetOrganizationLicensesOverview200Response) GetLicensedDeviceCountsOk() (*map[string]int32, bool)`

GetLicensedDeviceCountsOk returns a tuple with the LicensedDeviceCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensedDeviceCounts

`func (o *GetOrganizationLicensesOverview200Response) SetLicensedDeviceCounts(v map[string]int32)`

SetLicensedDeviceCounts sets LicensedDeviceCounts field to given value.

### HasLicensedDeviceCounts

`func (o *GetOrganizationLicensesOverview200Response) HasLicensedDeviceCounts() bool`

HasLicensedDeviceCounts returns a boolean if a field has been set.

### GetLicenseCount

`func (o *GetOrganizationLicensesOverview200Response) GetLicenseCount() int32`

GetLicenseCount returns the LicenseCount field if non-nil, zero value otherwise.

### GetLicenseCountOk

`func (o *GetOrganizationLicensesOverview200Response) GetLicenseCountOk() (*int32, bool)`

GetLicenseCountOk returns a tuple with the LicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseCount

`func (o *GetOrganizationLicensesOverview200Response) SetLicenseCount(v int32)`

SetLicenseCount sets LicenseCount field to given value.

### HasLicenseCount

`func (o *GetOrganizationLicensesOverview200Response) HasLicenseCount() bool`

HasLicenseCount returns a boolean if a field has been set.

### GetStates

`func (o *GetOrganizationLicensesOverview200Response) GetStates() GetOrganizationLicensesOverview200ResponseStates`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *GetOrganizationLicensesOverview200Response) GetStatesOk() (*GetOrganizationLicensesOverview200ResponseStates, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *GetOrganizationLicensesOverview200Response) SetStates(v GetOrganizationLicensesOverview200ResponseStates)`

SetStates sets States field to given value.

### HasStates

`func (o *GetOrganizationLicensesOverview200Response) HasStates() bool`

HasStates returns a boolean if a field has been set.

### GetLicenseTypes

`func (o *GetOrganizationLicensesOverview200Response) GetLicenseTypes() []GetOrganizationLicensesOverview200ResponseLicenseTypesInner`

GetLicenseTypes returns the LicenseTypes field if non-nil, zero value otherwise.

### GetLicenseTypesOk

`func (o *GetOrganizationLicensesOverview200Response) GetLicenseTypesOk() (*[]GetOrganizationLicensesOverview200ResponseLicenseTypesInner, bool)`

GetLicenseTypesOk returns a tuple with the LicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseTypes

`func (o *GetOrganizationLicensesOverview200Response) SetLicenseTypes(v []GetOrganizationLicensesOverview200ResponseLicenseTypesInner)`

SetLicenseTypes sets LicenseTypes field to given value.

### HasLicenseTypes

`func (o *GetOrganizationLicensesOverview200Response) HasLicenseTypes() bool`

HasLicenseTypes returns a boolean if a field has been set.

### GetSystemsManager

`func (o *GetOrganizationLicensesOverview200Response) GetSystemsManager() GetOrganizationLicensesOverview200ResponseSystemsManager`

GetSystemsManager returns the SystemsManager field if non-nil, zero value otherwise.

### GetSystemsManagerOk

`func (o *GetOrganizationLicensesOverview200Response) GetSystemsManagerOk() (*GetOrganizationLicensesOverview200ResponseSystemsManager, bool)`

GetSystemsManagerOk returns a tuple with the SystemsManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemsManager

`func (o *GetOrganizationLicensesOverview200Response) SetSystemsManager(v GetOrganizationLicensesOverview200ResponseSystemsManager)`

SetSystemsManager sets SystemsManager field to given value.

### HasSystemsManager

`func (o *GetOrganizationLicensesOverview200Response) HasSystemsManager() bool`

HasSystemsManager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


