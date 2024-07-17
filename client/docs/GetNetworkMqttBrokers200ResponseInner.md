# GetNetworkMqttBrokers200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the MQTT Broker. | [optional] 
**Name** | Pointer to **string** | Name of the MQTT Broker. | [optional] 
**Host** | Pointer to **string** | Host name/IP address where the MQTT broker runs. | [optional] 
**Port** | Pointer to **int32** | Host port though which the MQTT broker can be reached. | [optional] 
**Security** | Pointer to [**GetNetworkMqttBrokers200ResponseInnerSecurity**](GetNetworkMqttBrokers200ResponseInnerSecurity.md) |  | [optional] 
**Authentication** | Pointer to [**GetNetworkMqttBrokers200ResponseInnerAuthentication**](GetNetworkMqttBrokers200ResponseInnerAuthentication.md) |  | [optional] 

## Methods

### NewGetNetworkMqttBrokers200ResponseInner

`func NewGetNetworkMqttBrokers200ResponseInner() *GetNetworkMqttBrokers200ResponseInner`

NewGetNetworkMqttBrokers200ResponseInner instantiates a new GetNetworkMqttBrokers200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkMqttBrokers200ResponseInnerWithDefaults

`func NewGetNetworkMqttBrokers200ResponseInnerWithDefaults() *GetNetworkMqttBrokers200ResponseInner`

NewGetNetworkMqttBrokers200ResponseInnerWithDefaults instantiates a new GetNetworkMqttBrokers200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkMqttBrokers200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkMqttBrokers200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkMqttBrokers200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkMqttBrokers200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkMqttBrokers200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkMqttBrokers200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkMqttBrokers200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkMqttBrokers200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHost

`func (o *GetNetworkMqttBrokers200ResponseInner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GetNetworkMqttBrokers200ResponseInner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GetNetworkMqttBrokers200ResponseInner) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *GetNetworkMqttBrokers200ResponseInner) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *GetNetworkMqttBrokers200ResponseInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetNetworkMqttBrokers200ResponseInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetNetworkMqttBrokers200ResponseInner) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetNetworkMqttBrokers200ResponseInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSecurity

`func (o *GetNetworkMqttBrokers200ResponseInner) GetSecurity() GetNetworkMqttBrokers200ResponseInnerSecurity`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *GetNetworkMqttBrokers200ResponseInner) GetSecurityOk() (*GetNetworkMqttBrokers200ResponseInnerSecurity, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *GetNetworkMqttBrokers200ResponseInner) SetSecurity(v GetNetworkMqttBrokers200ResponseInnerSecurity)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *GetNetworkMqttBrokers200ResponseInner) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetAuthentication

`func (o *GetNetworkMqttBrokers200ResponseInner) GetAuthentication() GetNetworkMqttBrokers200ResponseInnerAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *GetNetworkMqttBrokers200ResponseInner) GetAuthenticationOk() (*GetNetworkMqttBrokers200ResponseInnerAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *GetNetworkMqttBrokers200ResponseInner) SetAuthentication(v GetNetworkMqttBrokers200ResponseInnerAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *GetNetworkMqttBrokers200ResponseInner) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


