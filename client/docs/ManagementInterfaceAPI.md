# \ManagementInterfaceAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceManagementInterface**](ManagementInterfaceAPI.md#GetDeviceManagementInterface) | **Get** /devices/{serial}/managementInterface | Return the management interface settings for a device
[**UpdateDeviceManagementInterface**](ManagementInterfaceAPI.md#UpdateDeviceManagementInterface) | **Put** /devices/{serial}/managementInterface | Update the management interface settings for a device



## GetDeviceManagementInterface

> GetDeviceManagementInterface200Response GetDeviceManagementInterface(ctx, serial).Execute()

Return the management interface settings for a device



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
	resp, r, err := apiClient.ManagementInterfaceAPI.GetDeviceManagementInterface(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementInterfaceAPI.GetDeviceManagementInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceManagementInterface`: GetDeviceManagementInterface200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementInterfaceAPI.GetDeviceManagementInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceManagementInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceManagementInterface200Response**](GetDeviceManagementInterface200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceManagementInterface

> GetDeviceManagementInterface200Response UpdateDeviceManagementInterface(ctx, serial).UpdateDeviceManagementInterfaceRequest(updateDeviceManagementInterfaceRequest).Execute()

Update the management interface settings for a device



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
	updateDeviceManagementInterfaceRequest := *openapiclient.NewUpdateDeviceManagementInterfaceRequest() // UpdateDeviceManagementInterfaceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementInterfaceAPI.UpdateDeviceManagementInterface(context.Background(), serial).UpdateDeviceManagementInterfaceRequest(updateDeviceManagementInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementInterfaceAPI.UpdateDeviceManagementInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceManagementInterface`: GetDeviceManagementInterface200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementInterfaceAPI.UpdateDeviceManagementInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceManagementInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceManagementInterfaceRequest** | [**UpdateDeviceManagementInterfaceRequest**](UpdateDeviceManagementInterfaceRequest.md) |  | 

### Return type

[**GetDeviceManagementInterface200Response**](GetDeviceManagementInterface200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

