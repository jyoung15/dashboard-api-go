# \SamlRolesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationSamlRole**](SamlRolesAPI.md#CreateOrganizationSamlRole) | **Post** /organizations/{organizationId}/samlRoles | Create a SAML role
[**DeleteOrganizationSamlRole**](SamlRolesAPI.md#DeleteOrganizationSamlRole) | **Delete** /organizations/{organizationId}/samlRoles/{samlRoleId} | Remove a SAML role
[**GetOrganizationSamlRole**](SamlRolesAPI.md#GetOrganizationSamlRole) | **Get** /organizations/{organizationId}/samlRoles/{samlRoleId} | Return a SAML role
[**GetOrganizationSamlRoles**](SamlRolesAPI.md#GetOrganizationSamlRoles) | **Get** /organizations/{organizationId}/samlRoles | List the SAML roles for this organization
[**UpdateOrganizationSamlRole**](SamlRolesAPI.md#UpdateOrganizationSamlRole) | **Put** /organizations/{organizationId}/samlRoles/{samlRoleId} | Update a SAML role



## CreateOrganizationSamlRole

> GetOrganizationSamlRoles200ResponseInner CreateOrganizationSamlRole(ctx, organizationId).CreateOrganizationSamlRoleRequest(createOrganizationSamlRoleRequest).Execute()

Create a SAML role



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
	createOrganizationSamlRoleRequest := *openapiclient.NewCreateOrganizationSamlRoleRequest("Role_example", "OrgAccess_example") // CreateOrganizationSamlRoleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SamlRolesAPI.CreateOrganizationSamlRole(context.Background(), organizationId).CreateOrganizationSamlRoleRequest(createOrganizationSamlRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamlRolesAPI.CreateOrganizationSamlRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationSamlRole`: GetOrganizationSamlRoles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SamlRolesAPI.CreateOrganizationSamlRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationSamlRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationSamlRoleRequest** | [**CreateOrganizationSamlRoleRequest**](CreateOrganizationSamlRoleRequest.md) |  | 

### Return type

[**GetOrganizationSamlRoles200ResponseInner**](GetOrganizationSamlRoles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationSamlRole

> DeleteOrganizationSamlRole(ctx, organizationId, samlRoleId).Execute()

Remove a SAML role



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
	samlRoleId := "samlRoleId_example" // string | Saml role ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamlRolesAPI.DeleteOrganizationSamlRole(context.Background(), organizationId, samlRoleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamlRolesAPI.DeleteOrganizationSamlRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**samlRoleId** | **string** | Saml role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationSamlRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSamlRole

> GetOrganizationSamlRoles200ResponseInner GetOrganizationSamlRole(ctx, organizationId, samlRoleId).Execute()

Return a SAML role



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
	samlRoleId := "samlRoleId_example" // string | Saml role ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SamlRolesAPI.GetOrganizationSamlRole(context.Background(), organizationId, samlRoleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamlRolesAPI.GetOrganizationSamlRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSamlRole`: GetOrganizationSamlRoles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SamlRolesAPI.GetOrganizationSamlRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**samlRoleId** | **string** | Saml role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSamlRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationSamlRoles200ResponseInner**](GetOrganizationSamlRoles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSamlRoles

> []GetOrganizationSamlRoles200ResponseInner GetOrganizationSamlRoles(ctx, organizationId).Execute()

List the SAML roles for this organization



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
	resp, r, err := apiClient.SamlRolesAPI.GetOrganizationSamlRoles(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamlRolesAPI.GetOrganizationSamlRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSamlRoles`: []GetOrganizationSamlRoles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SamlRolesAPI.GetOrganizationSamlRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSamlRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetOrganizationSamlRoles200ResponseInner**](GetOrganizationSamlRoles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationSamlRole

> GetOrganizationSamlRoles200ResponseInner UpdateOrganizationSamlRole(ctx, organizationId, samlRoleId).UpdateOrganizationSamlRoleRequest(updateOrganizationSamlRoleRequest).Execute()

Update a SAML role



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
	samlRoleId := "samlRoleId_example" // string | Saml role ID
	updateOrganizationSamlRoleRequest := *openapiclient.NewUpdateOrganizationSamlRoleRequest() // UpdateOrganizationSamlRoleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SamlRolesAPI.UpdateOrganizationSamlRole(context.Background(), organizationId, samlRoleId).UpdateOrganizationSamlRoleRequest(updateOrganizationSamlRoleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamlRolesAPI.UpdateOrganizationSamlRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationSamlRole`: GetOrganizationSamlRoles200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SamlRolesAPI.UpdateOrganizationSamlRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**samlRoleId** | **string** | Saml role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationSamlRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationSamlRoleRequest** | [**UpdateOrganizationSamlRoleRequest**](UpdateOrganizationSamlRoleRequest.md) |  | 

### Return type

[**GetOrganizationSamlRoles200ResponseInner**](GetOrganizationSamlRoles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

