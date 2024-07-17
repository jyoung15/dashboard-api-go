# GetNetworkMqttBrokers200ResponseInnerSecurityTls

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasCaCertificate** | Pointer to **bool** | Indicates whether the CA certificate is set | [optional] 
**VerifyHostnames** | Pointer to **bool** | Whether the TLS hostname verification is enabled for the MQTT broker. | [optional] 

## Methods

### NewGetNetworkMqttBrokers200ResponseInnerSecurityTls

`func NewGetNetworkMqttBrokers200ResponseInnerSecurityTls() *GetNetworkMqttBrokers200ResponseInnerSecurityTls`

NewGetNetworkMqttBrokers200ResponseInnerSecurityTls instantiates a new GetNetworkMqttBrokers200ResponseInnerSecurityTls object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkMqttBrokers200ResponseInnerSecurityTlsWithDefaults

`func NewGetNetworkMqttBrokers200ResponseInnerSecurityTlsWithDefaults() *GetNetworkMqttBrokers200ResponseInnerSecurityTls`

NewGetNetworkMqttBrokers200ResponseInnerSecurityTlsWithDefaults instantiates a new GetNetworkMqttBrokers200ResponseInnerSecurityTls object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasCaCertificate

`func (o *GetNetworkMqttBrokers200ResponseInnerSecurityTls) GetHasCaCertificate() bool`

GetHasCaCertificate returns the HasCaCertificate field if non-nil, zero value otherwise.

### GetHasCaCertificateOk

`func (o *GetNetworkMqttBrokers200ResponseInnerSecurityTls) GetHasCaCertificateOk() (*bool, bool)`

GetHasCaCertificateOk returns a tuple with the HasCaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCaCertificate

`func (o *GetNetworkMqttBrokers200ResponseInnerSecurityTls) SetHasCaCertificate(v bool)`

SetHasCaCertificate sets HasCaCertificate field to given value.

### HasHasCaCertificate

`func (o *GetNetworkMqttBrokers200ResponseInnerSecurityTls) HasHasCaCertificate() bool`

HasHasCaCertificate returns a boolean if a field has been set.

### GetVerifyHostnames

`func (o *GetNetworkMqttBrokers200ResponseInnerSecurityTls) GetVerifyHostnames() bool`

GetVerifyHostnames returns the VerifyHostnames field if non-nil, zero value otherwise.

### GetVerifyHostnamesOk

`func (o *GetNetworkMqttBrokers200ResponseInnerSecurityTls) GetVerifyHostnamesOk() (*bool, bool)`

GetVerifyHostnamesOk returns a tuple with the VerifyHostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyHostnames

`func (o *GetNetworkMqttBrokers200ResponseInnerSecurityTls) SetVerifyHostnames(v bool)`

SetVerifyHostnames sets VerifyHostnames field to given value.

### HasVerifyHostnames

`func (o *GetNetworkMqttBrokers200ResponseInnerSecurityTls) HasVerifyHostnames() bool`

HasVerifyHostnames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


