# CreateNetworkWirelessAirMarshalRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Indicates if this rule will allow, block, or alert. | 
**Match** | [**CreateNetworkWirelessAirMarshalRuleRequestMatch**](CreateNetworkWirelessAirMarshalRuleRequestMatch.md) |  | 

## Methods

### NewCreateNetworkWirelessAirMarshalRuleRequest

`func NewCreateNetworkWirelessAirMarshalRuleRequest(type_ string, match CreateNetworkWirelessAirMarshalRuleRequestMatch, ) *CreateNetworkWirelessAirMarshalRuleRequest`

NewCreateNetworkWirelessAirMarshalRuleRequest instantiates a new CreateNetworkWirelessAirMarshalRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkWirelessAirMarshalRuleRequestWithDefaults

`func NewCreateNetworkWirelessAirMarshalRuleRequestWithDefaults() *CreateNetworkWirelessAirMarshalRuleRequest`

NewCreateNetworkWirelessAirMarshalRuleRequestWithDefaults instantiates a new CreateNetworkWirelessAirMarshalRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateNetworkWirelessAirMarshalRuleRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateNetworkWirelessAirMarshalRuleRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateNetworkWirelessAirMarshalRuleRequest) SetType(v string)`

SetType sets Type field to given value.


### GetMatch

`func (o *CreateNetworkWirelessAirMarshalRuleRequest) GetMatch() CreateNetworkWirelessAirMarshalRuleRequestMatch`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *CreateNetworkWirelessAirMarshalRuleRequest) GetMatchOk() (*CreateNetworkWirelessAirMarshalRuleRequestMatch, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *CreateNetworkWirelessAirMarshalRuleRequest) SetMatch(v CreateNetworkWirelessAirMarshalRuleRequestMatch)`

SetMatch sets Match field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


