# \AssetsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationSplashThemeAsset**](AssetsAPI.md#CreateOrganizationSplashThemeAsset) | **Post** /organizations/{organizationId}/splash/themes/{themeIdentifier}/assets | Create a Splash Theme Asset
[**DeleteOrganizationSplashAsset**](AssetsAPI.md#DeleteOrganizationSplashAsset) | **Delete** /organizations/{organizationId}/splash/assets/{id} | Delete a Splash Theme Asset
[**GetOrganizationSplashAsset**](AssetsAPI.md#GetOrganizationSplashAsset) | **Get** /organizations/{organizationId}/splash/assets/{id} | Get a Splash Theme Asset



## CreateOrganizationSplashThemeAsset

> GetOrganizationSplashAsset200Response CreateOrganizationSplashThemeAsset(ctx, organizationId, themeIdentifier).CreateOrganizationSplashThemeAssetRequest(createOrganizationSplashThemeAssetRequest).Execute()

Create a Splash Theme Asset



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
	themeIdentifier := "themeIdentifier_example" // string | Theme identifier
	createOrganizationSplashThemeAssetRequest := *openapiclient.NewCreateOrganizationSplashThemeAssetRequest() // CreateOrganizationSplashThemeAssetRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.CreateOrganizationSplashThemeAsset(context.Background(), organizationId, themeIdentifier).CreateOrganizationSplashThemeAssetRequest(createOrganizationSplashThemeAssetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.CreateOrganizationSplashThemeAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationSplashThemeAsset`: GetOrganizationSplashAsset200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.CreateOrganizationSplashThemeAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**themeIdentifier** | **string** | Theme identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationSplashThemeAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createOrganizationSplashThemeAssetRequest** | [**CreateOrganizationSplashThemeAssetRequest**](CreateOrganizationSplashThemeAssetRequest.md) |  | 

### Return type

[**GetOrganizationSplashAsset200Response**](GetOrganizationSplashAsset200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationSplashAsset

> DeleteOrganizationSplashAsset(ctx, organizationId, id).Execute()

Delete a Splash Theme Asset



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
	id := "id_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AssetsAPI.DeleteOrganizationSplashAsset(context.Background(), organizationId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.DeleteOrganizationSplashAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationSplashAssetRequest struct via the builder pattern


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


## GetOrganizationSplashAsset

> GetOrganizationSplashAsset200Response GetOrganizationSplashAsset(ctx, organizationId, id).Execute()

Get a Splash Theme Asset



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
	id := "id_example" // string | ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetsAPI.GetOrganizationSplashAsset(context.Background(), organizationId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsAPI.GetOrganizationSplashAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSplashAsset`: GetOrganizationSplashAsset200Response
	fmt.Fprintf(os.Stdout, "Response from `AssetsAPI.GetOrganizationSplashAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSplashAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationSplashAsset200Response**](GetOrganizationSplashAsset200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

