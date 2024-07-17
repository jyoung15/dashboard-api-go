# \SchedulesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkCameraSchedules**](SchedulesAPI.md#GetNetworkCameraSchedules) | **Get** /networks/{networkId}/camera/schedules | Returns a list of all camera recording schedules.
[**GetNetworkWirelessSsidSchedules**](SchedulesAPI.md#GetNetworkWirelessSsidSchedules) | **Get** /networks/{networkId}/wireless/ssids/{number}/schedules | List the outage schedule for the SSID
[**UpdateNetworkWirelessSsidSchedules**](SchedulesAPI.md#UpdateNetworkWirelessSsidSchedules) | **Put** /networks/{networkId}/wireless/ssids/{number}/schedules | Update the outage schedule for the SSID



## GetNetworkCameraSchedules

> []GetNetworkCameraSchedules200ResponseInner GetNetworkCameraSchedules(ctx, networkId).Execute()

Returns a list of all camera recording schedules.



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
	resp, r, err := apiClient.SchedulesAPI.GetNetworkCameraSchedules(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchedulesAPI.GetNetworkCameraSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkCameraSchedules`: []GetNetworkCameraSchedules200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SchedulesAPI.GetNetworkCameraSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCameraSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkCameraSchedules200ResponseInner**](GetNetworkCameraSchedules200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidSchedules

> GetNetworkWirelessSsidSchedules200Response GetNetworkWirelessSsidSchedules(ctx, networkId, number).Execute()

List the outage schedule for the SSID



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
	number := "number_example" // string | Number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchedulesAPI.GetNetworkWirelessSsidSchedules(context.Background(), networkId, number).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchedulesAPI.GetNetworkWirelessSsidSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessSsidSchedules`: GetNetworkWirelessSsidSchedules200Response
	fmt.Fprintf(os.Stdout, "Response from `SchedulesAPI.GetNetworkWirelessSsidSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkWirelessSsidSchedules200Response**](GetNetworkWirelessSsidSchedules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidSchedules

> GetNetworkWirelessSsidSchedules200Response UpdateNetworkWirelessSsidSchedules(ctx, networkId, number).UpdateNetworkWirelessSsidSchedulesRequest(updateNetworkWirelessSsidSchedulesRequest).Execute()

Update the outage schedule for the SSID



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
	number := "number_example" // string | Number
	updateNetworkWirelessSsidSchedulesRequest := *openapiclient.NewUpdateNetworkWirelessSsidSchedulesRequest() // UpdateNetworkWirelessSsidSchedulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SchedulesAPI.UpdateNetworkWirelessSsidSchedules(context.Background(), networkId, number).UpdateNetworkWirelessSsidSchedulesRequest(updateNetworkWirelessSsidSchedulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SchedulesAPI.UpdateNetworkWirelessSsidSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkWirelessSsidSchedules`: GetNetworkWirelessSsidSchedules200Response
	fmt.Fprintf(os.Stdout, "Response from `SchedulesAPI.UpdateNetworkWirelessSsidSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidSchedulesRequest** | [**UpdateNetworkWirelessSsidSchedulesRequest**](UpdateNetworkWirelessSsidSchedulesRequest.md) |  | 

### Return type

[**GetNetworkWirelessSsidSchedules200Response**](GetNetworkWirelessSsidSchedules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

