# GetOrganizationAlertsProfiles200ResponseInnerRecipients

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | Pointer to **[]string** | A list of emails that will receive information about the alert | [optional] 
**HttpServerIds** | Pointer to **[]string** | A list base64 encoded urls of webhook endpoints that will receive information about the alert | [optional] 

## Methods

### NewGetOrganizationAlertsProfiles200ResponseInnerRecipients

`func NewGetOrganizationAlertsProfiles200ResponseInnerRecipients() *GetOrganizationAlertsProfiles200ResponseInnerRecipients`

NewGetOrganizationAlertsProfiles200ResponseInnerRecipients instantiates a new GetOrganizationAlertsProfiles200ResponseInnerRecipients object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationAlertsProfiles200ResponseInnerRecipientsWithDefaults

`func NewGetOrganizationAlertsProfiles200ResponseInnerRecipientsWithDefaults() *GetOrganizationAlertsProfiles200ResponseInnerRecipients`

NewGetOrganizationAlertsProfiles200ResponseInnerRecipientsWithDefaults instantiates a new GetOrganizationAlertsProfiles200ResponseInnerRecipients object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *GetOrganizationAlertsProfiles200ResponseInnerRecipients) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *GetOrganizationAlertsProfiles200ResponseInnerRecipients) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *GetOrganizationAlertsProfiles200ResponseInnerRecipients) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *GetOrganizationAlertsProfiles200ResponseInnerRecipients) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetHttpServerIds

`func (o *GetOrganizationAlertsProfiles200ResponseInnerRecipients) GetHttpServerIds() []string`

GetHttpServerIds returns the HttpServerIds field if non-nil, zero value otherwise.

### GetHttpServerIdsOk

`func (o *GetOrganizationAlertsProfiles200ResponseInnerRecipients) GetHttpServerIdsOk() (*[]string, bool)`

GetHttpServerIdsOk returns a tuple with the HttpServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServerIds

`func (o *GetOrganizationAlertsProfiles200ResponseInnerRecipients) SetHttpServerIds(v []string)`

SetHttpServerIds sets HttpServerIds field to given value.

### HasHttpServerIds

`func (o *GetOrganizationAlertsProfiles200ResponseInnerRecipients) HasHttpServerIds() bool`

HasHttpServerIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


