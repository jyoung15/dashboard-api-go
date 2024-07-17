# \NetflowAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkNetflow**](NetflowAPI.md#GetNetworkNetflow) | **Get** /networks/{networkId}/netflow | Return the NetFlow traffic reporting settings for a network
[**UpdateNetworkNetflow**](NetflowAPI.md#UpdateNetworkNetflow) | **Put** /networks/{networkId}/netflow | Update the NetFlow traffic reporting settings for a network



## GetNetworkNetflow

> GetNetworkNetflow200Response GetNetworkNetflow(ctx, networkId).Execute()

Return the NetFlow traffic reporting settings for a network



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
	resp, r, err := apiClient.NetflowAPI.GetNetworkNetflow(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetflowAPI.GetNetworkNetflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkNetflow`: GetNetworkNetflow200Response
	fmt.Fprintf(os.Stdout, "Response from `NetflowAPI.GetNetworkNetflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkNetflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkNetflow200Response**](GetNetworkNetflow200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkNetflow

> GetNetworkNetflow200Response UpdateNetworkNetflow(ctx, networkId).UpdateNetworkNetflowRequest(updateNetworkNetflowRequest).Execute()

Update the NetFlow traffic reporting settings for a network



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
	updateNetworkNetflowRequest := *openapiclient.NewUpdateNetworkNetflowRequest() // UpdateNetworkNetflowRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetflowAPI.UpdateNetworkNetflow(context.Background(), networkId).UpdateNetworkNetflowRequest(updateNetworkNetflowRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetflowAPI.UpdateNetworkNetflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkNetflow`: GetNetworkNetflow200Response
	fmt.Fprintf(os.Stdout, "Response from `NetflowAPI.UpdateNetworkNetflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkNetflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkNetflowRequest** | [**UpdateNetworkNetflowRequest**](UpdateNetworkNetflowRequest.md) |  | 

### Return type

[**GetNetworkNetflow200Response**](GetNetworkNetflow200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

