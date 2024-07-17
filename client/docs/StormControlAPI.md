# \StormControlAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSwitchStormControl**](StormControlAPI.md#GetNetworkSwitchStormControl) | **Get** /networks/{networkId}/switch/stormControl | Return the storm control configuration for a switch network
[**UpdateNetworkSwitchStormControl**](StormControlAPI.md#UpdateNetworkSwitchStormControl) | **Put** /networks/{networkId}/switch/stormControl | Update the storm control configuration for a switch network



## GetNetworkSwitchStormControl

> GetNetworkSwitchStormControl200Response GetNetworkSwitchStormControl(ctx, networkId).Execute()

Return the storm control configuration for a switch network



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
	resp, r, err := apiClient.StormControlAPI.GetNetworkSwitchStormControl(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StormControlAPI.GetNetworkSwitchStormControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchStormControl`: GetNetworkSwitchStormControl200Response
	fmt.Fprintf(os.Stdout, "Response from `StormControlAPI.GetNetworkSwitchStormControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStormControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSwitchStormControl200Response**](GetNetworkSwitchStormControl200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchStormControl

> GetNetworkSwitchStormControl200Response UpdateNetworkSwitchStormControl(ctx, networkId).UpdateNetworkSwitchStormControlRequest(updateNetworkSwitchStormControlRequest).Execute()

Update the storm control configuration for a switch network



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
	updateNetworkSwitchStormControlRequest := *openapiclient.NewUpdateNetworkSwitchStormControlRequest() // UpdateNetworkSwitchStormControlRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StormControlAPI.UpdateNetworkSwitchStormControl(context.Background(), networkId).UpdateNetworkSwitchStormControlRequest(updateNetworkSwitchStormControlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StormControlAPI.UpdateNetworkSwitchStormControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchStormControl`: GetNetworkSwitchStormControl200Response
	fmt.Fprintf(os.Stdout, "Response from `StormControlAPI.UpdateNetworkSwitchStormControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchStormControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchStormControlRequest** | [**UpdateNetworkSwitchStormControlRequest**](UpdateNetworkSwitchStormControlRequest.md) |  | 

### Return type

[**GetNetworkSwitchStormControl200Response**](GetNetworkSwitchStormControl200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

