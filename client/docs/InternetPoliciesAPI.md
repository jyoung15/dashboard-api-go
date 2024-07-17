# \InternetPoliciesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateNetworkApplianceSdwanInternetPolicies**](InternetPoliciesAPI.md#UpdateNetworkApplianceSdwanInternetPolicies) | **Put** /networks/{networkId}/appliance/sdwan/internetPolicies | Update SDWAN internet traffic preferences for an MX network



## UpdateNetworkApplianceSdwanInternetPolicies

> UpdateNetworkApplianceSdwanInternetPolicies200Response UpdateNetworkApplianceSdwanInternetPolicies(ctx, networkId).UpdateNetworkApplianceSdwanInternetPoliciesRequest(updateNetworkApplianceSdwanInternetPoliciesRequest).Execute()

Update SDWAN internet traffic preferences for an MX network



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
	updateNetworkApplianceSdwanInternetPoliciesRequest := *openapiclient.NewUpdateNetworkApplianceSdwanInternetPoliciesRequest() // UpdateNetworkApplianceSdwanInternetPoliciesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InternetPoliciesAPI.UpdateNetworkApplianceSdwanInternetPolicies(context.Background(), networkId).UpdateNetworkApplianceSdwanInternetPoliciesRequest(updateNetworkApplianceSdwanInternetPoliciesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternetPoliciesAPI.UpdateNetworkApplianceSdwanInternetPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceSdwanInternetPolicies`: UpdateNetworkApplianceSdwanInternetPolicies200Response
	fmt.Fprintf(os.Stdout, "Response from `InternetPoliciesAPI.UpdateNetworkApplianceSdwanInternetPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSdwanInternetPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceSdwanInternetPoliciesRequest** | [**UpdateNetworkApplianceSdwanInternetPoliciesRequest**](UpdateNetworkApplianceSdwanInternetPoliciesRequest.md) |  | 

### Return type

[**UpdateNetworkApplianceSdwanInternetPolicies200Response**](UpdateNetworkApplianceSdwanInternetPolicies200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

