# \LinkAggregationsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkSwitchLinkAggregation**](LinkAggregationsAPI.md#CreateNetworkSwitchLinkAggregation) | **Post** /networks/{networkId}/switch/linkAggregations | Create a link aggregation group
[**DeleteNetworkSwitchLinkAggregation**](LinkAggregationsAPI.md#DeleteNetworkSwitchLinkAggregation) | **Delete** /networks/{networkId}/switch/linkAggregations/{linkAggregationId} | Split a link aggregation group into separate ports
[**GetNetworkSwitchLinkAggregations**](LinkAggregationsAPI.md#GetNetworkSwitchLinkAggregations) | **Get** /networks/{networkId}/switch/linkAggregations | List link aggregation groups
[**UpdateNetworkSwitchLinkAggregation**](LinkAggregationsAPI.md#UpdateNetworkSwitchLinkAggregation) | **Put** /networks/{networkId}/switch/linkAggregations/{linkAggregationId} | Update a link aggregation group



## CreateNetworkSwitchLinkAggregation

> GetNetworkSwitchLinkAggregations200ResponseInner CreateNetworkSwitchLinkAggregation(ctx, networkId).CreateNetworkSwitchLinkAggregationRequest(createNetworkSwitchLinkAggregationRequest).Execute()

Create a link aggregation group



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
	createNetworkSwitchLinkAggregationRequest := *openapiclient.NewCreateNetworkSwitchLinkAggregationRequest() // CreateNetworkSwitchLinkAggregationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinkAggregationsAPI.CreateNetworkSwitchLinkAggregation(context.Background(), networkId).CreateNetworkSwitchLinkAggregationRequest(createNetworkSwitchLinkAggregationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinkAggregationsAPI.CreateNetworkSwitchLinkAggregation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSwitchLinkAggregation`: GetNetworkSwitchLinkAggregations200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `LinkAggregationsAPI.CreateNetworkSwitchLinkAggregation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchLinkAggregationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSwitchLinkAggregationRequest** | [**CreateNetworkSwitchLinkAggregationRequest**](CreateNetworkSwitchLinkAggregationRequest.md) |  | 

### Return type

[**GetNetworkSwitchLinkAggregations200ResponseInner**](GetNetworkSwitchLinkAggregations200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSwitchLinkAggregation

> DeleteNetworkSwitchLinkAggregation(ctx, networkId, linkAggregationId).Execute()

Split a link aggregation group into separate ports



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
	linkAggregationId := "linkAggregationId_example" // string | Link aggregation ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LinkAggregationsAPI.DeleteNetworkSwitchLinkAggregation(context.Background(), networkId, linkAggregationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinkAggregationsAPI.DeleteNetworkSwitchLinkAggregation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**linkAggregationId** | **string** | Link aggregation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchLinkAggregationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchLinkAggregations

> []GetNetworkSwitchLinkAggregations200ResponseInner GetNetworkSwitchLinkAggregations(ctx, networkId).Execute()

List link aggregation groups



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
	resp, r, err := apiClient.LinkAggregationsAPI.GetNetworkSwitchLinkAggregations(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinkAggregationsAPI.GetNetworkSwitchLinkAggregations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchLinkAggregations`: []GetNetworkSwitchLinkAggregations200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `LinkAggregationsAPI.GetNetworkSwitchLinkAggregations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchLinkAggregationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkSwitchLinkAggregations200ResponseInner**](GetNetworkSwitchLinkAggregations200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchLinkAggregation

> GetNetworkSwitchLinkAggregations200ResponseInner UpdateNetworkSwitchLinkAggregation(ctx, networkId, linkAggregationId).UpdateNetworkSwitchLinkAggregationRequest(updateNetworkSwitchLinkAggregationRequest).Execute()

Update a link aggregation group



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
	linkAggregationId := "linkAggregationId_example" // string | Link aggregation ID
	updateNetworkSwitchLinkAggregationRequest := *openapiclient.NewUpdateNetworkSwitchLinkAggregationRequest() // UpdateNetworkSwitchLinkAggregationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinkAggregationsAPI.UpdateNetworkSwitchLinkAggregation(context.Background(), networkId, linkAggregationId).UpdateNetworkSwitchLinkAggregationRequest(updateNetworkSwitchLinkAggregationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinkAggregationsAPI.UpdateNetworkSwitchLinkAggregation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchLinkAggregation`: GetNetworkSwitchLinkAggregations200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `LinkAggregationsAPI.UpdateNetworkSwitchLinkAggregation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**linkAggregationId** | **string** | Link aggregation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchLinkAggregationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkSwitchLinkAggregationRequest** | [**UpdateNetworkSwitchLinkAggregationRequest**](UpdateNetworkSwitchLinkAggregationRequest.md) |  | 

### Return type

[**GetNetworkSwitchLinkAggregations200ResponseInner**](GetNetworkSwitchLinkAggregations200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

