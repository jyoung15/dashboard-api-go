# GetNetworkSplashLoginAttempts200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | User name | [optional] 
**Login** | Pointer to **string** | User login identifier | [optional] 
**Ssid** | Pointer to **string** | SSID name | [optional] 
**LoginAt** | Pointer to **time.Time** | Login timestamp | [optional] 
**GatewayDeviceMac** | Pointer to **string** | Gateway device mac address | [optional] 
**ClientMac** | Pointer to **string** | Client mac address | [optional] 
**ClientId** | Pointer to **string** | Client ID | [optional] 
**Authorization** | Pointer to **string** | Authorization status | [optional] 

## Methods

### NewGetNetworkSplashLoginAttempts200ResponseInner

`func NewGetNetworkSplashLoginAttempts200ResponseInner() *GetNetworkSplashLoginAttempts200ResponseInner`

NewGetNetworkSplashLoginAttempts200ResponseInner instantiates a new GetNetworkSplashLoginAttempts200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSplashLoginAttempts200ResponseInnerWithDefaults

`func NewGetNetworkSplashLoginAttempts200ResponseInnerWithDefaults() *GetNetworkSplashLoginAttempts200ResponseInner`

NewGetNetworkSplashLoginAttempts200ResponseInnerWithDefaults instantiates a new GetNetworkSplashLoginAttempts200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLogin

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetSsid

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetLoginAt

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetLoginAt() time.Time`

GetLoginAt returns the LoginAt field if non-nil, zero value otherwise.

### GetLoginAtOk

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetLoginAtOk() (*time.Time, bool)`

GetLoginAtOk returns a tuple with the LoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAt

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) SetLoginAt(v time.Time)`

SetLoginAt sets LoginAt field to given value.

### HasLoginAt

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) HasLoginAt() bool`

HasLoginAt returns a boolean if a field has been set.

### GetGatewayDeviceMac

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetGatewayDeviceMac() string`

GetGatewayDeviceMac returns the GatewayDeviceMac field if non-nil, zero value otherwise.

### GetGatewayDeviceMacOk

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetGatewayDeviceMacOk() (*string, bool)`

GetGatewayDeviceMacOk returns a tuple with the GatewayDeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayDeviceMac

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) SetGatewayDeviceMac(v string)`

SetGatewayDeviceMac sets GatewayDeviceMac field to given value.

### HasGatewayDeviceMac

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) HasGatewayDeviceMac() bool`

HasGatewayDeviceMac returns a boolean if a field has been set.

### GetClientMac

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetClientMac() string`

GetClientMac returns the ClientMac field if non-nil, zero value otherwise.

### GetClientMacOk

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetClientMacOk() (*string, bool)`

GetClientMacOk returns a tuple with the ClientMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMac

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) SetClientMac(v string)`

SetClientMac sets ClientMac field to given value.

### HasClientMac

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) HasClientMac() bool`

HasClientMac returns a boolean if a field has been set.

### GetClientId

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetAuthorization

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetAuthorization() string`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) GetAuthorizationOk() (*string, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) SetAuthorization(v string)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *GetNetworkSplashLoginAttempts200ResponseInner) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


