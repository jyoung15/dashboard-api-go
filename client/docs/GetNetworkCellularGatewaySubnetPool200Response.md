# GetNetworkCellularGatewaySubnetPool200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentMode** | Pointer to **string** | Deployment mode for the cellular gateways in the network. (Passthrough/Routed) | [optional] 
**Cidr** | Pointer to **string** | CIDR of the pool of subnets. Each MG in this network will automatically pick a subnet from this pool. | [optional] 
**Mask** | Pointer to **int32** | Mask used for the subnet of all MGs in  this network. | [optional] 
**Subnets** | Pointer to [**[]GetNetworkCellularGatewaySubnetPool200ResponseSubnetsInner**](GetNetworkCellularGatewaySubnetPool200ResponseSubnetsInner.md) | List of subnets of all MGs in this network. | [optional] 

## Methods

### NewGetNetworkCellularGatewaySubnetPool200Response

`func NewGetNetworkCellularGatewaySubnetPool200Response() *GetNetworkCellularGatewaySubnetPool200Response`

NewGetNetworkCellularGatewaySubnetPool200Response instantiates a new GetNetworkCellularGatewaySubnetPool200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkCellularGatewaySubnetPool200ResponseWithDefaults

`func NewGetNetworkCellularGatewaySubnetPool200ResponseWithDefaults() *GetNetworkCellularGatewaySubnetPool200Response`

NewGetNetworkCellularGatewaySubnetPool200ResponseWithDefaults instantiates a new GetNetworkCellularGatewaySubnetPool200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentMode

`func (o *GetNetworkCellularGatewaySubnetPool200Response) GetDeploymentMode() string`

GetDeploymentMode returns the DeploymentMode field if non-nil, zero value otherwise.

### GetDeploymentModeOk

`func (o *GetNetworkCellularGatewaySubnetPool200Response) GetDeploymentModeOk() (*string, bool)`

GetDeploymentModeOk returns a tuple with the DeploymentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentMode

`func (o *GetNetworkCellularGatewaySubnetPool200Response) SetDeploymentMode(v string)`

SetDeploymentMode sets DeploymentMode field to given value.

### HasDeploymentMode

`func (o *GetNetworkCellularGatewaySubnetPool200Response) HasDeploymentMode() bool`

HasDeploymentMode returns a boolean if a field has been set.

### GetCidr

`func (o *GetNetworkCellularGatewaySubnetPool200Response) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *GetNetworkCellularGatewaySubnetPool200Response) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *GetNetworkCellularGatewaySubnetPool200Response) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *GetNetworkCellularGatewaySubnetPool200Response) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetMask

`func (o *GetNetworkCellularGatewaySubnetPool200Response) GetMask() int32`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *GetNetworkCellularGatewaySubnetPool200Response) GetMaskOk() (*int32, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *GetNetworkCellularGatewaySubnetPool200Response) SetMask(v int32)`

SetMask sets Mask field to given value.

### HasMask

`func (o *GetNetworkCellularGatewaySubnetPool200Response) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetSubnets

`func (o *GetNetworkCellularGatewaySubnetPool200Response) GetSubnets() []GetNetworkCellularGatewaySubnetPool200ResponseSubnetsInner`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *GetNetworkCellularGatewaySubnetPool200Response) GetSubnetsOk() (*[]GetNetworkCellularGatewaySubnetPool200ResponseSubnetsInner, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *GetNetworkCellularGatewaySubnetPool200Response) SetSubnets(v []GetNetworkCellularGatewaySubnetPool200ResponseSubnetsInner)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *GetNetworkCellularGatewaySubnetPool200Response) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


