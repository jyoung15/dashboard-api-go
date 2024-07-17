# GetDeviceCameraCustomAnalytics200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether custom analytics is enabled | [optional] 
**ArtifactId** | Pointer to **string** | Custom analytics artifact ID | [optional] 
**Parameters** | Pointer to [**[]GetDeviceCameraCustomAnalytics200ResponseParametersInner**](GetDeviceCameraCustomAnalytics200ResponseParametersInner.md) | Parameters for the custom analytics workload | [optional] 

## Methods

### NewGetDeviceCameraCustomAnalytics200Response

`func NewGetDeviceCameraCustomAnalytics200Response() *GetDeviceCameraCustomAnalytics200Response`

NewGetDeviceCameraCustomAnalytics200Response instantiates a new GetDeviceCameraCustomAnalytics200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceCameraCustomAnalytics200ResponseWithDefaults

`func NewGetDeviceCameraCustomAnalytics200ResponseWithDefaults() *GetDeviceCameraCustomAnalytics200Response`

NewGetDeviceCameraCustomAnalytics200ResponseWithDefaults instantiates a new GetDeviceCameraCustomAnalytics200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GetDeviceCameraCustomAnalytics200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetDeviceCameraCustomAnalytics200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetDeviceCameraCustomAnalytics200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetDeviceCameraCustomAnalytics200Response) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetArtifactId

`func (o *GetDeviceCameraCustomAnalytics200Response) GetArtifactId() string`

GetArtifactId returns the ArtifactId field if non-nil, zero value otherwise.

### GetArtifactIdOk

`func (o *GetDeviceCameraCustomAnalytics200Response) GetArtifactIdOk() (*string, bool)`

GetArtifactIdOk returns a tuple with the ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactId

`func (o *GetDeviceCameraCustomAnalytics200Response) SetArtifactId(v string)`

SetArtifactId sets ArtifactId field to given value.

### HasArtifactId

`func (o *GetDeviceCameraCustomAnalytics200Response) HasArtifactId() bool`

HasArtifactId returns a boolean if a field has been set.

### GetParameters

`func (o *GetDeviceCameraCustomAnalytics200Response) GetParameters() []GetDeviceCameraCustomAnalytics200ResponseParametersInner`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *GetDeviceCameraCustomAnalytics200Response) GetParametersOk() (*[]GetDeviceCameraCustomAnalytics200ResponseParametersInner, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *GetDeviceCameraCustomAnalytics200Response) SetParameters(v []GetDeviceCameraCustomAnalytics200ResponseParametersInner)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *GetDeviceCameraCustomAnalytics200Response) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


