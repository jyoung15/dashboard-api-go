# GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **string** | Unique ID of the RADIUS accounting server | [optional] 
**OrganizationRadiusServerId** | Pointer to **string** | Organization wide RADIUS server ID. This value will be empty if this RADIUS server is not an organization wide RADIUS server | [optional] 
**Host** | Pointer to **string** | Public IP address of the RADIUS accounting server | [optional] 
**Port** | Pointer to **int32** | UDP port that the RADIUS Accounting server listens on for access requests | [optional] 

## Methods

### NewGetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner

`func NewGetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner() *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner`

NewGetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner instantiates a new GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInnerWithDefaults

`func NewGetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInnerWithDefaults() *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner`

NewGetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInnerWithDefaults instantiates a new GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetOrganizationRadiusServerId

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner) GetOrganizationRadiusServerId() string`

GetOrganizationRadiusServerId returns the OrganizationRadiusServerId field if non-nil, zero value otherwise.

### GetOrganizationRadiusServerIdOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner) GetOrganizationRadiusServerIdOk() (*string, bool)`

GetOrganizationRadiusServerIdOk returns a tuple with the OrganizationRadiusServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRadiusServerId

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner) SetOrganizationRadiusServerId(v string)`

SetOrganizationRadiusServerId sets OrganizationRadiusServerId field to given value.

### HasOrganizationRadiusServerId

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner) HasOrganizationRadiusServerId() bool`

HasOrganizationRadiusServerId returns a boolean if a field has been set.

### GetHost

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetNetworkSwitchAccessPolicies200ResponseInnerRadiusAccountingServersInner) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


