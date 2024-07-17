# GetOrganizationPolicyObjectsGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Policy object ID | [optional] 
**Name** | Pointer to **string** | Name of the Policy object group. | [optional] 
**Category** | Pointer to **string** | Type of object groups. (NetworkObjectGroup, GeoLocationGroup, PortObjectGroup, ApplicationGroup) | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time Stamp of policy object creation. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time Stamp of policy object updation. | [optional] 
**ObjectIds** | Pointer to **[]int32** | Policy objects associated with Network Object Group or Port Object Group | [optional] 
**NetworkIds** | Pointer to **[]string** | Network ID&#39;s associated with the policy objects. | [optional] 

## Methods

### NewGetOrganizationPolicyObjectsGroups200Response

`func NewGetOrganizationPolicyObjectsGroups200Response() *GetOrganizationPolicyObjectsGroups200Response`

NewGetOrganizationPolicyObjectsGroups200Response instantiates a new GetOrganizationPolicyObjectsGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationPolicyObjectsGroups200ResponseWithDefaults

`func NewGetOrganizationPolicyObjectsGroups200ResponseWithDefaults() *GetOrganizationPolicyObjectsGroups200Response`

NewGetOrganizationPolicyObjectsGroups200ResponseWithDefaults instantiates a new GetOrganizationPolicyObjectsGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOrganizationPolicyObjectsGroups200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrganizationPolicyObjectsGroups200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrganizationPolicyObjectsGroups200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetOrganizationPolicyObjectsGroups200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetOrganizationPolicyObjectsGroups200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationPolicyObjectsGroups200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationPolicyObjectsGroups200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationPolicyObjectsGroups200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *GetOrganizationPolicyObjectsGroups200Response) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetOrganizationPolicyObjectsGroups200Response) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetOrganizationPolicyObjectsGroups200Response) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetOrganizationPolicyObjectsGroups200Response) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetOrganizationPolicyObjectsGroups200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetOrganizationPolicyObjectsGroups200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetOrganizationPolicyObjectsGroups200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetOrganizationPolicyObjectsGroups200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetOrganizationPolicyObjectsGroups200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetOrganizationPolicyObjectsGroups200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetOrganizationPolicyObjectsGroups200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetOrganizationPolicyObjectsGroups200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetObjectIds

`func (o *GetOrganizationPolicyObjectsGroups200Response) GetObjectIds() []int32`

GetObjectIds returns the ObjectIds field if non-nil, zero value otherwise.

### GetObjectIdsOk

`func (o *GetOrganizationPolicyObjectsGroups200Response) GetObjectIdsOk() (*[]int32, bool)`

GetObjectIdsOk returns a tuple with the ObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIds

`func (o *GetOrganizationPolicyObjectsGroups200Response) SetObjectIds(v []int32)`

SetObjectIds sets ObjectIds field to given value.

### HasObjectIds

`func (o *GetOrganizationPolicyObjectsGroups200Response) HasObjectIds() bool`

HasObjectIds returns a boolean if a field has been set.

### GetNetworkIds

`func (o *GetOrganizationPolicyObjectsGroups200Response) GetNetworkIds() []string`

GetNetworkIds returns the NetworkIds field if non-nil, zero value otherwise.

### GetNetworkIdsOk

`func (o *GetOrganizationPolicyObjectsGroups200Response) GetNetworkIdsOk() (*[]string, bool)`

GetNetworkIdsOk returns a tuple with the NetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIds

`func (o *GetOrganizationPolicyObjectsGroups200Response) SetNetworkIds(v []string)`

SetNetworkIds sets NetworkIds field to given value.

### HasNetworkIds

`func (o *GetOrganizationPolicyObjectsGroups200Response) HasNetworkIds() bool`

HasNetworkIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


