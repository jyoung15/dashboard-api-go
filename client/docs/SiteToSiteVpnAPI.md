# \SiteToSiteVpnAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkApplianceVpnSiteToSiteVpn**](SiteToSiteVpnAPI.md#GetNetworkApplianceVpnSiteToSiteVpn) | **Get** /networks/{networkId}/appliance/vpn/siteToSiteVpn | Return the site-to-site VPN settings of a network
[**UpdateNetworkApplianceVpnSiteToSiteVpn**](SiteToSiteVpnAPI.md#UpdateNetworkApplianceVpnSiteToSiteVpn) | **Put** /networks/{networkId}/appliance/vpn/siteToSiteVpn | Update the site-to-site VPN settings of a network



## GetNetworkApplianceVpnSiteToSiteVpn

> GetNetworkApplianceVpnSiteToSiteVpn200Response GetNetworkApplianceVpnSiteToSiteVpn(ctx, networkId).Execute()

Return the site-to-site VPN settings of a network



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
	resp, r, err := apiClient.SiteToSiteVpnAPI.GetNetworkApplianceVpnSiteToSiteVpn(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteToSiteVpnAPI.GetNetworkApplianceVpnSiteToSiteVpn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceVpnSiteToSiteVpn`: GetNetworkApplianceVpnSiteToSiteVpn200Response
	fmt.Fprintf(os.Stdout, "Response from `SiteToSiteVpnAPI.GetNetworkApplianceVpnSiteToSiteVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVpnSiteToSiteVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceVpnSiteToSiteVpn200Response**](GetNetworkApplianceVpnSiteToSiteVpn200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceVpnSiteToSiteVpn

> GetNetworkApplianceVpnSiteToSiteVpn200Response UpdateNetworkApplianceVpnSiteToSiteVpn(ctx, networkId).UpdateNetworkApplianceVpnSiteToSiteVpnRequest(updateNetworkApplianceVpnSiteToSiteVpnRequest).Execute()

Update the site-to-site VPN settings of a network



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
	updateNetworkApplianceVpnSiteToSiteVpnRequest := *openapiclient.NewUpdateNetworkApplianceVpnSiteToSiteVpnRequest("Mode_example") // UpdateNetworkApplianceVpnSiteToSiteVpnRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SiteToSiteVpnAPI.UpdateNetworkApplianceVpnSiteToSiteVpn(context.Background(), networkId).UpdateNetworkApplianceVpnSiteToSiteVpnRequest(updateNetworkApplianceVpnSiteToSiteVpnRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SiteToSiteVpnAPI.UpdateNetworkApplianceVpnSiteToSiteVpn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceVpnSiteToSiteVpn`: GetNetworkApplianceVpnSiteToSiteVpn200Response
	fmt.Fprintf(os.Stdout, "Response from `SiteToSiteVpnAPI.UpdateNetworkApplianceVpnSiteToSiteVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceVpnSiteToSiteVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceVpnSiteToSiteVpnRequest** | [**UpdateNetworkApplianceVpnSiteToSiteVpnRequest**](UpdateNetworkApplianceVpnSiteToSiteVpnRequest.md) |  | 

### Return type

[**GetNetworkApplianceVpnSiteToSiteVpn200Response**](GetNetworkApplianceVpnSiteToSiteVpn200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

