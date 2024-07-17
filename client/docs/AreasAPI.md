# \AreasAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationCameraBoundariesAreasByDevice**](AreasAPI.md#GetOrganizationCameraBoundariesAreasByDevice) | **Get** /organizations/{organizationId}/camera/boundaries/areas/byDevice | Returns all configured area boundaries of cameras



## GetOrganizationCameraBoundariesAreasByDevice

> []GetOrganizationCameraBoundariesAreasByDevice200ResponseInner GetOrganizationCameraBoundariesAreasByDevice(ctx, organizationId).Serials(serials).Execute()

Returns all configured area boundaries of cameras



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
	organizationId := "organizationId_example" // string | Organization ID
	serials := []string{"Inner_example"} // []string | A list of serial numbers. The returned cameras will be filtered to only include these serials. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AreasAPI.GetOrganizationCameraBoundariesAreasByDevice(context.Background(), organizationId).Serials(serials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AreasAPI.GetOrganizationCameraBoundariesAreasByDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationCameraBoundariesAreasByDevice`: []GetOrganizationCameraBoundariesAreasByDevice200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AreasAPI.GetOrganizationCameraBoundariesAreasByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraBoundariesAreasByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | A list of serial numbers. The returned cameras will be filtered to only include these serials. | 

### Return type

[**[]GetOrganizationCameraBoundariesAreasByDevice200ResponseInner**](GetOrganizationCameraBoundariesAreasByDevice200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

