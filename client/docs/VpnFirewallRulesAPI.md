# \VpnFirewallRulesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationApplianceVpnVpnFirewallRules**](VpnFirewallRulesAPI.md#GetOrganizationApplianceVpnVpnFirewallRules) | **Get** /organizations/{organizationId}/appliance/vpn/vpnFirewallRules | Return the firewall rules for an organization&#39;s site-to-site VPN
[**UpdateOrganizationApplianceVpnVpnFirewallRules**](VpnFirewallRulesAPI.md#UpdateOrganizationApplianceVpnVpnFirewallRules) | **Put** /organizations/{organizationId}/appliance/vpn/vpnFirewallRules | Update the firewall rules of an organization&#39;s site-to-site VPN



## GetOrganizationApplianceVpnVpnFirewallRules

> GetNetworkApplianceFirewallInboundCellularFirewallRules200Response GetOrganizationApplianceVpnVpnFirewallRules(ctx, organizationId).Execute()

Return the firewall rules for an organization's site-to-site VPN



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
	resp, r, err := apiClient.VpnFirewallRulesAPI.GetOrganizationApplianceVpnVpnFirewallRules(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpnFirewallRulesAPI.GetOrganizationApplianceVpnVpnFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationApplianceVpnVpnFirewallRules`: GetNetworkApplianceFirewallInboundCellularFirewallRules200Response
	fmt.Fprintf(os.Stdout, "Response from `VpnFirewallRulesAPI.GetOrganizationApplianceVpnVpnFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceVpnVpnFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceFirewallInboundCellularFirewallRules200Response**](GetNetworkApplianceFirewallInboundCellularFirewallRules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationApplianceVpnVpnFirewallRules

> GetNetworkApplianceFirewallInboundCellularFirewallRules200Response UpdateOrganizationApplianceVpnVpnFirewallRules(ctx, organizationId).UpdateOrganizationApplianceVpnVpnFirewallRulesRequest(updateOrganizationApplianceVpnVpnFirewallRulesRequest).Execute()

Update the firewall rules of an organization's site-to-site VPN



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
	updateOrganizationApplianceVpnVpnFirewallRulesRequest := *openapiclient.NewUpdateOrganizationApplianceVpnVpnFirewallRulesRequest() // UpdateOrganizationApplianceVpnVpnFirewallRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VpnFirewallRulesAPI.UpdateOrganizationApplianceVpnVpnFirewallRules(context.Background(), organizationId).UpdateOrganizationApplianceVpnVpnFirewallRulesRequest(updateOrganizationApplianceVpnVpnFirewallRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VpnFirewallRulesAPI.UpdateOrganizationApplianceVpnVpnFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationApplianceVpnVpnFirewallRules`: GetNetworkApplianceFirewallInboundCellularFirewallRules200Response
	fmt.Fprintf(os.Stdout, "Response from `VpnFirewallRulesAPI.UpdateOrganizationApplianceVpnVpnFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceVpnVpnFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationApplianceVpnVpnFirewallRulesRequest** | [**UpdateOrganizationApplianceVpnVpnFirewallRulesRequest**](UpdateOrganizationApplianceVpnVpnFirewallRulesRequest.md) |  | 

### Return type

[**GetNetworkApplianceFirewallInboundCellularFirewallRules200Response**](GetNetworkApplianceFirewallInboundCellularFirewallRules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

