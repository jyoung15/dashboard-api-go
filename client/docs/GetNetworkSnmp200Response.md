# GetNetworkSnmp200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **string** | The type of SNMP access. Can be one of &#39;none&#39; (disabled), &#39;community&#39; (V1/V2c), or &#39;users&#39; (V3). | [optional] 
**CommunityString** | Pointer to **string** | SNMP community string if access is &#39;community&#39;. | [optional] 
**Users** | Pointer to [**[]GetNetworkSnmp200ResponseUsersInner**](GetNetworkSnmp200ResponseUsersInner.md) | SNMP settings if access is &#39;users&#39;. | [optional] 

## Methods

### NewGetNetworkSnmp200Response

`func NewGetNetworkSnmp200Response() *GetNetworkSnmp200Response`

NewGetNetworkSnmp200Response instantiates a new GetNetworkSnmp200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSnmp200ResponseWithDefaults

`func NewGetNetworkSnmp200ResponseWithDefaults() *GetNetworkSnmp200Response`

NewGetNetworkSnmp200ResponseWithDefaults instantiates a new GetNetworkSnmp200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *GetNetworkSnmp200Response) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *GetNetworkSnmp200Response) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *GetNetworkSnmp200Response) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *GetNetworkSnmp200Response) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetCommunityString

`func (o *GetNetworkSnmp200Response) GetCommunityString() string`

GetCommunityString returns the CommunityString field if non-nil, zero value otherwise.

### GetCommunityStringOk

`func (o *GetNetworkSnmp200Response) GetCommunityStringOk() (*string, bool)`

GetCommunityStringOk returns a tuple with the CommunityString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityString

`func (o *GetNetworkSnmp200Response) SetCommunityString(v string)`

SetCommunityString sets CommunityString field to given value.

### HasCommunityString

`func (o *GetNetworkSnmp200Response) HasCommunityString() bool`

HasCommunityString returns a boolean if a field has been set.

### GetUsers

`func (o *GetNetworkSnmp200Response) GetUsers() []GetNetworkSnmp200ResponseUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GetNetworkSnmp200Response) GetUsersOk() (*[]GetNetworkSnmp200ResponseUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GetNetworkSnmp200Response) SetUsers(v []GetNetworkSnmp200ResponseUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GetNetworkSnmp200Response) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


