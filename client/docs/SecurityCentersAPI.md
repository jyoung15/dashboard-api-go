# \SecurityCentersAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSmDeviceSecurityCenters**](SecurityCentersAPI.md#GetNetworkSmDeviceSecurityCenters) | **Get** /networks/{networkId}/sm/devices/{deviceId}/securityCenters | List the security centers on a device



## GetNetworkSmDeviceSecurityCenters

> []GetNetworkSmDeviceSecurityCenters200ResponseInner GetNetworkSmDeviceSecurityCenters(ctx, networkId, deviceId).Execute()

List the security centers on a device



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
	resp, r, err := apiClient.SecurityCentersAPI.GetNetworkSmDeviceSecurityCenters(context.Background(), networkId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityCentersAPI.GetNetworkSmDeviceSecurityCenters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSmDeviceSecurityCenters`: []GetNetworkSmDeviceSecurityCenters200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SecurityCentersAPI.GetNetworkSmDeviceSecurityCenters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceSecurityCentersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceSecurityCenters200ResponseInner**](GetNetworkSmDeviceSecurityCenters200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

