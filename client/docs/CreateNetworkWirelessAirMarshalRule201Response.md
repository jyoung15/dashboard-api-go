# CreateNetworkWirelessAirMarshalRule201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**CreateNetworkWirelessAirMarshalRule201ResponseNetwork**](CreateNetworkWirelessAirMarshalRule201ResponseNetwork.md) |  | [optional] 
**RuleId** | Pointer to **string** | Indicates whether or not clients are allowed to       connect to rogue SSIDs by default. (blocked by default) | [optional] 
**Type** | Pointer to **string** | Indicates whether or not clients are allowed to       connect to rogue SSIDs by default. (blocked by default) | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Updated at timestamp | [optional] 
**CreatedAt** | Pointer to **time.Time** | Created at timestamp | [optional] 
**Match** | Pointer to [**CreateNetworkWirelessAirMarshalRule201ResponseMatch**](CreateNetworkWirelessAirMarshalRule201ResponseMatch.md) |  | [optional] 

## Methods

### NewCreateNetworkWirelessAirMarshalRule201Response

`func NewCreateNetworkWirelessAirMarshalRule201Response() *CreateNetworkWirelessAirMarshalRule201Response`

NewCreateNetworkWirelessAirMarshalRule201Response instantiates a new CreateNetworkWirelessAirMarshalRule201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkWirelessAirMarshalRule201ResponseWithDefaults

`func NewCreateNetworkWirelessAirMarshalRule201ResponseWithDefaults() *CreateNetworkWirelessAirMarshalRule201Response`

NewCreateNetworkWirelessAirMarshalRule201ResponseWithDefaults instantiates a new CreateNetworkWirelessAirMarshalRule201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *CreateNetworkWirelessAirMarshalRule201Response) GetNetwork() CreateNetworkWirelessAirMarshalRule201ResponseNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CreateNetworkWirelessAirMarshalRule201Response) GetNetworkOk() (*CreateNetworkWirelessAirMarshalRule201ResponseNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CreateNetworkWirelessAirMarshalRule201Response) SetNetwork(v CreateNetworkWirelessAirMarshalRule201ResponseNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CreateNetworkWirelessAirMarshalRule201Response) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetRuleId

`func (o *CreateNetworkWirelessAirMarshalRule201Response) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *CreateNetworkWirelessAirMarshalRule201Response) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *CreateNetworkWirelessAirMarshalRule201Response) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *CreateNetworkWirelessAirMarshalRule201Response) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetType

`func (o *CreateNetworkWirelessAirMarshalRule201Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateNetworkWirelessAirMarshalRule201Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateNetworkWirelessAirMarshalRule201Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateNetworkWirelessAirMarshalRule201Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CreateNetworkWirelessAirMarshalRule201Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateNetworkWirelessAirMarshalRule201Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateNetworkWirelessAirMarshalRule201Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreateNetworkWirelessAirMarshalRule201Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateNetworkWirelessAirMarshalRule201Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateNetworkWirelessAirMarshalRule201Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateNetworkWirelessAirMarshalRule201Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateNetworkWirelessAirMarshalRule201Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMatch

`func (o *CreateNetworkWirelessAirMarshalRule201Response) GetMatch() CreateNetworkWirelessAirMarshalRule201ResponseMatch`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *CreateNetworkWirelessAirMarshalRule201Response) GetMatchOk() (*CreateNetworkWirelessAirMarshalRule201ResponseMatch, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *CreateNetworkWirelessAirMarshalRule201Response) SetMatch(v CreateNetworkWirelessAirMarshalRule201ResponseMatch)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *CreateNetworkWirelessAirMarshalRule201Response) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


