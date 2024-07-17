# GetOrganizationPolicyObjects200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Policy object ID | [optional] 
**Name** | Pointer to **string** | Name of policy object (alphanumeric, space, dash, or underscore characters only). | [optional] 
**Category** | Pointer to **string** | Category of a policy object (one of: adaptivePolicy, network) | [optional] 
**Type** | Pointer to **string** | Type of a policy object (one of: adaptivePolicyIpv4Cidr, cidr, fqdn, ipAndMask) | [optional] 
**Cidr** | Pointer to **string** | CIDR Value of a policy object | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time Stamp of policy object creation. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time Stamp of policy object updation. | [optional] 
**GroupIds** | Pointer to **[]string** | The IDs of policy object groups the policy object belongs to. | [optional] 
**NetworkIds** | Pointer to **[]string** | The IDs of the networks that use the policy object. | [optional] 

## Methods

### NewGetOrganizationPolicyObjects200Response

`func NewGetOrganizationPolicyObjects200Response() *GetOrganizationPolicyObjects200Response`

NewGetOrganizationPolicyObjects200Response instantiates a new GetOrganizationPolicyObjects200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationPolicyObjects200ResponseWithDefaults

`func NewGetOrganizationPolicyObjects200ResponseWithDefaults() *GetOrganizationPolicyObjects200Response`

NewGetOrganizationPolicyObjects200ResponseWithDefaults instantiates a new GetOrganizationPolicyObjects200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOrganizationPolicyObjects200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrganizationPolicyObjects200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrganizationPolicyObjects200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetOrganizationPolicyObjects200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetOrganizationPolicyObjects200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationPolicyObjects200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationPolicyObjects200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationPolicyObjects200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *GetOrganizationPolicyObjects200Response) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetOrganizationPolicyObjects200Response) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetOrganizationPolicyObjects200Response) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetOrganizationPolicyObjects200Response) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetType

`func (o *GetOrganizationPolicyObjects200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetOrganizationPolicyObjects200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetOrganizationPolicyObjects200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetOrganizationPolicyObjects200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCidr

`func (o *GetOrganizationPolicyObjects200Response) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *GetOrganizationPolicyObjects200Response) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *GetOrganizationPolicyObjects200Response) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *GetOrganizationPolicyObjects200Response) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetOrganizationPolicyObjects200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetOrganizationPolicyObjects200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetOrganizationPolicyObjects200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetOrganizationPolicyObjects200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetOrganizationPolicyObjects200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetOrganizationPolicyObjects200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetOrganizationPolicyObjects200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetOrganizationPolicyObjects200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetGroupIds

`func (o *GetOrganizationPolicyObjects200Response) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *GetOrganizationPolicyObjects200Response) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *GetOrganizationPolicyObjects200Response) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *GetOrganizationPolicyObjects200Response) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetNetworkIds

`func (o *GetOrganizationPolicyObjects200Response) GetNetworkIds() []string`

GetNetworkIds returns the NetworkIds field if non-nil, zero value otherwise.

### GetNetworkIdsOk

`func (o *GetOrganizationPolicyObjects200Response) GetNetworkIdsOk() (*[]string, bool)`

GetNetworkIdsOk returns a tuple with the NetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIds

`func (o *GetOrganizationPolicyObjects200Response) SetNetworkIds(v []string)`

SetNetworkIds sets NetworkIds field to given value.

### HasNetworkIds

`func (o *GetOrganizationPolicyObjects200Response) HasNetworkIds() bool`

HasNetworkIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

