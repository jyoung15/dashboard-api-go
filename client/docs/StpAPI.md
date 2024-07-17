# \StpAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSwitchStp**](StpAPI.md#GetNetworkSwitchStp) | **Get** /networks/{networkId}/switch/stp | Returns STP settings
[**UpdateNetworkSwitchStp**](StpAPI.md#UpdateNetworkSwitchStp) | **Put** /networks/{networkId}/switch/stp | Updates STP settings



## GetNetworkSwitchStp

> GetNetworkSwitchStp200Response GetNetworkSwitchStp(ctx, networkId).Execute()

Returns STP settings



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
	resp, r, err := apiClient.StpAPI.GetNetworkSwitchStp(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StpAPI.GetNetworkSwitchStp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchStp`: GetNetworkSwitchStp200Response
	fmt.Fprintf(os.Stdout, "Response from `StpAPI.GetNetworkSwitchStp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSwitchStp200Response**](GetNetworkSwitchStp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchStp

> GetNetworkSwitchStp200Response UpdateNetworkSwitchStp(ctx, networkId).UpdateNetworkSwitchStpRequest(updateNetworkSwitchStpRequest).Execute()

Updates STP settings



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
	updateNetworkSwitchStpRequest := *openapiclient.NewUpdateNetworkSwitchStpRequest() // UpdateNetworkSwitchStpRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StpAPI.UpdateNetworkSwitchStp(context.Background(), networkId).UpdateNetworkSwitchStpRequest(updateNetworkSwitchStpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StpAPI.UpdateNetworkSwitchStp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchStp`: GetNetworkSwitchStp200Response
	fmt.Fprintf(os.Stdout, "Response from `StpAPI.UpdateNetworkSwitchStp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchStpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchStpRequest** | [**UpdateNetworkSwitchStpRequest**](UpdateNetworkSwitchStpRequest.md) |  | 

### Return type

[**GetNetworkSwitchStp200Response**](GetNetworkSwitchStp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

