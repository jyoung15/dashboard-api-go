# \ThemesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationSplashTheme**](ThemesAPI.md#CreateOrganizationSplashTheme) | **Post** /organizations/{organizationId}/splash/themes | Create a Splash Theme
[**CreateOrganizationSplashThemeAsset**](ThemesAPI.md#CreateOrganizationSplashThemeAsset) | **Post** /organizations/{organizationId}/splash/themes/{themeIdentifier}/assets | Create a Splash Theme Asset
[**DeleteOrganizationSplashTheme**](ThemesAPI.md#DeleteOrganizationSplashTheme) | **Delete** /organizations/{organizationId}/splash/themes/{id} | Delete a Splash Theme
[**GetOrganizationSplashThemes**](ThemesAPI.md#GetOrganizationSplashThemes) | **Get** /organizations/{organizationId}/splash/themes | List Splash Themes



## CreateOrganizationSplashTheme

> GetOrganizationSplashThemes200ResponseInner CreateOrganizationSplashTheme(ctx, organizationId).CreateOrganizationSplashThemeRequest(createOrganizationSplashThemeRequest).Execute()

Create a Splash Theme



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
	createOrganizationSplashThemeRequest := *openapiclient.NewCreateOrganizationSplashThemeRequest() // CreateOrganizationSplashThemeRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThemesAPI.CreateOrganizationSplashTheme(context.Background(), organizationId).CreateOrganizationSplashThemeRequest(createOrganizationSplashThemeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThemesAPI.CreateOrganizationSplashTheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationSplashTheme`: GetOrganizationSplashThemes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ThemesAPI.CreateOrganizationSplashTheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationSplashThemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationSplashThemeRequest** | [**CreateOrganizationSplashThemeRequest**](CreateOrganizationSplashThemeRequest.md) |  | 

### Return type

[**GetOrganizationSplashThemes200ResponseInner**](GetOrganizationSplashThemes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.ThemesAPI.CreateOrganizationSplashThemeAsset(context.Background(), organizationId, themeIdentifier).CreateOrganizationSplashThemeAssetRequest(createOrganizationSplashThemeAssetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThemesAPI.CreateOrganizationSplashThemeAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationSplashThemeAsset`: GetOrganizationSplashAsset200Response
	fmt.Fprintf(os.Stdout, "Response from `ThemesAPI.CreateOrganizationSplashThemeAsset`: %v\n", resp)
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


## DeleteOrganizationSplashTheme

> DeleteOrganizationSplashTheme(ctx, organizationId, id).Execute()

Delete a Splash Theme



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
	r, err := apiClient.ThemesAPI.DeleteOrganizationSplashTheme(context.Background(), organizationId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThemesAPI.DeleteOrganizationSplashTheme``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteOrganizationSplashThemeRequest struct via the builder pattern


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


## GetOrganizationSplashThemes

> []GetOrganizationSplashThemes200ResponseInner GetOrganizationSplashThemes(ctx, organizationId).Execute()

List Splash Themes



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
	resp, r, err := apiClient.ThemesAPI.GetOrganizationSplashThemes(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThemesAPI.GetOrganizationSplashThemes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSplashThemes`: []GetOrganizationSplashThemes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ThemesAPI.GetOrganizationSplashThemes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSplashThemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetOrganizationSplashThemes200ResponseInner**](GetOrganizationSplashThemes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

