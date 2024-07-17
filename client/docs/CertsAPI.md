# \CertsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSmDeviceCerts**](CertsAPI.md#GetNetworkSmDeviceCerts) | **Get** /networks/{networkId}/sm/devices/{deviceId}/certs | List the certs on a device



## GetNetworkSmDeviceCerts

> []GetNetworkSmDeviceCerts200ResponseInner GetNetworkSmDeviceCerts(ctx, networkId, deviceId).Execute()

List the certs on a device



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
	resp, r, err := apiClient.CertsAPI.GetNetworkSmDeviceCerts(context.Background(), networkId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertsAPI.GetNetworkSmDeviceCerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceCerts`: []GetNetworkSmDeviceCerts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CertsAPI.GetNetworkSmDeviceCerts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceCertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceCerts200ResponseInner**](GetNetworkSmDeviceCerts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

