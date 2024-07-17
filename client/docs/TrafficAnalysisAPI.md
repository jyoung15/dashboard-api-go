# \TrafficAnalysisAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkTrafficAnalysis**](TrafficAnalysisAPI.md#GetNetworkTrafficAnalysis) | **Get** /networks/{networkId}/trafficAnalysis | Return the traffic analysis settings for a network
[**UpdateNetworkTrafficAnalysis**](TrafficAnalysisAPI.md#UpdateNetworkTrafficAnalysis) | **Put** /networks/{networkId}/trafficAnalysis | Update the traffic analysis settings for a network



## GetNetworkTrafficAnalysis

> GetNetworkTrafficAnalysis200Response GetNetworkTrafficAnalysis(ctx, networkId).Execute()

Return the traffic analysis settings for a network



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
	resp, r, err := apiClient.TrafficAnalysisAPI.GetNetworkTrafficAnalysis(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrafficAnalysisAPI.GetNetworkTrafficAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkTrafficAnalysis`: GetNetworkTrafficAnalysis200Response
	fmt.Fprintf(os.Stdout, "Response from `TrafficAnalysisAPI.GetNetworkTrafficAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkTrafficAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkTrafficAnalysis200Response**](GetNetworkTrafficAnalysis200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkTrafficAnalysis

> GetNetworkTrafficAnalysis200Response UpdateNetworkTrafficAnalysis(ctx, networkId).UpdateNetworkTrafficAnalysisRequest(updateNetworkTrafficAnalysisRequest).Execute()

Update the traffic analysis settings for a network



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
	updateNetworkTrafficAnalysisRequest := *openapiclient.NewUpdateNetworkTrafficAnalysisRequest() // UpdateNetworkTrafficAnalysisRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrafficAnalysisAPI.UpdateNetworkTrafficAnalysis(context.Background(), networkId).UpdateNetworkTrafficAnalysisRequest(updateNetworkTrafficAnalysisRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrafficAnalysisAPI.UpdateNetworkTrafficAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkTrafficAnalysis`: GetNetworkTrafficAnalysis200Response
	fmt.Fprintf(os.Stdout, "Response from `TrafficAnalysisAPI.UpdateNetworkTrafficAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkTrafficAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkTrafficAnalysisRequest** | [**UpdateNetworkTrafficAnalysisRequest**](UpdateNetworkTrafficAnalysisRequest.md) |  | 

### Return type

[**GetNetworkTrafficAnalysis200Response**](GetNetworkTrafficAnalysis200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

