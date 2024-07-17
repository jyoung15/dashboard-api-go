# BindNetwork200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Network ID | [optional] 
**OrganizationId** | Pointer to **string** | Organization ID | [optional] 
**Name** | Pointer to **string** | Network name | [optional] 
**ProductTypes** | Pointer to **[]string** | List of the product types that the network supports | [optional] 
**TimeZone** | Pointer to **string** | Timezone of the network | [optional] 
**Tags** | Pointer to **[]string** | Network tags | [optional] 
**EnrollmentString** | Pointer to **string** | Enrollment string for the network | [optional] 
**Url** | Pointer to **string** | URL to the network Dashboard UI | [optional] 
**Notes** | Pointer to **string** | Notes for the network | [optional] 
**IsBoundToConfigTemplate** | Pointer to **bool** | If the network is bound to a config template | [optional] 
**ConfigTemplateId** | Pointer to **string** | ID of the config template the network is being bound to | [optional] 

## Methods

### NewBindNetwork200Response

`func NewBindNetwork200Response() *BindNetwork200Response`

NewBindNetwork200Response instantiates a new BindNetwork200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindNetwork200ResponseWithDefaults

`func NewBindNetwork200ResponseWithDefaults() *BindNetwork200Response`

NewBindNetwork200ResponseWithDefaults instantiates a new BindNetwork200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BindNetwork200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BindNetwork200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BindNetwork200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BindNetwork200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *BindNetwork200Response) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *BindNetwork200Response) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *BindNetwork200Response) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *BindNetwork200Response) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetName

`func (o *BindNetwork200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BindNetwork200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BindNetwork200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BindNetwork200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProductTypes

`func (o *BindNetwork200Response) GetProductTypes() []string`

GetProductTypes returns the ProductTypes field if non-nil, zero value otherwise.

### GetProductTypesOk

`func (o *BindNetwork200Response) GetProductTypesOk() (*[]string, bool)`

GetProductTypesOk returns a tuple with the ProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypes

`func (o *BindNetwork200Response) SetProductTypes(v []string)`

SetProductTypes sets ProductTypes field to given value.

### HasProductTypes

`func (o *BindNetwork200Response) HasProductTypes() bool`

HasProductTypes returns a boolean if a field has been set.

### GetTimeZone

`func (o *BindNetwork200Response) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *BindNetwork200Response) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *BindNetwork200Response) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *BindNetwork200Response) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetTags

`func (o *BindNetwork200Response) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BindNetwork200Response) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BindNetwork200Response) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BindNetwork200Response) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetEnrollmentString

`func (o *BindNetwork200Response) GetEnrollmentString() string`

GetEnrollmentString returns the EnrollmentString field if non-nil, zero value otherwise.

### GetEnrollmentStringOk

`func (o *BindNetwork200Response) GetEnrollmentStringOk() (*string, bool)`

GetEnrollmentStringOk returns a tuple with the EnrollmentString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentString

`func (o *BindNetwork200Response) SetEnrollmentString(v string)`

SetEnrollmentString sets EnrollmentString field to given value.

### HasEnrollmentString

`func (o *BindNetwork200Response) HasEnrollmentString() bool`

HasEnrollmentString returns a boolean if a field has been set.

### GetUrl

`func (o *BindNetwork200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BindNetwork200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BindNetwork200Response) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *BindNetwork200Response) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetNotes

`func (o *BindNetwork200Response) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BindNetwork200Response) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BindNetwork200Response) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BindNetwork200Response) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetIsBoundToConfigTemplate

`func (o *BindNetwork200Response) GetIsBoundToConfigTemplate() bool`

GetIsBoundToConfigTemplate returns the IsBoundToConfigTemplate field if non-nil, zero value otherwise.

### GetIsBoundToConfigTemplateOk

`func (o *BindNetwork200Response) GetIsBoundToConfigTemplateOk() (*bool, bool)`

GetIsBoundToConfigTemplateOk returns a tuple with the IsBoundToConfigTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBoundToConfigTemplate

`func (o *BindNetwork200Response) SetIsBoundToConfigTemplate(v bool)`

SetIsBoundToConfigTemplate sets IsBoundToConfigTemplate field to given value.

### HasIsBoundToConfigTemplate

`func (o *BindNetwork200Response) HasIsBoundToConfigTemplate() bool`

HasIsBoundToConfigTemplate returns a boolean if a field has been set.

### GetConfigTemplateId

`func (o *BindNetwork200Response) GetConfigTemplateId() string`

GetConfigTemplateId returns the ConfigTemplateId field if non-nil, zero value otherwise.

### GetConfigTemplateIdOk

`func (o *BindNetwork200Response) GetConfigTemplateIdOk() (*string, bool)`

GetConfigTemplateIdOk returns a tuple with the ConfigTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTemplateId

`func (o *BindNetwork200Response) SetConfigTemplateId(v string)`

SetConfigTemplateId sets ConfigTemplateId field to given value.

### HasConfigTemplateId

`func (o *BindNetwork200Response) HasConfigTemplateId() bool`

HasConfigTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


