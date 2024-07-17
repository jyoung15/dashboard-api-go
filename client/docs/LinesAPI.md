# \LinesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationCameraBoundariesLinesByDevice**](LinesAPI.md#GetOrganizationCameraBoundariesLinesByDevice) | **Get** /organizations/{organizationId}/camera/boundaries/lines/byDevice | Returns all configured crossingline boundaries of cameras



## GetOrganizationCameraBoundariesLinesByDevice

> []GetOrganizationCameraBoundariesLinesByDevice200ResponseInner GetOrganizationCameraBoundariesLinesByDevice(ctx, organizationId).Serials(serials).Execute()

Returns all configured crossingline boundaries of cameras



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
	resp, r, err := apiClient.LinesAPI.GetOrganizationCameraBoundariesLinesByDevice(context.Background(), organizationId).Serials(serials).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.GetOrganizationCameraBoundariesLinesByDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationCameraBoundariesLinesByDevice`: []GetOrganizationCameraBoundariesLinesByDevice200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `LinesAPI.GetOrganizationCameraBoundariesLinesByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraBoundariesLinesByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | A list of serial numbers. The returned cameras will be filtered to only include these serials. | 

### Return type

[**[]GetOrganizationCameraBoundariesLinesByDevice200ResponseInner**](GetOrganizationCameraBoundariesLinesByDevice200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

