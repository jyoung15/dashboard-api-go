# \RecentAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceCameraAnalyticsRecent**](RecentAPI.md#GetDeviceCameraAnalyticsRecent) | **Get** /devices/{serial}/camera/analytics/recent | Returns most recent record for analytics zones



## GetDeviceCameraAnalyticsRecent

> []GetDeviceCameraAnalyticsRecent200ResponseInner GetDeviceCameraAnalyticsRecent(ctx, serial).ObjectType(objectType).Execute()

Returns most recent record for analytics zones



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
	objectType := "objectType_example" // string | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecentAPI.GetDeviceCameraAnalyticsRecent(context.Background(), serial).ObjectType(objectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecentAPI.GetDeviceCameraAnalyticsRecent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraAnalyticsRecent`: []GetDeviceCameraAnalyticsRecent200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RecentAPI.GetDeviceCameraAnalyticsRecent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraAnalyticsRecentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectType** | **string** | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. | 

### Return type

[**[]GetDeviceCameraAnalyticsRecent200ResponseInner**](GetDeviceCameraAnalyticsRecent200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

