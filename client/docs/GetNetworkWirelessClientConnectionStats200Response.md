# GetNetworkWirelessClientConnectionStats200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | MAC address of the client | [optional] 
**ConnectionStats** | Pointer to [**GetNetworkWirelessClientConnectionStats200ResponseConnectionStats**](GetNetworkWirelessClientConnectionStats200ResponseConnectionStats.md) |  | [optional] 

## Methods

### NewGetNetworkWirelessClientConnectionStats200Response

`func NewGetNetworkWirelessClientConnectionStats200Response() *GetNetworkWirelessClientConnectionStats200Response`

NewGetNetworkWirelessClientConnectionStats200Response instantiates a new GetNetworkWirelessClientConnectionStats200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkWirelessClientConnectionStats200ResponseWithDefaults

`func NewGetNetworkWirelessClientConnectionStats200ResponseWithDefaults() *GetNetworkWirelessClientConnectionStats200Response`

NewGetNetworkWirelessClientConnectionStats200ResponseWithDefaults instantiates a new GetNetworkWirelessClientConnectionStats200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *GetNetworkWirelessClientConnectionStats200Response) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetNetworkWirelessClientConnectionStats200Response) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetNetworkWirelessClientConnectionStats200Response) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetNetworkWirelessClientConnectionStats200Response) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetConnectionStats

`func (o *GetNetworkWirelessClientConnectionStats200Response) GetConnectionStats() GetNetworkWirelessClientConnectionStats200ResponseConnectionStats`

GetConnectionStats returns the ConnectionStats field if non-nil, zero value otherwise.

### GetConnectionStatsOk

`func (o *GetNetworkWirelessClientConnectionStats200Response) GetConnectionStatsOk() (*GetNetworkWirelessClientConnectionStats200ResponseConnectionStats, bool)`

GetConnectionStatsOk returns a tuple with the ConnectionStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStats

`func (o *GetNetworkWirelessClientConnectionStats200Response) SetConnectionStats(v GetNetworkWirelessClientConnectionStats200ResponseConnectionStats)`

SetConnectionStats sets ConnectionStats field to given value.

### HasConnectionStats

`func (o *GetNetworkWirelessClientConnectionStats200Response) HasConnectionStats() bool`

HasConnectionStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


