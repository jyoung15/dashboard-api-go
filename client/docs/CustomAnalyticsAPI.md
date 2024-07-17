# \CustomAnalyticsAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationCameraCustomAnalyticsArtifact**](CustomAnalyticsAPI.md#CreateOrganizationCameraCustomAnalyticsArtifact) | **Post** /organizations/{organizationId}/camera/customAnalytics/artifacts | Create custom analytics artifact
[**DeleteOrganizationCameraCustomAnalyticsArtifact**](CustomAnalyticsAPI.md#DeleteOrganizationCameraCustomAnalyticsArtifact) | **Delete** /organizations/{organizationId}/camera/customAnalytics/artifacts/{artifactId} | Delete Custom Analytics Artifact
[**GetDeviceCameraCustomAnalytics**](CustomAnalyticsAPI.md#GetDeviceCameraCustomAnalytics) | **Get** /devices/{serial}/camera/customAnalytics | Return custom analytics settings for a camera
[**GetOrganizationCameraCustomAnalyticsArtifact**](CustomAnalyticsAPI.md#GetOrganizationCameraCustomAnalyticsArtifact) | **Get** /organizations/{organizationId}/camera/customAnalytics/artifacts/{artifactId} | Get Custom Analytics Artifact
[**GetOrganizationCameraCustomAnalyticsArtifacts**](CustomAnalyticsAPI.md#GetOrganizationCameraCustomAnalyticsArtifacts) | **Get** /organizations/{organizationId}/camera/customAnalytics/artifacts | List Custom Analytics Artifacts
[**UpdateDeviceCameraCustomAnalytics**](CustomAnalyticsAPI.md#UpdateDeviceCameraCustomAnalytics) | **Put** /devices/{serial}/camera/customAnalytics | Update custom analytics settings for a camera



## CreateOrganizationCameraCustomAnalyticsArtifact

> CreateOrganizationCameraCustomAnalyticsArtifact201Response CreateOrganizationCameraCustomAnalyticsArtifact(ctx, organizationId).CreateOrganizationCameraCustomAnalyticsArtifactRequest(createOrganizationCameraCustomAnalyticsArtifactRequest).Execute()

Create custom analytics artifact



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
	createOrganizationCameraCustomAnalyticsArtifactRequest := *openapiclient.NewCreateOrganizationCameraCustomAnalyticsArtifactRequest() // CreateOrganizationCameraCustomAnalyticsArtifactRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomAnalyticsAPI.CreateOrganizationCameraCustomAnalyticsArtifact(context.Background(), organizationId).CreateOrganizationCameraCustomAnalyticsArtifactRequest(createOrganizationCameraCustomAnalyticsArtifactRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomAnalyticsAPI.CreateOrganizationCameraCustomAnalyticsArtifact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationCameraCustomAnalyticsArtifact`: CreateOrganizationCameraCustomAnalyticsArtifact201Response
	fmt.Fprintf(os.Stdout, "Response from `CustomAnalyticsAPI.CreateOrganizationCameraCustomAnalyticsArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationCameraCustomAnalyticsArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationCameraCustomAnalyticsArtifactRequest** | [**CreateOrganizationCameraCustomAnalyticsArtifactRequest**](CreateOrganizationCameraCustomAnalyticsArtifactRequest.md) |  | 

### Return type

[**CreateOrganizationCameraCustomAnalyticsArtifact201Response**](CreateOrganizationCameraCustomAnalyticsArtifact201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationCameraCustomAnalyticsArtifact

> DeleteOrganizationCameraCustomAnalyticsArtifact(ctx, organizationId, artifactId).Execute()

Delete Custom Analytics Artifact



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
	artifactId := "artifactId_example" // string | Artifact ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomAnalyticsAPI.DeleteOrganizationCameraCustomAnalyticsArtifact(context.Background(), organizationId, artifactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomAnalyticsAPI.DeleteOrganizationCameraCustomAnalyticsArtifact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**artifactId** | **string** | Artifact ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationCameraCustomAnalyticsArtifactRequest struct via the builder pattern


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


## GetDeviceCameraCustomAnalytics

> GetDeviceCameraCustomAnalytics200Response GetDeviceCameraCustomAnalytics(ctx, serial).Execute()

Return custom analytics settings for a camera



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
	serial := "serial_example" // string | Serial

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomAnalyticsAPI.GetDeviceCameraCustomAnalytics(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomAnalyticsAPI.GetDeviceCameraCustomAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceCameraCustomAnalytics`: GetDeviceCameraCustomAnalytics200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomAnalyticsAPI.GetDeviceCameraCustomAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraCustomAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceCameraCustomAnalytics200Response**](GetDeviceCameraCustomAnalytics200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCameraCustomAnalyticsArtifact

> GetOrganizationCameraCustomAnalyticsArtifacts200ResponseInner GetOrganizationCameraCustomAnalyticsArtifact(ctx, organizationId, artifactId).Execute()

Get Custom Analytics Artifact



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
	artifactId := "artifactId_example" // string | Artifact ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomAnalyticsAPI.GetOrganizationCameraCustomAnalyticsArtifact(context.Background(), organizationId, artifactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomAnalyticsAPI.GetOrganizationCameraCustomAnalyticsArtifact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationCameraCustomAnalyticsArtifact`: GetOrganizationCameraCustomAnalyticsArtifacts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CustomAnalyticsAPI.GetOrganizationCameraCustomAnalyticsArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**artifactId** | **string** | Artifact ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraCustomAnalyticsArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationCameraCustomAnalyticsArtifacts200ResponseInner**](GetOrganizationCameraCustomAnalyticsArtifacts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCameraCustomAnalyticsArtifacts

> []GetOrganizationCameraCustomAnalyticsArtifacts200ResponseInner GetOrganizationCameraCustomAnalyticsArtifacts(ctx, organizationId).Execute()

List Custom Analytics Artifacts



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
	resp, r, err := apiClient.CustomAnalyticsAPI.GetOrganizationCameraCustomAnalyticsArtifacts(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomAnalyticsAPI.GetOrganizationCameraCustomAnalyticsArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationCameraCustomAnalyticsArtifacts`: []GetOrganizationCameraCustomAnalyticsArtifacts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CustomAnalyticsAPI.GetOrganizationCameraCustomAnalyticsArtifacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraCustomAnalyticsArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetOrganizationCameraCustomAnalyticsArtifacts200ResponseInner**](GetOrganizationCameraCustomAnalyticsArtifacts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceCameraCustomAnalytics

> GetDeviceCameraCustomAnalytics200Response UpdateDeviceCameraCustomAnalytics(ctx, serial).UpdateDeviceCameraCustomAnalyticsRequest(updateDeviceCameraCustomAnalyticsRequest).Execute()

Update custom analytics settings for a camera



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
	serial := "serial_example" // string | Serial
	updateDeviceCameraCustomAnalyticsRequest := *openapiclient.NewUpdateDeviceCameraCustomAnalyticsRequest() // UpdateDeviceCameraCustomAnalyticsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomAnalyticsAPI.UpdateDeviceCameraCustomAnalytics(context.Background(), serial).UpdateDeviceCameraCustomAnalyticsRequest(updateDeviceCameraCustomAnalyticsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomAnalyticsAPI.UpdateDeviceCameraCustomAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceCameraCustomAnalytics`: GetDeviceCameraCustomAnalytics200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomAnalyticsAPI.UpdateDeviceCameraCustomAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCameraCustomAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCameraCustomAnalyticsRequest** | [**UpdateDeviceCameraCustomAnalyticsRequest**](UpdateDeviceCameraCustomAnalyticsRequest.md) |  | 

### Return type

[**GetDeviceCameraCustomAnalytics200Response**](GetDeviceCameraCustomAnalytics200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

