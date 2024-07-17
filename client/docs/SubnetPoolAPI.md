# \SubnetPoolAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkCellularGatewaySubnetPool**](SubnetPoolAPI.md#GetNetworkCellularGatewaySubnetPool) | **Get** /networks/{networkId}/cellularGateway/subnetPool | Return the subnet pool and mask configured for MGs in the network.
[**UpdateNetworkCellularGatewaySubnetPool**](SubnetPoolAPI.md#UpdateNetworkCellularGatewaySubnetPool) | **Put** /networks/{networkId}/cellularGateway/subnetPool | Update the subnet pool and mask configuration for MGs in the network.



## GetNetworkCellularGatewaySubnetPool

> GetNetworkCellularGatewaySubnetPool200Response GetNetworkCellularGatewaySubnetPool(ctx, networkId).Execute()

Return the subnet pool and mask configured for MGs in the network.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetPoolAPI.GetNetworkCellularGatewaySubnetPool(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetPoolAPI.GetNetworkCellularGatewaySubnetPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkCellularGatewaySubnetPool`: GetNetworkCellularGatewaySubnetPool200Response
	fmt.Fprintf(os.Stdout, "Response from `SubnetPoolAPI.GetNetworkCellularGatewaySubnetPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCellularGatewaySubnetPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkCellularGatewaySubnetPool200Response**](GetNetworkCellularGatewaySubnetPool200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkCellularGatewaySubnetPool

> GetNetworkCellularGatewaySubnetPool200Response UpdateNetworkCellularGatewaySubnetPool(ctx, networkId).UpdateNetworkCellularGatewaySubnetPoolRequest(updateNetworkCellularGatewaySubnetPoolRequest).Execute()

Update the subnet pool and mask configuration for MGs in the network.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkCellularGatewaySubnetPoolRequest := *openapiclient.NewUpdateNetworkCellularGatewaySubnetPoolRequest() // UpdateNetworkCellularGatewaySubnetPoolRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetPoolAPI.UpdateNetworkCellularGatewaySubnetPool(context.Background(), networkId).UpdateNetworkCellularGatewaySubnetPoolRequest(updateNetworkCellularGatewaySubnetPoolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetPoolAPI.UpdateNetworkCellularGatewaySubnetPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkCellularGatewaySubnetPool`: GetNetworkCellularGatewaySubnetPool200Response
	fmt.Fprintf(os.Stdout, "Response from `SubnetPoolAPI.UpdateNetworkCellularGatewaySubnetPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkCellularGatewaySubnetPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkCellularGatewaySubnetPoolRequest** | [**UpdateNetworkCellularGatewaySubnetPoolRequest**](UpdateNetworkCellularGatewaySubnetPoolRequest.md) |  | 

### Return type

[**GetNetworkCellularGatewaySubnetPool200Response**](GetNetworkCellularGatewaySubnetPool200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

