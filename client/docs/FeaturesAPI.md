# \FeaturesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationEarlyAccessFeaturesOptIn**](FeaturesAPI.md#CreateOrganizationEarlyAccessFeaturesOptIn) | **Post** /organizations/{organizationId}/earlyAccess/features/optIns | Create a new early access feature opt-in for an organization
[**DeleteOrganizationEarlyAccessFeaturesOptIn**](FeaturesAPI.md#DeleteOrganizationEarlyAccessFeaturesOptIn) | **Delete** /organizations/{organizationId}/earlyAccess/features/optIns/{optInId} | Delete an early access feature opt-in
[**GetOrganizationEarlyAccessFeatures**](FeaturesAPI.md#GetOrganizationEarlyAccessFeatures) | **Get** /organizations/{organizationId}/earlyAccess/features | List the available early access features for organization
[**GetOrganizationEarlyAccessFeaturesOptIn**](FeaturesAPI.md#GetOrganizationEarlyAccessFeaturesOptIn) | **Get** /organizations/{organizationId}/earlyAccess/features/optIns/{optInId} | Show an early access feature opt-in for an organization
[**GetOrganizationEarlyAccessFeaturesOptIns**](FeaturesAPI.md#GetOrganizationEarlyAccessFeaturesOptIns) | **Get** /organizations/{organizationId}/earlyAccess/features/optIns | List the early access feature opt-ins for an organization
[**UpdateOrganizationEarlyAccessFeaturesOptIn**](FeaturesAPI.md#UpdateOrganizationEarlyAccessFeaturesOptIn) | **Put** /organizations/{organizationId}/earlyAccess/features/optIns/{optInId} | Update an early access feature opt-in for an organization



## CreateOrganizationEarlyAccessFeaturesOptIn

> GetOrganizationEarlyAccessFeaturesOptIns200Response CreateOrganizationEarlyAccessFeaturesOptIn(ctx, organizationId).CreateOrganizationEarlyAccessFeaturesOptInRequest(createOrganizationEarlyAccessFeaturesOptInRequest).Execute()

Create a new early access feature opt-in for an organization



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
	createOrganizationEarlyAccessFeaturesOptInRequest := *openapiclient.NewCreateOrganizationEarlyAccessFeaturesOptInRequest("ShortName_example") // CreateOrganizationEarlyAccessFeaturesOptInRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeaturesAPI.CreateOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId).CreateOrganizationEarlyAccessFeaturesOptInRequest(createOrganizationEarlyAccessFeaturesOptInRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.CreateOrganizationEarlyAccessFeaturesOptIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationEarlyAccessFeaturesOptIn`: GetOrganizationEarlyAccessFeaturesOptIns200Response
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.CreateOrganizationEarlyAccessFeaturesOptIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationEarlyAccessFeaturesOptInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationEarlyAccessFeaturesOptInRequest** | [**CreateOrganizationEarlyAccessFeaturesOptInRequest**](CreateOrganizationEarlyAccessFeaturesOptInRequest.md) |  | 

### Return type

[**GetOrganizationEarlyAccessFeaturesOptIns200Response**](GetOrganizationEarlyAccessFeaturesOptIns200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationEarlyAccessFeaturesOptIn

> DeleteOrganizationEarlyAccessFeaturesOptIn(ctx, organizationId, optInId).Execute()

Delete an early access feature opt-in



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
	optInId := "optInId_example" // string | Opt in ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FeaturesAPI.DeleteOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId, optInId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.DeleteOrganizationEarlyAccessFeaturesOptIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**optInId** | **string** | Opt in ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationEarlyAccessFeaturesOptInRequest struct via the builder pattern


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


## GetOrganizationEarlyAccessFeatures

> []GetOrganizationEarlyAccessFeatures200ResponseInner GetOrganizationEarlyAccessFeatures(ctx, organizationId).Execute()

List the available early access features for organization



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
	resp, r, err := apiClient.FeaturesAPI.GetOrganizationEarlyAccessFeatures(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.GetOrganizationEarlyAccessFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationEarlyAccessFeatures`: []GetOrganizationEarlyAccessFeatures200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.GetOrganizationEarlyAccessFeatures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationEarlyAccessFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetOrganizationEarlyAccessFeatures200ResponseInner**](GetOrganizationEarlyAccessFeatures200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationEarlyAccessFeaturesOptIn

> GetOrganizationEarlyAccessFeaturesOptIns200Response GetOrganizationEarlyAccessFeaturesOptIn(ctx, organizationId, optInId).Execute()

Show an early access feature opt-in for an organization



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
	optInId := "optInId_example" // string | Opt in ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeaturesAPI.GetOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId, optInId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.GetOrganizationEarlyAccessFeaturesOptIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationEarlyAccessFeaturesOptIn`: GetOrganizationEarlyAccessFeaturesOptIns200Response
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.GetOrganizationEarlyAccessFeaturesOptIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**optInId** | **string** | Opt in ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationEarlyAccessFeaturesOptInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationEarlyAccessFeaturesOptIns200Response**](GetOrganizationEarlyAccessFeaturesOptIns200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationEarlyAccessFeaturesOptIns

> GetOrganizationEarlyAccessFeaturesOptIns200Response GetOrganizationEarlyAccessFeaturesOptIns(ctx, organizationId).Execute()

List the early access feature opt-ins for an organization



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
	resp, r, err := apiClient.FeaturesAPI.GetOrganizationEarlyAccessFeaturesOptIns(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.GetOrganizationEarlyAccessFeaturesOptIns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationEarlyAccessFeaturesOptIns`: GetOrganizationEarlyAccessFeaturesOptIns200Response
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.GetOrganizationEarlyAccessFeaturesOptIns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationEarlyAccessFeaturesOptInsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrganizationEarlyAccessFeaturesOptIns200Response**](GetOrganizationEarlyAccessFeaturesOptIns200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationEarlyAccessFeaturesOptIn

> GetOrganizationEarlyAccessFeaturesOptIns200Response UpdateOrganizationEarlyAccessFeaturesOptIn(ctx, organizationId, optInId).UpdateOrganizationEarlyAccessFeaturesOptInRequest(updateOrganizationEarlyAccessFeaturesOptInRequest).Execute()

Update an early access feature opt-in for an organization



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
	optInId := "optInId_example" // string | Opt in ID
	updateOrganizationEarlyAccessFeaturesOptInRequest := *openapiclient.NewUpdateOrganizationEarlyAccessFeaturesOptInRequest() // UpdateOrganizationEarlyAccessFeaturesOptInRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeaturesAPI.UpdateOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId, optInId).UpdateOrganizationEarlyAccessFeaturesOptInRequest(updateOrganizationEarlyAccessFeaturesOptInRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeaturesAPI.UpdateOrganizationEarlyAccessFeaturesOptIn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationEarlyAccessFeaturesOptIn`: GetOrganizationEarlyAccessFeaturesOptIns200Response
	fmt.Fprintf(os.Stdout, "Response from `FeaturesAPI.UpdateOrganizationEarlyAccessFeaturesOptIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**optInId** | **string** | Opt in ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationEarlyAccessFeaturesOptInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationEarlyAccessFeaturesOptInRequest** | [**UpdateOrganizationEarlyAccessFeaturesOptInRequest**](UpdateOrganizationEarlyAccessFeaturesOptInRequest.md) |  | 

### Return type

[**GetOrganizationEarlyAccessFeaturesOptIns200Response**](GetOrganizationEarlyAccessFeaturesOptIns200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

