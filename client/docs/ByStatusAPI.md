# \ByStatusAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationSummaryTopNetworksByStatus**](ByStatusAPI.md#GetOrganizationSummaryTopNetworksByStatus) | **Get** /organizations/{organizationId}/summary/top/networks/byStatus | List the client and status overview information for the networks in an organization



## GetOrganizationSummaryTopNetworksByStatus

> []GetOrganizationSummaryTopNetworksByStatus200ResponseInner GetOrganizationSummaryTopNetworksByStatus(ctx, organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the client and status overview information for the networks in an organization



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
	networkTag := "networkTag_example" // string | Match result to an exact network tag (optional)
	deviceTag := "deviceTag_example" // string | Match result to an exact device tag (optional)
	networkId := "networkId_example" // string | Match result to an exact network id (optional)
	quantity := int32(56) // int32 | Set number of desired results to return. Default is 10. (optional)
	ssidName := "ssidName_example" // string | Filter results by ssid name (optional)
	usageUplink := "usageUplink_example" // string | Filter results by usage uplink (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 5000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ByStatusAPI.GetOrganizationSummaryTopNetworksByStatus(context.Background(), organizationId).NetworkTag(networkTag).DeviceTag(deviceTag).NetworkId(networkId).Quantity(quantity).SsidName(ssidName).UsageUplink(usageUplink).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ByStatusAPI.GetOrganizationSummaryTopNetworksByStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSummaryTopNetworksByStatus`: []GetOrganizationSummaryTopNetworksByStatus200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ByStatusAPI.GetOrganizationSummaryTopNetworksByStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopNetworksByStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkTag** | **string** | Match result to an exact network tag | 
 **deviceTag** | **string** | Match result to an exact device tag | 
 **networkId** | **string** | Match result to an exact network id | 
 **quantity** | **int32** | Set number of desired results to return. Default is 10. | 
 **ssidName** | **string** | Filter results by ssid name | 
 **usageUplink** | **string** | Filter results by usage uplink | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 5000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetOrganizationSummaryTopNetworksByStatus200ResponseInner**](GetOrganizationSummaryTopNetworksByStatus200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

