# \LoginSecurityAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationLoginSecurity**](LoginSecurityAPI.md#GetOrganizationLoginSecurity) | **Get** /organizations/{organizationId}/loginSecurity | Returns the login security settings for an organization.
[**UpdateOrganizationLoginSecurity**](LoginSecurityAPI.md#UpdateOrganizationLoginSecurity) | **Put** /organizations/{organizationId}/loginSecurity | Update the login security settings for an organization



## GetOrganizationLoginSecurity

> GetOrganizationLoginSecurity200Response GetOrganizationLoginSecurity(ctx, organizationId).Execute()

Returns the login security settings for an organization.



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
	resp, r, err := apiClient.LoginSecurityAPI.GetOrganizationLoginSecurity(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoginSecurityAPI.GetOrganizationLoginSecurity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationLoginSecurity`: GetOrganizationLoginSecurity200Response
	fmt.Fprintf(os.Stdout, "Response from `LoginSecurityAPI.GetOrganizationLoginSecurity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationLoginSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrganizationLoginSecurity200Response**](GetOrganizationLoginSecurity200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationLoginSecurity

> GetOrganizationLoginSecurity200Response UpdateOrganizationLoginSecurity(ctx, organizationId).UpdateOrganizationLoginSecurityRequest(updateOrganizationLoginSecurityRequest).Execute()

Update the login security settings for an organization



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
	updateOrganizationLoginSecurityRequest := *openapiclient.NewUpdateOrganizationLoginSecurityRequest() // UpdateOrganizationLoginSecurityRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoginSecurityAPI.UpdateOrganizationLoginSecurity(context.Background(), organizationId).UpdateOrganizationLoginSecurityRequest(updateOrganizationLoginSecurityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoginSecurityAPI.UpdateOrganizationLoginSecurity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationLoginSecurity`: GetOrganizationLoginSecurity200Response
	fmt.Fprintf(os.Stdout, "Response from `LoginSecurityAPI.UpdateOrganizationLoginSecurity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationLoginSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationLoginSecurityRequest** | [**UpdateOrganizationLoginSecurityRequest**](UpdateOrganizationLoginSecurityRequest.md) |  | 

### Return type

[**GetOrganizationLoginSecurity200Response**](GetOrganizationLoginSecurity200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

