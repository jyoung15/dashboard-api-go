# GetNetworkMqttBrokers200ResponseInnerSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Security protocol of the MQTT broker. | [optional] 
**Tls** | Pointer to [**GetNetworkMqttBrokers200ResponseInnerSecurityTls**](GetNetworkMqttBrokers200ResponseInnerSecurityTls.md) |  | [optional] 

## Methods

### NewGetNetworkMqttBrokers200ResponseInnerSecurity

`func NewGetNetworkMqttBrokers200ResponseInnerSecurity() *GetNetworkMqttBrokers200ResponseInnerSecurity`

NewGetNetworkMqttBrokers200ResponseInnerSecurity instantiates a new GetNetworkMqttBrokers200ResponseInnerSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkMqttBrokers200ResponseInnerSecurityWithDefaults

`func NewGetNetworkMqttBrokers200ResponseInnerSecurityWithDefaults() *GetNetworkMqttBrokers200ResponseInnerSecurity`

NewGetNetworkMqttBrokers200ResponseInnerSecurityWithDefaults instantiates a new GetNetworkMqttBrokers200ResponseInnerSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *GetNetworkMqttBrokers200ResponseInnerSecurity) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GetNetworkMqttBrokers200ResponseInnerSecurity) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GetNetworkMqttBrokers200ResponseInnerSecurity) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GetNetworkMqttBrokers200ResponseInnerSecurity) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetTls

`func (o *GetNetworkMqttBrokers200ResponseInnerSecurity) GetTls() GetNetworkMqttBrokers200ResponseInnerSecurityTls`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *GetNetworkMqttBrokers200ResponseInnerSecurity) GetTlsOk() (*GetNetworkMqttBrokers200ResponseInnerSecurityTls, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *GetNetworkMqttBrokers200ResponseInnerSecurity) SetTls(v GetNetworkMqttBrokers200ResponseInnerSecurityTls)`

SetTls sets Tls field to given value.

### HasTls

`func (o *GetNetworkMqttBrokers200ResponseInnerSecurity) HasTls() bool`

HasTls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


