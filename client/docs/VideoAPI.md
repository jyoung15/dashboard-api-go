# \VideoAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceCameraVideoSettings**](VideoAPI.md#GetDeviceCameraVideoSettings) | **Get** /devices/{serial}/camera/video/settings | Returns video settings for the given camera
[**UpdateDeviceCameraVideoSettings**](VideoAPI.md#UpdateDeviceCameraVideoSettings) | **Put** /devices/{serial}/camera/video/settings | Update video settings for the given camera



## GetDeviceCameraVideoSettings

> GetDeviceCameraVideoSettings200Response GetDeviceCameraVideoSettings(ctx, serial).Execute()

Returns video settings for the given camera



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
	resp, r, err := apiClient.VideoAPI.GetDeviceCameraVideoSettings(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.GetDeviceCameraVideoSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraVideoSettings`: GetDeviceCameraVideoSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `VideoAPI.GetDeviceCameraVideoSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraVideoSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceCameraVideoSettings200Response**](GetDeviceCameraVideoSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceCameraVideoSettings

> GetDeviceCameraVideoSettings200Response UpdateDeviceCameraVideoSettings(ctx, serial).UpdateDeviceCameraVideoSettingsRequest(updateDeviceCameraVideoSettingsRequest).Execute()

Update video settings for the given camera



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
	updateDeviceCameraVideoSettingsRequest := *openapiclient.NewUpdateDeviceCameraVideoSettingsRequest() // UpdateDeviceCameraVideoSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoAPI.UpdateDeviceCameraVideoSettings(context.Background(), serial).UpdateDeviceCameraVideoSettingsRequest(updateDeviceCameraVideoSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.UpdateDeviceCameraVideoSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceCameraVideoSettings`: GetDeviceCameraVideoSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `VideoAPI.UpdateDeviceCameraVideoSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCameraVideoSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCameraVideoSettingsRequest** | [**UpdateDeviceCameraVideoSettingsRequest**](UpdateDeviceCameraVideoSettingsRequest.md) |  | 

### Return type

[**GetDeviceCameraVideoSettings200Response**](GetDeviceCameraVideoSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

