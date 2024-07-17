# \RelationshipsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceSensorRelationships**](RelationshipsAPI.md#GetDeviceSensorRelationships) | **Get** /devices/{serial}/sensor/relationships | List the sensor roles for a given sensor or camera device.
[**GetNetworkSensorRelationships**](RelationshipsAPI.md#GetNetworkSensorRelationships) | **Get** /networks/{networkId}/sensor/relationships | List the sensor roles for devices in a given network
[**UpdateDeviceSensorRelationships**](RelationshipsAPI.md#UpdateDeviceSensorRelationships) | **Put** /devices/{serial}/sensor/relationships | Assign one or more sensor roles to a given sensor or camera device.



## GetDeviceSensorRelationships

> GetDeviceSensorRelationships200Response GetDeviceSensorRelationships(ctx, serial).Execute()

List the sensor roles for a given sensor or camera device.



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
	resp, r, err := apiClient.RelationshipsAPI.GetDeviceSensorRelationships(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationshipsAPI.GetDeviceSensorRelationships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceSensorRelationships`: GetDeviceSensorRelationships200Response
	fmt.Fprintf(os.Stdout, "Response from `RelationshipsAPI.GetDeviceSensorRelationships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSensorRelationshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceSensorRelationships200Response**](GetDeviceSensorRelationships200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSensorRelationships

> []GetNetworkSensorRelationships200ResponseInner GetNetworkSensorRelationships(ctx, networkId).Execute()

List the sensor roles for devices in a given network



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
	resp, r, err := apiClient.RelationshipsAPI.GetNetworkSensorRelationships(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationshipsAPI.GetNetworkSensorRelationships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSensorRelationships`: []GetNetworkSensorRelationships200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RelationshipsAPI.GetNetworkSensorRelationships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSensorRelationshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkSensorRelationships200ResponseInner**](GetNetworkSensorRelationships200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceSensorRelationships

> GetDeviceSensorRelationships200Response UpdateDeviceSensorRelationships(ctx, serial).UpdateDeviceSensorRelationshipsRequest(updateDeviceSensorRelationshipsRequest).Execute()

Assign one or more sensor roles to a given sensor or camera device.



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
	updateDeviceSensorRelationshipsRequest := *openapiclient.NewUpdateDeviceSensorRelationshipsRequest() // UpdateDeviceSensorRelationshipsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationshipsAPI.UpdateDeviceSensorRelationships(context.Background(), serial).UpdateDeviceSensorRelationshipsRequest(updateDeviceSensorRelationshipsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationshipsAPI.UpdateDeviceSensorRelationships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceSensorRelationships`: GetDeviceSensorRelationships200Response
	fmt.Fprintf(os.Stdout, "Response from `RelationshipsAPI.UpdateDeviceSensorRelationships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceSensorRelationshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceSensorRelationshipsRequest** | [**UpdateDeviceSensorRelationshipsRequest**](UpdateDeviceSensorRelationshipsRequest.md) |  | 

### Return type

[**GetDeviceSensorRelationships200Response**](GetDeviceSensorRelationships200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

