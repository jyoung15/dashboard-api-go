# UpdateNetworkWirelessAirMarshalRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Indicates if this rule will allow, block, or alert. | [optional] 
**Match** | Pointer to [**CreateNetworkWirelessAirMarshalRuleRequestMatch**](CreateNetworkWirelessAirMarshalRuleRequestMatch.md) |  | [optional] 

## Methods

### NewUpdateNetworkWirelessAirMarshalRuleRequest

`func NewUpdateNetworkWirelessAirMarshalRuleRequest() *UpdateNetworkWirelessAirMarshalRuleRequest`

NewUpdateNetworkWirelessAirMarshalRuleRequest instantiates a new UpdateNetworkWirelessAirMarshalRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessAirMarshalRuleRequestWithDefaults

`func NewUpdateNetworkWirelessAirMarshalRuleRequestWithDefaults() *UpdateNetworkWirelessAirMarshalRuleRequest`

NewUpdateNetworkWirelessAirMarshalRuleRequestWithDefaults instantiates a new UpdateNetworkWirelessAirMarshalRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateNetworkWirelessAirMarshalRuleRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateNetworkWirelessAirMarshalRuleRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateNetworkWirelessAirMarshalRuleRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateNetworkWirelessAirMarshalRuleRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMatch

`func (o *UpdateNetworkWirelessAirMarshalRuleRequest) GetMatch() CreateNetworkWirelessAirMarshalRuleRequestMatch`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *UpdateNetworkWirelessAirMarshalRuleRequest) GetMatchOk() (*CreateNetworkWirelessAirMarshalRuleRequestMatch, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *UpdateNetworkWirelessAirMarshalRuleRequest) SetMatch(v CreateNetworkWirelessAirMarshalRuleRequestMatch)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *UpdateNetworkWirelessAirMarshalRuleRequest) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


