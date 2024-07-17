# \RequestsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkPiiRequest**](RequestsAPI.md#CreateNetworkPiiRequest) | **Post** /networks/{networkId}/pii/requests | Submit a new delete or restrict processing PII request
[**DeleteNetworkPiiRequest**](RequestsAPI.md#DeleteNetworkPiiRequest) | **Delete** /networks/{networkId}/pii/requests/{requestId} | Delete a restrict processing PII request
[**GetNetworkPiiRequest**](RequestsAPI.md#GetNetworkPiiRequest) | **Get** /networks/{networkId}/pii/requests/{requestId} | Return a PII request
[**GetNetworkPiiRequests**](RequestsAPI.md#GetNetworkPiiRequests) | **Get** /networks/{networkId}/pii/requests | List the PII requests for this network or organization



## CreateNetworkPiiRequest

> GetNetworkPiiRequests200ResponseInner CreateNetworkPiiRequest(ctx, networkId).CreateNetworkPiiRequestRequest(createNetworkPiiRequestRequest).Execute()

Submit a new delete or restrict processing PII request



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
	createNetworkPiiRequestRequest := *openapiclient.NewCreateNetworkPiiRequestRequest() // CreateNetworkPiiRequestRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.CreateNetworkPiiRequest(context.Background(), networkId).CreateNetworkPiiRequestRequest(createNetworkPiiRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.CreateNetworkPiiRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkPiiRequest`: GetNetworkPiiRequests200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.CreateNetworkPiiRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkPiiRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkPiiRequestRequest** | [**CreateNetworkPiiRequestRequest**](CreateNetworkPiiRequestRequest.md) |  | 

### Return type

[**GetNetworkPiiRequests200ResponseInner**](GetNetworkPiiRequests200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkPiiRequest

> DeleteNetworkPiiRequest(ctx, networkId, requestId).Execute()

Delete a restrict processing PII request



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
	requestId := "requestId_example" // string | Request ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RequestsAPI.DeleteNetworkPiiRequest(context.Background(), networkId, requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.DeleteNetworkPiiRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkPiiRequestRequest struct via the builder pattern


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


## GetNetworkPiiRequest

> GetNetworkPiiRequests200ResponseInner GetNetworkPiiRequest(ctx, networkId, requestId).Execute()

Return a PII request



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
	requestId := "requestId_example" // string | Request ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.GetNetworkPiiRequest(context.Background(), networkId, requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.GetNetworkPiiRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPiiRequest`: GetNetworkPiiRequests200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.GetNetworkPiiRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**requestId** | **string** | Request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkPiiRequests200ResponseInner**](GetNetworkPiiRequests200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPiiRequests

> []GetNetworkPiiRequests200ResponseInner GetNetworkPiiRequests(ctx, networkId).Execute()

List the PII requests for this network or organization



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
	resp, r, err := apiClient.RequestsAPI.GetNetworkPiiRequests(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.GetNetworkPiiRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkPiiRequests`: []GetNetworkPiiRequests200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.GetNetworkPiiRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkPiiRequests200ResponseInner**](GetNetworkPiiRequests200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

