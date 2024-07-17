# \PermissionsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationCameraPermission**](PermissionsAPI.md#GetOrganizationCameraPermission) | **Get** /organizations/{organizationId}/camera/permissions/{permissionScopeId} | Retrieve a single permission scope
[**GetOrganizationCameraPermissions**](PermissionsAPI.md#GetOrganizationCameraPermissions) | **Get** /organizations/{organizationId}/camera/permissions | List the permissions scopes for this organization



## GetOrganizationCameraPermission

> GetOrganizationCameraPermissions200ResponseInner GetOrganizationCameraPermission(ctx, organizationId, permissionScopeId).Execute()

Retrieve a single permission scope



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
	permissionScopeId := "permissionScopeId_example" // string | Permission scope ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.GetOrganizationCameraPermission(context.Background(), organizationId, permissionScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.GetOrganizationCameraPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationCameraPermission`: GetOrganizationCameraPermissions200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.GetOrganizationCameraPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**permissionScopeId** | **string** | Permission scope ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationCameraPermissions200ResponseInner**](GetOrganizationCameraPermissions200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCameraPermissions

> []GetOrganizationCameraPermissions200ResponseInner GetOrganizationCameraPermissions(ctx, organizationId).Execute()

List the permissions scopes for this organization



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
	resp, r, err := apiClient.PermissionsAPI.GetOrganizationCameraPermissions(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.GetOrganizationCameraPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationCameraPermissions`: []GetOrganizationCameraPermissions200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.GetOrganizationCameraPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetOrganizationCameraPermissions200ResponseInner**](GetOrganizationCameraPermissions200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

