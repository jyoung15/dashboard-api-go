# ProvisionNetworkClients201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clients** | Pointer to [**[]ProvisionNetworkClients201ResponseClientsInner**](ProvisionNetworkClients201ResponseClientsInner.md) | The list of clients to provision | [optional] 
**DevicePolicy** | Pointer to **string** | The name of the client&#39;s policy | [optional] 
**GroupPolicyId** | Pointer to **string** | The group policy identifier of the client | [optional] 

## Methods

### NewProvisionNetworkClients201Response

`func NewProvisionNetworkClients201Response() *ProvisionNetworkClients201Response`

NewProvisionNetworkClients201Response instantiates a new ProvisionNetworkClients201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionNetworkClients201ResponseWithDefaults

`func NewProvisionNetworkClients201ResponseWithDefaults() *ProvisionNetworkClients201Response`

NewProvisionNetworkClients201ResponseWithDefaults instantiates a new ProvisionNetworkClients201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClients

`func (o *ProvisionNetworkClients201Response) GetClients() []ProvisionNetworkClients201ResponseClientsInner`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *ProvisionNetworkClients201Response) GetClientsOk() (*[]ProvisionNetworkClients201ResponseClientsInner, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *ProvisionNetworkClients201Response) SetClients(v []ProvisionNetworkClients201ResponseClientsInner)`

SetClients sets Clients field to given value.

### HasClients

`func (o *ProvisionNetworkClients201Response) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetDevicePolicy

`func (o *ProvisionNetworkClients201Response) GetDevicePolicy() string`

GetDevicePolicy returns the DevicePolicy field if non-nil, zero value otherwise.

### GetDevicePolicyOk

`func (o *ProvisionNetworkClients201Response) GetDevicePolicyOk() (*string, bool)`

GetDevicePolicyOk returns a tuple with the DevicePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePolicy

`func (o *ProvisionNetworkClients201Response) SetDevicePolicy(v string)`

SetDevicePolicy sets DevicePolicy field to given value.

### HasDevicePolicy

`func (o *ProvisionNetworkClients201Response) HasDevicePolicy() bool`

HasDevicePolicy returns a boolean if a field has been set.

### GetGroupPolicyId

`func (o *ProvisionNetworkClients201Response) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *ProvisionNetworkClients201Response) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *ProvisionNetworkClients201Response) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *ProvisionNetworkClients201Response) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


