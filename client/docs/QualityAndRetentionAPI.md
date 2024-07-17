# \QualityAndRetentionAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceCameraQualityAndRetention**](QualityAndRetentionAPI.md#GetDeviceCameraQualityAndRetention) | **Get** /devices/{serial}/camera/qualityAndRetention | Returns quality and retention settings for the given camera
[**UpdateDeviceCameraQualityAndRetention**](QualityAndRetentionAPI.md#UpdateDeviceCameraQualityAndRetention) | **Put** /devices/{serial}/camera/qualityAndRetention | Update quality and retention settings for the given camera



## GetDeviceCameraQualityAndRetention

> map[string]interface{} GetDeviceCameraQualityAndRetention(ctx, serial).Execute()

Returns quality and retention settings for the given camera



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
	resp, r, err := apiClient.QualityAndRetentionAPI.GetDeviceCameraQualityAndRetention(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityAndRetentionAPI.GetDeviceCameraQualityAndRetention``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraQualityAndRetention`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `QualityAndRetentionAPI.GetDeviceCameraQualityAndRetention`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraQualityAndRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceCameraQualityAndRetention

> map[string]interface{} UpdateDeviceCameraQualityAndRetention(ctx, serial).UpdateDeviceCameraQualityAndRetentionRequest(updateDeviceCameraQualityAndRetentionRequest).Execute()

Update quality and retention settings for the given camera



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
	updateDeviceCameraQualityAndRetentionRequest := *openapiclient.NewUpdateDeviceCameraQualityAndRetentionRequest() // UpdateDeviceCameraQualityAndRetentionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QualityAndRetentionAPI.UpdateDeviceCameraQualityAndRetention(context.Background(), serial).UpdateDeviceCameraQualityAndRetentionRequest(updateDeviceCameraQualityAndRetentionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityAndRetentionAPI.UpdateDeviceCameraQualityAndRetention``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceCameraQualityAndRetention`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `QualityAndRetentionAPI.UpdateDeviceCameraQualityAndRetention`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCameraQualityAndRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCameraQualityAndRetentionRequest** | [**UpdateDeviceCameraQualityAndRetentionRequest**](UpdateDeviceCameraQualityAndRetentionRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

