# GetNetworkClientPolicy200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | The MAC address of the client | [optional] 
**DevicePolicy** | Pointer to **string** | The name of the client&#39;s policy | [optional] 
**GroupPolicyId** | Pointer to **string** | The group policy identifier of the client | [optional] 

## Methods

### NewGetNetworkClientPolicy200Response

`func NewGetNetworkClientPolicy200Response() *GetNetworkClientPolicy200Response`

NewGetNetworkClientPolicy200Response instantiates a new GetNetworkClientPolicy200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkClientPolicy200ResponseWithDefaults

`func NewGetNetworkClientPolicy200ResponseWithDefaults() *GetNetworkClientPolicy200Response`

NewGetNetworkClientPolicy200ResponseWithDefaults instantiates a new GetNetworkClientPolicy200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *GetNetworkClientPolicy200Response) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetNetworkClientPolicy200Response) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetNetworkClientPolicy200Response) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetNetworkClientPolicy200Response) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetDevicePolicy

`func (o *GetNetworkClientPolicy200Response) GetDevicePolicy() string`

GetDevicePolicy returns the DevicePolicy field if non-nil, zero value otherwise.

### GetDevicePolicyOk

`func (o *GetNetworkClientPolicy200Response) GetDevicePolicyOk() (*string, bool)`

GetDevicePolicyOk returns a tuple with the DevicePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePolicy

`func (o *GetNetworkClientPolicy200Response) SetDevicePolicy(v string)`

SetDevicePolicy sets DevicePolicy field to given value.

### HasDevicePolicy

`func (o *GetNetworkClientPolicy200Response) HasDevicePolicy() bool`

HasDevicePolicy returns a boolean if a field has been set.

### GetGroupPolicyId

`func (o *GetNetworkClientPolicy200Response) GetGroupPolicyId() string`

GetGroupPolicyId returns the GroupPolicyId field if non-nil, zero value otherwise.

### GetGroupPolicyIdOk

`func (o *GetNetworkClientPolicy200Response) GetGroupPolicyIdOk() (*string, bool)`

GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPolicyId

`func (o *GetNetworkClientPolicy200Response) SetGroupPolicyId(v string)`

SetGroupPolicyId sets GroupPolicyId field to given value.

### HasGroupPolicyId

`func (o *GetNetworkClientPolicy200Response) HasGroupPolicyId() bool`

HasGroupPolicyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


