# \DscpTaggingOptionsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkTrafficShapingDscpTaggingOptions**](DscpTaggingOptionsAPI.md#GetNetworkTrafficShapingDscpTaggingOptions) | **Get** /networks/{networkId}/trafficShaping/dscpTaggingOptions | Returns the available DSCP tagging options for your traffic shaping rules.



## GetNetworkTrafficShapingDscpTaggingOptions

> []map[string]interface{} GetNetworkTrafficShapingDscpTaggingOptions(ctx, networkId).Execute()

Returns the available DSCP tagging options for your traffic shaping rules.



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
	resp, r, err := apiClient.DscpTaggingOptionsAPI.GetNetworkTrafficShapingDscpTaggingOptions(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DscpTaggingOptionsAPI.GetNetworkTrafficShapingDscpTaggingOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkTrafficShapingDscpTaggingOptions`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DscpTaggingOptionsAPI.GetNetworkTrafficShapingDscpTaggingOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkTrafficShapingDscpTaggingOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

