# GetNetworkAlertsSettings200ResponseDefaultDestinations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | Pointer to **[]string** | A list of emails that will receive the alert(s). | [optional] 
**AllAdmins** | Pointer to **bool** | If true, then all network admins will receive emails. | [optional] 
**Snmp** | Pointer to **bool** | If true, then an SNMP trap will be sent if there is an SNMP trap server configured for this network. | [optional] 
**HttpServerIds** | Pointer to **[]string** | A list of HTTP server IDs to send a Webhook to | [optional] 

## Methods

### NewGetNetworkAlertsSettings200ResponseDefaultDestinations

`func NewGetNetworkAlertsSettings200ResponseDefaultDestinations() *GetNetworkAlertsSettings200ResponseDefaultDestinations`

NewGetNetworkAlertsSettings200ResponseDefaultDestinations instantiates a new GetNetworkAlertsSettings200ResponseDefaultDestinations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkAlertsSettings200ResponseDefaultDestinationsWithDefaults

`func NewGetNetworkAlertsSettings200ResponseDefaultDestinationsWithDefaults() *GetNetworkAlertsSettings200ResponseDefaultDestinations`

NewGetNetworkAlertsSettings200ResponseDefaultDestinationsWithDefaults instantiates a new GetNetworkAlertsSettings200ResponseDefaultDestinations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetAllAdmins

`func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) GetAllAdmins() bool`

GetAllAdmins returns the AllAdmins field if non-nil, zero value otherwise.

### GetAllAdminsOk

`func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) GetAllAdminsOk() (*bool, bool)`

GetAllAdminsOk returns a tuple with the AllAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAdmins

`func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) SetAllAdmins(v bool)`

SetAllAdmins sets AllAdmins field to given value.

### HasAllAdmins

`func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) HasAllAdmins() bool`

HasAllAdmins returns a boolean if a field has been set.

### GetSnmp

`func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) GetSnmp() bool`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) GetSnmpOk() (*bool, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) SetSnmp(v bool)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetHttpServerIds

`func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) GetHttpServerIds() []string`

GetHttpServerIds returns the HttpServerIds field if non-nil, zero value otherwise.

### GetHttpServerIdsOk

`func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) GetHttpServerIdsOk() (*[]string, bool)`

GetHttpServerIdsOk returns a tuple with the HttpServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServerIds

`func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) SetHttpServerIds(v []string)`

SetHttpServerIds sets HttpServerIds field to given value.

### HasHttpServerIds

`func (o *GetNetworkAlertsSettings200ResponseDefaultDestinations) HasHttpServerIds() bool`

HasHttpServerIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


