# GetOrganizationAlertsProfiles200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The alert config ID | [optional] 
**Type** | Pointer to **string** | The alert type | [optional] 
**Enabled** | Pointer to **bool** | Is the alert config enabled | [optional] 
**AlertCondition** | Pointer to [**GetOrganizationAlertsProfiles200ResponseInnerAlertCondition**](GetOrganizationAlertsProfiles200ResponseInnerAlertCondition.md) |  | [optional] 
**Recipients** | Pointer to [**GetOrganizationAlertsProfiles200ResponseInnerRecipients**](GetOrganizationAlertsProfiles200ResponseInnerRecipients.md) |  | [optional] 
**NetworkTags** | Pointer to **[]string** | Networks with these tags will be monitored for the alert | [optional] 
**Description** | Pointer to **string** | User supplied description of the alert | [optional] 

## Methods

### NewGetOrganizationAlertsProfiles200ResponseInner

`func NewGetOrganizationAlertsProfiles200ResponseInner() *GetOrganizationAlertsProfiles200ResponseInner`

NewGetOrganizationAlertsProfiles200ResponseInner instantiates a new GetOrganizationAlertsProfiles200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationAlertsProfiles200ResponseInnerWithDefaults

`func NewGetOrganizationAlertsProfiles200ResponseInnerWithDefaults() *GetOrganizationAlertsProfiles200ResponseInner`

NewGetOrganizationAlertsProfiles200ResponseInnerWithDefaults instantiates a new GetOrganizationAlertsProfiles200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOrganizationAlertsProfiles200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrganizationAlertsProfiles200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrganizationAlertsProfiles200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetOrganizationAlertsProfiles200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetOrganizationAlertsProfiles200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetOrganizationAlertsProfiles200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetOrganizationAlertsProfiles200ResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetOrganizationAlertsProfiles200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *GetOrganizationAlertsProfiles200ResponseInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetOrganizationAlertsProfiles200ResponseInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetOrganizationAlertsProfiles200ResponseInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetOrganizationAlertsProfiles200ResponseInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAlertCondition

`func (o *GetOrganizationAlertsProfiles200ResponseInner) GetAlertCondition() GetOrganizationAlertsProfiles200ResponseInnerAlertCondition`

GetAlertCondition returns the AlertCondition field if non-nil, zero value otherwise.

### GetAlertConditionOk

`func (o *GetOrganizationAlertsProfiles200ResponseInner) GetAlertConditionOk() (*GetOrganizationAlertsProfiles200ResponseInnerAlertCondition, bool)`

GetAlertConditionOk returns a tuple with the AlertCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCondition

`func (o *GetOrganizationAlertsProfiles200ResponseInner) SetAlertCondition(v GetOrganizationAlertsProfiles200ResponseInnerAlertCondition)`

SetAlertCondition sets AlertCondition field to given value.

### HasAlertCondition

`func (o *GetOrganizationAlertsProfiles200ResponseInner) HasAlertCondition() bool`

HasAlertCondition returns a boolean if a field has been set.

### GetRecipients

`func (o *GetOrganizationAlertsProfiles200ResponseInner) GetRecipients() GetOrganizationAlertsProfiles200ResponseInnerRecipients`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *GetOrganizationAlertsProfiles200ResponseInner) GetRecipientsOk() (*GetOrganizationAlertsProfiles200ResponseInnerRecipients, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *GetOrganizationAlertsProfiles200ResponseInner) SetRecipients(v GetOrganizationAlertsProfiles200ResponseInnerRecipients)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *GetOrganizationAlertsProfiles200ResponseInner) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetNetworkTags

`func (o *GetOrganizationAlertsProfiles200ResponseInner) GetNetworkTags() []string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *GetOrganizationAlertsProfiles200ResponseInner) GetNetworkTagsOk() (*[]string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *GetOrganizationAlertsProfiles200ResponseInner) SetNetworkTags(v []string)`

SetNetworkTags sets NetworkTags field to given value.

### HasNetworkTags

`func (o *GetOrganizationAlertsProfiles200ResponseInner) HasNetworkTags() bool`

HasNetworkTags returns a boolean if a field has been set.

### GetDescription

`func (o *GetOrganizationAlertsProfiles200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetOrganizationAlertsProfiles200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetOrganizationAlertsProfiles200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetOrganizationAlertsProfiles200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


