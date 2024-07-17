# \IntrusionAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkApplianceSecurityIntrusion**](IntrusionAPI.md#GetNetworkApplianceSecurityIntrusion) | **Get** /networks/{networkId}/appliance/security/intrusion | Returns all supported intrusion settings for an MX network
[**GetOrganizationApplianceSecurityIntrusion**](IntrusionAPI.md#GetOrganizationApplianceSecurityIntrusion) | **Get** /organizations/{organizationId}/appliance/security/intrusion | Returns all supported intrusion settings for an organization
[**UpdateNetworkApplianceSecurityIntrusion**](IntrusionAPI.md#UpdateNetworkApplianceSecurityIntrusion) | **Put** /networks/{networkId}/appliance/security/intrusion | Set the supported intrusion settings for an MX network
[**UpdateOrganizationApplianceSecurityIntrusion**](IntrusionAPI.md#UpdateOrganizationApplianceSecurityIntrusion) | **Put** /organizations/{organizationId}/appliance/security/intrusion | Sets supported intrusion settings for an organization



## GetNetworkApplianceSecurityIntrusion

> GetNetworkApplianceSecurityIntrusion200Response GetNetworkApplianceSecurityIntrusion(ctx, networkId).Execute()

Returns all supported intrusion settings for an MX network



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
	resp, r, err := apiClient.IntrusionAPI.GetNetworkApplianceSecurityIntrusion(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntrusionAPI.GetNetworkApplianceSecurityIntrusion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceSecurityIntrusion`: GetNetworkApplianceSecurityIntrusion200Response
	fmt.Fprintf(os.Stdout, "Response from `IntrusionAPI.GetNetworkApplianceSecurityIntrusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSecurityIntrusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceSecurityIntrusion200Response**](GetNetworkApplianceSecurityIntrusion200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceSecurityIntrusion

> map[string]interface{} GetOrganizationApplianceSecurityIntrusion(ctx, organizationId).Execute()

Returns all supported intrusion settings for an organization



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
	resp, r, err := apiClient.IntrusionAPI.GetOrganizationApplianceSecurityIntrusion(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntrusionAPI.GetOrganizationApplianceSecurityIntrusion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationApplianceSecurityIntrusion`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntrusionAPI.GetOrganizationApplianceSecurityIntrusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceSecurityIntrusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceSecurityIntrusion

> GetNetworkApplianceSecurityIntrusion200Response UpdateNetworkApplianceSecurityIntrusion(ctx, networkId).UpdateNetworkApplianceSecurityIntrusionRequest(updateNetworkApplianceSecurityIntrusionRequest).Execute()

Set the supported intrusion settings for an MX network



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
	updateNetworkApplianceSecurityIntrusionRequest := *openapiclient.NewUpdateNetworkApplianceSecurityIntrusionRequest() // UpdateNetworkApplianceSecurityIntrusionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntrusionAPI.UpdateNetworkApplianceSecurityIntrusion(context.Background(), networkId).UpdateNetworkApplianceSecurityIntrusionRequest(updateNetworkApplianceSecurityIntrusionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntrusionAPI.UpdateNetworkApplianceSecurityIntrusion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceSecurityIntrusion`: GetNetworkApplianceSecurityIntrusion200Response
	fmt.Fprintf(os.Stdout, "Response from `IntrusionAPI.UpdateNetworkApplianceSecurityIntrusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSecurityIntrusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceSecurityIntrusionRequest** | [**UpdateNetworkApplianceSecurityIntrusionRequest**](UpdateNetworkApplianceSecurityIntrusionRequest.md) |  | 

### Return type

[**GetNetworkApplianceSecurityIntrusion200Response**](GetNetworkApplianceSecurityIntrusion200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationApplianceSecurityIntrusion

> map[string]interface{} UpdateOrganizationApplianceSecurityIntrusion(ctx, organizationId).UpdateOrganizationApplianceSecurityIntrusionRequest(updateOrganizationApplianceSecurityIntrusionRequest).Execute()

Sets supported intrusion settings for an organization



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
	updateOrganizationApplianceSecurityIntrusionRequest := *openapiclient.NewUpdateOrganizationApplianceSecurityIntrusionRequest([]openapiclient.UpdateOrganizationApplianceSecurityIntrusionRequestAllowedRulesInner{*openapiclient.NewUpdateOrganizationApplianceSecurityIntrusionRequestAllowedRulesInner("RuleId_example")}) // UpdateOrganizationApplianceSecurityIntrusionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntrusionAPI.UpdateOrganizationApplianceSecurityIntrusion(context.Background(), organizationId).UpdateOrganizationApplianceSecurityIntrusionRequest(updateOrganizationApplianceSecurityIntrusionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntrusionAPI.UpdateOrganizationApplianceSecurityIntrusion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationApplianceSecurityIntrusion`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `IntrusionAPI.UpdateOrganizationApplianceSecurityIntrusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceSecurityIntrusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationApplianceSecurityIntrusionRequest** | [**UpdateOrganizationApplianceSecurityIntrusionRequest**](UpdateOrganizationApplianceSecurityIntrusionRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

