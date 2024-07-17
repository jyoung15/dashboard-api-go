# GetOrganizationAssuranceAlerts200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the health alert | 
**CategoryType** | **string** | Category type that the health alert belongs to | 
**Network** | [**GetOrganizationAssuranceAlerts200ResponseInnerNetwork**](GetOrganizationAssuranceAlerts200ResponseInnerNetwork.md) |  | 
**StartedAt** | **time.Time** | Time when the alert started | 
**ResolvedAt** | Pointer to **time.Time** | Time when the alert was resolved | [optional] 
**DismissedAt** | Pointer to **time.Time** | Time when the alert was dismissed | [optional] 
**DeviceType** | Pointer to **string** | Device Type that the alert occurred on | [optional] 
**Type** | **string** | Alert Type | 
**Title** | **string** | Human Readable Title for Alert type | 
**Description** | Pointer to **string** | Description containing additional details | [optional] 
**Severity** | **string** | Alert severity | 
**Scope** | Pointer to [**GetOrganizationAssuranceAlerts200ResponseInnerScope**](GetOrganizationAssuranceAlerts200ResponseInnerScope.md) |  | [optional] 

## Methods

### NewGetOrganizationAssuranceAlerts200ResponseInner

`func NewGetOrganizationAssuranceAlerts200ResponseInner(id string, categoryType string, network GetOrganizationAssuranceAlerts200ResponseInnerNetwork, startedAt time.Time, type_ string, title string, severity string, ) *GetOrganizationAssuranceAlerts200ResponseInner`

NewGetOrganizationAssuranceAlerts200ResponseInner instantiates a new GetOrganizationAssuranceAlerts200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationAssuranceAlerts200ResponseInnerWithDefaults

`func NewGetOrganizationAssuranceAlerts200ResponseInnerWithDefaults() *GetOrganizationAssuranceAlerts200ResponseInner`

NewGetOrganizationAssuranceAlerts200ResponseInnerWithDefaults instantiates a new GetOrganizationAssuranceAlerts200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetCategoryType

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetCategoryType() string`

GetCategoryType returns the CategoryType field if non-nil, zero value otherwise.

### GetCategoryTypeOk

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetCategoryTypeOk() (*string, bool)`

GetCategoryTypeOk returns a tuple with the CategoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryType

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) SetCategoryType(v string)`

SetCategoryType sets CategoryType field to given value.


### GetNetwork

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetNetwork() GetOrganizationAssuranceAlerts200ResponseInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetNetworkOk() (*GetOrganizationAssuranceAlerts200ResponseInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) SetNetwork(v GetOrganizationAssuranceAlerts200ResponseInnerNetwork)`

SetNetwork sets Network field to given value.


### GetStartedAt

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetResolvedAt

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetResolvedAt() time.Time`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetResolvedAtOk() (*time.Time, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) SetResolvedAt(v time.Time)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### GetDismissedAt

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetDismissedAt() time.Time`

GetDismissedAt returns the DismissedAt field if non-nil, zero value otherwise.

### GetDismissedAtOk

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetDismissedAtOk() (*time.Time, bool)`

GetDismissedAtOk returns a tuple with the DismissedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissedAt

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) SetDismissedAt(v time.Time)`

SetDismissedAt sets DismissedAt field to given value.

### HasDismissedAt

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) HasDismissedAt() bool`

HasDismissedAt returns a boolean if a field has been set.

### GetDeviceType

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetType

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSeverity

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetScope

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetScope() GetOrganizationAssuranceAlerts200ResponseInnerScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) GetScopeOk() (*GetOrganizationAssuranceAlerts200ResponseInnerScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) SetScope(v GetOrganizationAssuranceAlerts200ResponseInnerScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *GetOrganizationAssuranceAlerts200ResponseInner) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


