# \SplashAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationSplashTheme**](SplashAPI.md#CreateOrganizationSplashTheme) | **Post** /organizations/{organizationId}/splash/themes | Create a Splash Theme
[**CreateOrganizationSplashThemeAsset**](SplashAPI.md#CreateOrganizationSplashThemeAsset) | **Post** /organizations/{organizationId}/splash/themes/{themeIdentifier}/assets | Create a Splash Theme Asset
[**DeleteOrganizationSplashAsset**](SplashAPI.md#DeleteOrganizationSplashAsset) | **Delete** /organizations/{organizationId}/splash/assets/{id} | Delete a Splash Theme Asset
[**DeleteOrganizationSplashTheme**](SplashAPI.md#DeleteOrganizationSplashTheme) | **Delete** /organizations/{organizationId}/splash/themes/{id} | Delete a Splash Theme
[**GetNetworkWirelessSsidSplashSettings**](SplashAPI.md#GetNetworkWirelessSsidSplashSettings) | **Get** /networks/{networkId}/wireless/ssids/{number}/splash/settings | Display the splash page settings for the given SSID
[**GetOrganizationSplashAsset**](SplashAPI.md#GetOrganizationSplashAsset) | **Get** /organizations/{organizationId}/splash/assets/{id} | Get a Splash Theme Asset
[**GetOrganizationSplashThemes**](SplashAPI.md#GetOrganizationSplashThemes) | **Get** /organizations/{organizationId}/splash/themes | List Splash Themes
[**UpdateNetworkWirelessSsidSplashSettings**](SplashAPI.md#UpdateNetworkWirelessSsidSplashSettings) | **Put** /networks/{networkId}/wireless/ssids/{number}/splash/settings | Modify the splash page settings for the given SSID



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
	resp, r, err := apiClient.SplashAPI.CreateOrganizationSplashTheme(context.Background(), organizationId).CreateOrganizationSplashThemeRequest(createOrganizationSplashThemeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SplashAPI.CreateOrganizationSplashTheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationSplashTheme`: GetOrganizationSplashThemes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SplashAPI.CreateOrganizationSplashTheme`: %v\n", resp)
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
	resp, r, err := apiClient.SplashAPI.CreateOrganizationSplashThemeAsset(context.Background(), organizationId, themeIdentifier).CreateOrganizationSplashThemeAssetRequest(createOrganizationSplashThemeAssetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SplashAPI.CreateOrganizationSplashThemeAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationSplashThemeAsset`: GetOrganizationSplashAsset200Response
	fmt.Fprintf(os.Stdout, "Response from `SplashAPI.CreateOrganizationSplashThemeAsset`: %v\n", resp)
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
	r, err := apiClient.SplashAPI.DeleteOrganizationSplashAsset(context.Background(), organizationId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SplashAPI.DeleteOrganizationSplashAsset``: %v\n", err)
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
	r, err := apiClient.SplashAPI.DeleteOrganizationSplashTheme(context.Background(), organizationId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SplashAPI.DeleteOrganizationSplashTheme``: %v\n", err)
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


## GetNetworkWirelessSsidSplashSettings

> GetNetworkWirelessSsidSplashSettings200Response GetNetworkWirelessSsidSplashSettings(ctx, networkId, number).Execute()

Display the splash page settings for the given SSID



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
	number := "number_example" // string | Number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SplashAPI.GetNetworkWirelessSsidSplashSettings(context.Background(), networkId, number).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SplashAPI.GetNetworkWirelessSsidSplashSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkWirelessSsidSplashSettings`: GetNetworkWirelessSsidSplashSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `SplashAPI.GetNetworkWirelessSsidSplashSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidSplashSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkWirelessSsidSplashSettings200Response**](GetNetworkWirelessSsidSplashSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
	resp, r, err := apiClient.SplashAPI.GetOrganizationSplashAsset(context.Background(), organizationId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SplashAPI.GetOrganizationSplashAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSplashAsset`: GetOrganizationSplashAsset200Response
	fmt.Fprintf(os.Stdout, "Response from `SplashAPI.GetOrganizationSplashAsset`: %v\n", resp)
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
	resp, r, err := apiClient.SplashAPI.GetOrganizationSplashThemes(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SplashAPI.GetOrganizationSplashThemes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationSplashThemes`: []GetOrganizationSplashThemes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SplashAPI.GetOrganizationSplashThemes`: %v\n", resp)
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


## UpdateNetworkWirelessSsidSplashSettings

> GetNetworkWirelessSsidSplashSettings200Response UpdateNetworkWirelessSsidSplashSettings(ctx, networkId, number).UpdateNetworkWirelessSsidSplashSettingsRequest(updateNetworkWirelessSsidSplashSettingsRequest).Execute()

Modify the splash page settings for the given SSID



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
	number := "number_example" // string | Number
	updateNetworkWirelessSsidSplashSettingsRequest := *openapiclient.NewUpdateNetworkWirelessSsidSplashSettingsRequest() // UpdateNetworkWirelessSsidSplashSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SplashAPI.UpdateNetworkWirelessSsidSplashSettings(context.Background(), networkId, number).UpdateNetworkWirelessSsidSplashSettingsRequest(updateNetworkWirelessSsidSplashSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SplashAPI.UpdateNetworkWirelessSsidSplashSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkWirelessSsidSplashSettings`: GetNetworkWirelessSsidSplashSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `SplashAPI.UpdateNetworkWirelessSsidSplashSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidSplashSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidSplashSettingsRequest** | [**UpdateNetworkWirelessSsidSplashSettingsRequest**](UpdateNetworkWirelessSsidSplashSettingsRequest.md) |  | 

### Return type

[**GetNetworkWirelessSsidSplashSettings200Response**](GetNetworkWirelessSsidSplashSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

