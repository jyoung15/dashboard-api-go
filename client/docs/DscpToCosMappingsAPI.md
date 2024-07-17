# \DscpToCosMappingsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSwitchDscpToCosMappings**](DscpToCosMappingsAPI.md#GetNetworkSwitchDscpToCosMappings) | **Get** /networks/{networkId}/switch/dscpToCosMappings | Return the DSCP to CoS mappings
[**UpdateNetworkSwitchDscpToCosMappings**](DscpToCosMappingsAPI.md#UpdateNetworkSwitchDscpToCosMappings) | **Put** /networks/{networkId}/switch/dscpToCosMappings | Update the DSCP to CoS mappings



## GetNetworkSwitchDscpToCosMappings

> GetNetworkSwitchDscpToCosMappings200Response GetNetworkSwitchDscpToCosMappings(ctx, networkId).Execute()

Return the DSCP to CoS mappings



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
	resp, r, err := apiClient.DscpToCosMappingsAPI.GetNetworkSwitchDscpToCosMappings(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DscpToCosMappingsAPI.GetNetworkSwitchDscpToCosMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkSwitchDscpToCosMappings`: GetNetworkSwitchDscpToCosMappings200Response
	fmt.Fprintf(os.Stdout, "Response from `DscpToCosMappingsAPI.GetNetworkSwitchDscpToCosMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchDscpToCosMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkSwitchDscpToCosMappings200Response**](GetNetworkSwitchDscpToCosMappings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchDscpToCosMappings

> GetNetworkSwitchDscpToCosMappings200Response UpdateNetworkSwitchDscpToCosMappings(ctx, networkId).UpdateNetworkSwitchDscpToCosMappingsRequest(updateNetworkSwitchDscpToCosMappingsRequest).Execute()

Update the DSCP to CoS mappings



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
	updateNetworkSwitchDscpToCosMappingsRequest := *openapiclient.NewUpdateNetworkSwitchDscpToCosMappingsRequest([]openapiclient.UpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner{*openapiclient.NewUpdateNetworkSwitchDscpToCosMappingsRequestMappingsInner(int32(123), int32(123))}) // UpdateNetworkSwitchDscpToCosMappingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DscpToCosMappingsAPI.UpdateNetworkSwitchDscpToCosMappings(context.Background(), networkId).UpdateNetworkSwitchDscpToCosMappingsRequest(updateNetworkSwitchDscpToCosMappingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DscpToCosMappingsAPI.UpdateNetworkSwitchDscpToCosMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkSwitchDscpToCosMappings`: GetNetworkSwitchDscpToCosMappings200Response
	fmt.Fprintf(os.Stdout, "Response from `DscpToCosMappingsAPI.UpdateNetworkSwitchDscpToCosMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchDscpToCosMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchDscpToCosMappingsRequest** | [**UpdateNetworkSwitchDscpToCosMappingsRequest**](UpdateNetworkSwitchDscpToCosMappingsRequest.md) |  | 

### Return type

[**GetNetworkSwitchDscpToCosMappings200Response**](GetNetworkSwitchDscpToCosMappings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

