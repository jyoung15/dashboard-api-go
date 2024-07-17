# \BillingAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkWirelessBilling**](BillingAPI.md#GetNetworkWirelessBilling) | **Get** /networks/{networkId}/wireless/billing | Return the billing settings of this network
[**UpdateNetworkWirelessBilling**](BillingAPI.md#UpdateNetworkWirelessBilling) | **Put** /networks/{networkId}/wireless/billing | Update the billing settings



## GetNetworkWirelessBilling

> GetNetworkWirelessBilling200Response GetNetworkWirelessBilling(ctx, networkId).Execute()

Return the billing settings of this network



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
	resp, r, err := apiClient.BillingAPI.GetNetworkWirelessBilling(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.GetNetworkWirelessBilling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessBilling`: GetNetworkWirelessBilling200Response
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.GetNetworkWirelessBilling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessBillingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkWirelessBilling200Response**](GetNetworkWirelessBilling200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessBilling

> GetNetworkWirelessBilling200Response UpdateNetworkWirelessBilling(ctx, networkId).UpdateNetworkWirelessBillingRequest(updateNetworkWirelessBillingRequest).Execute()

Update the billing settings



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
	updateNetworkWirelessBillingRequest := *openapiclient.NewUpdateNetworkWirelessBillingRequest() // UpdateNetworkWirelessBillingRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillingAPI.UpdateNetworkWirelessBilling(context.Background(), networkId).UpdateNetworkWirelessBillingRequest(updateNetworkWirelessBillingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillingAPI.UpdateNetworkWirelessBilling``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkWirelessBilling`: GetNetworkWirelessBilling200Response
	fmt.Fprintf(os.Stdout, "Response from `BillingAPI.UpdateNetworkWirelessBilling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessBillingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkWirelessBillingRequest** | [**UpdateNetworkWirelessBillingRequest**](UpdateNetworkWirelessBillingRequest.md) |  | 

### Return type

[**GetNetworkWirelessBilling200Response**](GetNetworkWirelessBilling200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

