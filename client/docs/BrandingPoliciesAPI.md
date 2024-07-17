# \BrandingPoliciesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationBrandingPolicy**](BrandingPoliciesAPI.md#CreateOrganizationBrandingPolicy) | **Post** /organizations/{organizationId}/brandingPolicies | Add a new branding policy to an organization
[**DeleteOrganizationBrandingPolicy**](BrandingPoliciesAPI.md#DeleteOrganizationBrandingPolicy) | **Delete** /organizations/{organizationId}/brandingPolicies/{brandingPolicyId} | Delete a branding policy
[**GetOrganizationBrandingPolicies**](BrandingPoliciesAPI.md#GetOrganizationBrandingPolicies) | **Get** /organizations/{organizationId}/brandingPolicies | List the branding policies of an organization
[**GetOrganizationBrandingPoliciesPriorities**](BrandingPoliciesAPI.md#GetOrganizationBrandingPoliciesPriorities) | **Get** /organizations/{organizationId}/brandingPolicies/priorities | Return the branding policy IDs of an organization in priority order
[**GetOrganizationBrandingPolicy**](BrandingPoliciesAPI.md#GetOrganizationBrandingPolicy) | **Get** /organizations/{organizationId}/brandingPolicies/{brandingPolicyId} | Return a branding policy
[**UpdateOrganizationBrandingPoliciesPriorities**](BrandingPoliciesAPI.md#UpdateOrganizationBrandingPoliciesPriorities) | **Put** /organizations/{organizationId}/brandingPolicies/priorities | Update the priority ordering of an organization&#39;s branding policies.
[**UpdateOrganizationBrandingPolicy**](BrandingPoliciesAPI.md#UpdateOrganizationBrandingPolicy) | **Put** /organizations/{organizationId}/brandingPolicies/{brandingPolicyId} | Update a branding policy



## CreateOrganizationBrandingPolicy

> CreateOrganizationBrandingPolicy201Response CreateOrganizationBrandingPolicy(ctx, organizationId).CreateOrganizationBrandingPolicyRequest(createOrganizationBrandingPolicyRequest).Execute()

Add a new branding policy to an organization



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
	createOrganizationBrandingPolicyRequest := *openapiclient.NewCreateOrganizationBrandingPolicyRequest() // CreateOrganizationBrandingPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BrandingPoliciesAPI.CreateOrganizationBrandingPolicy(context.Background(), organizationId).CreateOrganizationBrandingPolicyRequest(createOrganizationBrandingPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandingPoliciesAPI.CreateOrganizationBrandingPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationBrandingPolicy`: CreateOrganizationBrandingPolicy201Response
	fmt.Fprintf(os.Stdout, "Response from `BrandingPoliciesAPI.CreateOrganizationBrandingPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationBrandingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationBrandingPolicyRequest** | [**CreateOrganizationBrandingPolicyRequest**](CreateOrganizationBrandingPolicyRequest.md) |  | 

### Return type

[**CreateOrganizationBrandingPolicy201Response**](CreateOrganizationBrandingPolicy201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationBrandingPolicy

> DeleteOrganizationBrandingPolicy(ctx, organizationId, brandingPolicyId).Execute()

Delete a branding policy



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
	brandingPolicyId := "brandingPolicyId_example" // string | Branding policy ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BrandingPoliciesAPI.DeleteOrganizationBrandingPolicy(context.Background(), organizationId, brandingPolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandingPoliciesAPI.DeleteOrganizationBrandingPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**brandingPolicyId** | **string** | Branding policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationBrandingPolicyRequest struct via the builder pattern


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


## GetOrganizationBrandingPolicies

> []GetOrganizationBrandingPolicies200ResponseInner GetOrganizationBrandingPolicies(ctx, organizationId).Execute()

List the branding policies of an organization



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
	resp, r, err := apiClient.BrandingPoliciesAPI.GetOrganizationBrandingPolicies(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandingPoliciesAPI.GetOrganizationBrandingPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationBrandingPolicies`: []GetOrganizationBrandingPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `BrandingPoliciesAPI.GetOrganizationBrandingPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationBrandingPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetOrganizationBrandingPolicies200ResponseInner**](GetOrganizationBrandingPolicies200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationBrandingPoliciesPriorities

> GetOrganizationBrandingPoliciesPriorities200Response GetOrganizationBrandingPoliciesPriorities(ctx, organizationId).Execute()

Return the branding policy IDs of an organization in priority order



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
	resp, r, err := apiClient.BrandingPoliciesAPI.GetOrganizationBrandingPoliciesPriorities(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandingPoliciesAPI.GetOrganizationBrandingPoliciesPriorities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationBrandingPoliciesPriorities`: GetOrganizationBrandingPoliciesPriorities200Response
	fmt.Fprintf(os.Stdout, "Response from `BrandingPoliciesAPI.GetOrganizationBrandingPoliciesPriorities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationBrandingPoliciesPrioritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrganizationBrandingPoliciesPriorities200Response**](GetOrganizationBrandingPoliciesPriorities200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationBrandingPolicy

> GetOrganizationBrandingPolicies200ResponseInner GetOrganizationBrandingPolicy(ctx, organizationId, brandingPolicyId).Execute()

Return a branding policy



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
	brandingPolicyId := "brandingPolicyId_example" // string | Branding policy ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BrandingPoliciesAPI.GetOrganizationBrandingPolicy(context.Background(), organizationId, brandingPolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandingPoliciesAPI.GetOrganizationBrandingPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationBrandingPolicy`: GetOrganizationBrandingPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `BrandingPoliciesAPI.GetOrganizationBrandingPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**brandingPolicyId** | **string** | Branding policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationBrandingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationBrandingPolicies200ResponseInner**](GetOrganizationBrandingPolicies200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationBrandingPoliciesPriorities

> GetOrganizationBrandingPoliciesPriorities200Response UpdateOrganizationBrandingPoliciesPriorities(ctx, organizationId).UpdateOrganizationBrandingPoliciesPrioritiesRequest(updateOrganizationBrandingPoliciesPrioritiesRequest).Execute()

Update the priority ordering of an organization's branding policies.



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
	updateOrganizationBrandingPoliciesPrioritiesRequest := *openapiclient.NewUpdateOrganizationBrandingPoliciesPrioritiesRequest() // UpdateOrganizationBrandingPoliciesPrioritiesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BrandingPoliciesAPI.UpdateOrganizationBrandingPoliciesPriorities(context.Background(), organizationId).UpdateOrganizationBrandingPoliciesPrioritiesRequest(updateOrganizationBrandingPoliciesPrioritiesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandingPoliciesAPI.UpdateOrganizationBrandingPoliciesPriorities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationBrandingPoliciesPriorities`: GetOrganizationBrandingPoliciesPriorities200Response
	fmt.Fprintf(os.Stdout, "Response from `BrandingPoliciesAPI.UpdateOrganizationBrandingPoliciesPriorities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationBrandingPoliciesPrioritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationBrandingPoliciesPrioritiesRequest** | [**UpdateOrganizationBrandingPoliciesPrioritiesRequest**](UpdateOrganizationBrandingPoliciesPrioritiesRequest.md) |  | 

### Return type

[**GetOrganizationBrandingPoliciesPriorities200Response**](GetOrganizationBrandingPoliciesPriorities200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationBrandingPolicy

> GetOrganizationBrandingPolicies200ResponseInner UpdateOrganizationBrandingPolicy(ctx, organizationId, brandingPolicyId).UpdateOrganizationBrandingPolicyRequest(updateOrganizationBrandingPolicyRequest).Execute()

Update a branding policy



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
	brandingPolicyId := "brandingPolicyId_example" // string | Branding policy ID
	updateOrganizationBrandingPolicyRequest := *openapiclient.NewUpdateOrganizationBrandingPolicyRequest() // UpdateOrganizationBrandingPolicyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BrandingPoliciesAPI.UpdateOrganizationBrandingPolicy(context.Background(), organizationId, brandingPolicyId).UpdateOrganizationBrandingPolicyRequest(updateOrganizationBrandingPolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandingPoliciesAPI.UpdateOrganizationBrandingPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationBrandingPolicy`: GetOrganizationBrandingPolicies200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `BrandingPoliciesAPI.UpdateOrganizationBrandingPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**brandingPolicyId** | **string** | Branding policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationBrandingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationBrandingPolicyRequest** | [**UpdateOrganizationBrandingPolicyRequest**](UpdateOrganizationBrandingPolicyRequest.md) |  | 

### Return type

[**GetOrganizationBrandingPolicies200ResponseInner**](GetOrganizationBrandingPolicies200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

