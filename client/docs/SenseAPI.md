# \SenseAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceCameraSense**](SenseAPI.md#GetDeviceCameraSense) | **Get** /devices/{serial}/camera/sense | Returns sense settings for a given camera
[**GetDeviceCameraSenseObjectDetectionModels**](SenseAPI.md#GetDeviceCameraSenseObjectDetectionModels) | **Get** /devices/{serial}/camera/sense/objectDetectionModels | Returns the MV Sense object detection model list for the given camera
[**UpdateDeviceCameraSense**](SenseAPI.md#UpdateDeviceCameraSense) | **Put** /devices/{serial}/camera/sense | Update sense settings for the given camera



## GetDeviceCameraSense

> map[string]interface{} GetDeviceCameraSense(ctx, serial).Execute()

Returns sense settings for a given camera



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
	resp, r, err := apiClient.SenseAPI.GetDeviceCameraSense(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SenseAPI.GetDeviceCameraSense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraSense`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SenseAPI.GetDeviceCameraSense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraSenseRequest struct via the builder pattern


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
	resp, r, err := apiClient.SenseAPI.GetDeviceCameraSenseObjectDetectionModels(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SenseAPI.GetDeviceCameraSenseObjectDetectionModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraSenseObjectDetectionModels`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SenseAPI.GetDeviceCameraSenseObjectDetectionModels`: %v\n", resp)
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


## UpdateDeviceCameraSense

> map[string]interface{} UpdateDeviceCameraSense(ctx, serial).UpdateDeviceCameraSenseRequest(updateDeviceCameraSenseRequest).Execute()

Update sense settings for the given camera



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
	updateDeviceCameraSenseRequest := *openapiclient.NewUpdateDeviceCameraSenseRequest() // UpdateDeviceCameraSenseRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SenseAPI.UpdateDeviceCameraSense(context.Background(), serial).UpdateDeviceCameraSenseRequest(updateDeviceCameraSenseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SenseAPI.UpdateDeviceCameraSense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceCameraSense`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SenseAPI.UpdateDeviceCameraSense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCameraSenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCameraSenseRequest** | [**UpdateDeviceCameraSenseRequest**](UpdateDeviceCameraSenseRequest.md) |  | 

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

