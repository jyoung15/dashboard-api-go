# \ThirdPartyVPNPeersAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationApplianceVpnThirdPartyVPNPeers**](ThirdPartyVPNPeersAPI.md#GetOrganizationApplianceVpnThirdPartyVPNPeers) | **Get** /organizations/{organizationId}/appliance/vpn/thirdPartyVPNPeers | Return the third party VPN peers for an organization
[**UpdateOrganizationApplianceVpnThirdPartyVPNPeers**](ThirdPartyVPNPeersAPI.md#UpdateOrganizationApplianceVpnThirdPartyVPNPeers) | **Put** /organizations/{organizationId}/appliance/vpn/thirdPartyVPNPeers | Update the third party VPN peers for an organization



## GetOrganizationApplianceVpnThirdPartyVPNPeers

> GetOrganizationApplianceVpnThirdPartyVPNPeers200Response GetOrganizationApplianceVpnThirdPartyVPNPeers(ctx, organizationId).Execute()

Return the third party VPN peers for an organization



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThirdPartyVPNPeersAPI.GetOrganizationApplianceVpnThirdPartyVPNPeers(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThirdPartyVPNPeersAPI.GetOrganizationApplianceVpnThirdPartyVPNPeers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationApplianceVpnThirdPartyVPNPeers`: GetOrganizationApplianceVpnThirdPartyVPNPeers200Response
	fmt.Fprintf(os.Stdout, "Response from `ThirdPartyVPNPeersAPI.GetOrganizationApplianceVpnThirdPartyVPNPeers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceVpnThirdPartyVPNPeersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrganizationApplianceVpnThirdPartyVPNPeers200Response**](GetOrganizationApplianceVpnThirdPartyVPNPeers200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationApplianceVpnThirdPartyVPNPeers

> GetOrganizationApplianceVpnThirdPartyVPNPeers200Response UpdateOrganizationApplianceVpnThirdPartyVPNPeers(ctx, organizationId).UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest(updateOrganizationApplianceVpnThirdPartyVPNPeersRequest).Execute()

Update the third party VPN peers for an organization



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
	updateOrganizationApplianceVpnThirdPartyVPNPeersRequest := *openapiclient.NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest([]openapiclient.UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner{*openapiclient.NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner("Name_example", []string{"PrivateSubnets_example"}, "Secret_example")}) // UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThirdPartyVPNPeersAPI.UpdateOrganizationApplianceVpnThirdPartyVPNPeers(context.Background(), organizationId).UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest(updateOrganizationApplianceVpnThirdPartyVPNPeersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThirdPartyVPNPeersAPI.UpdateOrganizationApplianceVpnThirdPartyVPNPeers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationApplianceVpnThirdPartyVPNPeers`: GetOrganizationApplianceVpnThirdPartyVPNPeers200Response
	fmt.Fprintf(os.Stdout, "Response from `ThirdPartyVPNPeersAPI.UpdateOrganizationApplianceVpnThirdPartyVPNPeers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationApplianceVpnThirdPartyVPNPeersRequest** | [**UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest**](UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest.md) |  | 

### Return type

[**GetOrganizationApplianceVpnThirdPartyVPNPeers200Response**](GetOrganizationApplianceVpnThirdPartyVPNPeers200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

