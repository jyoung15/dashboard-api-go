# \QualityRetentionProfilesAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkCameraQualityRetentionProfile**](QualityRetentionProfilesAPI.md#CreateNetworkCameraQualityRetentionProfile) | **Post** /networks/{networkId}/camera/qualityRetentionProfiles | Creates new quality retention profile for this network.
[**DeleteNetworkCameraQualityRetentionProfile**](QualityRetentionProfilesAPI.md#DeleteNetworkCameraQualityRetentionProfile) | **Delete** /networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId} | Delete an existing quality retention profile for this network.
[**GetNetworkCameraQualityRetentionProfile**](QualityRetentionProfilesAPI.md#GetNetworkCameraQualityRetentionProfile) | **Get** /networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId} | Retrieve a single quality retention profile
[**GetNetworkCameraQualityRetentionProfiles**](QualityRetentionProfilesAPI.md#GetNetworkCameraQualityRetentionProfiles) | **Get** /networks/{networkId}/camera/qualityRetentionProfiles | List the quality retention profiles for this network
[**UpdateNetworkCameraQualityRetentionProfile**](QualityRetentionProfilesAPI.md#UpdateNetworkCameraQualityRetentionProfile) | **Put** /networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId} | Update an existing quality retention profile for this network.



## CreateNetworkCameraQualityRetentionProfile

> map[string]interface{} CreateNetworkCameraQualityRetentionProfile(ctx, networkId).CreateNetworkCameraQualityRetentionProfileRequest(createNetworkCameraQualityRetentionProfileRequest).Execute()

Creates new quality retention profile for this network.



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
	createNetworkCameraQualityRetentionProfileRequest := *openapiclient.NewCreateNetworkCameraQualityRetentionProfileRequest("Name_example") // CreateNetworkCameraQualityRetentionProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QualityRetentionProfilesAPI.CreateNetworkCameraQualityRetentionProfile(context.Background(), networkId).CreateNetworkCameraQualityRetentionProfileRequest(createNetworkCameraQualityRetentionProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityRetentionProfilesAPI.CreateNetworkCameraQualityRetentionProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkCameraQualityRetentionProfile`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `QualityRetentionProfilesAPI.CreateNetworkCameraQualityRetentionProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkCameraQualityRetentionProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkCameraQualityRetentionProfileRequest** | [**CreateNetworkCameraQualityRetentionProfileRequest**](CreateNetworkCameraQualityRetentionProfileRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkCameraQualityRetentionProfile

> DeleteNetworkCameraQualityRetentionProfile(ctx, networkId, qualityRetentionProfileId).Execute()

Delete an existing quality retention profile for this network.



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
	qualityRetentionProfileId := "qualityRetentionProfileId_example" // string | Quality retention profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.QualityRetentionProfilesAPI.DeleteNetworkCameraQualityRetentionProfile(context.Background(), networkId, qualityRetentionProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityRetentionProfilesAPI.DeleteNetworkCameraQualityRetentionProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**qualityRetentionProfileId** | **string** | Quality retention profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkCameraQualityRetentionProfileRequest struct via the builder pattern


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


## GetNetworkCameraQualityRetentionProfile

> map[string]interface{} GetNetworkCameraQualityRetentionProfile(ctx, networkId, qualityRetentionProfileId).Execute()

Retrieve a single quality retention profile



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
	qualityRetentionProfileId := "qualityRetentionProfileId_example" // string | Quality retention profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QualityRetentionProfilesAPI.GetNetworkCameraQualityRetentionProfile(context.Background(), networkId, qualityRetentionProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityRetentionProfilesAPI.GetNetworkCameraQualityRetentionProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkCameraQualityRetentionProfile`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `QualityRetentionProfilesAPI.GetNetworkCameraQualityRetentionProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**qualityRetentionProfileId** | **string** | Quality retention profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCameraQualityRetentionProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkCameraQualityRetentionProfiles

> []map[string]interface{} GetNetworkCameraQualityRetentionProfiles(ctx, networkId).Execute()

List the quality retention profiles for this network



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QualityRetentionProfilesAPI.GetNetworkCameraQualityRetentionProfiles(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityRetentionProfilesAPI.GetNetworkCameraQualityRetentionProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkCameraQualityRetentionProfiles`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `QualityRetentionProfilesAPI.GetNetworkCameraQualityRetentionProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCameraQualityRetentionProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkCameraQualityRetentionProfile

> map[string]interface{} UpdateNetworkCameraQualityRetentionProfile(ctx, networkId, qualityRetentionProfileId).UpdateNetworkCameraQualityRetentionProfileRequest(updateNetworkCameraQualityRetentionProfileRequest).Execute()

Update an existing quality retention profile for this network.



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
	qualityRetentionProfileId := "qualityRetentionProfileId_example" // string | Quality retention profile ID
	updateNetworkCameraQualityRetentionProfileRequest := *openapiclient.NewUpdateNetworkCameraQualityRetentionProfileRequest() // UpdateNetworkCameraQualityRetentionProfileRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QualityRetentionProfilesAPI.UpdateNetworkCameraQualityRetentionProfile(context.Background(), networkId, qualityRetentionProfileId).UpdateNetworkCameraQualityRetentionProfileRequest(updateNetworkCameraQualityRetentionProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QualityRetentionProfilesAPI.UpdateNetworkCameraQualityRetentionProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkCameraQualityRetentionProfile`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `QualityRetentionProfilesAPI.UpdateNetworkCameraQualityRetentionProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**qualityRetentionProfileId** | **string** | Quality retention profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkCameraQualityRetentionProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkCameraQualityRetentionProfileRequest** | [**UpdateNetworkCameraQualityRetentionProfileRequest**](UpdateNetworkCameraQualityRetentionProfileRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

