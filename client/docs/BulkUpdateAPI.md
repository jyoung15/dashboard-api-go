# \BulkUpdateAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkUpdateOrganizationDevicesDetails**](BulkUpdateAPI.md#BulkUpdateOrganizationDevicesDetails) | **Post** /organizations/{organizationId}/devices/details/bulkUpdate | Updating device details (currently only used for Catalyst devices)



## BulkUpdateOrganizationDevicesDetails

> BulkUpdateOrganizationDevicesDetails200Response BulkUpdateOrganizationDevicesDetails(ctx, organizationId).BulkUpdateOrganizationDevicesDetailsRequest(bulkUpdateOrganizationDevicesDetailsRequest).Execute()

Updating device details (currently only used for Catalyst devices)



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
	bulkUpdateOrganizationDevicesDetailsRequest := *openapiclient.NewBulkUpdateOrganizationDevicesDetailsRequest([]string{"Serials_example"}, []openapiclient.BulkUpdateOrganizationDevicesDetailsRequestDetailsInner{*openapiclient.NewBulkUpdateOrganizationDevicesDetailsRequestDetailsInner("Name_example")}) // BulkUpdateOrganizationDevicesDetailsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkUpdateAPI.BulkUpdateOrganizationDevicesDetails(context.Background(), organizationId).BulkUpdateOrganizationDevicesDetailsRequest(bulkUpdateOrganizationDevicesDetailsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkUpdateAPI.BulkUpdateOrganizationDevicesDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkUpdateOrganizationDevicesDetails`: BulkUpdateOrganizationDevicesDetails200Response
	fmt.Fprintf(os.Stdout, "Response from `BulkUpdateAPI.BulkUpdateOrganizationDevicesDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateOrganizationDevicesDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkUpdateOrganizationDevicesDetailsRequest** | [**BulkUpdateOrganizationDevicesDetailsRequest**](BulkUpdateOrganizationDevicesDetailsRequest.md) |  | 

### Return type

[**BulkUpdateOrganizationDevicesDetails200Response**](BulkUpdateOrganizationDevicesDetails200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

