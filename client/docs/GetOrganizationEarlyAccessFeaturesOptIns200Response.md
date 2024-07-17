# GetOrganizationEarlyAccessFeaturesOptIns200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of Early Access Feature | [optional] 
**ShortName** | Pointer to **string** | Name of Early Access Feature | [optional] 
**LimitScopeToNetworks** | Pointer to [**[]GetOrganizationEarlyAccessFeaturesOptIns200ResponseLimitScopeToNetworksInner**](GetOrganizationEarlyAccessFeaturesOptIns200ResponseLimitScopeToNetworksInner.md) | Networks assigned to the Early Access Feature | [optional] 
**OptOutEligibility** | Pointer to [**GetOrganizationEarlyAccessFeaturesOptIns200ResponseOptOutEligibility**](GetOrganizationEarlyAccessFeaturesOptIns200ResponseOptOutEligibility.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when Early Access Feature was created | [optional] 

## Methods

### NewGetOrganizationEarlyAccessFeaturesOptIns200Response

`func NewGetOrganizationEarlyAccessFeaturesOptIns200Response() *GetOrganizationEarlyAccessFeaturesOptIns200Response`

NewGetOrganizationEarlyAccessFeaturesOptIns200Response instantiates a new GetOrganizationEarlyAccessFeaturesOptIns200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationEarlyAccessFeaturesOptIns200ResponseWithDefaults

`func NewGetOrganizationEarlyAccessFeaturesOptIns200ResponseWithDefaults() *GetOrganizationEarlyAccessFeaturesOptIns200Response`

NewGetOrganizationEarlyAccessFeaturesOptIns200ResponseWithDefaults instantiates a new GetOrganizationEarlyAccessFeaturesOptIns200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetShortName

`func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetLimitScopeToNetworks

`func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) GetLimitScopeToNetworks() []GetOrganizationEarlyAccessFeaturesOptIns200ResponseLimitScopeToNetworksInner`

GetLimitScopeToNetworks returns the LimitScopeToNetworks field if non-nil, zero value otherwise.

### GetLimitScopeToNetworksOk

`func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) GetLimitScopeToNetworksOk() (*[]GetOrganizationEarlyAccessFeaturesOptIns200ResponseLimitScopeToNetworksInner, bool)`

GetLimitScopeToNetworksOk returns a tuple with the LimitScopeToNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitScopeToNetworks

`func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) SetLimitScopeToNetworks(v []GetOrganizationEarlyAccessFeaturesOptIns200ResponseLimitScopeToNetworksInner)`

SetLimitScopeToNetworks sets LimitScopeToNetworks field to given value.

### HasLimitScopeToNetworks

`func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) HasLimitScopeToNetworks() bool`

HasLimitScopeToNetworks returns a boolean if a field has been set.

### GetOptOutEligibility

`func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) GetOptOutEligibility() GetOrganizationEarlyAccessFeaturesOptIns200ResponseOptOutEligibility`

GetOptOutEligibility returns the OptOutEligibility field if non-nil, zero value otherwise.

### GetOptOutEligibilityOk

`func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) GetOptOutEligibilityOk() (*GetOrganizationEarlyAccessFeaturesOptIns200ResponseOptOutEligibility, bool)`

GetOptOutEligibilityOk returns a tuple with the OptOutEligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptOutEligibility

`func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) SetOptOutEligibility(v GetOrganizationEarlyAccessFeaturesOptIns200ResponseOptOutEligibility)`

SetOptOutEligibility sets OptOutEligibility field to given value.

### HasOptOutEligibility

`func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) HasOptOutEligibility() bool`

HasOptOutEligibility returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetOrganizationEarlyAccessFeaturesOptIns200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


