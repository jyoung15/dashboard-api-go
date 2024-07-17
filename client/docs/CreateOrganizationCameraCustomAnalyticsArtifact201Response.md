# CreateOrganizationCameraCustomAnalyticsArtifact201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactId** | Pointer to **string** | Custom analytics artifact ID | [optional] 
**OrganizationId** | Pointer to **string** | Organization ID | [optional] 
**Name** | Pointer to **string** | Custom analytics artifact name | [optional] 
**Status** | Pointer to [**GetOrganizationCameraCustomAnalyticsArtifacts200ResponseInnerStatus**](GetOrganizationCameraCustomAnalyticsArtifacts200ResponseInnerStatus.md) |  | [optional] 
**UploadId** | Pointer to **string** | Upload ID | [optional] 
**UploadUrl** | Pointer to **string** | Upload URL | [optional] 
**UploadUrlExpiry** | Pointer to **time.Time** | Upload URL expiry time | [optional] 

## Methods

### NewCreateOrganizationCameraCustomAnalyticsArtifact201Response

`func NewCreateOrganizationCameraCustomAnalyticsArtifact201Response() *CreateOrganizationCameraCustomAnalyticsArtifact201Response`

NewCreateOrganizationCameraCustomAnalyticsArtifact201Response instantiates a new CreateOrganizationCameraCustomAnalyticsArtifact201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationCameraCustomAnalyticsArtifact201ResponseWithDefaults

`func NewCreateOrganizationCameraCustomAnalyticsArtifact201ResponseWithDefaults() *CreateOrganizationCameraCustomAnalyticsArtifact201Response`

NewCreateOrganizationCameraCustomAnalyticsArtifact201ResponseWithDefaults instantiates a new CreateOrganizationCameraCustomAnalyticsArtifact201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactId

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) GetArtifactId() string`

GetArtifactId returns the ArtifactId field if non-nil, zero value otherwise.

### GetArtifactIdOk

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) GetArtifactIdOk() (*string, bool)`

GetArtifactIdOk returns a tuple with the ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactId

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) SetArtifactId(v string)`

SetArtifactId sets ArtifactId field to given value.

### HasArtifactId

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) HasArtifactId() bool`

HasArtifactId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetName

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) GetStatus() GetOrganizationCameraCustomAnalyticsArtifacts200ResponseInnerStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) GetStatusOk() (*GetOrganizationCameraCustomAnalyticsArtifacts200ResponseInnerStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) SetStatus(v GetOrganizationCameraCustomAnalyticsArtifacts200ResponseInnerStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUploadId

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) GetUploadId() string`

GetUploadId returns the UploadId field if non-nil, zero value otherwise.

### GetUploadIdOk

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) GetUploadIdOk() (*string, bool)`

GetUploadIdOk returns a tuple with the UploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadId

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) SetUploadId(v string)`

SetUploadId sets UploadId field to given value.

### HasUploadId

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) HasUploadId() bool`

HasUploadId returns a boolean if a field has been set.

### GetUploadUrl

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.

### HasUploadUrl

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) HasUploadUrl() bool`

HasUploadUrl returns a boolean if a field has been set.

### GetUploadUrlExpiry

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) GetUploadUrlExpiry() time.Time`

GetUploadUrlExpiry returns the UploadUrlExpiry field if non-nil, zero value otherwise.

### GetUploadUrlExpiryOk

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) GetUploadUrlExpiryOk() (*time.Time, bool)`

GetUploadUrlExpiryOk returns a tuple with the UploadUrlExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrlExpiry

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) SetUploadUrlExpiry(v time.Time)`

SetUploadUrlExpiry sets UploadUrlExpiry field to given value.

### HasUploadUrlExpiry

`func (o *CreateOrganizationCameraCustomAnalyticsArtifact201Response) HasUploadUrlExpiry() bool`

HasUploadUrlExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


