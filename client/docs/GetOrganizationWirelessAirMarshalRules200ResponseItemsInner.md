# GetOrganizationWirelessAirMarshalRules200ResponseItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**CreateNetworkWirelessAirMarshalRule201ResponseNetwork**](CreateNetworkWirelessAirMarshalRule201ResponseNetwork.md) |  | [optional] 
**RuleId** | Pointer to **string** | Indicates whether or not clients are allowed to        connect to rogue SSIDs by default. (blocked by default) | [optional] 
**Type** | Pointer to **string** | Indicates whether or not clients are allowed to        connect to rogue SSIDs by default. (blocked by default) | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Updated at timestamp | [optional] 
**CreatedAt** | Pointer to **time.Time** | Created at timestamp | [optional] 
**Match** | Pointer to [**GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch**](GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch.md) |  | [optional] 

## Methods

### NewGetOrganizationWirelessAirMarshalRules200ResponseItemsInner

`func NewGetOrganizationWirelessAirMarshalRules200ResponseItemsInner() *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner`

NewGetOrganizationWirelessAirMarshalRules200ResponseItemsInner instantiates a new GetOrganizationWirelessAirMarshalRules200ResponseItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationWirelessAirMarshalRules200ResponseItemsInnerWithDefaults

`func NewGetOrganizationWirelessAirMarshalRules200ResponseItemsInnerWithDefaults() *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner`

NewGetOrganizationWirelessAirMarshalRules200ResponseItemsInnerWithDefaults instantiates a new GetOrganizationWirelessAirMarshalRules200ResponseItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) GetNetwork() CreateNetworkWirelessAirMarshalRule201ResponseNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) GetNetworkOk() (*CreateNetworkWirelessAirMarshalRule201ResponseNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) SetNetwork(v CreateNetworkWirelessAirMarshalRule201ResponseNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetRuleId

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetType

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMatch

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) GetMatch() GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) GetMatchOk() (*GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) SetMatch(v GetOrganizationWirelessAirMarshalRules200ResponseItemsInnerMatch)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *GetOrganizationWirelessAirMarshalRules200ResponseItemsInner) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


