# \ClaimAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VmxNetworkDevicesClaim**](ClaimAPI.md#VmxNetworkDevicesClaim) | **Post** /networks/{networkId}/devices/claim/vmx | Claim a vMX into a network



## VmxNetworkDevicesClaim

> VmxNetworkDevicesClaim200Response VmxNetworkDevicesClaim(ctx, networkId).VmxNetworkDevicesClaimRequest(vmxNetworkDevicesClaimRequest).Execute()

Claim a vMX into a network



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
	vmxNetworkDevicesClaimRequest := *openapiclient.NewVmxNetworkDevicesClaimRequest("Size_example") // VmxNetworkDevicesClaimRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClaimAPI.VmxNetworkDevicesClaim(context.Background(), networkId).VmxNetworkDevicesClaimRequest(vmxNetworkDevicesClaimRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClaimAPI.VmxNetworkDevicesClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VmxNetworkDevicesClaim`: VmxNetworkDevicesClaim200Response
	fmt.Fprintf(os.Stdout, "Response from `ClaimAPI.VmxNetworkDevicesClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmxNetworkDevicesClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vmxNetworkDevicesClaimRequest** | [**VmxNetworkDevicesClaimRequest**](VmxNetworkDevicesClaimRequest.md) |  | 

### Return type

[**VmxNetworkDevicesClaim200Response**](VmxNetworkDevicesClaim200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

