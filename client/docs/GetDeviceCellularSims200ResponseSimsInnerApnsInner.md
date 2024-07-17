# GetDeviceCellularSims200ResponseSimsInnerApnsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | APN name. | 
**AllowedIpTypes** | **[]string** | IP versions to support (permitted values include &#39;ipv4&#39;, &#39;ipv6&#39;). | 
**Authentication** | Pointer to [**GetDeviceCellularSims200ResponseSimsInnerApnsInnerAuthentication**](GetDeviceCellularSims200ResponseSimsInnerApnsInnerAuthentication.md) |  | [optional] 

## Methods

### NewGetDeviceCellularSims200ResponseSimsInnerApnsInner

`func NewGetDeviceCellularSims200ResponseSimsInnerApnsInner(name string, allowedIpTypes []string, ) *GetDeviceCellularSims200ResponseSimsInnerApnsInner`

NewGetDeviceCellularSims200ResponseSimsInnerApnsInner instantiates a new GetDeviceCellularSims200ResponseSimsInnerApnsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceCellularSims200ResponseSimsInnerApnsInnerWithDefaults

`func NewGetDeviceCellularSims200ResponseSimsInnerApnsInnerWithDefaults() *GetDeviceCellularSims200ResponseSimsInnerApnsInner`

NewGetDeviceCellularSims200ResponseSimsInnerApnsInnerWithDefaults instantiates a new GetDeviceCellularSims200ResponseSimsInnerApnsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetDeviceCellularSims200ResponseSimsInnerApnsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDeviceCellularSims200ResponseSimsInnerApnsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDeviceCellularSims200ResponseSimsInnerApnsInner) SetName(v string)`

SetName sets Name field to given value.


### GetAllowedIpTypes

`func (o *GetDeviceCellularSims200ResponseSimsInnerApnsInner) GetAllowedIpTypes() []string`

GetAllowedIpTypes returns the AllowedIpTypes field if non-nil, zero value otherwise.

### GetAllowedIpTypesOk

`func (o *GetDeviceCellularSims200ResponseSimsInnerApnsInner) GetAllowedIpTypesOk() (*[]string, bool)`

GetAllowedIpTypesOk returns a tuple with the AllowedIpTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIpTypes

`func (o *GetDeviceCellularSims200ResponseSimsInnerApnsInner) SetAllowedIpTypes(v []string)`

SetAllowedIpTypes sets AllowedIpTypes field to given value.


### GetAuthentication

`func (o *GetDeviceCellularSims200ResponseSimsInnerApnsInner) GetAuthentication() GetDeviceCellularSims200ResponseSimsInnerApnsInnerAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *GetDeviceCellularSims200ResponseSimsInnerApnsInner) GetAuthenticationOk() (*GetDeviceCellularSims200ResponseSimsInnerApnsInnerAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *GetDeviceCellularSims200ResponseSimsInnerApnsInner) SetAuthentication(v GetDeviceCellularSims200ResponseSimsInnerApnsInnerAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *GetDeviceCellularSims200ResponseSimsInnerApnsInner) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


