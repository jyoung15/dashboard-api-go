# UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **string** | Unique ID of the RADIUS accounting server. When provided, the existing RADIUS server will be updated instead of creating a new one | [optional] 
**OrganizationRadiusServerId** | Pointer to **string** | Organization wide RADIUS server ID. If this field is provided, the host, port and secret field will be ignored | [optional] 
**Host** | Pointer to **string** | Public IP address of the RADIUS accounting server | [optional] 
**Port** | Pointer to **int32** | UDP port that the RADIUS Accounting server listens on for access requests | [optional] 
**Secret** | Pointer to **string** | RADIUS client shared secret | [optional] 

## Methods

### NewUpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner

`func NewUpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner() *UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner`

NewUpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner instantiates a new UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInnerWithDefaults

`func NewUpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInnerWithDefaults() *UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner`

NewUpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInnerWithDefaults instantiates a new UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetOrganizationRadiusServerId

`func (o *UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) GetOrganizationRadiusServerId() string`

GetOrganizationRadiusServerId returns the OrganizationRadiusServerId field if non-nil, zero value otherwise.

### GetOrganizationRadiusServerIdOk

`func (o *UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) GetOrganizationRadiusServerIdOk() (*string, bool)`

GetOrganizationRadiusServerIdOk returns a tuple with the OrganizationRadiusServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRadiusServerId

`func (o *UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) SetOrganizationRadiusServerId(v string)`

SetOrganizationRadiusServerId sets OrganizationRadiusServerId field to given value.

### HasOrganizationRadiusServerId

`func (o *UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) HasOrganizationRadiusServerId() bool`

HasOrganizationRadiusServerId returns a boolean if a field has been set.

### GetHost

`func (o *UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecret

`func (o *UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *UpdateNetworkSwitchAccessPolicyRequestRadiusAccountingServersInner) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


