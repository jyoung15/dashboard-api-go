# \EventTypesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkEventsEventTypes**](EventTypesAPI.md#GetNetworkEventsEventTypes) | **Get** /networks/{networkId}/events/eventTypes | List the event type to human-readable description



## GetNetworkEventsEventTypes

> []GetNetworkEventsEventTypes200ResponseInner GetNetworkEventsEventTypes(ctx, networkId).Execute()

List the event type to human-readable description



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
	resp, r, err := apiClient.EventTypesAPI.GetNetworkEventsEventTypes(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventTypesAPI.GetNetworkEventsEventTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkEventsEventTypes`: []GetNetworkEventsEventTypes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `EventTypesAPI.GetNetworkEventsEventTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkEventsEventTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkEventsEventTypes200ResponseInner**](GetNetworkEventsEventTypes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

