# \LiveAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceCameraAnalyticsLive**](LiveAPI.md#GetDeviceCameraAnalyticsLive) | **Get** /devices/{serial}/camera/analytics/live | Returns live state from camera analytics zones



## GetDeviceCameraAnalyticsLive

> GetDeviceCameraAnalyticsLive200Response GetDeviceCameraAnalyticsLive(ctx, serial).Execute()

Returns live state from camera analytics zones



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
	serial := "serial_example" // string | Serial

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveAPI.GetDeviceCameraAnalyticsLive(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveAPI.GetDeviceCameraAnalyticsLive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraAnalyticsLive`: GetDeviceCameraAnalyticsLive200Response
	fmt.Fprintf(os.Stdout, "Response from `LiveAPI.GetDeviceCameraAnalyticsLive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraAnalyticsLiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceCameraAnalyticsLive200Response**](GetDeviceCameraAnalyticsLive200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

