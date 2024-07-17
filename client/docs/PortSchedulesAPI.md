# \PortSchedulesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkSwitchPortSchedule**](PortSchedulesAPI.md#CreateNetworkSwitchPortSchedule) | **Post** /networks/{networkId}/switch/portSchedules | Add a switch port schedule
[**DeleteNetworkSwitchPortSchedule**](PortSchedulesAPI.md#DeleteNetworkSwitchPortSchedule) | **Delete** /networks/{networkId}/switch/portSchedules/{portScheduleId} | Delete a switch port schedule
[**GetNetworkSwitchPortSchedules**](PortSchedulesAPI.md#GetNetworkSwitchPortSchedules) | **Get** /networks/{networkId}/switch/portSchedules | List switch port schedules
[**UpdateNetworkSwitchPortSchedule**](PortSchedulesAPI.md#UpdateNetworkSwitchPortSchedule) | **Put** /networks/{networkId}/switch/portSchedules/{portScheduleId} | Update a switch port schedule



## CreateNetworkSwitchPortSchedule

> GetNetworkSwitchPortSchedules200ResponseInner CreateNetworkSwitchPortSchedule(ctx, networkId).CreateNetworkSwitchPortScheduleRequest(createNetworkSwitchPortScheduleRequest).Execute()

Add a switch port schedule



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
	createNetworkSwitchPortScheduleRequest := *openapiclient.NewCreateNetworkSwitchPortScheduleRequest("Name_example") // CreateNetworkSwitchPortScheduleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortSchedulesAPI.CreateNetworkSwitchPortSchedule(context.Background(), networkId).CreateNetworkSwitchPortScheduleRequest(createNetworkSwitchPortScheduleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortSchedulesAPI.CreateNetworkSwitchPortSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkSwitchPortSchedule`: GetNetworkSwitchPortSchedules200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PortSchedulesAPI.CreateNetworkSwitchPortSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchPortScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSwitchPortScheduleRequest** | [**CreateNetworkSwitchPortScheduleRequest**](CreateNetworkSwitchPortScheduleRequest.md) |  | 

### Return type

[**GetNetworkSwitchPortSchedules200ResponseInner**](GetNetworkSwitchPortSchedules200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSwitchPortSchedule

> DeleteNetworkSwitchPortSchedule(ctx, networkId, portScheduleId).Execute()

Delete a switch port schedule



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
	portScheduleId := "portScheduleId_example" // string | Port schedule ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PortSchedulesAPI.DeleteNetworkSwitchPortSchedule(context.Background(), networkId, portScheduleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortSchedulesAPI.DeleteNetworkSwitchPortSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**portScheduleId** | **string** | Port schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchPortScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchPortSchedules

> []GetNetworkSwitchPortSchedules200ResponseInner GetNetworkSwitchPortSchedules(ctx, networkId).Execute()

List switch port schedules



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
	resp, r, err := apiClient.PortSchedulesAPI.GetNetworkSwitchPortSchedules(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortSchedulesAPI.GetNetworkSwitchPortSchedules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchPortSchedules`: []GetNetworkSwitchPortSchedules200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PortSchedulesAPI.GetNetworkSwitchPortSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchPortSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkSwitchPortSchedules200ResponseInner**](GetNetworkSwitchPortSchedules200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchPortSchedule

> GetNetworkSwitchPortSchedules200ResponseInner UpdateNetworkSwitchPortSchedule(ctx, networkId, portScheduleId).UpdateNetworkSwitchPortScheduleRequest(updateNetworkSwitchPortScheduleRequest).Execute()

Update a switch port schedule



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
	portScheduleId := "portScheduleId_example" // string | Port schedule ID
	updateNetworkSwitchPortScheduleRequest := *openapiclient.NewUpdateNetworkSwitchPortScheduleRequest() // UpdateNetworkSwitchPortScheduleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortSchedulesAPI.UpdateNetworkSwitchPortSchedule(context.Background(), networkId, portScheduleId).UpdateNetworkSwitchPortScheduleRequest(updateNetworkSwitchPortScheduleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortSchedulesAPI.UpdateNetworkSwitchPortSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchPortSchedule`: GetNetworkSwitchPortSchedules200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PortSchedulesAPI.UpdateNetworkSwitchPortSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**portScheduleId** | **string** | Port schedule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchPortScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkSwitchPortScheduleRequest** | [**UpdateNetworkSwitchPortScheduleRequest**](UpdateNetworkSwitchPortScheduleRequest.md) |  | 

### Return type

[**GetNetworkSwitchPortSchedules200ResponseInner**](GetNetworkSwitchPortSchedules200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

