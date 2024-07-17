# \FieldsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateNetworkSmDevicesFields**](FieldsAPI.md#UpdateNetworkSmDevicesFields) | **Put** /networks/{networkId}/sm/devices/fields | Modify the fields of a device



## UpdateNetworkSmDevicesFields

> []UpdateNetworkSmDevicesFields200ResponseInner UpdateNetworkSmDevicesFields(ctx, networkId).UpdateNetworkSmDevicesFieldsRequest(updateNetworkSmDevicesFieldsRequest).Execute()

Modify the fields of a device



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
	updateNetworkSmDevicesFieldsRequest := *openapiclient.NewUpdateNetworkSmDevicesFieldsRequest(*openapiclient.NewUpdateNetworkSmDevicesFieldsRequestDeviceFields()) // UpdateNetworkSmDevicesFieldsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldsAPI.UpdateNetworkSmDevicesFields(context.Background(), networkId).UpdateNetworkSmDevicesFieldsRequest(updateNetworkSmDevicesFieldsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldsAPI.UpdateNetworkSmDevicesFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSmDevicesFields`: []UpdateNetworkSmDevicesFields200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `FieldsAPI.UpdateNetworkSmDevicesFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSmDevicesFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSmDevicesFieldsRequest** | [**UpdateNetworkSmDevicesFieldsRequest**](UpdateNetworkSmDevicesFieldsRequest.md) |  | 

### Return type

[**[]UpdateNetworkSmDevicesFields200ResponseInner**](UpdateNetworkSmDevicesFields200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

