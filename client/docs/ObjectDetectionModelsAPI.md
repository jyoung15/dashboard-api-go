# \ObjectDetectionModelsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceCameraSenseObjectDetectionModels**](ObjectDetectionModelsAPI.md#GetDeviceCameraSenseObjectDetectionModels) | **Get** /devices/{serial}/camera/sense/objectDetectionModels | Returns the MV Sense object detection model list for the given camera



## GetDeviceCameraSenseObjectDetectionModels

> []map[string]interface{} GetDeviceCameraSenseObjectDetectionModels(ctx, serial).Execute()

Returns the MV Sense object detection model list for the given camera



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
	resp, r, err := apiClient.ObjectDetectionModelsAPI.GetDeviceCameraSenseObjectDetectionModels(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectDetectionModelsAPI.GetDeviceCameraSenseObjectDetectionModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraSenseObjectDetectionModels`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ObjectDetectionModelsAPI.GetDeviceCameraSenseObjectDetectionModels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraSenseObjectDetectionModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

