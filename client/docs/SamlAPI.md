# \SamlAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationSamlIdp**](SamlAPI.md#CreateOrganizationSamlIdp) | **Post** /organizations/{organizationId}/saml/idps | Create a SAML IdP for your organization.
[**DeleteOrganizationSamlIdp**](SamlAPI.md#DeleteOrganizationSamlIdp) | **Delete** /organizations/{organizationId}/saml/idps/{idpId} | Remove a SAML IdP in your organization.
[**GetOrganizationSaml**](SamlAPI.md#GetOrganizationSaml) | **Get** /organizations/{organizationId}/saml | Returns the SAML SSO enabled settings for an organization.
[**GetOrganizationSamlIdp**](SamlAPI.md#GetOrganizationSamlIdp) | **Get** /organizations/{organizationId}/saml/idps/{idpId} | Get a SAML IdP from your organization.
[**GetOrganizationSamlIdps**](SamlAPI.md#GetOrganizationSamlIdps) | **Get** /organizations/{organizationId}/saml/idps | List the SAML IdPs in your organization.
[**UpdateOrganizationSaml**](SamlAPI.md#UpdateOrganizationSaml) | **Put** /organizations/{organizationId}/saml | Updates the SAML SSO enabled settings for an organization.
[**UpdateOrganizationSamlIdp**](SamlAPI.md#UpdateOrganizationSamlIdp) | **Put** /organizations/{organizationId}/saml/idps/{idpId} | Update a SAML IdP in your organization



## CreateOrganizationSamlIdp

> []GetOrganizationSamlIdps200ResponseInner CreateOrganizationSamlIdp(ctx, organizationId).CreateOrganizationSamlIdpRequest(createOrganizationSamlIdpRequest).Execute()

Create a SAML IdP for your organization.



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
	createOrganizationSamlIdpRequest := *openapiclient.NewCreateOrganizationSamlIdpRequest("X509certSha1Fingerprint_example") // CreateOrganizationSamlIdpRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SamlAPI.CreateOrganizationSamlIdp(context.Background(), organizationId).CreateOrganizationSamlIdpRequest(createOrganizationSamlIdpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamlAPI.CreateOrganizationSamlIdp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationSamlIdp`: []GetOrganizationSamlIdps200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SamlAPI.CreateOrganizationSamlIdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationSamlIdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationSamlIdpRequest** | [**CreateOrganizationSamlIdpRequest**](CreateOrganizationSamlIdpRequest.md) |  | 

### Return type

[**[]GetOrganizationSamlIdps200ResponseInner**](GetOrganizationSamlIdps200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationSamlIdp

> DeleteOrganizationSamlIdp(ctx, organizationId, idpId).Execute()

Remove a SAML IdP in your organization.



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
	idpId := "idpId_example" // string | Idp ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SamlAPI.DeleteOrganizationSamlIdp(context.Background(), organizationId, idpId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamlAPI.DeleteOrganizationSamlIdp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**idpId** | **string** | Idp ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationSamlIdpRequest struct via the builder pattern


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


## GetOrganizationSaml

> GetOrganizationSaml200Response GetOrganizationSaml(ctx, organizationId).Execute()

Returns the SAML SSO enabled settings for an organization.



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
	resp, r, err := apiClient.SamlAPI.GetOrganizationSaml(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamlAPI.GetOrganizationSaml``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSaml`: GetOrganizationSaml200Response
	fmt.Fprintf(os.Stdout, "Response from `SamlAPI.GetOrganizationSaml`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSamlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrganizationSaml200Response**](GetOrganizationSaml200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSamlIdp

> GetOrganizationSamlIdps200ResponseInner GetOrganizationSamlIdp(ctx, organizationId, idpId).Execute()

Get a SAML IdP from your organization.



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
	idpId := "idpId_example" // string | Idp ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SamlAPI.GetOrganizationSamlIdp(context.Background(), organizationId, idpId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamlAPI.GetOrganizationSamlIdp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSamlIdp`: GetOrganizationSamlIdps200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SamlAPI.GetOrganizationSamlIdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**idpId** | **string** | Idp ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSamlIdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationSamlIdps200ResponseInner**](GetOrganizationSamlIdps200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSamlIdps

> []GetOrganizationSamlIdps200ResponseInner GetOrganizationSamlIdps(ctx, organizationId).Execute()

List the SAML IdPs in your organization.



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
	resp, r, err := apiClient.SamlAPI.GetOrganizationSamlIdps(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamlAPI.GetOrganizationSamlIdps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSamlIdps`: []GetOrganizationSamlIdps200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SamlAPI.GetOrganizationSamlIdps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSamlIdpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetOrganizationSamlIdps200ResponseInner**](GetOrganizationSamlIdps200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationSaml

> GetOrganizationSaml200Response UpdateOrganizationSaml(ctx, organizationId).UpdateOrganizationSamlRequest(updateOrganizationSamlRequest).Execute()

Updates the SAML SSO enabled settings for an organization.



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
	updateOrganizationSamlRequest := *openapiclient.NewUpdateOrganizationSamlRequest() // UpdateOrganizationSamlRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SamlAPI.UpdateOrganizationSaml(context.Background(), organizationId).UpdateOrganizationSamlRequest(updateOrganizationSamlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamlAPI.UpdateOrganizationSaml``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationSaml`: GetOrganizationSaml200Response
	fmt.Fprintf(os.Stdout, "Response from `SamlAPI.UpdateOrganizationSaml`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationSamlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationSamlRequest** | [**UpdateOrganizationSamlRequest**](UpdateOrganizationSamlRequest.md) |  | 

### Return type

[**GetOrganizationSaml200Response**](GetOrganizationSaml200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationSamlIdp

> []GetOrganizationSamlIdps200ResponseInner UpdateOrganizationSamlIdp(ctx, organizationId, idpId).UpdateOrganizationSamlIdpRequest(updateOrganizationSamlIdpRequest).Execute()

Update a SAML IdP in your organization



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
	idpId := "idpId_example" // string | Idp ID
	updateOrganizationSamlIdpRequest := *openapiclient.NewUpdateOrganizationSamlIdpRequest() // UpdateOrganizationSamlIdpRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SamlAPI.UpdateOrganizationSamlIdp(context.Background(), organizationId, idpId).UpdateOrganizationSamlIdpRequest(updateOrganizationSamlIdpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamlAPI.UpdateOrganizationSamlIdp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationSamlIdp`: []GetOrganizationSamlIdps200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SamlAPI.UpdateOrganizationSamlIdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**idpId** | **string** | Idp ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationSamlIdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationSamlIdpRequest** | [**UpdateOrganizationSamlIdpRequest**](UpdateOrganizationSamlIdpRequest.md) |  | 

### Return type

[**[]GetOrganizationSamlIdps200ResponseInner**](GetOrganizationSamlIdps200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

