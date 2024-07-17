# \MtuAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSwitchMtu**](MtuAPI.md#GetNetworkSwitchMtu) | **Get** /networks/{networkId}/switch/mtu | Return the MTU configuration
[**UpdateNetworkSwitchMtu**](MtuAPI.md#UpdateNetworkSwitchMtu) | **Put** /networks/{networkId}/switch/mtu | Update the MTU configuration



## GetNetworkSwitchMtu

> GetNetworkSwitchMtu200Response GetNetworkSwitchMtu(ctx, networkId).Execute()

Return the MTU configuration



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
	resp, r, err := apiClient.MtuAPI.GetNetworkSwitchMtu(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MtuAPI.GetNetworkSwitchMtu``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchMtu`: GetNetworkSwitchMtu200Response
	fmt.Fprintf(os.Stdout, "Response from `MtuAPI.GetNetworkSwitchMtu`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchMtuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSwitchMtu200Response**](GetNetworkSwitchMtu200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchMtu

> GetNetworkSwitchMtu200Response UpdateNetworkSwitchMtu(ctx, networkId).UpdateNetworkSwitchMtuRequest(updateNetworkSwitchMtuRequest).Execute()

Update the MTU configuration



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
	updateNetworkSwitchMtuRequest := *openapiclient.NewUpdateNetworkSwitchMtuRequest() // UpdateNetworkSwitchMtuRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MtuAPI.UpdateNetworkSwitchMtu(context.Background(), networkId).UpdateNetworkSwitchMtuRequest(updateNetworkSwitchMtuRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MtuAPI.UpdateNetworkSwitchMtu``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchMtu`: GetNetworkSwitchMtu200Response
	fmt.Fprintf(os.Stdout, "Response from `MtuAPI.UpdateNetworkSwitchMtu`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchMtuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchMtuRequest** | [**UpdateNetworkSwitchMtuRequest**](UpdateNetworkSwitchMtuRequest.md) |  | 

### Return type

[**GetNetworkSwitchMtu200Response**](GetNetworkSwitchMtu200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

