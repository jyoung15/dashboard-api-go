# \DeviceProfilesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSmDeviceDeviceProfiles**](DeviceProfilesAPI.md#GetNetworkSmDeviceDeviceProfiles) | **Get** /networks/{networkId}/sm/devices/{deviceId}/deviceProfiles | Get the installed profiles associated with a device
[**GetNetworkSmUserDeviceProfiles**](DeviceProfilesAPI.md#GetNetworkSmUserDeviceProfiles) | **Get** /networks/{networkId}/sm/users/{userId}/deviceProfiles | Get the profiles associated with a user



## GetNetworkSmDeviceDeviceProfiles

> []GetNetworkSmDeviceDeviceProfiles200ResponseInner GetNetworkSmDeviceDeviceProfiles(ctx, networkId, deviceId).Execute()

Get the installed profiles associated with a device



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
	deviceId := "deviceId_example" // string | Device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceProfilesAPI.GetNetworkSmDeviceDeviceProfiles(context.Background(), networkId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceProfilesAPI.GetNetworkSmDeviceDeviceProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceDeviceProfiles`: []GetNetworkSmDeviceDeviceProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DeviceProfilesAPI.GetNetworkSmDeviceDeviceProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceDeviceProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceDeviceProfiles200ResponseInner**](GetNetworkSmDeviceDeviceProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmUserDeviceProfiles

> []GetNetworkSmDeviceDeviceProfiles200ResponseInner GetNetworkSmUserDeviceProfiles(ctx, networkId, userId).Execute()

Get the profiles associated with a user



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
	userId := "userId_example" // string | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceProfilesAPI.GetNetworkSmUserDeviceProfiles(context.Background(), networkId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceProfilesAPI.GetNetworkSmUserDeviceProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmUserDeviceProfiles`: []GetNetworkSmDeviceDeviceProfiles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DeviceProfilesAPI.GetNetworkSmUserDeviceProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**userId** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmUserDeviceProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceDeviceProfiles200ResponseInner**](GetNetworkSmDeviceDeviceProfiles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

